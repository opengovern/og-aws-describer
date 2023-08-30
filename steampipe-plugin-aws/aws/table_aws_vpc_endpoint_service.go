package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsVpcEndpointService(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_endpoint_service",
		Description: "AWS VPC Endpoint Service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("service_name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidServiceName"}),
			},
			Hydrate: kaytu.GetEC2VPCEndpointService,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2VPCEndpointService,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "service_name",
				Description: "The Amazon Resource Name (ARN) of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpointService.ServiceName")},
			{
				Name:        "service_id",
				Description: "The ID of the endpoint service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpointService.ServiceId")},
			{
				Name:        "owner",
				Description: "The AWS account ID of the service owner.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpointService.Owner")},
			{
				Name:        "acceptance_required",
				Description: "Indicates whether VPC endpoint connection requests to the service must be accepted by the service owner.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VpcEndpointService.AcceptanceRequired")},
			{
				Name:        "manages_vpc_endpoints",
				Description: "Indicates whether the service manages its VPC endpoints. Management of the service VPC endpoints using the VPC endpoint API is restricted.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VpcEndpointService.ManagesVpcEndpoints")},
			{
				Name:        "private_dns_name",
				Description: "The private DNS name for the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpointService.PrivateDnsName")},
			{
				Name:        "private_dns_name_verification_state",
				Description: "The verification state of the VPC endpoint service. Consumers of the endpoint service cannot use the private name when the state is not verified.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpointService.PrivateDnsNameVerificationState")},
			{
				Name:        "vpc_endpoint_policy_supported",
				Description: "Indicates whether the service supports endpoint policies.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.VpcEndpointService.VpcEndpointPolicySupported")},
			{
				Name:        "availability_zones",
				Description: "The Availability Zones in which the service is available.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpointService.AvailabilityZones")},
			{
				Name:        "base_endpoint_dns_names",
				Description: "The DNS names for the service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpointService.BaseEndpointDnsNames")},
			{
				Name:        "service_type",
				Description: "The type of service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpointService.ServiceType")},
			{
				Name:        "vpc_endpoint_connections",
				Description: "Information about one or more VPC endpoint connections.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpointConnections"),
			},
			{
				Name:        "vpc_endpoint_service_permissions",
				Description: "Information about one or more allowed principals.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AllowedPrincipals"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags assigned to the service.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VpcEndpointService.Tags")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(getVpcEndpointServiceTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VpcEndpointService.ServiceName")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getVpcEndpointServiceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	endpointService := d.HydrateItem.(kaytu.EC2VPCEndpointService).Description.VpcEndpointService
	return ec2V2TagsToMap(endpointService.Tags)
}
