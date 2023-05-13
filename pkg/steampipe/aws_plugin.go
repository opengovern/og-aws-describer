package steampipe

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/context_key"
	"gitlab.com/keibiengine/steampipe-plugin-aws/aws"
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

func GetResourceTypeByTableName(tableName string) string {
	tableName = strings.ToLower(tableName)
	for k, v := range awsMap {
		if tableName == strings.ToLower(v) {
			return k
		}
	}

	return ""
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
