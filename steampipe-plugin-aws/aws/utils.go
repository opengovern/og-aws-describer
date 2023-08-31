package aws

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	acmpca "github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	cloudformation "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	cloudfront "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	codeartifact "github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	codedeploy "github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	directconnect "github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	ecrpublictypes "github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	elasticbeanstalk "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	fms "github.com/aws/aws-sdk-go-v2/service/fms/types"
	fsxtypes "github.com/aws/aws-sdk-go-v2/service/fsx/types"
	keyspacestypes "github.com/aws/aws-sdk-go-v2/service/keyspaces/types"
	memorydbtypes "github.com/aws/aws-sdk-go-v2/service/memorydb/types"
	neptunetypes "github.com/aws/aws-sdk-go-v2/service/neptune/types"
	networkfirewall "github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	opensearchtypes "github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	opsworkscm "github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	rdstypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	redshiftServerlesstypes "github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	sesv2types "github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	shield "github.com/aws/aws-sdk-go-v2/service/shield/types"
	storagegateway "github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	wafregional "github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	workspacestypes "github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/ghodss/yaml"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func ec2V2TagsToMap(tags []ec2types.Tag) (map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		if i.Key == nil || i.Value == nil {
			continue
		}
		turbotTagsMap[*i.Key] = *i.Value
	}

	return turbotTagsMap, nil
}

func ecrpublicV2TagsToMap(tags []ecrpublictypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func ecsV2TagsToMap(tags []ecstypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func fsxV2TagsToMap(tags []fsxtypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func keyspacesV2TagsToMap(tags []keyspacestypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func neptuneV2TagsToMap(tags []neptunetypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func memoryDbV2TagsToMap(tags []memorydbtypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func opensearchV2TagsToMap(tags []opensearchtypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func rdsV2TagsToMap(tags []rdstypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func redshiftServerlessV2TagsToMap(tags []redshiftServerlesstypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func sesV2TagsToMap(tags []sesv2types.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func workspacesV2TagsToMap(tags []workspacestypes.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func acmpcaV2TagsToMap(tags []acmpca.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func cloudFormationV2TagsToMap(tags []cloudformation.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func codeArtifactV2TagsToMap(tags []codeartifact.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func codeDeployV2TagsToMap(tags []codedeploy.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func directConnectV2TagsToMap(tags []directconnect.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func elasticBeanstalkV2TagsToMap(tags []elasticbeanstalk.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func fmsV2TagsToMap(tags []fms.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func networkFirewallV2TagsToMap(tags []networkfirewall.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func opsWorksCMV2TagsToMap(tags []opsworkscm.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func shieldV2TagsToMap(tags []shield.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func storageGatewayV2TagsToMap(tags []storagegateway.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func wafRegionalV2TagsToMap(tags []wafregional.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func cloudfrontV2TagsToMap(tags []cloudfront.Tag) (*map[string]string, error) {
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func arnToAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	arn := types.SafeString(d.Value)
	return []string{arn}, nil
}

func arnToTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	arn := types.SafeString(d.Value)

	// get the resource title
	title := arn[strings.LastIndex(arn, ":")+1:]

	return title, nil
}

func convertTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	epochTime := d.Value.(*int64)

	if epochTime != nil {
		timeInSec := math.Floor(float64(*epochTime) / 1000)
		unixTimestamp := time.Unix(int64(timeInSec), 0)
		timestampRFC3339Format := unixTimestamp.Format(time.RFC3339)
		return timestampRFC3339Format, nil
	}
	return nil, nil
}

func extractNameFromSqsQueueURL(queue string) (string, error) {
	//http://sqs.us-west-2.amazonaws.com/123456789012/queueName
	u, err := url.Parse(queue)
	if err != nil {
		return "", err
	}
	segments := strings.Split(u.Path, "/")
	if len(segments) != 3 {
		return "", fmt.Errorf("SQS Url not parsed correctly")
	}

	return segments[2], nil
}

func handleNilString(_ context.Context, d *transform.TransformData) (interface{}, error) {
	value := types.SafeString(fmt.Sprintf("%v", d.Value))
	if value == "" {
		return "false", nil
	}
	return value, nil
}

func resourceInterfaceDescription(key string) string {
	switch key {
	case "akas":
		return "Array of globally unique identifier strings (also known as) for the resource."
	case "tags":
		return "A map of tags for the resource."
	case "title":
		return "Title of the resource."
	}
	return ""
}

func getLastPathElement(path string) string {
	if path == "" {
		return ""
	}
	pathItems := strings.Split(path, "/")
	return pathItems[len(pathItems)-1]
}

func lastPathElement(_ context.Context, d *transform.TransformData) (interface{}, error) {
	return getLastPathElement(types.SafeString(d.Value)), nil
}

func base64DecodedData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data, err := base64.StdEncoding.DecodeString(types.SafeString(d.Value))
	// check if CorruptInputError or invalid UTF-8
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-add-user-data.html
	if err != nil {
		return nil, nil
	} else if !utf8.Valid(data) {
		return types.SafeString(d.Value), nil
	}
	return data, nil
}

func UnmarshalYAMLorJSONNoUnescape(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	inputStr := types.SafeString(d.Value)
	var result interface{}
	if inputStr != "" {
		err := yaml.Unmarshal([]byte(inputStr), &result)
		if err != nil {
			err = json.Unmarshal([]byte(inputStr), &result)
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

func getQualsValueByColumn(equalQuals plugin.KeyColumnQualMap, columnName string, dataType string) interface{} {
	var value interface{}
	for _, q := range equalQuals[columnName].Quals {
		if dataType == "string" {
			if q.Value.GetStringValue() != "" {
				value = q.Value.GetStringValue()
			} else {
				valList := getListValues(q.Value.GetListValue())
				if len(valList) > 0 {
					value = valList
				}
			}
		}
		if dataType == "boolean" {
			switch q.Operator {
			case "<>":
				value = "false"
			case "=":
				value = "true"
			}
		}
		if dataType == "int64" {
			value = q.Value.GetInt64Value()
			if q.Value.GetInt64Value() == 0 {
				valueSlice := make([]*string, 0)
				for _, value := range q.Value.GetListValue().Values {
					val := strconv.FormatInt(value.GetInt64Value(), 10)
					valueSlice = append(valueSlice, &val)
				}
				value = valueSlice
			}
		}
		if dataType == "double" {
			value = q.Value.GetDoubleValue()
			if q.Value.GetDoubleValue() == 0 {
				valueSlice := make([]*string, 0)
				for _, value := range q.Value.GetListValue().Values {
					val := strconv.FormatFloat(value.GetDoubleValue(), 'f', 4, 64)
					valueSlice = append(valueSlice, &val)
				}
				value = valueSlice
			}

		}
		if dataType == "ipaddr" {
			value = q.Value.GetInetValue().Addr
			if q.Value.GetInetValue().Addr == "" {
				valueSlice := make([]*string, 0)
				for _, value := range q.Value.GetListValue().Values {
					val := value.GetInetValue().Addr
					valueSlice = append(valueSlice, &val)
				}
				value = valueSlice
			}
		}
		if dataType == "cidr" {
			value = q.Value.GetInetValue().Cidr
			if q.Value.GetInetValue().Addr == "" {
				valueSlice := make([]*string, 0)
				for _, value := range q.Value.GetListValue().Values {
					val := value.GetInetValue().Cidr
					valueSlice = append(valueSlice, &val)
				}
				value = valueSlice
			}
		}
		if dataType == "time" {
			value = getListValues(q.Value.GetListValue())
			if len(getListValues(q.Value.GetListValue())) == 0 {
				value = q.Value.GetTimestampValue().AsTime()
			}
		}
	}
	return value
}
