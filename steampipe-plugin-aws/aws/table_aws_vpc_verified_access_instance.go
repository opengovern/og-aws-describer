package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsVpcVerifiedAccessInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_verified_access_instance",
		Description: "AWS VPC Verified Access Instance",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2VerifiedAccessInstance,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "verified_access_instance_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "verified_access_instance_id",
				Description: "The ID of the AWS Verified Access instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountInstance.VerifiedAccessInstanceId")},
			{
				Name:        "creation_time",
				Description: "The creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountInstance.CreationTime")},
			{
				Name:        "description",
				Description: "A description for the AWS Verified Access instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountInstance.Description")},
			{
				Name:        "last_updated_time",
				Description: "The last updated time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountInstance.LastUpdatedTime")},
			{
				Name:        "verified_access_trust_providers",
				Description: "The IDs of the AWS Verified Access trusted providers.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VerifiedAccountInstance.VerifiedAccessTrustProviders")},
			{
				Name:        "tags_src",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VerifiedAccountInstance.Tags")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(verifiedAccessInstanceTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(verifiedAccessInstanceTurbotTags),
			},
		}),
	}
}

func verifiedAccessInstanceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	accessPoint := d.HydrateItem.(kaytu.EC2VerifiedAccessInstance).Description.VerifiedAccountInstance

	// Get the resource tags
	var turbotTagsMap map[string]string
	if accessPoint.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range accessPoint.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func verifiedAccessInstanceTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	accessPoint := d.HydrateItem.(kaytu.EC2VerifiedAccessInstance).Description.VerifiedAccountInstance
	title := accessPoint.VerifiedAccessInstanceId

	if accessPoint.Tags != nil {
		for _, i := range accessPoint.Tags {
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}

	return title, nil
}
