package describer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	s3controltypes "github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

const (
	s3NoSuchBucketPolicy                             = "NoSuchBucketPolicy"
	s3NoSuchLifecycleConfiguration                   = "NoSuchLifecycleConfiguration"
	s3NoSuchPublicAccessBlockConfiguration           = "NoSuchPublicAccessBlockConfiguration"
	s3ObjectLockConfigurationNotFoundError           = "ObjectLockConfigurationNotFoundError"
	s3ReplicationConfigurationNotFoundError          = "ReplicationConfigurationNotFoundError"
	s3ServerSideEncryptionConfigurationNotFoundError = "ServerSideEncryptionConfigurationNotFoundError"
	s3BucketNoOfWorkers                              = 8
)

type s3bucketResult struct {
	Bucket   types.Bucket
	Resource Resource
	Region   string
	Err      error
}

// S3Bucket describe S3 buckets.
// ListBuckets returns buckets in all regions. However, this function categorizes the buckets based
// on their location constaint, aka the regions they reside in.
func S3Bucket(ctx context.Context, cfg aws.Config, regions []string, stream *StreamSender) (map[string][]Resource, error) {
	regionalValues := make(map[string][]Resource, len(regions))
	for _, r := range regions {
		regionalValues[r] = make([]Resource, 0)
	}

	describeCtx := GetDescribeContext(ctx)

	client := s3.NewFromConfig(cfg)
	output, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		fmt.Println("Error listing buckets")
		return nil, err
	}

	done := make(chan interface{})
	jobChan := make(chan types.Bucket, len(output.Buckets)+1)
	resultChan := make(chan s3bucketResult, len(output.Buckets)+1)

	describer := func(bucket types.Bucket) {
		region, err := getBucketLocation(ctx, client, bucket)
		if err != nil {
			resultChan <- s3bucketResult{
				Err:    err,
				Bucket: bucket,
			}
			fmt.Println("S3Bucket Error", "get location error", bucket.Name)
			fmt.Println(err)
			return
		}

		if !isIncludedInRegions(regions, region) {
			resultChan <- s3bucketResult{
				Bucket: bucket,
			}
			fmt.Println("S3Bucket Error", "not included in regions", bucket.Name)
			return
		}

		desc, err := getBucketDescription(ctx, cfg, bucket, region)
		if err != nil {
			resultChan <- s3bucketResult{
				Bucket: bucket,
				Err:    err,
			}
			fmt.Println("S3Bucket Error", "describing error", bucket.Name)
			fmt.Println(err)
			return
		}

		arn := "arn:" + describeCtx.Partition + ":s3:::" + *bucket.Name
		resultChan <- s3bucketResult{
			Resource: Resource{
				Region:      describeCtx.Region,
				ARN:         arn,
				Name:        *bucket.Name,
				Description: desc,
			},
			Region: region,
			Bucket: bucket,
		}
	}

	worker := func() {
		for {
			select {
			case j := <-jobChan:
				describer(j)
			case <-done:
				return
			}
		}
	}

	for i := 0; i < s3BucketNoOfWorkers; i++ {
		go worker()
	}

	for _, bucket := range output.Buckets {
		jobChan <- bucket
	}

	var globalErr error
	for range output.Buckets {
		res := <-resultChan
		if res.Err != nil {
			globalErr = res.Err
			fmt.Println("S3Bucket Error", "global", res)
		} else if res.Region != "" {
			if _, ok := regionalValues[res.Region]; ok {
				regionalValues[res.Region] = append(regionalValues[res.Region], res.Resource)
			}
		}
	}

	for i := 0; i < s3BucketNoOfWorkers; i++ {
		done <- i
	}
	if stream != nil {
		for _, v := range regionalValues {
			for _, resource := range v {
				if err := (*stream)(resource); err != nil {
					return regionalValues, err
				}
			}
		}
		return map[string][]Resource{}, nil
	}
	return regionalValues, globalErr
}

func GetS3Bucket(ctx context.Context, cfg aws.Config, regions []string, bucketName string) (map[string][]Resource, error) {
	regionalValues := make(map[string][]Resource, len(regions))
	for _, r := range regions {
		regionalValues[r] = make([]Resource, 0)
	}

	describeCtx := GetDescribeContext(ctx)

	client := s3.NewFromConfig(cfg)
	output, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		fmt.Println("Error listing buckets")
		return nil, err
	}

	for _, bucket := range output.Buckets {
		if *bucket.Name != bucketName {
			continue
		}

		region, err := getBucketLocation(ctx, client, bucket)
		if err != nil {
			return nil, err
		}

		if !isIncludedInRegions(regions, region) {
			return nil, fmt.Errorf("not included in regions %s", *bucket.Name)
		}

		desc, err := getBucketDescription(ctx, cfg, bucket, region)
		if err != nil {
			return nil, err
		}

		arn := "arn:" + describeCtx.Partition + ":s3:::" + *bucket.Name

		regionalValues[region] = append(regionalValues[region], Resource{
			Region:      describeCtx.Region,
			ARN:         arn,
			Name:        *bucket.Name,
			Description: desc,
		})
	}
	return regionalValues, nil
}

func getBucketLocation(ctx context.Context, client *s3.Client, bucket types.Bucket) (string, error) {
	output, err := client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		return "", err
	}

	region := string(output.LocationConstraint)
	if region == "" {
		// Buckets in Region us-east-1 have a LocationConstraint of null.
		region = "us-east-1"
	}

	return region, nil
}

func getBucketDescription(ctx context.Context, cfg aws.Config, bucket types.Bucket, region string) (*model.S3BucketDescription, error) {
	rClient := s3.NewFromConfig(cfg, func(o *s3.Options) { o.Region = region })
	o1, err := getBucketIsPublic(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o2, err := getBucketVersioning(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o3, err := getBucketEncryption(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o4, err := getBucketPublicAccessBlock(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o5, err := getBucketACL(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o6, err := getBucketLifecycle(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o7, err := getBucketLogging(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o8, err := getBucketPolicy(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o9, err := getBucketReplication(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o10, err := getObjectLockConfiguration(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	o11, err := getBucketTagging(ctx, rClient, bucket)
	if err != nil {
		return nil, err
	}

	rulesJson, err := json.Marshal(o6.Rules)
	if err != nil {
		return nil, err
	}

	replicationJson, err := json.Marshal(o9.ReplicationConfiguration)
	if err != nil {
		return nil, err
	}

	return &model.S3BucketDescription{
		Bucket: bucket,
		BucketAcl: struct {
			Grants []types.Grant
			Owner  *types.Owner
		}{
			Grants: o5.Grants,
			Owner:  o5.Owner,
		},
		Policy:                         o8.Policy,
		PolicyStatus:                   o1.PolicyStatus,
		PublicAccessBlockConfiguration: o4.PublicAccessBlockConfiguration,
		Versioning: struct {
			MFADelete types.MFADeleteStatus
			Status    types.BucketVersioningStatus
		}{
			MFADelete: o2.MFADelete,
			Status:    o2.Status,
		},
		LifecycleRules:                    string(rulesJson),
		LoggingEnabled:                    o7.LoggingEnabled,
		ServerSideEncryptionConfiguration: o3.ServerSideEncryptionConfiguration,
		ObjectLockConfiguration:           o10.ObjectLockConfiguration,
		ReplicationConfiguration:          string(replicationJson),
		Tags:                              o11.TagSet,
	}, nil
}

func getBucketIsPublic(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketPolicyStatusOutput, error) {
	output, err := client.GetBucketPolicyStatus(ctx, &s3.GetBucketPolicyStatusInput{
		Bucket: bucket.Name,
	})

	if err != nil {
		if isErr(err, s3NoSuchBucketPolicy) {
			return &s3.GetBucketPolicyStatusOutput{}, nil
		}

		return nil, err
	}

	return output, nil
}

func getBucketVersioning(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketVersioningOutput, error) {
	output, err := client.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}

func getBucketEncryption(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketEncryptionOutput, error) {
	output, err := client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		if isErr(err, s3ServerSideEncryptionConfigurationNotFoundError) {
			return &s3.GetBucketEncryptionOutput{}, nil
		}

		return nil, err
	}

	return output, nil
}

func getBucketPublicAccessBlock(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetPublicAccessBlockOutput, error) {
	output, err := client.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		// If the GetPublicAccessBlock is called on buckets which were created before Public Access Block setting was
		// introduced, sometime it fails with error NoSuchPublicAccessBlockConfiguration
		if isErr(err, s3NoSuchPublicAccessBlockConfiguration) {
			return &s3.GetPublicAccessBlockOutput{
				PublicAccessBlockConfiguration: &types.PublicAccessBlockConfiguration{
					BlockPublicAcls:       false,
					BlockPublicPolicy:     false,
					IgnorePublicAcls:      false,
					RestrictPublicBuckets: false,
				},
			}, nil
		}

		return nil, err
	}

	return output, nil
}

func getBucketACL(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketAclOutput, error) {
	output, err := client.GetBucketAcl(ctx, &s3.GetBucketAclInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}

func getBucketLifecycle(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	output, err := client.GetBucketLifecycleConfiguration(ctx, &s3.GetBucketLifecycleConfigurationInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		if isErr(err, s3NoSuchLifecycleConfiguration) {
			return &s3.GetBucketLifecycleConfigurationOutput{}, nil
		}

		return nil, err
	}

	return output, nil
}

func getBucketLogging(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketLoggingOutput, error) {
	output, err := client.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		return nil, err
	}

	return output, nil
}

func getBucketPolicy(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketPolicyOutput, error) {
	output, err := client.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		if isErr(err, s3NoSuchBucketPolicy) {
			return &s3.GetBucketPolicyOutput{}, nil
		}

		return nil, err
	}

	return output, nil
}

func getBucketReplication(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketReplicationOutput, error) {
	output, err := client.GetBucketReplication(ctx, &s3.GetBucketReplicationInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		if isErr(err, s3ReplicationConfigurationNotFoundError) {
			return &s3.GetBucketReplicationOutput{}, nil
		}

		return nil, err
	}

	return output, nil
}

func getObjectLockConfiguration(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetObjectLockConfigurationOutput, error) {
	output, err := client.GetObjectLockConfiguration(ctx, &s3.GetObjectLockConfigurationInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		if isErr(err, s3ObjectLockConfigurationNotFoundError) {
			return &s3.GetObjectLockConfigurationOutput{}, nil
		}

		return nil, err
	}

	return output, nil
}

func getBucketTagging(ctx context.Context, client *s3.Client, bucket types.Bucket) (*s3.GetBucketTaggingOutput, error) {
	output, err := client.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{
		Bucket: bucket.Name,
	})
	if err != nil {
		if isErr(err, "NoSuchTagSet") {
			return &s3.GetBucketTaggingOutput{}, nil
		}
		return nil, err
	}

	return output, nil
}

func isIncludedInRegions(regions []string, region string) bool {
	for _, region := range regions {
		if strings.EqualFold(region, region) {
			return true
		}
	}

	return false
}

func isErr(err error, code string) bool {
	var ae smithy.APIError
	return errors.As(err, &ae) && ae.ErrorCode() == code
}

func S3AccessPoint(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	stsClient := sts.NewFromConfig(cfg)
	output, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, err
	}

	client := s3control.NewFromConfig(cfg)
	paginator := s3control.NewListAccessPointsPaginator(client, &s3control.ListAccessPointsInput{
		AccountId: output.Account,
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.AccessPointList {
			ap, err := client.GetAccessPoint(ctx, &s3control.GetAccessPointInput{
				Name:      v.Name,
				AccountId: output.Account,
			})
			if err != nil {
				return nil, err
			}

			params := &s3control.GetAccessPointPolicyInput{
				Name:      v.Name,
				AccountId: output.Account,
			}
			app, err := client.GetAccessPointPolicy(ctx, params)
			if err != nil {
				if !isErr(err, "NoSuchAccessPointPolicy") {
					return nil, err
				}
				app = &s3control.GetAccessPointPolicyOutput{}
			}

			appsParams := &s3control.GetAccessPointPolicyStatusInput{
				Name:      v.Name,
				AccountId: output.Account,
			}
			apps, err := client.GetAccessPointPolicyStatus(ctx, appsParams)
			if err != nil {
				if !isErr(err, "NoSuchAccessPointPolicy") {
					return nil, err
				}
				apps = &s3control.GetAccessPointPolicyStatusOutput{}
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.Region,
				ARN:    *v.AccessPointArn,
				Name:   *v.Name,
				Description: model.S3AccessPointDescription{
					AccessPoint:  ap,
					Policy:       app.Policy,
					PolicyStatus: apps.PolicyStatus,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func S3StorageLens(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	stsClient := sts.NewFromConfig(cfg)
	output, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return nil, err
	}

	client := s3control.NewFromConfig(cfg)
	paginator := s3control.NewListStorageLensConfigurationsPaginator(client, &s3control.ListStorageLensConfigurationsInput{
		AccountId: output.Account,
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.StorageLensConfigurationList {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.StorageLensArn,
				Name:        *v.Id,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func S3AccountSetting(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	accountId, err := STSAccount(ctx, cfg)
	if err != nil {
		return nil, err
	}

	client := s3control.NewFromConfig(cfg)
	output, err := client.GetPublicAccessBlock(ctx, &s3control.GetPublicAccessBlockInput{
		AccountId: &accountId,
	})
	if err != nil {
		if !isErr(err, s3NoSuchPublicAccessBlockConfiguration) {
			return nil, err
		}

		output = &s3control.GetPublicAccessBlockOutput{
			PublicAccessBlockConfiguration: &s3controltypes.PublicAccessBlockConfiguration{
				BlockPublicAcls:       false,
				BlockPublicPolicy:     false,
				IgnorePublicAcls:      false,
				RestrictPublicBuckets: false,
			},
		}
	}

	var values []Resource
	resource := Resource{
		Region: describeCtx.Region,
		// No ARN or ID. Account level setting
		Name: accountId + " S3 Account Setting",
		Description: model.S3AccountSettingDescription{
			PublicAccessBlockConfiguration: *output.PublicAccessBlockConfiguration,
		},
	}
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}
