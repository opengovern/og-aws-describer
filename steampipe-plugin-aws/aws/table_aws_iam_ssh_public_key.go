package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsIamSshPublicKey(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_ssh_public_key",
		Description: "AWS IAM User Access Key",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMSSHPublicKey,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "user_name", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "ssh_public_key_id",
				Description: "The ID for this ssh public key key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKeyKey.SSHPublicKeyId"),
			},
			{
				Name:        "user_name",
				Description: "The name of the IAM user that the key is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKeyKey.UserName"),
			},
			{
				Name:        "status",
				Description: "The status of the ssh public key. Active means that the key is valid for API calls; Inactive means it is not.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKeyKey.Status"),
			},
			{
				Name:        "update_date",
				Description: "The date when the ssh public key was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.SSHPublicKeyKey.UploadDate"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.SSHPublicKeyKey.SSHPublicKeyId"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getIamSSHPublicKeyAka),
			},
		}),
	}
}

func getIamSSHPublicKeyAka(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	arn := d.HydrateItem.(kaytu.IAMSSHPublicKey).ARN

	aka := []string{arn}
	return aka, nil
}
