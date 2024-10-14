package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsS3Object(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_s3_object",
		Description: "List AWS S3 Objects in S3 buckets by bucket name.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListS3Object,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "bucket_name", Require: plugin.Required, CacheMatch: "exact"},
				{Name: "prefix", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuAccountColumns([]*plugin.Column{
			{
				Name:        "key",
				Description: "The name that you assign to an object. You use the object key to retrieve the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ObjectSummary.Key")},
			{
				Name:        "arn",
				Description: "The ARN of the AWS S3 Object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "bucket_name",
				Description: "The name of the container bucket of this object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.BucketName")},
			{
				Name:        "last_modified",
				Description: "Last modified time of the object.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ObjectSummary.LastModified")},
			{
				Name:        "storage_class",
				Description: "The class of storage used to store the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.StorageClass")},
			{
				Name:        "version_id",
				Description: "The version ID of the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.VersionId")},
			{
				Name:        "accept_ranges",
				Description: "Indicates that a range of bytes was specified.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.AcceptRanges"),
			},
			{
				Name:        "body",
				Description: "The raw bytes of the object data as a string. If the bytes entirely consists of valid UTF8 runes, an UTF8 is sent otherwise the bas64 encoding of the bytes is sent.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.Body"),
			},
			{
				Name:        "bucket_key_enabled",
				Description: "Indicates whether the object uses an S3 Bucket Key for server-side encryption with Amazon Web Services KMS (SSE-KMS)",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Object.BucketKeyEnabled"),
			},
			{
				Name:        "cache_control",
				Description: "Specifies caching behavior along the request/reply chain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.CacheControl"),
			},
			{
				Name:        "checksum_crc32",
				Description: "The base64-encoded, 32-bit CRC32 checksum of the object. This will only be present if it was uploaded with the object. With multipart uploads, this may not be a checksum value of the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ChecksumCRC32")},
			{
				Name:        "checksum_crc32c",
				Description: "The base64-encoded, 32-bit CRC32C checksum of the object. This will only be present if it was uploaded with the object. With multipart uploads, this may not be a checksum value of the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ChecksumCRC32C")},
			{
				Name:        "checksum_sha1",
				Description: "The base64-encoded, 160-bit SHA-1 digest of the object. This will only be present if it was uploaded with the object. With multipart uploads, this may not be a checksum value of the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ChecksumSHA1")},
			{
				Name:        "checksum_sha256",
				Description: "The base64-encoded, 256-bit SHA-256 digest of the object. This will only be present if it was uploaded with the object. With multipart uploads, this may not be a checksum value of the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ChecksumSHA256")},
			{
				Name:        "content_disposition",
				Description: "Specifies presentational information for the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ContentDisposition"),
			},
			{
				Name:        "content_encoding",
				Description: "Specifies what content encodings have been applied to the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ContentEncoding"),
			},
			{
				Name:        "content_language",
				Description: "The language the content is in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ContentLanguage"),
			},
			{
				Name:        "content_length",
				Description: "Size of the body in bytes.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ContentLength"),
			},
			{
				Name:        "content_range",
				Description: "The portion of the object returned in the response.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ContentRange"),
			},
			{
				Name:        "content_type",
				Description: "A standard MIME type describing the format of the object data.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ContentType"),
			},
			{
				Name:        "delete_marker",
				Description: "Specifies whether the object retrieved was (true) or was not (false) a delete marker.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Object.DeleteMarker"),
			},
			{
				Name:        "expiration",
				Description: "If the object expiration is configured (see PUT Bucket lifecycle), the response includes this header. It includes the expiry-date and rule-id key-value pairs providing object expiration information. The value of the rule-id is URL-encoded.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.Expiration"),
			},
			{
				Name:        "expires",
				Description: "The date and time at which the object is no longer cacheable.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Object.Expires"),
			},
			{
				Name:        "etag",
				Description: "The entity tag of the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ETag"),
			},
			{
				Name:        "object_lock_legal_hold_status",
				Description: "Like a retention period, a legal hold prevents an object version from being overwritten or deleted. A legal hold remains in effect until removed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ObjectLockLegalHoldStatus")},
			{
				Name:        "object_lock_mode",
				Description: "The Object Lock mode currently in place for this object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ObjectLockMode")},
			{
				Name:        "object_lock_retain_until_date",
				Description: "The date and time when this object's Object Lock will expire.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Object.ObjectLockRetainUntilDate")},
			{
				Name:        "parts_count",
				Description: "The count of parts this object has. This value is only returned if you specify partNumber in your request and the object was uploaded as a multipart upload.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Object.PartsCount"),
			},
			{
				Name:        "prefix",
				Description: "The prefix of the key of the object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "replication_status",
				Description: "Amazon S3 can return this if your request involves a bucket that is either a source or destination in a replication rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("escription.Object.ReplicationStatus"),
			},
			{
				Name:        "request_charged",
				Description: "If present, indicates that the requester was successfully charged for the request.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("escription.Object.RequestCharged"),
			},
			{
				Name:        "restore",
				Description: "Provides information about object restoration action and expiration time of the restored object copy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("escription.Object.Restore"),
			},
			{
				Name:        "server_side_encryption",
				Description: "The server-side encryption algorithm used when storing this object in Amazon S3.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.ServerSideEncryption")},
			{
				Name:        "size",
				Description: "Size in bytes of the object.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ObjectSummary.Size")},
			{
				Name:        "sse_customer_algorithm",
				Description: "If server-side encryption with a customer-provided encryption key was requested, the response will include this header confirming the encryption algorithm used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.SSECustomerAlgorithm"),
			},
			{
				Name:        "sse_customer_key_md5",
				Description: "If server-side encryption with a customer-provided encryption key was requested, the response will include this header to provide round-trip message integrity verification of the customer-provided encryption key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.SSECustomerKeyMD5"),
			},
			{
				Name:        "sse_kms_key_id",
				Description: "If present, specifies the ID of the Amazon Web Services Key Management Service(Amazon Web Services KMS) symmetric customer managed key that was used for the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.SSEKMSKeyId"),
			},
			{
				Name:        "tag_count",
				Description: "The number of tags, if any, on the object.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.TagCount"),
			},
			{
				Name:        "website_redirection_location",
				Description: "If the bucket is configured as a website, redirects requests for this object  to another object in the same bucket or to an external URL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Object.WebsiteRedirectLocation"),
			},
			{
				Name:        "acl",
				Description: "ACLs define which AWS accounts or groups are granted access along with the type of access.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ObjectAcl"),
			},
			{
				Name:        "checksum",
				Description: "The checksum or digest of the object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ObjectAttributes.Checksum"),
			},
			{
				Name:        "metadata",
				Description: "A map of metadata to store with the object in S3.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Object.Metadata"),
			},
			{
				Name:        "object_parts",
				Description: "A collection of parts associated with a multipart upload.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ObjectAttributes.ObjectParts"),
			},
			{
				Name:        "owner",
				Description: "The owner of the object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ObjectSummary.Owner"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ObjectTags.TagSet"),
			},

			// Steampipe standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ObjectTags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ObjectSummary.Key"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
			{
				Name:        "region",
				Description: "The AWS Region in which the object is located.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getBucketLocationForObjects,
				Transform:   transform.FromValue(),
			},
		}),
	}
}

func getBucketLocationForObjects(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	c, err := getCommonColumns(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("aws_s3_object.getBucketLocationForObjects", "get_common_columns_error", err)
		return nil, err
	}
	commonColumnData := c.(*awsCommonColumnData)

	bucketName := d.EqualsQuals["bucket_name"].GetStringValue()

	// have we already created and cached the session?
	cacheKey := "getBucketLocationForObjects" + bucketName + commonColumnData.AccountId

	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(string), nil
	}

	// Unlike most services, S3 buckets are a global list. They can be retrieved
	// from any single region. It's best to use the client region of the user
	// (e.g. closest to them).
	clientRegion, err := getDefaultRegion(ctx, d, h)
	if err != nil {
		return "", err
	}
	svc, err := S3Client(ctx, d, clientRegion)
	if err != nil {
		plugin.Logger(ctx).Error("aws_s3_object.getBucketLocationForObjects", "get_client_error", err, "clientRegion", clientRegion)
		return "", err
	}

	params := &s3.GetBucketLocationInput{Bucket: aws.String(bucketName), ExpectedBucketOwner: aws.String(commonColumnData.AccountId)}

	// Specifies the Region where the bucket resides. For a list of all the Amazon
	// S3 supported location constraints by Region, see Regions and Endpoints (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region).
	location, err := svc.GetBucketLocation(ctx, params)
	if err != nil {
		plugin.Logger(ctx).Error("aws_s3_object.getBucketLocationForObjects", "bucket_name", bucketName, "clientRegion", clientRegion, "api_error", err)
		return "", err
	}
	var locationConstraint string
	if location != nil && location.LocationConstraint != "" {
		// Buckets in eu-west-1 created through the AWS CLI or other API driven methods can return a location of "EU",
		// so we need to convert back
		if location.LocationConstraint == "EU" {
			locationConstraint = "eu-west-1"
			d.ConnectionManager.Cache.Set(cacheKey, locationConstraint)
			return locationConstraint, nil
		}
		d.ConnectionManager.Cache.Set(cacheKey, string(location.LocationConstraint))
		return string(location.LocationConstraint), nil
	}

	// Buckets in us-east-1 have a LocationConstraint of null
	locationConstraint = "us-east-1"
	d.ConnectionManager.Cache.Set(cacheKey, locationConstraint)
	return locationConstraint, nil
}
