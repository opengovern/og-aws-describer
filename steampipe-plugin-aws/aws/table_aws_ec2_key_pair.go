package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2KeyPair(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_key_pair",
		Description: "AWS EC2 Key Pair",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("key_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidKeyPair.NotFound", "InvalidKeyPair.Unavailable", "InvalidKeyPair.Malformed"}),
			},
			Hydrate: kaytu.GetEC2KeyPair,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2KeyPair,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "key_pair_id", Require: plugin.Optional},
				{Name: "key_fingerprint", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "key_name",
				Description: "The name of the key pair",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.KeyPair.KeyName")},
			{
				Name:        "key_pair_id",
				Description: "The ID of the key pair",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.KeyPair.KeyPairId")},
			{
				Name:        "key_fingerprint",
				Description: "If key pair was created using CreateKeyPair, this is the SHA-1 digest of the DER encoded private key. If key pair was created using ImportKeyPair to provide AWS the public key, this is the MD5 public key fingerprint as specified in section 4 of RFC4716",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.KeyPair.KeyFingerprint")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the key pair",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.KeyPair.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2KeyPairTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.KeyPair.KeyName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEC2KeyPairAkas)},
			{
				Name:        "create_time",
				Description: "The time that the keypair was created or imported",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.KeyPair.CreateTime"),
			},
		}),
	}
}

// Create Session

//// HYDRATE FUNCTIONS

// create service

// Get data for turbot defined properties

//// TRANSFORM FUNCTIONS

func getEC2KeyPairAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	keyPair := d.HydrateItem.(kaytu.EC2KeyPair).Description.KeyPair
	metadata := d.HydrateItem.(kaytu.EC2KeyPair).Metadata

	akas := []string{"arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":key-pair/" + *keyPair.KeyName}
	return akas, nil
}

func getEc2KeyPairTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	keyPair := d.HydrateItem.(kaytu.EC2KeyPair)
	return ec2V2TagsToMap(keyPair.Description.KeyPair.Tags)
}
