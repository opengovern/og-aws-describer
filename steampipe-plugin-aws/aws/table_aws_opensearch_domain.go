package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsOpenSearchDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_opensearch_domain",
		Description: "AWS OpenSearch Domain",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("domain_name"),
			Hydrate:    kaytu.GetOpenSearchDomain,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListOpenSearchDomain,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "domain_id",
				Description: "The id of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainId")},
			{
				Name:        "domain_name",
				Description: "The name of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the domain",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.ARN")},
			{
				Name:        "access_policies",
				Description: "The IAM access policies of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.AccessPolicies")},
			{
				Name:        "created",
				Description: "The domain creation status.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Created")},
			{
				Name:        "deleted",
				Description: "The domain deletion status.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Deleted")},
			{
				Name:        "endpoint",
				Description: "The domain endpoint that is used to submit index and search requests.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Endpoint")},
			{
				Name:        "engine_version",
				Description: "The domain's OpenSearch version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.EngineVersion")},
			{
				Name:        "processing",
				Description: "The status of the domain configuration.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.Processing")},
			{
				Name:        "upgrade_processing",
				Description: "The status of the domain version upgrade.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.UpgradeProcessing")},
			{
				Name:        "node_to_node_encryption_options_enabled",
				Description: "Specifies the status of the node to node encryption status.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Domain.NodeToNodeEncryptionOptions.Enabled")},
			{
				Name:        "advanced_options",
				Description: "Specifies the status of the advanced options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.AdvancedOptions")},
			{
				Name:        "advanced_security_options",
				Description: "Specifies The current status of the OpenSearch domain's advanced security options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.AdvancedSecurityOptions")},
			{
				Name:        "auto_tune_options",
				Description: "The current status of the domain's auto-tune options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.AutoTuneOptions")},
			{
				Name:        "cluster_config",
				Description: "The type and number of instances in the domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.ClusterConfig")},
			{
				Name:        "cognito_options",
				Description: "The cognito options for the specified domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.CognitoOptions")},
			{
				Name:        "domain_endpoint_options",
				Description: "The current status of the domain's endpoint options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.DomainEndpointOptions")},
			{
				Name:        "ebs_options",
				Description: "The EBSOptions for the specified domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.EBSOptions")},
			{
				Name:        "encryption_at_rest_options",
				Description: "The status of the encryption at rest options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.EncryptionAtRestOptions")},
			{
				Name:        "endpoints",
				Description: "Map containing the domain endpoints used to submit index and search requests.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.Endpoints")},
			{
				Name:        "log_publishing_options",
				Description: "Log publishing options for the given domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.LogPublishingOptions")},
			{
				Name:        "service_software_options",
				Description: "The current status of the domain's service software.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.ServiceSoftwareOptions")},
			{
				Name:        "snapshot_options",
				Description: "Specifies the status of the snapshot options.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.SnapshotOptions")},
			{
				Name:        "vpc_options",
				Description: "The vpc options for the specified domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.VPCOptions")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getOpenSearchDomainTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getOpenSearchDomainTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.OpenSearchDomain).Description.Tags
	return opensearchV2TagsToMap(tags)
}
