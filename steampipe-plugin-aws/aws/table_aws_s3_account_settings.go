package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsS3AccountSettings(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_s3_account_settings",
		Description: "AWS S3 Account Block Public Access Settings",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListS3AccountSetting,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "block_public_acls",
				Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicAccessBlockConfiguration.BlockPublicAcls"),
			},
			{
				Name:        "block_public_policy",
				Description: "Specifies whether Amazon S3 should block public bucket policies for this bucket. If TRUE it causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicAccessBlockConfiguration.BlockPublicPolicy"),
			},
			{
				Name:        "ignore_public_acls",
				Description: "Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicAccessBlockConfiguration.IgnorePublicAcls"),
			},
			{
				Name:        "restrict_public_buckets",
				Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS service principals and authorized users within this account if the bucket has a public policy",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.PublicAccessBlockConfiguration.RestrictPublicBuckets"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Default:     "S3 Account Level Block Public Access Settings",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(s3AccountDataToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// Transform Functions

func s3AccountDataToAkas(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	plugin.Logger(ctx).Trace("s3AccountDataToAkas")
	metadata := d.HydrateItem.(kaytu.S3AccountSetting).Metadata

	akas := []string{"arn:" + metadata.Partition + ":s3::" + metadata.AccountID + ":account"}

	return akas, nil
}
