package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsVpcVerifiedAccessGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_verified_access_group",
		Description: "AWS VPC Verified Access Group",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("verified_access_group_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue", "InvalidVerifiedAccessGroupId.NotFound", "InvalidAction"}),
			},
			Hydrate: opengovernance.GetEC2VerifiedAccessGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2VerifiedAccessGroup,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			// DescribeVerifiedAccessGroups API accept group id as input param.
			// We are passing MaxResults value as DescribeVerifiedAccessGroups api input
			// We cannot pass both MaxResults and VerifiedAccessGroupId at a time in the same input, we have to pass either one. So verified_access_group_id cannot be added as optional quals. Added get config for filtering out the group by its id.
			KeyColumns: []*plugin.KeyColumn{
				{Name: "verified_access_instance_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "verified_access_group_id",
				Description: "The ID of the verified access group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.VerifiedAccessGroupId")},
			{
				Name:        "arn",
				Description: "The ARN of the verified access group.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VerifiedAccountGroup.VerifiedAccessGroupArn")},
			{
				Name:        "verified_access_instance_id",
				Description: "The ID of the AWS Verified Access instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.VerifiedAccessInstanceId")},
			{
				Name:        "creation_time",
				Description: "The creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.CreationTime")},
			{
				Name:        "deletion_time",
				Description: "The deleteion time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.DeletionTime")},
			{
				Name:        "description",
				Description: "A description for the AWS verified access group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.Description")},
			{
				Name:        "last_updated_time",
				Description: "The last updated time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.LastUpdatedTime")},
			{
				Name:        "owner",
				Description: "The AWS account number that owns the group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountGroup.Owner")},
			{
				Name:        "tags_src",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VerifiedAccountGroup.Tags")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(verifiedAccessGroupTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(verifiedAccessGroupTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VerifiedAccountGroup.VerifiedAccessGroupArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func verifiedAccessGroupTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	group := d.HydrateItem.(opengovernance.EC2VerifiedAccessGroup).Description.VerifiedAccountGroup

	// Get the resource tags
	var turbotTagsMap map[string]string
	if group.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range group.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func verifiedAccessGroupTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	group := d.HydrateItem.(opengovernance.EC2VerifiedAccessGroup).Description.VerifiedAccountGroup
	title := group.VerifiedAccessGroupId

	if group.Tags != nil {
		for _, i := range group.Tags {
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}

	return title, nil
}
