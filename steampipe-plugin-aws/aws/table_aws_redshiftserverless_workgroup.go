package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRedshiftServerlessWorkgroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_redshiftserverless_workgroup",
		Description: "AWS Redshift Serverless Workgroup",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("workgroup_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: kaytu.GetRedshiftServerlessWorkgroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRedshiftServerlessWorkgroup,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "workgroup_name",
				Description: "The name of the workgroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workgroup.WorkgroupName")},
			{
				Name:        "workgroup_id",
				Description: "The unique identifier of the workgroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workgroup.WorkgroupId")},
			{
				Name:        "workgroup_arn",
				Description: "The Amazon Resource Name (ARN) that links to the workgroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workgroup.WorkgroupArn")},
			{
				Name:        "status",
				Description: "The status of the workgroup.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workgroup.Status")},
			{
				Name:        "base_capacity",
				Description: "The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Workgroup.BaseCapacity")},
			{
				Name:        "creation_date",
				Description: "The creation date of the workgroup.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Workgroup.CreationDate")},
			{
				Name:        "enhanced_vpc_routing",
				Description: "The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Workgroup.EnhancedVpcRouting")},
			{
				Name:        "namespace_name",
				Description: "The namespace the workgroup is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workgroup.NamespaceName")},
			{
				Name:        "publicly_accessible",
				Description: "A value that specifies whether the workgroup can be accessible from a public network.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Workgroup.PubliclyAccessible")},
			{
				Name:        "config_parameters",
				Description: "An array of parameters to set for finer control over a database.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workgroup.ConfigParameters")},
			{
				Name:        "endpoint",
				Description: "The endpoint that is created from the workgroup.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workgroup.Endpoint")},
			{
				Name:        "security_group_ids",
				Description: "An array of security group IDs to associate with the workgroup.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workgroup.SecurityGroupIds")},
			{
				Name:        "subnet_ids",
				Description: "An array of subnet IDs the workgroup is associated with.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workgroup.SubnetIds")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the workgroup.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Workgroup.WorkgroupName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getWorkgroupTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Workgroup.WorkgroupArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getWorkgroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	op := d.HydrateItem.(*redshiftserverless.ListTagsForResourceOutput)

	if op.Tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range op.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
