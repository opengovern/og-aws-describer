package aws

import (
	"context"
	elbv2types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2ApplicationLoadBalancer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_application_load_balancer",
		Description: "AWS EC2 Application Load Balancer",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"LoadBalancerNotFound", "ValidationError"}),
			},
			Hydrate: getELBV2LoadBalancer(string(elbv2types.LoadBalancerTypeEnumApplication)),
		},
		List: &plugin.ListConfig{
			Hydrate: listELBV2LoadBalancer(string(elbv2types.LoadBalancerTypeEnumApplication)),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"LoadBalancerNotFound", "ValidationError"}),
			}, KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "arn",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name of the Load Balancer that was provided during resource creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.LoadBalancerName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.LoadBalancerArn"),
			},
			{
				Name:        "type",
				Description: "The type of load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Type")},
			{
				Name:        "scheme",
				Description: "The load balancing scheme of load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Scheme")},
			{
				Name:        "canonical_hosted_zone_id",
				Description: "The ID of the Amazon Route 53 hosted zone associated with the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.CanonicalHostedZoneId")},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.VpcId")},
			{
				Name:        "created_time",
				Description: "The date and time the load balancer was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LoadBalancer.CreatedTime")},
			{
				Name:        "customer_owned_ipv4_pool",
				Description: "The ID of the customer-owned address pool.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.CustomerOwnedIpv4Pool")},
			{
				Name:        "dns_name",
				Description: "The public DNS name of the load balancer.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancer.DNSName")},
			{
				Name:        "ip_address_type",
				Description: "The type of IP addresses used by the subnets for your load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.IpAddressType")},
			{
				Name:        "state_code",
				Description: "Current state of the load balancer.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancer.State.Code")},
			{
				Name:        "state_reason",
				Description: "A description of the state.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.LoadBalancer.State.Reason")},
			{
				Name:        "availability_zones",
				Description: "The subnets for the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.AvailabilityZones")},
			{
				Name:        "security_groups",
				Description: "The IDs of the security groups for the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.SecurityGroups")},
			{
				Name:        "load_balancer_attributes",
				Description: "The AWS account ID of the image owner.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes")},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Standard columns
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getELBV2LoadBalancerTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.LoadBalancerName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.LoadBalancerArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

func listELBV2LoadBalancer(t string) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		d.EqualsQuals["type"] = &proto.QualValue{
			Value: &proto.QualValue_StringValue{
				StringValue: t,
			},
		}

		return kaytu.ListElasticLoadBalancingV2LoadBalancer(ctx, d, h)
	}
}

func getELBV2LoadBalancer(t string) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		d.EqualsQuals["type"] = &proto.QualValue{
			Value: &proto.QualValue_StringValue{
				StringValue: t,
			},
		}

		return kaytu.GetElasticLoadBalancingV2LoadBalancer(ctx, d, h)
	}
}

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS ////

func getELBV2LoadBalancerTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	applicationLoadBalancerTags := d.HydrateItem.(kaytu.ElasticLoadBalancingV2LoadBalancer).Description.Tags

	if applicationLoadBalancerTags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range applicationLoadBalancerTags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
