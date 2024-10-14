package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEmrBlockPublicAccessConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_emr_block_public_access_configuration",
		Description: "AWS EMR Block Public Access Configuration",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEMRBlockPublicAccessConfiguration,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "block_public_security_group_rules",
				Description: "Indicates whether Amazon EMR block public access is enabled (true) or disabled (false).",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Configuration.BlockPublicSecurityGroupRules"),
			},
			{
				Name:        "classification",
				Description: "The classification within a configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Configuration.Classification"),
			},
			{
				Name:        "created_by_arn",
				Description: "The Amazon Resource Name that created or last modified the configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationMetadata.CreatedByArn"),
			},
			{
				Name:        "creation_date",
				Description: "The date and time that the configuration was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ConfigurationMetadata.CreationDateTime"),
			},
			{
				Name:        "permitted_public_security_group_rule_ranges",
				Description: "Specifies ports and port ranges that are permitted to have security group rules that allow inbound traffic from all public sources.",
				Transform:   transform.FromField("Description.Configuration.PermittedPublicSecurityGroupRuleRanges"),
				Type:        proto.ColumnType_JSON,
			},

			// The API always returns null value for the columns configurations and properties.
			// The GO SDK provides these properties but the AWS CLI doesn't, so following columns have been commented out.
			// Raised a discussion with AWS SDK repo. Ref: https://github.com/aws/aws-sdk-go-v2/discussions/2045

			// {
			// 	Name:        "configurations",
			// 	Description: "A list of additional configurations to apply within a configuration object.",
			// 	Transform:   transform.FromField("BlockPublicAccessConfiguration.Configurations"),
			// 	Type:        proto.ColumnType_JSON,
			// },
			// {
			// 	Name:        "properties",
			// 	Description: "A set of properties specified within a configuration classification.",
			// 	Transform:   transform.FromField("BlockPublicAccessConfiguration.Properties"),
			// 	Type:        proto.ColumnType_JSON,
			// },
		}),
	}
}
