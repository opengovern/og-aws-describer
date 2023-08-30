package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsIamAccessKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_access_key",
		Description: "AWS IAM User Access Key",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMAccessKey,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "user_name", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "access_key_id",
				Description: "The ID for this access key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessKey.AccessKeyId"),
			},
			{
				Name:        "user_name",
				Description: "The name of the IAM user that the key is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessKey.UserName"),
			},
			{
				Name:        "status",
				Description: "The status of the access key. Active means that the key is valid for API calls; Inactive means it is not.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessKey.Status"),
			},
			{
				Name:        "create_date",
				Description: "The date when the access key was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AccessKey.CreateDate"),
			},
			{
				Name:        "access_key_last_used_date",
				Description: "The date when the access key was last used.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AccessKeyLastUsed.LastUsedData"),
			},
			{
				Name:        "access_key_last_used_service",
				Description: "The service last used by the access key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessKeyLastUsed.ServiceName"),
			},
			{
				Name:        "access_key_last_used_region",
				Description: "The region in which the access key was last used.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessKeyLastUsed.Region"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessKey.AccessKeyId"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getIamAccessKeyAka),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

func getIamAccessKeyAka(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	accessKey := d.HydrateItem.(kaytu.IAMAccessKey).Description.AccessKey
	metadata := d.HydrateItem.(kaytu.IAMAccessKey).Metadata

	aka := []string{"arn:" + metadata.Partition + ":iam::" + metadata.AccountID + ":user/" + *accessKey.UserName + "/accesskey/" + *accessKey.AccessKeyId}
	return aka, nil
}
