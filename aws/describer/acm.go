package describer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func CertificateManagerAccount(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := acm.NewFromConfig(cfg)
	output, err := client.GetAccountConfiguration(ctx, &acm.GetAccountConfigurationInput{})
	if err != nil {
		return nil, err
	}

	return []Resource{
		{
			Region: describeCtx.KaytuRegion,
			// No ID or ARN. Per Account Configuration
			Description: output,
		}}, nil
}

func CertificateManagerCertificate(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := acm.NewFromConfig(cfg)
	paginator := acm.NewListCertificatesPaginator(client, &acm.ListCertificatesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.CertificateSummaryList {
			getOutput, err := client.GetCertificate(ctx, &acm.GetCertificateInput{
				CertificateArn: v.CertificateArn,
			})
			if err != nil {
				return nil, err
			}

			describeOutput, err := client.DescribeCertificate(ctx, &acm.DescribeCertificateInput{
				CertificateArn: v.CertificateArn,
			})
			if err != nil {
				return nil, err
			}

			tagsOutput, err := client.ListTagsForCertificate(ctx, &acm.ListTagsForCertificateInput{
				CertificateArn: v.CertificateArn,
			})
			if err != nil {
				return nil, err
			}

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.CertificateArn,
				Name:   nameFromArn(*v.CertificateArn),
				Description: model.CertificateManagerCertificateDescription{
					Certificate: *describeOutput.Certificate,
					Attributes: struct {
						Certificate      *string
						CertificateChain *string
					}{
						Certificate:      getOutput.Certificate,
						CertificateChain: getOutput.CertificateChain,
					},
					Tags: tagsOutput.Tags,
				},
			}
			if stream != nil {
				m := *stream
				err := m(resource)
				if err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func ACMPCACertificateAuthority(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := acmpca.NewFromConfig(cfg)
	paginator := acmpca.NewListCertificateAuthoritiesPaginator(client, &acmpca.ListCertificateAuthoritiesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.CertificateAuthorities {
			tags, err := client.ListTags(ctx, &acmpca.ListTagsInput{
				CertificateAuthorityArn: v.Arn,
			})
			if err != nil {
				return nil, err
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.Arn,
				Name:   nameFromArn(*v.Arn),
				Description: model.ACMPCACertificateAuthorityDescription{
					CertificateAuthority: v,
					Tags:                 tags.Tags,
				},
			}
			if stream != nil {
				m := *stream
				err := m(resource)
				if err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}
