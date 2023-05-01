package kaytu_aws_describer

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"io"

	"github.com/kaytu-io/kaytu-aws-describer/proto/src/golang"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

type ResourceSender struct {
	authToken        string
	workspaceName    string
	logger           *zap.Logger
	resourceChannel  chan *golang.AWSResource
	resourceIDs      []string
	doneChannel      chan interface{}
	conn             *grpc.ClientConn
	describeEndpoint string
	jobID            uint
	client           golang.DescribeServiceClient
}

func NewResourceSender(workspaceName string, describeEndpoint string, describeToken string, jobID uint, logger *zap.Logger) (*ResourceSender, error) {
	rs := ResourceSender{
		authToken:        describeToken,
		workspaceName:    workspaceName,
		logger:           logger,
		resourceChannel:  make(chan *golang.AWSResource, 1000),
		resourceIDs:      nil,
		doneChannel:      make(chan interface{}),
		conn:             nil,
		describeEndpoint: describeEndpoint,
		jobID:            jobID,
	}
	if err := rs.Connect(); err != nil {
		return nil, err
	}

	go rs.ResourceHandler()
	return &rs, nil
}

func (s *ResourceSender) Connect() error {
	conn, err := grpc.Dial(
		s.describeEndpoint,
		grpc.WithTransportCredentials(credentials.NewTLS(nil)),
		grpc.WithPerRPCCredentials(oauth.TokenSource{
			TokenSource: oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: s.authToken,
			}),
		}),
	)
	if err != nil {
		return err
	}
	s.conn = conn

	client := golang.NewDescribeServiceClient(conn)
	s.client = client
	return nil
}

func (s *ResourceSender) ResourceHandler() {
	for resource := range s.resourceChannel {
		if resource == nil {
			s.doneChannel <- struct{}{}
			return
		}

		s.resourceIDs = append(s.resourceIDs, resource.UniqueId)

		grpcCtx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
			"workspace-name":  s.workspaceName,
			"resource-job-id": fmt.Sprintf("%d", s.jobID),
		}))
		_, err := s.client.DeliverAWSResources(grpcCtx, resource)
		if err != nil {
			s.logger.Error("failed to send resource", zap.Error(err))
			if errors.Is(err, io.EOF) {
				err = s.Connect()
				if err != nil {
					s.logger.Error("failed to reconnect", zap.Error(err))
				} else {
					s.resourceChannel <- resource
				}
			}
			continue
		}
	}
}

func (s *ResourceSender) Finish() {
	s.resourceChannel <- nil
	_ = <-s.doneChannel
	s.conn.Close()
}

func (s *ResourceSender) GetResourceIDs() []string {
	return s.resourceIDs
}

func (s *ResourceSender) Send(resource *golang.AWSResource) {
	s.resourceChannel <- resource
}
