package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsElasticsearchDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_elasticsearch_domain",
		Description: "AWS Elasticsearch Domain",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("domain_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException"}),
			},
			Hydrate: opengovernance.GetESDomain,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListESDomain,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "domain_name",
				Description: "The name of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainName"),
			},
			{
				Name:        "engine_type",
				Description: "Specifies the EngineType of the domain.",
				Type:        proto.ColumnType_STRING,
				Default:     "elasticsearch",
			},
			{
				Name:        "domain_id",
				Description: "The id of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.ARN"),
			},
			{
				Name:        "elasticsearch_version",
				Description: "The version for the Elasticsearch domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.ElasticsearchVersion"),
			},
			{
				Name:        "endpoint",
				Description: "The Elasticsearch domain endpoint that use to submit index and search requests.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Endpoint"),
			},
			{
				Name:        "endpoints",
				Description: "The Elasticsearch domain endpoints that use to submit index and search requests.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.Endpoints"),
			},
			{
				Name:        "access_policies",
				Description: "IAM access policy as a JSON-formatted string.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.AccessPolicies"),
			},
			{
				Name:        "created",
				Description: "The domain creation status.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Created"),
			},
			{
				Name:        "deleted",
				Description: "The domain deletion status.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Deleted"),
			},
			{
				Name:        "processing",
				Description: "The status of the Elasticsearch domain configuration.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Processing"),
			},
			{
				Name:        "upgrade_processing",
				Description: "The status of an Elasticsearch domain version upgrade.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.UpgradeProcessing"),
			},
			{
				Name:        "enabled",
				Description: "Specifies the status of the NodeToNodeEncryptionOptions.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.NodeToNodeEncryptionOptions.Enabled"),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.AccessPolicies").Transform(unescape).Transform(policyToCanonical),
			},
			{
				Name:        "ebs_options",
				Description: "Specifies whether EBS-based storage is enabled.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.EBSOptions"),
			},
			{
				Name:        "advanced_options",
				Description: "Specifies the status of the AdvancedOptions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.AdvancedOptions"),
			},
			{
				Name:        "advanced_security_options",
				Description: "Specifies The current status of the Elasticsearch domain's advanced security options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.AdvancedSecurityOptions"),
			},
			{
				Name:        "auto_tune_options",
				Description: "The current status of the Elasticsearch domain's Auto-Tune options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.AutoTuneOptions"),
			},
			{
				Name:        "cognito_options",
				Description: "The CognitoOptions for the specified domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.CognitoOptions"),
			},
			{
				Name:        "domain_endpoint_options",
				Description: "The current status of the Elasticsearch domain's endpoint options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.DomainEndpointOptions"),
			},
			{
				Name:        "elasticsearch_cluster_config",
				Description: "The type and number of instances in the domain cluster.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.ElasticsearchClusterConfig"),
			},
			{
				Name:        "encryption_at_rest_options",
				Description: "Specifies the status of the EncryptionAtRestOptions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.EncryptionAtRestOptions"),
			},
			{
				Name:        "log_publishing_options",
				Description: "Log publishing options for the given domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.LogPublishingOptions"),
			},
			{
				Name:        "service_software_options",
				Description: "The current status of the Elasticsearch domain's service software.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.ServiceSoftwareOptions"),
			},
			{
				Name:        "snapshot_options",
				Description: "Specifies the status of the SnapshotOptions.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.SnapshotOptions"),
			},
			{
				Name:        "vpc_options",
				Description: "The VPCOptions for the specified domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.VPCOptions"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(getAwsElasticsearchDomaintagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     opengovernance.GetESDomain,
				Transform:   transform.FromField("Description.Domain.ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTION

func getAwsElasticsearchDomaintagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAwsElasticsearchDomaintagListToTurbotTags")
	tagList := d.HydrateItem.(opengovernance.ESDomain).Description.Tags

	if tagList == nil {
		return nil, nil
	}
	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
