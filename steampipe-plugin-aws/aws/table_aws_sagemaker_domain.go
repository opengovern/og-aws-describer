package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSageMakerDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_sagemaker_domain",
		Description: "AWS Sagemaker Domain",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "NotFoundException", "ResourceNotFound"}),
			},
			Hydrate: opengovernance.GetSageMakerDomain,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSageMakerDomain,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The domain ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainId"),
			},
			{
				Name:        "name",
				Description: "The domain name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainArn"),
			},
			{
				Name:        "creation_time",
				Description: "A timestamp that indicates when the domain was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Domain.CreationTime")},

			{
				Name:        "app_network_access_type",
				Description: "Specifies the VPC used for non-EFS traffic.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.AppNetworkAccessType")},
			{
				Name:        "app_security_group_management",
				Description: "The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.AppSecurityGroupManagement")},
			{
				Name:        "auth_mode",
				Description: "The domain's authentication mode.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.AuthMode")},
			{
				Name:        "failure_reason",
				Description: "The domain's failure reason.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.FailureReason")},
			{
				Name:        "home_efs_file_system_id",
				Description: "The ID of the Amazon Elastic File System (EFS) managed by this domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.HomeEfsFileSystemId")},
			{
				Name:        "kms_key_id",
				Description: "The Amazon Web Services KMS customer managed key used to encrypt the EFS volume attached to the domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.KmsKeyId")},
			{
				Name:        "last_modified_time",
				Description: "The domain's last modified time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Domain.LastModifiedTime")},
			{
				Name:        "security_group_id_for_domain_boundary",
				Description: "The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.SecurityGroupIdForDomainBoundary")},
			{
				Name:        "single_sign_on_managed_application_instance_id",
				Description: "The SSO managed application instance ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.SingleSignOnManagedApplicationInstanceId")},
			{
				Name:        "status",
				Description: "The domain's status.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.Status")},
			{
				Name:        "tags_src",
				Description: "The list of tags for the domain.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "url",
				Description: "The domain's URL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DomainItem.Url")},
			{
				Name:        "default_user_settings",
				Description: "Settings which are applied to UserProfiles in this domain if settings are not explicitly specified in a given UserProfile.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.DefaultUserSettings")},
			{
				Name:        "domain_settings",
				Description: "A collection of domain settings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.DomainSettings")},
			{
				Name:        "subnet_ids",
				Description: "The VPC subnets that studio uses for communication.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.SubnetIds")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Domain.DomainName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(sageMakerDomainTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Domain.DomainArn").Transform(arnToAkas),
			},
		}),
	}
}

func sageMakerDomainTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(opengovernance.SageMakerDomain).Description.Tags

	if tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}

	return nil, nil
}
