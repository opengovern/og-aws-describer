package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSDBOptionGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_db_option_group",
		Description: "AWS RDS DB Option Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"OptionGroupNotFoundFault"}),
			},
			Hydrate: kaytu.GetRDSOptionGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRDSOptionGroup,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "engine_name", Require: plugin.Optional},
				{Name: "major_engine_version", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name to identify the option group.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.OptionGroup.OptionGroupName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the option group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OptionGroup.OptionGroupArn"),
			},
			{
				Name:        "description",
				Description: "Provides a description of the option group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OptionGroup.OptionGroupDescription")},
			{
				Name:        "allows_vpc_and_non_vpc_instance_memberships",
				Description: "Specifies whether this option group can be applied to both VPC and non-VPC instances.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.OptionGroup.AllowsVpcAndNonVpcInstanceMemberships")},
			{
				Name:        "engine_name",
				Description: "Indicates the name of the engine that this option group can be applied to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OptionGroup.EngineName")},
			{
				Name:        "major_engine_version",
				Description: "Indicates the major engine version associated with this option group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OptionGroup.MajorEngineVersion")},
			{
				Name:        "vpc_id",
				Description: "Indicates the ID of the VPC, option group can be applied.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OptionGroup.VpcId")},
			{
				Name:        "options",
				Description: "Indicates what options are available in the option group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.OptionGroup.Options")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the option group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags.TagList")},

			// Standard columns

			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRDSDBOptionGroupTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OptionGroup.OptionGroupName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.OptionGroup.OptionGroupArn").Transform(arnToAkas),

				//// LIST FUNCTION
			},
		}),
	}
}

func getRDSDBOptionGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	optionGroup := d.HydrateItem.(kaytu.RDSOptionGroup).Description.Tags

	if optionGroup != nil && optionGroup.TagList != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range optionGroup.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
