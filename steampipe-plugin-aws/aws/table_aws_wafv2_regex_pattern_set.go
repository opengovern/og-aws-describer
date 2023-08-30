package aws

import (
	"context"
	"strings"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWafv2RegexPatternSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wafv2_regex_pattern_set",
		Description: "AWS WAFv2 Regex Pattern Set",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "name", "scope"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"WAFInvalidParameterException", "WAFNonexistentItemException", "ValidationException", "InvalidParameter"}),
			},
			Hydrate: kaytu.GetWAFv2RegexPatternSet,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWAFv2RegexPatternSet,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Regex Pattern set.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RegexPatternSet.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the entity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RegexPatternSet.ARN")},
			{
				Name:        "id",
				Description: "A unique identifier for the Regex Pattern set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RegexPatternSet.Id")},
			{
				Name:        "scope",
				Description: "Specifies the scope of the Regex Pattern Set. Possible values are: 'REGIONAL' and 'CLOUDFRONT'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(regexPatternSetLocation),
			},
			{
				Name:        "description",
				Description: "A description of the Regex Pattern set that helps with identification.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RegexPatternSet.Description")},
			{
				Name:        "lock_token",
				Description: "A token used for optimistic locking.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RegexPatternSetSummary.LockToken")},
			{
				Name:        "regular_expressions",
				Description: "The list of regular expression patterns in the set.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.RegexPatternSet.RegularExpressionList").Transform(regularExpressionObjectListToRegularExpressionList),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with the resource.",
				Type:        proto.ColumnType_JSON,

				Transform:

				// steampipe standard columns
				transform.FromField("Description.Tags.TagInfoForResource.TagList")},

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RegexPatternSet.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags.TagInfoForResource.TagList").Transform(regexPatternSetTagListToTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform:

				// AWS standard columns
				transform.FromField("Description.RegexPatternSet.ARN").Transform(arnToAkas),
			},

			{
				Name:        "partition",
				Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata.Partition")},
			{
				Name:        "region",
				Description: "The AWS Region in which the resource is located.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RegexPatternSet").Transform(regexPatternSetRegion),
			},
			{
				Name:        "account_id",
				Description: "The AWS Account ID in which the resource is located.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Metadata.AccountID")},
		},
	}
}

// Create session

// unsupported region check

// List all Regex Pattern Sets

// Reduce the basic request limit down if the user has only requested a small number of rows

// ListRegexPatternSets API doesn't support aws-sdk-go-v2 paginator yet

// Context may get cancelled due to manual cancellation or if the limit has been reached

//// HYDRATE FUNCTIONS

/*
 * The region endpoint is same for both Global Regex Pattern Set and the Regional Regex Pattern Set created in us-east-1.
 * The following checks are required to remove duplicate resource entries due to above mentioned condition, when performing GET operation.
 * To work with CloudFront, you must specify the Region US East (N. Virginia) or us-east-1
 * For the Regional Regex Pattern Set, region value should not be 'global', as 'global' region is only used to get Global Regex Pattern Sets.
 * For any other region, region value will be same as working region.
 */

// Create Session

// unsupported region check

// ListTagsForResource.NextMarker return empty string in API call
// due to which pagination will not work properly
// https://github.com/aws/aws-sdk-go/issues/3513

// To work with CloudFront, you must specify the Region US East (N. Virginia)

// Create session

// unsupported region check

// Build param with maximum limit set

//// TRANSFORM FUNCTIONS

func regexPatternSetLocation(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2RegexPatternSet)
	loc := strings.Split(strings.Split(*data.Description.RegexPatternSet.ARN, ":")[5], "/")[0]
	if loc == "regional" {
		return "REGIONAL", nil
	}
	return "CLOUDFRONT", nil
}

func regexPatternSetTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2RegexPatternSet).Description.Tags

	if data.TagInfoForResource.TagList == nil || len(data.TagInfoForResource.TagList) < 1 {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if data.TagInfoForResource.TagList != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range data.TagInfoForResource.TagList {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func regularExpressionObjectListToRegularExpressionList(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2RegexPatternSet).Description

	if data.RegexPatternSet.RegularExpressionList == nil || len(data.RegexPatternSet.RegularExpressionList) < 1 {
		return nil, nil
	}

	// Fetching the regex patterns from the array of object & storing as a list
	var regexList []string
	if data.RegexPatternSet.RegularExpressionList != nil {
		for _, i := range data.RegexPatternSet.RegularExpressionList {
			regexList = append(regexList, *i.RegexString)
		}
	}

	return regexList, nil
}

func regexPatternSetRegion(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.WAFv2RegexPatternSet)
	loc := strings.Split(strings.Split(*data.Description.RegexPatternSet.ARN, ":")[5], "/")[0]

	if loc == "global" {
		return "global", nil
	}
	return data.Metadata.Region, nil
}
