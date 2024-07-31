package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsEC2LaunchTemplate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_launch_template",
		Description: "AWS EC2 LaunchTemplate",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:    kaytu.GetEC2LaunchTemplate,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2LaunchTemplate,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the launchtemplate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplate.LaunchTemplateId")},
			{
				Name:        "name",
				Description: "The name of the launchtemplate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplate.LaunchTemplateName")},
			{
				Name:        "launch_template_name",
				Description: "The name of the launch template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplate.LaunchTemplateName")},
			{
				Name:        "launch_template_id",
				Description: "The ID of the launch template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplate.LaunchTemplateId")},
			{
				Name:        "create_time",
				Description: "The time launch template was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CreateTime")},
			{
				Name:        "created_by",
				Description: "The principal that created the launch template.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CreatedBy")},
			{
				Name:        "default_version_number",
				Description: "The version number of the default version of the launch template.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.DefaultVersionNumber")},
			{
				Name:        "latest_version_number",
				Description: "The name of the Application-Layer Protocol Negotiation (ALPN) policy.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LatestVersionNumber")},
			{
				Name:        "tags_src",
				Description: "The tags for the launch template.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LaunchTemplate.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEC2LaunchTemplateTurbotTags), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(launchTemplateAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEC2LaunchTemplateTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.EC2LaunchTemplate).Description.LaunchTemplate.Tags
	return ec2V2TagsToMap(tags)
}

func launchTemplateAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	launchTemplate := d.HydrateItem.(kaytu.EC2LaunchTemplate).Description.LaunchTemplate
	metadata := d.HydrateItem.(kaytu.EC2AMI).Metadata

	// Get data for Turbot defined properties
	akas := []string{"arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":launch-template/" + *launchTemplate.LaunchTemplateId}

	return akas, nil
}
