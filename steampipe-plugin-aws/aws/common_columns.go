package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// Columns defined on every account-level resource (e.g. aws_iam_access_key)
func commonColumnsForAccountResource() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "partition",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Transform:   transform.FromCamel(),
			Description: "The AWS Account ID in which the resource is located.",
		},
	}
}

func commonKaytuColumnsForAccountResource() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "partition",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Transform:   transform.FromCamel(),
			Description: "The AWS Account ID in which the resource is located.",
		},
		{
			Name:        "kaytu_account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The Kaytu Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.SourceID"),
		},
		{
			Name:        "kaytu_resource_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique ID of the resource in opengovernance.",
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "kaytu_metadata",
			Type:        proto.ColumnType_STRING,
			Description: "Kaytu Metadata of the AWS resource.",
			Transform:   transform.FromField("Metadata").Transform(marshalJSON),
		},
	}
}

func commonAwsKaytuRegionalColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "partition",
			Type:        proto.ColumnType_STRING,
			Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
			Transform:   transform.FromField("Metadata.Partition"),
		},
		{
			Name:        "region",
			Type:        proto.ColumnType_STRING,
			Description: "The AWS Region in which the resource is located.",
			Transform:   transform.FromField("Metadata.Region"),
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The AWS Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.AccountID"),
		},
		{
			Name:        "kaytu_account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The Kaytu Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.SourceID"),
		},
		{
			Name:        "kaytu_resource_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique ID of the resource in opengovernance.",
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "kaytu_metadata",
			Type:        proto.ColumnType_STRING,
			Description: "kaytu Metadata of the AWS resource.",
			Transform:   transform.FromField("Metadata").Transform(marshalJSON),
		},
		{
			Name:        "kaytu_description",
			Type:        proto.ColumnType_JSON,
			Description: "The full model description of the resource",
			Transform:   transform.FromField("Description").Transform(marshalJSON),
		},
	}
}

// Columns defined on every region-level resource (e.g. aws_ec2_instance)
func commonColumnsForRegionalResource() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "partition",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
		},
		{
			Name:        "region",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS Region in which the resource is located.",
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS Account ID in which the resource is located.",
			Transform:   transform.FromCamel(),
		},
	}
}

// column definitions for the common columns
func commonKaytuColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "partition",
			Type:        proto.ColumnType_STRING,
			Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
			Transform:   transform.FromField("Metadata.Partition"),
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The AWS Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.AccountID"),
		},
		{
			Name:        "kaytu_account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The Kaytu Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.SourceID"),
		},
		{
			Name:        "kaytu_resource_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique ID of the resource in opengovernance.",
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "kaytu_metadata",
			Type:        proto.ColumnType_STRING,
			Description: "Kaytu Metadata of the AWS resource.",
			Transform:   transform.FromField("Metadata").Transform(marshalJSON),
		},
	}
}

// Columns defined on every global-region-level resource (e.g. aws_waf_rule)
func commonColumnsForGlobalRegionResource() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "partition",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
		},
		{
			Name: "region",
			Type: proto.ColumnType_STRING,
			// Region is hard-coded to special global region
			Transform:   transform.FromConstant("global"),
			Description: "The AWS Region in which the resource is located.",
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS Account ID in which the resource is located.",
			Transform:   transform.FromCamel(),
		},
	}
}

// Columns defined on every global-region-level resource (e.g. aws_waf_rule)
func commonKaytuColumnsForGlobalRegionResource() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "partition",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
		},
		{
			Name: "region",
			Type: proto.ColumnType_STRING,
			// Region is hard-coded to special global region
			Transform:   transform.FromConstant("global"),
			Description: "The AWS Region in which the resource is located.",
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCommonColumns,
			Description: "The AWS Account ID in which the resource is located.",
			Transform:   transform.FromCamel(),
		},
		{
			Name:        "kaytu_account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The Kaytu Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.SourceID"),
		},
		{
			Name:        "kaytu_resource_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique ID of the resource in opengovernance.",
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "kaytu_metadata",
			Type:        proto.ColumnType_STRING,
			Description: "Kaytu Metadata of the AWS resource.",
			Transform:   transform.FromField("Metadata").Transform(marshalJSON),
		},
	}
}
func commonAwsKaytuColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "partition",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Metadata.Partition"),
			Description: "The AWS partition in which the resource is located (aws, aws-cn, or aws-us-gov).",
		},
		{
			Name:        "region",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromConstant("global"),
			Description: "The AWS Region in which the resource is located.",
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("Metadata.AccountID"),
			Description: "The AWS Account ID in which the resource is located.",
		},
		{
			Name:        "kaytu_account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The Kaytu Account ID in which the resource is located.",
			Transform:   transform.FromField("Metadata.SourceID"),
		},
		{
			Name:        "kaytu_resource_id",
			Type:        proto.ColumnType_STRING,
			Description: "The unique ID of the resource in opengovernance.",
			Transform:   transform.FromField("ID"),
		},
		{
			Name:        "kaytu_metadata",
			Type:        proto.ColumnType_STRING,
			Description: "Kaytu Metadata of the AWS resource.",
			Transform:   transform.FromField("Metadata").Transform(marshalJSON),
		},
		{
			Name:        "kaytu_description",
			Type:        proto.ColumnType_JSON,
			Description: "The full model description of the resource",
			Transform:   transform.FromField("Description").Transform(marshalJSON),
		},
	}
}

// Append columns for account-level resource (e.g. aws_iam_access_key)
func awsRegionalColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumnsForRegionalResource()...)
}
func awsKaytuRegionalColumns(columns []*plugin.Column) []*plugin.Column {
	commonCols := commonAwsKaytuRegionalColumns()
	filteredCommonCols := make([]*plugin.Column, 0)
	for _, col := range commonCols {
		found := false
		for _, c := range columns {
			if col.Name == c.Name {
				found = true
				break
			}
		}
		if !found {
			filteredCommonCols = append(filteredCommonCols, col)
		}
	}
	return append(columns, filteredCommonCols...)
}

// Append columns for region-level resource (e.g. aws_ec2_instance)
func awsGlobalRegionColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumnsForGlobalRegionResource()...)
}
func awsKaytuGlobalRegionColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonKaytuColumnsForGlobalRegionResource()...)
}

func awsKaytuColumns(columns []*plugin.Column) []*plugin.Column {
	commonCols := commonAwsKaytuColumns()
	filteredCommonCols := make([]*plugin.Column, 0)
	for _, col := range commonCols {
		found := false
		for _, c := range columns {
			if col.Name == c.Name {
				found = true
				break
			}
		}
		if !found {
			filteredCommonCols = append(filteredCommonCols, col)
		}
	}
	return append(columns, filteredCommonCols...)
}

// Append columns for global-region-level resource (e.g. aws_waf_rule)
func awsAccountColumns(columns []*plugin.Column) []*plugin.Column {
	return append(columns, commonColumnsForAccountResource()...)
}

func awsKaytuAccountColumns(columns []*plugin.Column) []*plugin.Column {
	commonCols := commonKaytuColumnsForAccountResource()
	filteredCommonCols := make([]*plugin.Column, 0)
	for _, col := range commonCols {
		found := false
		for _, c := range columns {
			if col.Name == c.Name {
				found = true
				break
			}
		}
		if !found {
			filteredCommonCols = append(filteredCommonCols, col)
		}
	}
	return append(columns, filteredCommonCols...)
}

func awsKaytuDefaultColumns(columns []*plugin.Column) []*plugin.Column {
	commonCols := commonKaytuColumns()
	filteredCommonCols := make([]*plugin.Column, 0)
	for _, col := range commonCols {
		found := false
		for _, c := range columns {
			if col.Name == c.Name {
				found = true
				break
			}
		}
		if !found {
			filteredCommonCols = append(filteredCommonCols, col)
		}
	}
	return append(columns, filteredCommonCols...)
}

// struct to store the common column data
type awsCommonColumnData struct {
	Partition, Region, AccountId string
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize
// since getCommonColumns is a multi-region call, caching should be per connection per region
var getCommonColumns = plugin.HydrateFunc(getCommonColumnsUncached).Memoize(memoize.WithCacheKeyFunction(getCommonColumnsCacheKey))

// Build a cache key for the call to getCommonColumns, including the region since this is a multi-region call.
// Notably, this may be called WITHOUT a region. In that case we just share a cache for non-region data.
func getCommonColumnsCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	key := fmt.Sprintf("getCommonColumns-%s", region)
	return key, nil
}

// get columns which are returned with all tables: region, partition and account
func getCommonColumnsUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	if region == "" {
		region = "global"
	}

	// Trace logging to debug cache and execution flows
	plugin.Logger(ctx).Trace("getCommonColumnsUncached", "status", "starting", "connection_name", d.Connection.Name, "region", region)

	// use the cached version of the getCallerIdentity to reduce the number of request
	var commonColumnData *awsCommonColumnData
	getCallerIdentityData, err := getCallerIdentity(ctx, d, h)
	if err != nil {
		plugin.Logger(ctx).Error("getCommonColumnsUncached", "status", "failed", "connection_name", d.Connection.Name, "region", region, "error", err)
		return nil, err
	}

	callerIdentity := getCallerIdentityData.(*sts.GetCallerIdentityOutput)
	commonColumnData = &awsCommonColumnData{
		// extract partition from arn
		Partition: strings.Split(*callerIdentity.Arn, ":")[1],
		AccountId: *callerIdentity.Account,
		Region:    region,
	}

	plugin.Logger(ctx).Trace("getCommonColumnsUncached", "status", "starting", "connection_name", d.Connection.Name, "common_column_data", *commonColumnData)

	return commonColumnData, nil
}

// define cached version of getCallerIdentity and getCommonColumns
// by default, Memoize cached the data per connection
// if no argument is passed in Memoize, the cache key will be in the format of <function_name>-<connection_name>
var getCallerIdentity = plugin.HydrateFunc(getCallerIdentityUncached).Memoize()

// returns details about the IAM user or role whose credentials are used to call the operation
func getCallerIdentityUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	cacheKey := "GetCallerIdentity"

	// if found in cache, return the result
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*sts.GetCallerIdentityOutput), nil
	}

	// Trace logging to debug cache and execution flows
	plugin.Logger(ctx).Trace("getCallerIdentityUncached", "status", "starting", "connection_name", d.Connection.Name)

	// get the service connection for the service
	svc, err := STSClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getCallerIdentityUncached", "status", "failed", "connection_name", d.Connection.Name, "client_error", err)
		return nil, err
	}

	callerIdentity, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		plugin.Logger(ctx).Error("getCallerIdentityUncached", "status", "failed", "connection_name", d.Connection.Name, "api_error", err)
		return nil, err
	}

	d.ConnectionManager.Cache.Set(cacheKey, callerIdentity)
	plugin.Logger(ctx).Trace("getCallerIdentityUncached", "status", "finished", "connection_name", d.Connection.Name)
	return callerIdentity, nil
}

func marshalJSON(_ context.Context, d *transform.TransformData) (interface{}, error) {
	b, err := json.Marshal(d.Value)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
