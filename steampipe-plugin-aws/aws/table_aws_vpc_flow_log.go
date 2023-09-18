package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsVpcFlowlog(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_flow_log",
		Description: "AWS VPC Flowlog",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("flow_log_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"Client.InvalidInstanceID.NotFound", "InvalidParameterValue"}),
			},
			Hydrate: kaytu.GetEC2FlowLog,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2FlowLog,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "deliver_logs_status", Require: plugin.Optional},
				{Name: "log_destination_type", Require: plugin.Optional},
				{Name: "log_group_name", Require: plugin.Optional},
				{Name: "resource_id", Require: plugin.Optional},
				{Name: "traffic_type", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "flow_log_id",
				Description: "The ID of the flow log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.FlowLogId"),
			},
			{
				Name:        "creation_time",
				Description: "The date and time the flow log was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.FlowLog.CreationTime"),
			},
			{
				Name:        "deliver_logs_error_message",
				Description: "Information about the error that occurred.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.DeliverLogsErrorMessage"),
			},
			{
				Name:        "deliver_logs_permission_arn",
				Description: "The ARN of the IAM role that posts logs to CloudWatch Logs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.DeliverLogsPermissionArn"),
			},
			{
				Name:        "deliver_logs_status",
				Description: "The status of the logs delivery (SUCCESS | FAILED).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.DeliverLogsStatus"),
			},
			{
				Name:        "flow_log_status",
				Description: "The status of the flow log (ACTIVE).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.FlowLogStatus"),
			},
			{
				Name:        "log_group_name",
				Description: "The name of the flow log group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.LogGroupName"),
			},
			{
				Name:        "resource_id",
				Description: "The ID of the VPC, subnet, or network interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.ResourceId"),
			},
			{
				Name:        "traffic_type",
				Description: "The type of traffic. Valid values are: 'ACCEPT', 'REJECT',  'ALL'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.TrafficType"),
			},
			{
				Name:        "log_destination_type",
				Description: "Specifies the type of destination to which the flow log data is published.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.LogDestinationType"),
			},
			{
				Name:        "log_destination",
				Description: "Specifies the destination to which the flow log data is published.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.LogDestination"),
			},
			// bucket_name is a simpler view of the log destination bucket name, without the full path
			{
				Name:        "bucket_name",
				Description: "The name of the destination bucket to which the flow log data is published.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(logDestinationBucketName),
			},
			{
				Name:        "log_format",
				Description: "The format of the flow log record.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.FlowLog.LogFormat"),
			},
			{
				Name:        "max_aggregation_interval",
				Description: "The maximum interval of time, in seconds, during which a flow of packets is captured and aggregated into a flow log record.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.FlowLog.MaxAggregationInterval"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the VPC flowlog.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.FlowLog.Tags"),
			},

			//standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromP(vpcFlowlogTurbotData, "Tags"),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromP(vpcFlowlogTurbotData, "Title"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcFlowlogAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

func getVpcFlowlogAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpcFlowlog := d.HydrateItem.(kaytu.EC2FlowLog).Description.FlowLog
	metadata := d.HydrateItem.(kaytu.EC2FlowLog).Metadata

	akas := []string{"arn:" + metadata.Partition + ":ec2:" + metadata.Region + ":" + metadata.AccountID + ":vpc-flow-log/" + *vpcFlowlog.FlowLogId}

	return akas, nil
}

//// TRANSFORM FUNCTIONS

func vpcFlowlogTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	vpcFlowlog := d.HydrateItem.(kaytu.EC2FlowLog).Description.FlowLog
	param := d.Param.(string)

	// Get resource title
	title := vpcFlowlog.FlowLogId

	// Get the resource tags
	var turbotTagsMap map[string]string
	if vpcFlowlog.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range vpcFlowlog.Tags {
			turbotTagsMap[*i.Key] = *i.Value
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}

	if param == "Tags" {
		return turbotTagsMap, nil
	}

	return title, nil
}

func logDestinationBucketName(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.EC2FlowLog).Description.FlowLog
	if data.LogDestination == nil {
		return nil, nil
	}
	logDestination := *(data.LogDestination)
	if logDestination == "" {
		return nil, nil
	}
	splitData := strings.Split(logDestination, ":")
	return splitData[len(splitData)-1], nil
}
