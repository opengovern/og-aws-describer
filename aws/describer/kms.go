package describer

import (
	"context"
	"errors"

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/turbot/go-kit/helpers"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func KMSAlias(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	client := kms.NewFromConfig(cfg)
	paginator := kms.NewListAliasesPaginator(client, &kms.ListAliasesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Aliases {
			resource := KMSAliasHandle(ctx, v)
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
func KMSAliasHandle(ctx context.Context, v types.AliasListEntry) Resource {
	describeCtx := GetDescribeContext(ctx)
	resource := Resource{
		Region:      describeCtx.Region,
		ARN:         *v.AliasArn,
		Name:        *v.AliasName,
		Description: v,
	}
	return resource
}
func GetKMSAlias(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	keyId := fields["keyId"]
	client := kms.NewFromConfig(cfg)

	out, err := client.ListAliases(ctx, &kms.ListAliasesInput{
		KeyId: &keyId,
	})
	if err != nil {
		if isErr(err, "ListAliasesNotFound") || isErr(err, "InvalidParameterValue") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource
	for _, v := range out.Aliases {
		values = append(values, KMSAliasHandle(ctx, v))
	}
	return values, nil
}

func KMSKey(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := kms.NewFromConfig(cfg)
	paginator := kms.NewListKeysPaginator(client, &kms.ListKeysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Keys {
			key, err := client.DescribeKey(ctx, &kms.DescribeKeyInput{
				KeyId: v.KeyId,
			})
			if err != nil {
				if isErr(err, "AccessDeniedException") {
					return nil, nil
				} else {
					return nil, err
				}
			}

			aliasPaginator := kms.NewListAliasesPaginator(client, &kms.ListAliasesInput{
				KeyId: v.KeyId,
			})

			var keyAlias []types.AliasListEntry
			for aliasPaginator.HasMorePages() {
				aliasPage, err := aliasPaginator.NextPage(ctx)
				if err != nil {
					return nil, err
				}

				keyAlias = append(keyAlias, aliasPage.Aliases...)
			}

			rotationStatus, err := client.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{
				KeyId: v.KeyId,
			})
			if err != nil {
				// For AWS managed KMS keys GetKeyRotationStatus API generates exceptions
				var ae smithy.APIError
				if errors.As(err, &ae) &&
					helpers.StringSliceContains([]string{"AccessDeniedException", "UnsupportedOperationException"}, ae.ErrorCode()) {
					rotationStatus = &kms.GetKeyRotationStatusOutput{}
					err = nil
				}

				if a, ok := err.(awserr.Error); ok {
					if helpers.StringSliceContains([]string{"AccessDeniedException", "UnsupportedOperationException"}, a.Code()) {
						rotationStatus = &kms.GetKeyRotationStatusOutput{}
						err = nil
					}
				}

				if err != nil {
					return nil, err
				}
			}

			var defaultPolicy = "default"
			policy, err := client.GetKeyPolicy(ctx, &kms.GetKeyPolicyInput{
				KeyId:      v.KeyId,
				PolicyName: &defaultPolicy,
			})
			if err != nil {
				if isErr(err, "AccessDeniedException") {
					policy = &kms.GetKeyPolicyOutput{}
				} else {
					return nil, err
				}
			}

			tags, err := client.ListResourceTags(ctx, &kms.ListResourceTagsInput{
				KeyId: v.KeyId,
			})
			if err != nil {
				if isErr(err, "AccessDeniedException") {
					tags = &kms.ListResourceTagsOutput{}
				} else {
					return nil, err
				}
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.KeyArn,
				Name:   *v.KeyId,
				Description: model.KMSKeyDescription{
					Metadata:           key.KeyMetadata,
					Aliases:            keyAlias,
					KeyRotationEnabled: rotationStatus.KeyRotationEnabled,
					Policy:             policy.Policy,
					Tags:               tags.Tags,
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

func GetKMSKey(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	id := fields["id"]
	client := kms.NewFromConfig(cfg)

	var values []Resource
	key, err := client.DescribeKey(ctx, &kms.DescribeKeyInput{
		KeyId: &id,
	})
	if err != nil {
		if isErr(err, "AccessDeniedException") {
			return nil, nil
		} else {
			return nil, err
		}
	}
	v := key.KeyMetadata

	aliasPaginator := kms.NewListAliasesPaginator(client, &kms.ListAliasesInput{
		KeyId: v.KeyId,
	})

	var keyAlias []types.AliasListEntry
	for aliasPaginator.HasMorePages() {
		aliasPage, err := aliasPaginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		keyAlias = append(keyAlias, aliasPage.Aliases...)
	}

	rotationStatus, err := client.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{
		KeyId: v.KeyId,
	})
	if err != nil {
		// For AWS managed KMS keys GetKeyRotationStatus API generates exceptions
		var ae smithy.APIError
		if errors.As(err, &ae) &&
			helpers.StringSliceContains([]string{"AccessDeniedException", "UnsupportedOperationException"}, ae.ErrorCode()) {
			rotationStatus = &kms.GetKeyRotationStatusOutput{}
			err = nil
		}

		if a, ok := err.(awserr.Error); ok {
			if helpers.StringSliceContains([]string{"AccessDeniedException", "UnsupportedOperationException"}, a.Code()) {
				rotationStatus = &kms.GetKeyRotationStatusOutput{}
				err = nil
			}
		}

		if err != nil {
			return nil, err
		}
	}

	var defaultPolicy = "default"
	policy, err := client.GetKeyPolicy(ctx, &kms.GetKeyPolicyInput{
		KeyId:      v.KeyId,
		PolicyName: &defaultPolicy,
	})
	if err != nil {
		return nil, err
	}

	tags, err := client.ListResourceTags(ctx, &kms.ListResourceTagsInput{
		KeyId: v.KeyId,
	})
	if err != nil {
		if isErr(err, "AccessDeniedException") {
			tags = &kms.ListResourceTagsOutput{}
		} else {
			return nil, err
		}
	}

	values = append(values, Resource{
		Region: describeCtx.KaytuRegion,
		ARN:    *key.KeyMetadata.Arn,
		Name:   *v.KeyId,
		Description: model.KMSKeyDescription{
			Metadata:           key.KeyMetadata,
			Aliases:            keyAlias,
			KeyRotationEnabled: rotationStatus.KeyRotationEnabled,
			Policy:             policy.Policy,
			Tags:               tags.Tags,
		},
	})

	return values, nil
}
