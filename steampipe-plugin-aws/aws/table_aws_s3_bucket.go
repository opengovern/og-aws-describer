package aws

import (
	"context"
	"encoding/json"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsS3Bucket(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_s3_bucket",
		Description: "AWS S3 Bucket",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetS3Bucket,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListS3Bucket,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The user friendly name of the bucket.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bucket.Name"),
			},
			{
				Name:        "arn",
				Description: "The ARN of the AWS S3 Bucket.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "creation_date",
				Description: "The date and tiem when bucket was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Bucket.CreationDate"),
			},
			{
				Name:        "bucket_policy_is_public",
				Description: "The policy status for an Amazon S3 bucket, indicating whether the bucket is public.",
				Type:        proto.ColumnType_BOOL,
				Default:     false,
				Transform:   transform.FromField("Description.PolicyStatus.IsPublic"),
			},
			{
				Name:        "versioning_enabled",
				Description: "The versioning state of a bucket.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Versioning.Status").Transform(handleNilString).Transform(transform.ToBool),
			},
			{
				Name:        "versioning_mfa_delete",
				Description: "The MFA Delete status of the versioning state.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Versioning.MFADelete").Transform(handleNilString).Transform(transform.ToBool),
			},
			{
				Name:        "block_public_acls",
				Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicAccessBlockConfiguration.BlockPublicAcls"),
			},
			{
				Name:        "block_public_policy",
				Description: "Specifies whether Amazon S3 should block public bucket policies for this bucket. If TRUE it causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicAccessBlockConfiguration.BlockPublicPolicy"),
			},
			{
				Name:        "ignore_public_acls",
				Description: "Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicAccessBlockConfiguration.IgnorePublicAcls"),
			},
			{
				Name:        "restrict_public_buckets",
				Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS service principals and authorized users within this account if the bucket has a public policy.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicAccessBlockConfiguration.RestrictPublicBuckets"),
			},
			{
				Name:        "event_notification_configuration",
				Description: "A container for specifying the notification configuration of the bucket. If this element is empty, notifications are turned off for the bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventNotificationConfiguration"),
			},
			{
				Name:        "server_side_encryption_configuration",
				Description: "The default encryption configuration for an Amazon S3 bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ServerSideEncryptionConfiguration"),
			},
			{
				Name:        "acl",
				Description: "The access control list (ACL) of a bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BucketAcl"),
			},
			{
				Name:        "lifecycle_rules",
				Description: "The lifecycle configuration information of the bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getS3BucketLifecycleRules),
			},
			{
				Name:        "logging",
				Description: "The logging status of a bucket and the permissions users have to view and modify that status.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoggingEnabled"),
			},
			{
				Name:        "object_lock_configuration",
				Description: "The specified bucket's object lock configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ObjectLockConfiguration"),
			},

			{
				Name:        "object_ownership_controls",
				Description: "The Ownership Controls for an Amazon S3 bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BucketOwnershipControls.OwnershipControls"),
			},
			{
				Name:        "policy",
				Description: "The resource IAM access document for the bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy"),
			},
			{
				Name:        "policy_std",
				Description: "Contains the policy in a canonical form for easier searching.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Policy").Transform(policyToCanonical),
			},
			{
				Name:        "replication",
				Description: "The replication configuration of a bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReplicationConfiguration"),
			},
			{
				Name:        "website_configuration",
				Description: "The website configuration information of the bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.BucketWebsite"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to bucket.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(s3TagsToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Bucket.Name"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(transform.EnsureStringArray),
			},
			{
				Name:        "region",
				Description: "The AWS Region in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Region"),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

func getS3BucketLifecycleRules(_ context.Context, d *transform.TransformData) (any, error) {
	bucket := d.HydrateItem.(opengovernance.S3Bucket).Description

	var p any
	err := json.Unmarshal([]byte(bucket.LifecycleRules), &p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

//// TRANSFORM FUNCTIONS

func s3TagsToTurbotTags(_ context.Context, d *transform.TransformData) (any, error) {
	tags := d.HydrateItem.(opengovernance.S3Bucket).Description.Tags

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
