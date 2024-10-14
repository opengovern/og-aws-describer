package aws

import (
	"context"

	"github.com/kaytu-io/og-aws-describer/pkg/opengovernance-es-sdk"
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
			Hydrate: opengovernance.ListEC2RegionalSettings,
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
			{
				Name:        "snapshot_block_public_access_state",
				Description: "Gets the current state of block public access for snapshots setting for the account and Region.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SnapshotBlockPublicAccessState"),
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
	metadata := d.HydrateItem.(opengovernance.EC2RegionalSettings).Metadata

	title := metadata.Region + " EC2 Settings"
	return title, nil
}
