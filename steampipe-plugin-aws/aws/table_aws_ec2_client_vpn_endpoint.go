package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	ec2v1 "github.com/aws/aws-sdk-go/service/ec2"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEC2ClientVPNEndpoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ec2_client_vpn_endpoint",
		Description: "AWS EC2 Client VPN Endpoint",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("client_vpn_endpoint_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationError", "InvalidQueryParameter", "InvalidParameterValue", "InvalidClientVpnEndpointId.NotFound"}),
			},
			Hydrate: kaytu.GetEC2ClientVpnEndpoint,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListEC2ClientVpnEndpoint,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "transport_protocol", Require: plugin.Optional},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(ec2v1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "client_vpn_endpoint_id",
				Description: "The ID of the client VPN endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.ClientVpnEndpointId"),
			},
			{
				Name:        "transport_protocol",
				Description: "The transport protocol.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.TransportProtocol"),
			},
			{
				Name:        "client_cidr_block",
				Description: "The IPv4 address range, in CIDR notation, from which client IP addresses are assigned.",
				Type:        proto.ColumnType_CIDR,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.ClientCidrBlock"),
			},
			{
				Name:        "creation_time",
				Description: "The date and time when the Client VPN endpoint was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.CreationTime"),
			},
			{
				Name:        "deletion_time",
				Description: "The date and time when the Client VPN endpoint was deleted.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.DeletionTime"),
			},
			{
				Name:        "description",
				Description: "A brief description of the endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.Description"),
			},
			{
				Name:        "dns_name",
				Description: "The DNS name to be used by clients when connecting to the Client VPN endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.DnsName"),
			},
			{
				Name:        "self_service_portal_url",
				Description: "The URL of the self-service portal.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.SelfServicePortalUrl"),
			},
			{
				Name:        "server_certificate_arn",
				Description: "The ARN of the server certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.ServerCertificateArn"),
			},
			{
				Name:        "session_timeout_hours",
				Description: "The maximum VPN session duration time in hours. Valid values: 8, 10, 12, 24. Defaults to 24.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.SessionTimeoutHours"),
			},
			{
				Name:        "split_tunnel",
				Description: "Indicates whether split-tunnel is enabled in the Client VPN endpoint.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.SplitTunnel"),
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.VpcId"),
			},
			{
				Name:        "vpn_port",
				Description: "The port number for the Client VPN endpoint.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.VpnPort"),
			},
			{
				Name:        "authentication_options",
				Description: "Information about the authentication method used by the Client VPN endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.AuthenticationOptions"),
			},
			{
				Name:        "client_connect_options",
				Description: "The options for managing connection authorization for new client connections.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.ClientConnectOptions"),
			},
			{
				Name:        "client_log_options",
				Description: "Information about the client connection logging options for the Client VPN endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.ConnectionLogOptions"),
			},
			{
				Name:        "client_login_banner_options",
				Description: "Options for enabling a customizable text banner that will be displayed on Amazon Web Services provided clients when a VPN session is established.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.ClientLoginBannerOptions"),
			},
			{
				Name:        "dns_servers",
				Description: "Information about the DNS servers to be used for DNS resolution.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.DnsServers"),
			},
			{
				Name:        "security_group_ids",
				Description: "The IDs of the security groups for the target network.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.SecurityGroupIds"),
			},
			{
				Name:        "status",
				Description: "The current state of the Client VPN endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.Status"),
			},
			{
				Name:        "tags_src",
				Description: "Any tags assigned to the Client VPN endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.Tags"),
			},
			{
				Name:        "vpn_protocol",
				Description: "The protocol used by the VPN session.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.VpnProtocol"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.Tags").Transform(getEC2ClientVPNEndpointTurbotTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ClientVpnEndpoint.Tags").Transform(getEC2ClientVPNEndpointTurbotTags),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func getEC2ClientVPNEndpointTurbotTitle(_ context.Context, d *transform.TransformData) (any, error) {
	data := d.HydrateItem.(types.ClientVpnEndpoint)
	title := data.ClientVpnEndpointId
	if data.Tags != nil {
		for _, i := range data.Tags {
			if i.Key == nil {
				continue
			}
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}
	return title, nil
}

func getEC2ClientVPNEndpointTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(types.ClientVpnEndpoint)
	var turbotTagsMap map[string]string
	if data.Tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range data.Tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}
