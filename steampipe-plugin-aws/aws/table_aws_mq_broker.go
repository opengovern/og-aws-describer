package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsMQBroker(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_mq_broker",
		Description: "AWS MQ Broker",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("broker_name"),
			Hydrate:    kaytu.GetMQBroker,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListMQBroker,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "broker_id",
				Description: "The id of the broker.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Broker.BrokerId")},
			{
				Name:        "broker_name",
				Description: "The name of the broker.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Broker.BrokerName")},
			{
				Name:        "broker_arn",
				Description: "The Amazon Resource Name (ARN) of the broker",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Broker.BrokerArn")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Broker.BrokerName")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Broker.BrokerArn").Transform(arnToAkas),
			},
		}),
	}
}
