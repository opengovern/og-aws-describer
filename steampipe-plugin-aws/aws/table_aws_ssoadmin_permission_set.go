package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSsoAdminPermissionSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssoadmin_permission_set",
		Description: "AWS SSO Permission Set",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSSOAdminPermissionSet,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "instance_arn", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the permission set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PermissionSet.Name"),
			},
			{
				Name:        "arn",
				Description: "The ARN of the permission set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PermissionSet.PermissionSetArn"),
			},
			{
				Name:        "created_date",
				Description: "The date that the permission set was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.PermissionSet.CreatedDate"),
			},
			{
				Name:        "description",
				Description: "The description of the permission set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PermissionSet.Description"),
			},
			{
				Name:        "instance_arn",
				Description: "The Amazon Resource Name (ARN) of the SSO Instance under which the operation will be executed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.InstanceArn"),
			},
			{
				Name:        "relay_state",
				Description: "Used to redirect users within the application during the federation authentication process.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PermissionSet.RelayState"),
			},
			{
				Name:        "session_duration",
				Description: "The length of time that the application user sessions are valid for in the ISO-8601 standard.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PermissionSet.SessionDuration"),
			},
			{
				Name:        "tags_src",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.PermissionSet.Name"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.PermissionSet.PermissionSetArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
