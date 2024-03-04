package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsIamAccessAdvisor(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:             "aws_iam_access_advisor",
		Description:      "AWS IAM Access Advisor",
		DefaultTransform: transform.FromGo(),
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMAccessAdvisor,
		},
		Columns: awsKaytuGlobalRegionColumns([]*plugin.Column{
			{
				Name:        "principal_arn",
				Description: "The ARN of the IAM resource (user, group, role, or managed policy) used to generate information about when the resource was last used in an attempt to access an AWS service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PrincipalArn"),
			},
			{
				Name:        "service_name",
				Description: "The name of the service in which access was attempted.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceLastAccessed.ServiceName"),
			},
			{
				Name:        "service_namespace",
				Description: "The namespace of the service in which access was attempted.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceLastAccessed.ServiceNamespace"),
			},

			{
				Name:        "last_authenticated",
				Description: "The date and time when an authenticated entity most recently attempted to access the service. AWS does not report unauthenticated requests.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ServiceLastAccessed.LastAuthenticated"),
			},
			{
				Name:        "last_authenticated_entity",
				Description: "The ARN of the authenticated entity (user or role) that last attempted to access the service. AWS does not report unauthenticated requests.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceLastAccessed.LastAuthenticatedEntity"),
			},
			{
				Name:        "last_authenticated_region",
				Description: "The Region from which the authenticated entity (user or role) last attempted to access the service. AWS does not report unauthenticated requests.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceLastAccessed.LastAuthenticatedRegion"),
			},

			{
				Name:        "total_authenticated_entities",
				Description: "The total number of authenticated principals (root user, IAM users, or IAM roles) that have attempted to access the service.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ServiceLastAccessed.TotalAuthenticatedEntities"),
			},
			{
				Name:        "tracked_actions_last_accessed",
				Description: "An array of objects that contains details about the most recent attempt to access a tracked action within the service.  Currently, only S3 supports action level tracking.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServiceLastAccessed.TrackedActionsLastAccessed"),
			},
		}),
	}
}
