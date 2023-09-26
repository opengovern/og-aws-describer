package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSSMParameter(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_parameter",
		Description: "AWS SSM Parameter",
		Get: &plugin.GetConfig{
			//KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException"}),
			},
			Hydrate: kaytu.GetSSMParameter,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSSMParameter,
			//KeyColumns: []*plugin.KeyColumn{
			//	{Name: "type", Require: plugin.Optional},
			//	{Name: "key_id", Require: plugin.Optional},
			//	{Name: "tier", Require: plugin.Optional},
			//	{Name: "data_type", Require: plugin.Optional},
			//},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The parameter name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.Name")},
			{
				Name:        "type",
				Description: "The type of parameter. Valid parameter types include the following: String, StringList, and SecureString.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.Type")},
			{
				Name:        "value",
				Description: "The value of parameter.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Parameter.Value")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the parameter.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Parameter.ARN")},
			{
				Name:        "data_type",
				Description: "The data type of the parameter, such as text or aws:ec2:image. The default is text.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Parameter.DataType")},
			{
				Name:        "key_id",
				Description: "The ID of the query key used for this parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterMetadata.KeyId")},
			{
				Name:        "last_modified_date",
				Description: "Date the parameter was last changed or updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Parameter.LastModifiedDate")},
			{
				Name:        "last_modified_user",
				Description: "Amazon Resource Name (ARN) of the AWS user who last changed the parameter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterMetadata.LastModifiedUser")},
			{
				Name:        "version",
				Description: "The parameter version.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Parameter.Version")},
			{
				Name:        "selector",
				Description: "Either the version number or the label used to retrieve the parameter value.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Parameter.Selector")},
			{
				Name:        "source_result",
				Description: "SourceResult is the raw result or response from the source. Applies to parameters that reference information in other AWS services.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Parameter.SourceResult")},
			{
				Name:        "tier",
				Description: "The parameter tier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ParameterMetadata.Tier")},
			{
				Name:        "policies",
				Description: "A list of policies associated with a parameter. Parameter policies help you manage a growing set of parameters by enabling you to assign specific criteria to a parameter such as an expiration date or time to live. Parameter policies are especially helpful in forcing you to update or delete passwords and configuration data stored in Parameter Store.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ParameterMetadata.Policies")},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the parameter.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags")},

			/// Standard columns for all tables

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Parameter.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags").Transform(ssmTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Parameter.ARN").Transform(makeJson)},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func makeJson(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	arn := d.Value.(*string)
	return []string{*arn}, nil
}

func ssmTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.Value.([]types.Tag)

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if len(tagList) > 0 {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}
