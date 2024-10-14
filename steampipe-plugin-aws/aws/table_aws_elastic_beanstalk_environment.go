package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsElasticBeanstalkEnvironment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elastic_beanstalk_environment",
		Description: "AWS ElasticBeanstalk Environment",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("environment_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetElasticBeanstalkEnvironment,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListElasticBeanstalkEnvironment,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "environment_id", Require: plugin.Optional},
				{Name: "application_name", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "environment_name",
				Description: "The name of this environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.EnvironmentName")},
			{
				Name:        "environment_id",
				Description: "The ID of this environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.EnvironmentId")},
			{
				Name:        "arn",
				Description: "The environment's Amazon Resource Name (ARN).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.EnvironmentArn")},
			{
				Name:        "description",
				Description: "Describes this environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription")},
			{
				Name:        "date_created",
				Description: "The creation date for this environment.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.EnvironmentDescription.DateCreated")},
			{
				Name:        "abortable_operation_in_progress",
				Description: "Indicates if there is an in-progress environment configuration update or application version deployment that you can cancel.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.EnvironmentDescription.AbortableOperationInProgress")},
			{
				Name:        "application_name",
				Description: "The name of the application associated with this environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.ApplicationName")},
			{
				Name:        "cname",
				Description: "The URL to the CNAME for this environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.CNAME")},
			{
				Name:        "date_updated",
				Description: "The last modified date for this environment.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.EnvironmentDescription.DateUpdated")},
			{
				Name:        "endpoint_url",
				Description: "The URL to the LoadBalancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.EndpointURL")},
			{
				Name:        "health",
				Description: "The health status of the environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.Health")},
			{
				Name:        "health_status",
				Description: "Returns the health status of the application running in your environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.HealthStatus")},
			{
				Name:        "operations_role",
				Description: "The Amazon Resource Name (ARN) of the environment's operations role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.OperationsRole")},
			{
				Name:        "platform_arn",
				Description: "The ARN of the platform version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.PlatformArn")},
			{
				Name:        "solution_stack_name",
				Description: "The name of the SolutionStack deployed with this environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.SolutionStackName")},
			{
				Name:        "status",
				Description: "The current operational status of the environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.Status")},
			{
				Name:        "template_name",
				Description: "The name of the configuration template used to originally launch this environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.TemplateName")},
			{
				Name:        "version_label",
				Description: "The application version deployed in this environment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.EnvironmentDescription.VersionLabel")},
			{
				Name:        "configuration_settings",
				Description: "Returns a description of the settings for the specified configuration set, that is, either a configuration template or the configuration set associated with a running environment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConfigurationSetting")},
			{
				Name:        "environment_links",
				Description: "A list of links to other environments in the same group.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EnvironmentDescription.EnvironmentLinks")},
			{
				Name:        "managed_actions",
				Description: "A list of upcoming and in-progress managed actions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ManagedAction")},
			{
				Name:        "resources",
				Description: "The description of the AWS resources used by this environment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EnvironmentDescription.Resources")},
			{
				Name:        "tier",
				Description: "Describes the current tier of this environment.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EnvironmentDescription.Tier")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the Repository",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.EnvironmentDescription.EnvironmentName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("ResourceTags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(elasticBeanstalkEnvironmentTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EnvironmentDescription.EnvironmentArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

// // TRANSFORM FUNCTIONS
func elasticBeanstalkEnvironmentTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("elasticBeanstalkEnvironmentTagListToTurbotTags")
	tags := d.HydrateItem.(opengovernance.ElasticBeanstalkEnvironment).Description.Tags

	// Mapping the resource tags inside turbotTags
	if tags == nil {
		return nil, nil
	}
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
