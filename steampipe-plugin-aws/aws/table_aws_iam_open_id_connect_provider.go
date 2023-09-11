package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamOpenIdConnectProvider(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_open_id_connect_provider",
		Description: "AWS IAM OpenID Connect Provider",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMOpenIdConnectProvider,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"arn"}),
			Hydrate:    kaytu.GetIAMOpenIdConnectProvider,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchEntity", "InvalidInput"}),
			},
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the OIDC provider resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "client_id_list",
				Description: "A list of client IDs (also known as audiences) that are associated with the specified IAM OIDC provider resource object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientIDList"),
			},
			{
				Name:        "create_date",
				Description: "The date and time when the IAM OIDC provider resource object was created in the Amazon Web Services account.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CreateDate"),
			},
			{
				Name:        "thumbprint_list",
				Description: "A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ThumbprintList"),
			},
			{
				Name:        "url",
				Description: "The URL that the IAM OIDC provider resource object is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.URL"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags that are attached to the specified IAM OIDC provider.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags").Transform(openIDConnectTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

type OpenIDConnectProvider struct {
	Arn string `type:"string"`
	iam.GetOpenIDConnectProviderOutput
}

//// TRANSFORM FUNCTION

func openIDConnectTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(OpenIDConnectProvider)
	var turbotTagsMap map[string]string
	if len(tags.Tags) > 0 {
		turbotTagsMap = map[string]string{}
		for _, i := range tags.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}
