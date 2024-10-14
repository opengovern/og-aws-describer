package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRedshiftServerlessNamespace(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_redshiftserverless_namespace",
		Description: "AWS Redshift Serverless Namespace",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("namespace_name"),
			Hydrate:    opengovernance.GetRedshiftServerlessNamespace,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRedshiftServerlessNamespace,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "namespace_id",
				Description: "The id of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.NamespaceId")},
			{
				Name:        "namespace_name",
				Description: "The name of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.NamespaceName")},
			{
				Name:        "namespace_arn",
				Description: "The Amazon Resource Name (ARN) of the namespace",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.NamespaceArn")},
			{
				Name:        "status",
				Description: "The status of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.Status")},
			{
				Name:        "admin_username",
				Description: "The username of the administrator for the first database created in the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.AdminUsername")},
			{
				Name:        "creation_date",
				Description: "The creation date of the namespace.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Namespace.CreationDate")},
			{
				Name:        "db_name",
				Description: "The name of the first database created in the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.DbName")},
			{
				Name:        "default_iam_role_arn",
				Description: "The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.DefaultIamRoleArn")},
			{
				Name:        "kms_key_id",
				Description: "The ID of the Amazon Web Services Key Management Service key used to encrypt your data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.KmsKeyId")},
			{
				Name:        "iam_roles",
				Description: "A list of IAM roles to associate with the namespace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Namespace.IamRoles")},
			{
				Name:        "log_exports",
				Description: "The types of logs the namespace can export. Available export types are User log, Connection log, and User activity log.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Namespace.LogExports")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the namespace.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Namespace.NamespaceName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getRedshiftServerlessNamespaceTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Namespace.NamespaceArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getRedshiftServerlessNamespaceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.RedshiftServerlessNamespace).Description.Tags
	return redshiftServerlessV2TagsToMap(tags)
}
