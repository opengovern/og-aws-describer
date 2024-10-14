package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsVpcVerifiedAccessEndpoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_vpc_verified_access_endpoint",
		Description: "AWS VPC verified access Endpoint",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("verified_access_endpoint_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue", "InvalidVerifiedAccessEndpointId.NotFound", "InvalidAction"}),
			},
			Hydrate: opengovernance.GetEC2VerifiedAccessEndpoint,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEC2VerifiedAccessEndpoint,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
			// DescribeVerifiedAccessEndpoints API accept endpoint id as input param.
			// We have to pass MaxResults value to DescribeVerifiedAccessEndpoints as input to perform pagination.
			// We can not pass both MaxResults and VerifiedAccessEndpointId at a time in the same input, we have to pass either one. So verified_access_endpoint_id can not be added as optional quals and added get config for filtering out the endpoint by their id.
			KeyColumns: []*plugin.KeyColumn{
				{Name: "verified_access_group_id", Require: plugin.Optional},
				{Name: "verified_access_instance_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "verified_access_endpoint_id",
				Description: "The ID of the AWS verified access endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.VerifiedAccessEndpointId")},
			{
				Name:        "verified_access_group_id",
				Description: "The ID of the AWS verified access group.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.VerifiedAccessGroupId")},
			{
				Name:        "verified_access_instance_id",
				Description: "The ID of the AWS verified access instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.VerifiedAccessInstanceId")},
			{
				Name:        "creation_time",
				Description: "The creation time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.CreationTime")},
			{
				Name:        "status_code",
				Description: "The endpoint status code. Possible values are pending, active, updating, deleting or deleted.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VerifiedAccountEndpoint.Status.Code")},
			{
				Name:        "application_domain",
				Description: "The DNS name for users to reach your application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.ApplicationDomain")},
			{
				Name:        "attachment_type",
				Description: "The type of attachment used to provide connectivity between the AWS verified access endpoint and the application.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.AttachmentType")},
			{
				Name:        "deletion_time",
				Description: "The deletion time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.DeletionTime")},
			{
				Name:        "description",
				Description: "A description for the AWS verified access endpoint.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.Description")},
			{
				Name:        "device_validation_domain",
				Description: "Returned if endpoint has a device trust provider attached.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.DeviceValidationDomain")},
			{
				Name:        "domain_certificate_arn",
				Description: "The ARN of a public TLS/SSL certificate imported into or created with ACM.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.DomainCertificateArn")},
			{
				Name:        "endpoint_domain",
				Description: "A DNS name that is generated for the endpoint..",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.EndpointDomain")},
			{
				Name:        "endpoint_type",
				Description: "The type of AWS verified access endpoint. Incoming application requests will be sent to an IP address, load balancer or a network interface depending on the endpoint type specified. Possible values are load-balancer or network-interface.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.EndpointType")},
			{
				Name:        "last_updated_time",
				Description: "The last updated time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.LastUpdatedTime")},
			{
				Name:        "load_balancer_options",
				Description: "The load balancer details if creating the AWS verified access endpoint as load-balancertype.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.LoadBalancerOptions")},
			{
				Name:        "network_interface_options",
				Description: "The options for network-interface type endpoint.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.NetworkInterfaceOptions")},
			{
				Name:        "status",
				Description: "The endpoint status.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VerifiedAccountEndpoint.Status")},
			{
				Name:        "tags_src",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.VerifiedAccountEndpoint.Tags")},

			// Steampipe standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(endpointTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(endpointTurbotTags),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func endpointTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	accessPoint := d.HydrateItem.(opengovernance.EC2VerifiedAccessEndpoint).Description.VerifiedAccountEndpoint

	// Get the resource tags
	var turbotTagsMap map[string]string
	if accessPoint.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range accessPoint.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func endpointTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	accessPoint := d.HydrateItem.(types.VerifiedAccessEndpoint)
	title := accessPoint.VerifiedAccessEndpointId

	if accessPoint.Tags != nil {
		for _, i := range accessPoint.Tags {
			if *i.Key == "Name" {
				title = i.Value
			}
		}
	}

	return title, nil
}
