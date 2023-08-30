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

func tableAwsEc2ApplicationLoadBalancerListener(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_load_balancer_listener",
		Description: "AWS EC2 Load Balancer Listener",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ListenerNotFound", "LoadBalancerNotFound"}),
			},
			Hydrate: kaytu.GetElasticLoadBalancingV2Listener,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListElasticLoadBalancingV2Listener,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the listener.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Listener.ListenerArn"),
			},
			{
				Name:        "load_balancer_arn",
				Description: "The Amazon Resource Name (ARN) of the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Listener.LoadBalancerArn")},
			{
				Name:        "port",
				Description: "The port on which the load balancer is listening.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Listener.Port")},
			{
				Name:        "protocol",
				Description: "The protocol for connections from clients to the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Listener.Protocol")},
			{
				Name:        "ssl_policy",
				Description: "The security policy that defines which protocols and ciphers are supported.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Listener.SslPolicy")},
			{
				Name:        "alpn_policy",
				Description: "The name of the Application-Layer Protocol Negotiation (ALPN) policy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Listener.AlpnPolicy")},
			{
				Name:        "certificates",
				Description: "The default certificate for the listener.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Listener.Certificates")},
			{
				Name:        "default_actions",
				Description: "The default actions for the listener.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Listener.DefaultActions")},

			// Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEc2ApplicationLoadBalancerListenerTurbotTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Listener.ListenerArn").Transform(arnToAkas),
			},
		}),
	}
}

//// PARENT LIST FUNCTION

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS ////

func getEc2ApplicationLoadBalancerListenerTurbotTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.ElasticLoadBalancingV2Listener).Description.Listener
	splitID := strings.Split(*data.ListenerArn, "/")
	title := splitID[2] + "-" + splitID[4]
	return title, nil
}
