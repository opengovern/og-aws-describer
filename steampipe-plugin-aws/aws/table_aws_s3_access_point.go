package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsS3AccessPoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_s3_access_point",
		Description: "AWS S3 Access Point",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListS3AccessPoint,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchAccessPoint", "InvalidParameter", "InvalidRequest"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "bucket_name", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "region"}),
			Hydrate:    opengovernance.GetS3AccessPoint,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchAccessPoint\", \"AccessDenied\", \"InvalidParameter"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Specifies the name of the access point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.Name")},
			{
				Name:        "access_point_arn",
				Description: "Amazon Resource Name (ARN) of the access point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "bucket_name",
				Description: "The name of the bucket associated with this access point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.Bucket")},
			{
				Name:        "access_point_policy_is_public",
				Description: "Indicates whether this access point policy is public, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PolicyStatus.IsPublic"), Default: false,
			},
			{
				Name:        "block_public_acls",
				Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AccessPoint.PublicAccessBlockConfiguration.BlockPublicAcls")},
			{
				Name:        "block_public_policy",
				Description: "Specifies whether Amazon S3 should block public bucket policies for buckets in this account.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AccessPoint.PublicAccessBlockConfiguration.BlockPublicPolicy")},
			{
				Name:        "ignore_public_acls",
				Description: "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AccessPoint.PublicAccessBlockConfiguration.IgnorePublicAcls")},
			{
				Name:        "restrict_public_buckets",
				Description: "Specifies whether Amazon S3 should restrict public bucket policies for buckets in this account.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.AccessPoint.PublicAccessBlockConfiguration.RestrictPublicBuckets")},
			{
				Name:        "creation_date",
				Description: "The date and time when the specified access point was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AccessPoint.CreationDate")},
			{
				Name:        "network_origin",
				Description: "Indicates whether this access point allows access from the public internet. If VpcConfiguration is specified for this access point, then NetworkOrigin is VPC, and the access point doesn't allow access from the public internet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.NetworkOrigin")},
			{
				Name:        "vpc_id",
				Description: "Specifies the VPC ID from which the access point will only allow connections.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.VpcConfiguration.VpcId")},
			{
				Name:        "policy",
				Description: "The access point policy associated with the specified access point.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(transform.UnmarshalYAML),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(policyToCanonical),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}
