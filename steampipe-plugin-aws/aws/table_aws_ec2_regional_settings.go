package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2RegionalSettings(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_regional_settings",
		Description: "AWS EC2 Regional Settings",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2RegionalSettings,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "default_ebs_encryption_enabled",
				Description: "Indicates whether encryption by default is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EbsEncryptionByDefault"),
			},
			{
				Name:        "default_ebs_encryption_key",
				Description: "The Amazon Resource Name (ARN) or alias of the default CMK for encryption by default.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.KmsKeyId"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEc2SettingTitle),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getEc2SettingTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	metadata := d.HydrateItem.(kaytu.EC2RegionalSettings).Metadata

	title := metadata.Region + " EC2 Settings"
	return title, nil
}
