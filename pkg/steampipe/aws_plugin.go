package steampipe

import (
	"context"
	"strings"

	"github.com/hashicorp/go-hclog"

	"github.com/kaytu-io/kaytu-aws-describer/steampipe-plugin-aws/aws"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/context_key"
)

import (
	"fmt"

	"github.com/kaytu-io/kaytu-util/pkg/steampipe"
)

func buildContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, context_key.Logger, hclog.New(nil))
	return ctx
}

func AWSDescriptionToRecord(resource interface{}, indexName string) (map[string]*proto.Column, error) {
	return steampipe.DescriptionToRecord(aws.Plugin(buildContext()), resource, indexName)
}

func AWSCells(indexName string) ([]string, error) {
	return steampipe.Cells(aws.Plugin(buildContext()), indexName)
}

func ExtractTableName(resourceType string) string {
	resourceType = strings.ToLower(resourceType)
	for k, v := range awsMap {
		if resourceType == strings.ToLower(k) {
			return v
		}
	}
	return ""
}

func ExtractResourceType(tableName string) string {
	tableName = strings.ToLower(tableName)
	return strings.ToLower(awsReverseMap[tableName])
}

func GetResourceTypeByTableName(tableName string) string {
	return ExtractResourceType(tableName)
}

func Plugin() *plugin.Plugin {
	return aws.Plugin(buildContext())
}

func ExtractTagsAndNames(plg *plugin.Plugin, resourceType string, source interface{}) (map[string]string, string, error) {
	pluginTableName := ExtractTableName(resourceType)
	if pluginTableName == "" {
		return nil, "", fmt.Errorf("cannot find table name for resourceType: %s", resourceType)
	}
	return steampipe.ExtractTagsAndNames(plg, pluginTableName, resourceType, source, AWSDescriptionMap)
}
