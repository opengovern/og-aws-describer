package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRegion(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_region",
		Description: "AWS Region",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			Hydrate: kaytu.GetEC2Region,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2Region,
		},
		Columns: awsKaytuDefaultColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the region",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region.RegionName"),
			},
			{
				Name:        "opt_in_status",
				Description: "The Region opt-in status. The possible values are opt-in-not-required, opted-in, and not-opted-in",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region.OptInStatus"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region.RegionName"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getAwsRegionAkas),
			},
			{
				Name:        "partition",
				Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata.Partition"),
			},
			{
				Name:        "region",
				Description: "The AWS Region in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region.RegionName"), // Can't use awskaytuRegionalColumn function because of this
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getAwsRegionAkas(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAwsRegionAkas")
	region := d.HydrateItem.(kaytu.EC2Region).Description.Region
	metadata := d.HydrateItem.(kaytu.EC2Region).Metadata

	akas := []string{"arn:" + metadata.Partition + "::" + *region.RegionName + ":" + metadata.AccountID}
	return akas, nil
}
