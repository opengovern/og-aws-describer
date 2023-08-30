package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAWSResourceExplorerSupportedResourceType(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_resource_explorer_supported_resource_type",
		Description: "AWS Resource Explorer Supported Resource Type",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListResourceExplorer2SupportedResourceType,
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "resource_type",
				Description: "The unique identifier of the resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SupportedResourceType.ResourceType")},
			{
				Name:        "service",
				Description: "The Amazon Web Service that is associated with the resource type. This is the primary service that lets you create and interact with resources of this type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SupportedResourceType.Service")},
		}),
	}
}
