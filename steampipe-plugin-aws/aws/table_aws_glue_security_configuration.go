package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGlueSecurityConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_glue_security_configuration",
		Description: "AWS Glue Security Configuration",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"EntityNotFoundException"}),
			},
			Hydrate: kaytu.GetGlueSecurityConfiguration,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGlueSecurityConfiguration,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the security configuration.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityConfiguration.Name")},
			{
				Name:        "created_time_stamp",
				Description: "The time at which this security configuration was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.SecurityConfiguration.CreatedTimeStamp")},
			{
				Name:        "cloud_watch_encryption",
				Description: "The encryption configuration for Amazon CloudWatch.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityConfiguration.EncryptionConfiguration.CloudWatchEncryption"),
			},
			{
				Name:        "job_bookmarks_encryption",
				Description: "The encryption configuration for job bookmarks.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityConfiguration.EncryptionConfiguration.JobBookmarksEncryption"),
			},
			{
				Name:        "s3_encryption",
				Description: "The encryption configuration for Amazon Simple Storage Service (Amazon S3) data.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SecurityConfiguration.EncryptionConfiguration.S3Encryption"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SecurityConfiguration.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
