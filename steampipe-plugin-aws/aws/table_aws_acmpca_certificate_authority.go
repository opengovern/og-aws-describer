package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsACMPCACertificateAuthority(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_acmpca_certificate_authority",
		Description: "AWS ACMPCA CertificateAuthority",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    kaytu.GetACMPCACertificateAuthority,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListACMPCACertificateAuthority,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the certificate authority",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.Arn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.Arn")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getACMPCACertificateAuthorityTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CertificateAuthority.Arn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getACMPCACertificateAuthorityTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.ACMPCACertificateAuthority).Description.Tags
	return acmpcaV2TagsToMap(tags)
}
