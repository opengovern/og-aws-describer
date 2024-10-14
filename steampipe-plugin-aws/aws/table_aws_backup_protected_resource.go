package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsBackupProtectedResource(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_protected_resource",
		Description: "AWS Backup Protected Resource",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("resource_arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "InvalidParameter", "InvalidParameterValueException"}),
			},
			Hydrate: opengovernance.GetBackupProtectedResource,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBackupProtectedResource,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "resource_arn",
				Description: "An Amazon Resource Name (ARN) that uniquely identifies a resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProtectedResource.ResourceArn"),
			},
			{
				Name:        "resource_type",
				Description: "The type of Amazon Web Services resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ProtectedResource.ResourceType"),
			},
			{
				Name:        "last_backup_time",
				Description: "The date and time a resource was last backed up.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ProtectedResource.LastBackupTime"),
			},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ProtectedResource.ResourceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTION
