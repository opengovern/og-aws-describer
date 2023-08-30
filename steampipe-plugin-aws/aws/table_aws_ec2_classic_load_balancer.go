package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEc2ClassicLoadBalancer(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_classic_load_balancer",
		Description: "AWS EC2 Classic Load Balancer",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"LoadBalancerNotFound"}),
			},
			Hydrate: kaytu.GetElasticLoadBalancingLoadBalancer,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListElasticLoadBalancingLoadBalancer,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The friendly name of the Load Balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.LoadBalancerName"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the classic load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getEc2ClassicLoadBalancerARN),
			},
			{
				Name:        "scheme",
				Description: "The load balancing scheme of load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.Scheme"),
			},
			{
				Name:        "created_time",
				Description: "The date and time the load balancer was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LoadBalancer.CreatedTime"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.VPCId"),
			},
			{
				Name:        "access_log_emit_interval",
				Description: "The interval for publishing the access logs.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Attributes.AccessLog.EmitInterval"),
			},
			{
				Name:        "access_log_enabled",
				Description: "Specifies whether access logs are enabled for the load balancer.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.AccessLog.Enabled"),
			},
			{
				Name:        "access_log_s3_bucket_name",
				Description: "The name of the Amazon S3 bucket where the access logs are stored.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.AccessLog.S3BucketName"),
			},
			{
				Name:        "access_log_s3_bucket_prefix",
				Description: "The logical hierarchy you created for your Amazon S3 bucket.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.AccessLog.S3BucketPrefix"),
			},
			{
				Name:        "canonical_hosted_zone_name",
				Description: "The name of the Amazon Route 53 hosted zone for the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.CanonicalHostedZoneName"),
			},
			{
				Name:        "canonical_hosted_zone_name_id",
				Description: "The ID of the Amazon Route 53 hosted zone for the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.CanonicalHostedZoneNameID"),
			},
			{
				Name:        "connection_draining_enabled",
				Description: "Specifies whether connection draining is enabled for the load balancer.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.ConnectionDraining.Enabled"),
			},
			{
				Name:        "connection_draining_timeout",
				Description: "The maximum time, in seconds, to keep the existing connections open before deregistering the instances.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Attributes.ConnectionDraining.Timeout"),
			},
			{
				Name:        "connection_settings_idle_timeout",
				Description: "The time, in seconds, that the connection is allowed to be idle (no data has been sent over the connection) before it is closed by the load balancer.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Attributes.ConnectionSettings.IdleTimeout"),
			},
			{
				Name:        "cross_zone_load_balancing_enabled",
				Description: "Specifies whether cross-zone load balancing is enabled for the load balancer.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Attributes.CrossZoneLoadBalancing.Enabled"),
			},
			{
				Name:        "dns_name",
				Description: "The DNS name of the load balancer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.DNSName"),
			},
			{
				Name:        "health_check_interval",
				Description: "The approximate interval, in seconds, between health checks of an individual instance.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LoadBalancer.HealthCheck.Interval"),
			},
			{
				Name:        "health_check_timeout",
				Description: "The amount of time, in seconds, during which no response means a failed health check.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LoadBalancer.HealthCheck.Timeout"),
			},
			{
				Name:        "healthy_threshold",
				Description: "The number of consecutive health checks successes required before moving the instance to the Healthy state.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LoadBalancer.HealthCheck.HealthyThreshold"),
			},
			{
				Name:        "health_check_target",
				Description: "The instance being checked. The protocol is either TCP, HTTP, HTTPS, or SSL. The range of valid ports is one (1) through 65535.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.HealthCheck.Target"),
			},
			{
				Name:        "source_security_group_name",
				Description: "The name of the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.SourceSecurityGroup.GroupName"),
			},
			{
				Name:        "source_security_group_owner_alias",
				Description: "The owner of the security group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.SourceSecurityGroup.OwnerAlias"),
			},
			{
				Name:        "unhealthy_threshold",
				Description: "The number of consecutive health check failures required before moving the instance to the Unhealthy state.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.LoadBalancer.HealthCheck.UnhealthyThreshold"),
			},
			{
				Name:        "additional_attributes",
				Description: "A list of additional attributes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Attributes.AdditionalAttributes"),
			},
			{
				Name:        "app_cookie_stickiness_policies",
				Description: "A list of the stickiness policies created using CreateAppCookieStickinessPolicy.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Policies.AppCookieStickinessPolicies"),
			},
			{
				Name:        "availability_zones",
				Description: "A list of the Availability Zones for the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.AvailabilityZones"),
			},
			{
				Name:        "backend_server_descriptions",
				Description: "A list of information about your EC2 instances.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.BackendServerDescriptions"),
			},
			{
				Name:        "instances",
				Description: "A list of the IDs of the instances for the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Instances"),
			},
			{
				Name:        "lb_cookie_stickiness_policies",
				Description: "A list of the stickiness policies created using CreateLBCookieStickinessPolicy.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.LoadBalancer.Policies.LBCookieStickinessPolicies"),
			},
			{
				Name:        "listener_descriptions",
				Description: "A list of the listeners for the load balancer",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.ListenerDescriptions"),
			},
			{
				Name:        "other_policies",
				Description: "A list of policies other than the stickiness policies.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Policies.OtherPolicies"),
			},
			{
				Name:        "security_groups",
				Description: "A list of the security groups for the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.SecurityGroups"),
			},
			{
				Name:        "subnets",
				Description: "A list of the IDs of the subnets for the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LoadBalancer.Subnets"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached to the load balancer.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LoadBalancer.LoadBalancerName"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2ClassicLoadBalancerTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getEc2ClassicLoadBalancerARN).Transform(transform.EnsureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getEc2ClassicLoadBalancerARN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	classicLoadBalancer := d.HydrateItem.(kaytu.ElasticLoadBalancingLoadBalancer).Description.LoadBalancer
	metadata := d.HydrateItem.(kaytu.ElasticLoadBalancingLoadBalancer).Metadata

	// Build ARN
	arn := "arn:" + metadata.Partition + ":elasticloadbalancing:" + metadata.Region + ":" + metadata.AccountID + ":loadbalancer/" + *classicLoadBalancer.LoadBalancerName

	return arn, nil
}

//// TRANSFORM FUNCTIONS

func getEc2ClassicLoadBalancerTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	classicLoadBalancerTags := d.HydrateItem.(kaytu.ElasticLoadBalancingLoadBalancer).Description.Tags

	if classicLoadBalancerTags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range classicLoadBalancerTags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
