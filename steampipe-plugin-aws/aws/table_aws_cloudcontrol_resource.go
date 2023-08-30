package aws

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsCloudControlResource(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_cloudcontrol_resource",
		Description: "AWS Cloud Control Resource",
		List: &plugin.ListConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "type_name", Require: plugin.Required},
				{Name: "resource_model", Require: plugin.Optional},
			},
			//Hydrate: kaytu.ListCloudControlResource,
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{Name: "type_name"},
				{Name: "identifier"},
			},
			//Hydrate: kaytu.GetCloudControlResource,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "type_name",
				Description: "The name of the resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("type_name"),
			},
			{
				Name:        "identifier",
				Description: "The identifier for the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Resource.Identifier")},
			{
				Name:        "resource_model",
				Description: "The resource model to use to select the resources to return.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("resource_model"),
			},
			{
				Name:        "properties",
				Description: "Represents information about a provisioned resource.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Resource.Properties")},
		}),
	}
}
