package describer

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaytu-io/kaytu-util/proto/src/golang"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"

	"github.com/kaytu-io/kaytu-util/pkg/describe"
	"github.com/kaytu-io/kaytu-util/pkg/vault"
	"go.uber.org/zap"
)

const (
	DescribeResourceJobFailed    string = "FAILED"
	DescribeResourceJobSucceeded string = "SUCCEEDED"
)

func getJWTAuthToken(workspaceId string) (string, error) {
	privateKey, ok := os.LookupEnv("JWT_PRIVATE_KEY")
	if !ok {
		return "", fmt.Errorf("JWT_PRIVATE_KEY not set")
	}

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("JWT_PRIVATE_KEY not base64 encoded")
	}

	pk, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("JWT_PRIVATE_KEY not valid")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"https://app.keibi.io/workspaceAccess": map[string]string{
			workspaceId: "admin",
		},
		"https://app.keibi.io/email": "lambda-worker@kaytu.io",
	}).SignedString(pk)
	if err != nil {
		return "", fmt.Errorf("JWT token generation failed %v", err)
	}
	return token, nil
}

func DescribeHandler(ctx context.Context, input describe.LambdaDescribeWorkerInput) error {
	fmt.Printf("Input: %v", input)

	var err error
	logger := zap.NewNop()

	if debug := os.Getenv("DEBUG"); debug == "true" {
		logger, err = zap.NewProduction()
		if err != nil {
			return err
		}
	}

	if input.WorkspaceName == "" {
		return fmt.Errorf("workspace name is required")
	}

	kmsVault, err := vault.NewKMSVaultSourceConfig(ctx, "", "", input.KeyRegion)
	if err != nil {
		return fmt.Errorf("failed to initialize KMS vault: %w", err)
	}

	token, err := getJWTAuthToken(input.WorkspaceId)
	if err != nil {
		return fmt.Errorf("failed to get JWT token: %w", err)
	}

	resourceIds, err := Do(
		ctx,
		kmsVault,
		logger,
		input.DescribeJob,
		input.KeyARN,
		input.DescribeEndpoint,
		token,
		input.WorkspaceName,
	)

	errMsg := ""
	errCode := ""
	status := DescribeResourceJobSucceeded
	if err != nil {
		errMsg = err.Error()
		var kerr KaytuError
		if errors.As(err, &kerr) {
			errCode = kerr.ErrCode
		}
		status = DescribeResourceJobFailed
	}

	for retry := 0; retry < 5; retry++ {
		conn, err := grpc.Dial(
			input.DescribeEndpoint,
			grpc.WithTransportCredentials(credentials.NewTLS(nil)),
			grpc.WithPerRPCCredentials(oauth.TokenSource{
				TokenSource: oauth2.StaticTokenSource(&oauth2.Token{
					AccessToken: token,
				}),
			}),
		)
		if err != nil {
			logger.Error("[result delivery] connection failure:", zap.Error(err))
			time.Sleep(1 * time.Second)
			continue
		}
		client := golang.NewDescribeServiceClient(conn)

		grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"workspace-name": input.WorkspaceName,
		}))
		_, err = client.DeliverResult(grpcCtx, &golang.DeliverResultRequest{
			JobId:       uint32(input.DescribeJob.JobID),
			ParentJobId: uint32(input.DescribeJob.ParentJobID),
			Status:      status,
			Error:       errMsg,
			ErrorCode:   errCode,
			DescribeJob: &golang.DescribeJob{
				JobId:         uint32(input.DescribeJob.JobID),
				ScheduleJobId: uint32(input.DescribeJob.ScheduleJobID),
				ParentJobId:   uint32(input.DescribeJob.ParentJobID),
				ResourceType:  input.DescribeJob.ResourceType,
				SourceId:      input.DescribeJob.SourceID,
				AccountId:     input.DescribeJob.AccountID,
				DescribedAt:   input.DescribeJob.DescribedAt,
				SourceType:    string(input.DescribeJob.SourceType),
				ConfigReg:     input.DescribeJob.CipherText,
				TriggerType:   string(input.DescribeJob.TriggerType),
				RetryCounter:  uint32(input.DescribeJob.RetryCounter),
			},
			DescribedResourceIds: resourceIds,
		})
		if err != nil {
			logger.Error("[result delivery] rpc failed:", zap.Error(err))
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	return err
}
