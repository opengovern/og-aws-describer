package describer

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go"
	"github.com/kaytu-io/kaytu-aws-describer/aws/model"
)

func EC2VolumeSnapshot(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	var values []Resource
	client := ec2.NewFromConfig(cfg)

	paginator := ec2.NewDescribeSnapshotsPaginator(client, &ec2.DescribeSnapshotsInput{
		OwnerIds: []string{"self"},
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, snapshot := range page.Snapshots {
			// This prevents Implicit memory aliasing in for loop
			snapshot := snapshot
			attrs, err := client.DescribeSnapshotAttribute(ctx, &ec2.DescribeSnapshotAttributeInput{
				Attribute:  types.SnapshotAttributeNameCreateVolumePermission,
				SnapshotId: snapshot.SnapshotId,
			})
			if err != nil {
				return nil, err
			}

			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":snapshot/" + *snapshot.SnapshotId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *snapshot.SnapshotId,
				Description: model.EC2VolumeSnapshotDescription{
					Snapshot:                &snapshot,
					CreateVolumePermissions: attrs.CreateVolumePermissions,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
	}

	return values, nil
}

func EC2Volume(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	var values []Resource
	client := ec2.NewFromConfig(cfg)

	paginator := ec2.NewDescribeVolumesPaginator(client, &ec2.DescribeVolumesInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, volume := range page.Volumes {
			volume := volume
			var description model.EC2VolumeDescription
			description.Volume = &volume

			attrs := []types.VolumeAttributeName{
				types.VolumeAttributeNameAutoEnableIO,
				types.VolumeAttributeNameProductCodes,
			}

			for _, attr := range attrs {
				attrs, err := client.DescribeVolumeAttribute(ctx, &ec2.DescribeVolumeAttributeInput{
					Attribute: attr,
					VolumeId:  volume.VolumeId,
				})
				if err != nil {
					return nil, err
				}

				switch attr {
				case types.VolumeAttributeNameAutoEnableIO:
					description.Attributes.AutoEnableIO = *attrs.AutoEnableIO.Value
				case types.VolumeAttributeNameProductCodes:
					description.Attributes.ProductCodes = attrs.ProductCodes
				}
			}

			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":volume/" + *volume.VolumeId
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         arn,
				Name:        *volume.VolumeId,
				Description: description,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2CapacityReservation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeCapacityReservationsPaginator(client, &ec2.DescribeCapacityReservationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if isErr(err, "InvalidCapacityReservationId.NotFound") || isErr(err, "InvalidCapacityReservationId.Unavailable") || isErr(err, "InvalidCapacityReservationId.Malformed") {
				continue
			}
			return nil, err
		}

		for _, v := range page.CapacityReservations {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.CapacityReservationArn,
				Name:   *v.CapacityReservationId,
				Description: model.EC2CapacityReservationDescription{
					CapacityReservation: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2CapacityReservationFleet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeCapacityReservationFleetsPaginator(client, &ec2.DescribeCapacityReservationFleetsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.CapacityReservationFleets {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.CapacityReservationFleetArn,
				Name:   *v.CapacityReservationFleetId,
				Description: model.EC2CapacityReservationFleetDescription{
					CapacityReservationFleet: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2CarrierGateway(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeCarrierGatewaysPaginator(client, &ec2.DescribeCarrierGatewaysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.CarrierGateways {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.CarrierGatewayId,
				Name:        *v.CarrierGatewayId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2ClientVpnAuthorizationRule(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	endpoints, err := EC2ClientVpnEndpoint(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

	var values []Resource
	for _, e := range endpoints {
		endpoint := e.Description.(types.ClientVpnEndpoint)
		paginator := ec2.NewDescribeClientVpnAuthorizationRulesPaginator(client, &ec2.DescribeClientVpnAuthorizationRulesInput{
			ClientVpnEndpointId: endpoint.ClientVpnEndpointId,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.AuthorizationRules {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*v.ClientVpnEndpointId, *v.DestinationCidr, *v.GroupId),
					Name:        *v.ClientVpnEndpointId,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2ClientVpnEndpoint(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeClientVpnEndpointsPaginator(client, &ec2.DescribeClientVpnEndpointsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ClientVpnEndpoints {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.ClientVpnEndpointId,
				Name:        *v.ClientVpnEndpointId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2ClientVpnRoute(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	endpoints, err := EC2ClientVpnEndpoint(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

	var values []Resource
	for _, e := range endpoints {
		endpoint := e.Description.(types.ClientVpnEndpoint)
		paginator := ec2.NewDescribeClientVpnRoutesPaginator(client, &ec2.DescribeClientVpnRoutesInput{
			ClientVpnEndpointId: endpoint.ClientVpnEndpointId,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.Routes {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*v.ClientVpnEndpointId, *v.DestinationCidr, *v.TargetSubnet),
					Name:        *v.ClientVpnEndpointId,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2ClientVpnTargetNetworkAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	endpoints, err := EC2ClientVpnEndpoint(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

	var values []Resource
	for _, e := range endpoints {
		endpoint := e.Description.(types.ClientVpnEndpoint)
		paginator := ec2.NewDescribeClientVpnTargetNetworksPaginator(client, &ec2.DescribeClientVpnTargetNetworksInput{
			ClientVpnEndpointId: endpoint.ClientVpnEndpointId,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.ClientVpnTargetNetworks {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          *v.AssociationId,
					Name:        *v.AssociationId,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2CustomerGateway(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeCustomerGateways(ctx, &ec2.DescribeCustomerGatewaysInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.CustomerGateways {
		resource := Resource{
			Region: describeCtx.Region,
			ID:     *v.CustomerGatewayId,
			Name:   *v.DeviceName,
			Description: model.EC2CustomerGatewayDescription{
				CustomerGateway: v,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func EC2VerifiedAccessInstance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeVerifiedAccessInstancesInput{}

	var values []Resource
	for {
		resp, err := client.DescribeVerifiedAccessInstances(ctx, input)
		if err != nil {
			return nil, nil
		}

		for _, instance := range resp.VerifiedAccessInstances {
			resource := Resource{
				Region: describeCtx.Region,
				ID:     *instance.VerifiedAccessInstanceId,
				Name:   *instance.VerifiedAccessInstanceId,
				Description: model.EC2VerifiedAccessInstanceDescription{
					VerifiedAccountInstance: instance,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
		if resp.NextToken == nil {
			break
		} else {
			input.NextToken = resp.NextToken
		}
	}

	return values, nil
}

func EC2DHCPOptions(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeDhcpOptionsPaginator(client, &ec2.DescribeDhcpOptionsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "InvalidDhcpOptionID.NotFound") {
				return nil, err
			}
			continue
		}

		for _, v := range page.DhcpOptions {
			arn := fmt.Sprintf("arn:%s:ec2:%s:%s:dhcp-options/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.DhcpOptionsId)

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.DhcpOptionsId,
				Description: model.EC2DhcpOptionsDescription{
					DhcpOptions: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2Fleet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeFleetsPaginator(client, &ec2.DescribeFleetsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Fleets {
			arn := fmt.Sprintf("arn:%s:ec2:%s:%s:fleet/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.FleetId)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     arn,
				Name:   *v.FleetId,
				Description: model.EC2FleetDescription{
					Fleet: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2EgressOnlyInternetGateway(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeEgressOnlyInternetGatewaysPaginator(client, &ec2.DescribeEgressOnlyInternetGatewaysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "InvalidEgressOnlyInternetGatewayId.NotFound") && !isErr(err, "InvalidEgressOnlyInternetGatewayId.Malformed") {
				return nil, err
			}
			continue
		}

		for _, v := range page.EgressOnlyInternetGateways {
			arn := fmt.Sprintf("arn:%s:ec2:%s:%s:egress-only-internet-gateway/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.EgressOnlyInternetGatewayId)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     arn,
				Name:   *v.EgressOnlyInternetGatewayId,
				Description: model.EC2EgressOnlyInternetGatewayDescription{
					EgressOnlyInternetGateway: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2EIP(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{})
	if err != nil {
		if !isErr(err, "InvalidAllocationID.NotFound") && !isErr(err, "InvalidAllocationID.Malformed") {
			return nil, err
		}
		return nil, nil
	}

	var values []Resource
	for _, v := range output.Addresses {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":eip/" + *v.AllocationId
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.AllocationId,
			Description: model.EC2EIPDescription{
				Address: v,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func GetEC2EIP(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	allocationId := fields["id"]

	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{
		AllocationIds: []string{allocationId},
	})
	if err != nil {
		if !isErr(err, "InvalidAllocationID.NotFound") && !isErr(err, "InvalidAllocationID.Malformed") {
			return nil, err
		}
		return nil, nil
	}

	var values []Resource
	for _, v := range output.Addresses {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":eip/" + *v.AllocationId
		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.AllocationId,
			Description: model.EC2EIPDescription{
				Address: v,
			},
		})
	}

	return values, nil
}

func EC2EnclaveCertificateIamRoleAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	certs, err := CertificateManagerCertificate(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

	var values []Resource
	for _, c := range certs {
		cert := c.Description.(model.CertificateManagerCertificateDescription)

		output, err := client.GetAssociatedEnclaveCertificateIamRoles(ctx, &ec2.GetAssociatedEnclaveCertificateIamRolesInput{
			CertificateArn: cert.Certificate.CertificateArn,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.AssociatedRoles {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.AssociatedRoleArn, // Don't set to ARN since that will be the same for the role itself and this association
				Name:        *v.AssociatedRoleArn,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
	}

	return values, nil
}

func EC2FlowLog(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeFlowLogsPaginator(client, &ec2.DescribeFlowLogsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.FlowLogs {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":vpc-flow-log/" + *v.FlowLogId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.FlowLogId,
				Description: model.EC2FlowLogDescription{
					FlowLog: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2Host(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeHostsPaginator(client, &ec2.DescribeHostsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Hosts {
			arn := fmt.Sprintf("arn:%s:ec2:%s:%s:dedicated-host/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.HostId)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     arn,
				Name:   *v.HostId,
				Description: model.EC2HostDescription{
					Host: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2Instance(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeInstancesPaginator(client, &ec2.DescribeInstancesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, r := range page.Reservations {
			for _, v := range r.Instances {
				var desc model.EC2InstanceDescription

				in := v // Do this to avoid the pointer being replaced by the for loop
				desc.Instance = &in

				statusOutput, err := client.DescribeInstanceStatus(ctx, &ec2.DescribeInstanceStatusInput{
					InstanceIds:         []string{*v.InstanceId},
					IncludeAllInstances: aws.Bool(true),
				})
				if err != nil {
					return nil, err
				}
				if len(statusOutput.InstanceStatuses) > 0 {
					desc.InstanceStatus = &statusOutput.InstanceStatuses[0]
				}

				attrs := []types.InstanceAttributeName{
					types.InstanceAttributeNameUserData,
					types.InstanceAttributeNameInstanceInitiatedShutdownBehavior,
					types.InstanceAttributeNameDisableApiTermination,
				}

				for _, attr := range attrs {
					output, err := client.DescribeInstanceAttribute(ctx, &ec2.DescribeInstanceAttributeInput{
						InstanceId: v.InstanceId,
						Attribute:  attr,
					})
					if err != nil {
						return nil, err
					}

					switch attr {
					case types.InstanceAttributeNameUserData:
						desc.Attributes.UserData = aws.ToString(output.UserData.Value)
					case types.InstanceAttributeNameInstanceInitiatedShutdownBehavior:
						desc.Attributes.InstanceInitiatedShutdownBehavior = aws.ToString(output.InstanceInitiatedShutdownBehavior.Value)
					case types.InstanceAttributeNameDisableApiTermination:
						desc.Attributes.DisableApiTermination = aws.ToBool(output.DisableApiTermination.Value)
					}
				}
				arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":instance/" + *v.InstanceId

				resource := Resource{
					Region:      describeCtx.Region,
					ARN:         arn,
					Name:        *v.InstanceId,
					Description: desc,
				}
				if stream != nil {
					m := *stream
					err = m(resource)
					if err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func GetEC2Instance(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	instanceID := fields["id"]
	client := ec2.NewFromConfig(cfg)

	out, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		return nil, err
	}

	var values []Resource

	for _, r := range out.Reservations {
		for _, v := range r.Instances {
			var desc model.EC2InstanceDescription

			in := v // Do this to avoid the pointer being replaced by the for loop
			desc.Instance = &in

			statusOutput, err := client.DescribeInstanceStatus(ctx, &ec2.DescribeInstanceStatusInput{
				InstanceIds:         []string{*v.InstanceId},
				IncludeAllInstances: aws.Bool(true),
			})
			if err != nil {
				return nil, err
			}
			if len(statusOutput.InstanceStatuses) > 0 {
				desc.InstanceStatus = &statusOutput.InstanceStatuses[0]
			}

			attrs := []types.InstanceAttributeName{
				types.InstanceAttributeNameUserData,
				types.InstanceAttributeNameInstanceInitiatedShutdownBehavior,
				types.InstanceAttributeNameDisableApiTermination,
			}

			for _, attr := range attrs {
				output, err := client.DescribeInstanceAttribute(ctx, &ec2.DescribeInstanceAttributeInput{
					InstanceId: v.InstanceId,
					Attribute:  attr,
				})
				if err != nil {
					return nil, err
				}

				switch attr {
				case types.InstanceAttributeNameUserData:
					desc.Attributes.UserData = aws.ToString(output.UserData.Value)
				case types.InstanceAttributeNameInstanceInitiatedShutdownBehavior:
					desc.Attributes.InstanceInitiatedShutdownBehavior = aws.ToString(output.InstanceInitiatedShutdownBehavior.Value)
				case types.InstanceAttributeNameDisableApiTermination:
					desc.Attributes.DisableApiTermination = aws.ToBool(output.DisableApiTermination.Value)
				}
			}
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":instance/" + *v.InstanceId
			values = append(values, Resource{
				Region:      describeCtx.Region,
				ARN:         arn,
				Name:        *v.InstanceId,
				Description: desc,
			})
		}
	}

	return values, nil
}

func EC2InternetGateway(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeInternetGatewaysPaginator(client, &ec2.DescribeInternetGatewaysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.InternetGateways {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":internet-gateway/" + *v.InternetGatewayId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.InternetGatewayId,
				Description: model.EC2InternetGatewayDescription{
					InternetGateway: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2LaunchTemplate(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeLaunchTemplatesPaginator(client, &ec2.DescribeLaunchTemplatesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.LaunchTemplates {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.LaunchTemplateId,
				Name:        *v.LaunchTemplateName,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2NatGateway(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeNatGatewaysPaginator(client, &ec2.DescribeNatGatewaysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.NatGateways {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":natgateway/" + *v.NatGatewayId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.NatGatewayId,
				Description: model.EC2NatGatewayDescription{
					NatGateway: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2NetworkAcl(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeNetworkAclsPaginator(client, &ec2.DescribeNetworkAclsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.NetworkAcls {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":network-acl/" + *v.NetworkAclId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.NetworkAclId,
				Description: model.EC2NetworkAclDescription{
					NetworkAcl: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2NetworkInsightsAnalysis(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeNetworkInsightsAnalysesPaginator(client, &ec2.DescribeNetworkInsightsAnalysesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.NetworkInsightsAnalyses {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.NetworkInsightsAnalysisArn,
				Name:        *v.NetworkInsightsAnalysisArn,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2NetworkInsightsPath(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeNetworkInsightsPathsPaginator(client, &ec2.DescribeNetworkInsightsPathsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.NetworkInsightsPaths {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.NetworkInsightsPathArn,
				Name:        *v.NetworkInsightsPathArn,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2NetworkInterface(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeNetworkInterfacesPaginator(client, &ec2.DescribeNetworkInterfacesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.NetworkInterfaces {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":network-interface/" + *v.NetworkInterfaceId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.NetworkInterfaceId,
				Description: model.EC2NetworkInterfaceDescription{
					NetworkInterface: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GetEC2NetworkInterface(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	networkInterfaceID := fields["id"]
	client := ec2.NewFromConfig(cfg)

	out, err := client.DescribeNetworkInterfaces(ctx, &ec2.DescribeNetworkInterfacesInput{
		NetworkInterfaceIds: []string{networkInterfaceID},
	})
	if err != nil {
		return nil, err
	}

	var values []Resource

	for _, v := range out.NetworkInterfaces {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":network-interface/" + *v.NetworkInterfaceId
		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.NetworkInterfaceId,
			Description: model.EC2NetworkInterfaceDescription{
				NetworkInterface: v,
			},
		})
	}

	return values, nil
}

func EC2NetworkInterfacePermission(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeNetworkInterfacePermissionsPaginator(client, &ec2.DescribeNetworkInterfacePermissionsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.NetworkInterfacePermissions {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.NetworkInterfacePermissionId,
				Name:        *v.NetworkInterfacePermissionId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2PlacementGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribePlacementGroups(ctx, &ec2.DescribePlacementGroupsInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.PlacementGroups {
		arn := fmt.Sprintf("arn:%s:ec2:%s:%s:placement-group/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.GroupName)
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ID:     arn,
			Name:   *v.GroupName,
			Description: model.EC2PlacementGroupDescription{
				PlacementGroup: v,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func EC2PrefixList(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribePrefixListsPaginator(client, &ec2.DescribePrefixListsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.PrefixLists {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.PrefixListId,
				Name:        *v.PrefixListName,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2RegionalSettings(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	out, err := client.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{})
	if err != nil {
		return nil, err
	}

	outkey, err := client.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{})
	if err != nil {
		return nil, err
	}

	resource := Resource{
		Region: describeCtx.KaytuRegion,
		// No ID or ARN. Per Account Configuration
		Name: cfg.Region + " EC2 Settings", // Based on Steampipe
		Description: model.EC2RegionalSettingsDescription{
			EbsEncryptionByDefault: out.EbsEncryptionByDefault,
			KmsKeyId:               outkey.KmsKeyId,
		},
	}
	var values []Resource
	if stream != nil {
		if err := (*stream)(resource); err != nil {
			return nil, err
		}
	} else {
		values = append(values, resource)
	}

	return values, nil
}

func EC2RouteTable(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeRouteTablesPaginator(client, &ec2.DescribeRouteTablesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.RouteTables {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":route-table/" + *v.RouteTableId

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.RouteTableId,
				Description: model.EC2RouteTableDescription{
					RouteTable: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GetEC2RouteTable(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)

	routeTableID := fields["id"]

	out, err := client.DescribeRouteTables(ctx, &ec2.DescribeRouteTablesInput{RouteTableIds: []string{routeTableID}})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range out.RouteTables {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":route-table/" + *v.RouteTableId

		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.RouteTableId,
			Description: model.EC2RouteTableDescription{
				RouteTable: v,
			},
		})
	}

	return values, nil
}

func EC2LocalGatewayRouteTable(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeLocalGatewayRouteTablesPaginator(client, &ec2.DescribeLocalGatewayRouteTablesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.LocalGatewayRouteTables {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.LocalGatewayRouteTableArn,
				Name:        *v.LocalGatewayRouteTableId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2LocalGatewayRouteTableVPCAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeLocalGatewayRouteTableVpcAssociationsPaginator(client, &ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.LocalGatewayRouteTableVpcAssociations {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.LocalGatewayRouteTableVpcAssociationId,
				Name:        *v.LocalGatewayRouteTableVpcAssociationId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayRouteTable(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTransitGatewayRouteTablesPaginator(client, &ec2.DescribeTransitGatewayRouteTablesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "InvalidRouteTableID.NotFound") && !isErr(err, "InvalidRouteTableId.Unavailable") && !isErr(err, "InvalidRouteTableId.Malformed") {
				return nil, err
			}
			continue
		}

		for _, v := range page.TransitGatewayRouteTables {
			arn := fmt.Sprintf("arn:%s:ec2:%s:%s:transit-gateway-route-table/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.TransitGatewayRouteTableId)

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.TransitGatewayRouteTableId,
				Description: model.EC2TransitGatewayRouteTableDescription{
					TransitGatewayRouteTable: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayRouteTableAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	rts, err := EC2TransitGatewayRouteTable(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

	var values []Resource
	for _, r := range rts {
		routeTable := r.Description.(types.TransitGatewayRouteTable)
		paginator := ec2.NewGetTransitGatewayRouteTableAssociationsPaginator(client, &ec2.GetTransitGatewayRouteTableAssociationsInput{
			TransitGatewayRouteTableId: routeTable.TransitGatewayRouteTableId,
		})
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.Associations {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          *v.TransitGatewayAttachmentId,
					Name:        *v.TransitGatewayAttachmentId,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayRouteTablePropagation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	rts, err := EC2TransitGatewayRouteTable(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

	var values []Resource
	for _, r := range rts {
		routeTable := r.Description.(types.TransitGatewayRouteTable)
		paginator := ec2.NewGetTransitGatewayRouteTablePropagationsPaginator(client, &ec2.GetTransitGatewayRouteTablePropagationsInput{
			TransitGatewayRouteTableId: routeTable.TransitGatewayRouteTableId,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.TransitGatewayRouteTablePropagations {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*routeTable.TransitGatewayRouteTableId, *v.TransitGatewayAttachmentId),
					Name:        *routeTable.TransitGatewayRouteTableId,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2SecurityGroup(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeSecurityGroupsPaginator(client, &ec2.DescribeSecurityGroupsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.SecurityGroups {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":security-group/" + *v.GroupId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.GroupName,
				Description: model.EC2SecurityGroupDescription{
					SecurityGroup: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GetEC2SecurityGroup(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	groupID := fields["group_id"]
	client := ec2.NewFromConfig(cfg)

	out, err := client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
		GroupIds: []string{groupID},
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range out.SecurityGroups {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":security-group/" + *v.GroupId
		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.GroupName,
			Description: model.EC2SecurityGroupDescription{
				SecurityGroup: v,
			},
		})
	}

	return values, nil
}

func getEC2SecurityGroupRuleDescriptionFromIPPermission(group types.SecurityGroup, permission types.IpPermission, groupType string) []model.EC2SecurityGroupRuleDescription {
	var descArr []model.EC2SecurityGroupRuleDescription

	// create 1 row per ip-range
	if permission.IpRanges != nil {
		for _, r := range permission.IpRanges {
			descArr = append(descArr, model.EC2SecurityGroupRuleDescription{
				Group:           group,
				Permission:      permission,
				IPRange:         &r,
				Ipv6Range:       nil,
				UserIDGroupPair: nil,
				PrefixListId:    nil,
				Type:            groupType,
			})
		}
	}

	// create 1 row per prefix-list Id
	if permission.PrefixListIds != nil {
		for _, r := range permission.PrefixListIds {
			descArr = append(descArr, model.EC2SecurityGroupRuleDescription{
				Group:           group,
				Permission:      permission,
				IPRange:         nil,
				Ipv6Range:       nil,
				UserIDGroupPair: nil,
				PrefixListId:    &r,
				Type:            groupType,
			})
		}
	}

	// create 1 row per ipv6-range
	if permission.Ipv6Ranges != nil {
		for _, r := range permission.Ipv6Ranges {
			descArr = append(descArr, model.EC2SecurityGroupRuleDescription{
				Group:           group,
				Permission:      permission,
				IPRange:         nil,
				Ipv6Range:       &r,
				UserIDGroupPair: nil,
				PrefixListId:    nil,
				Type:            groupType,
			})
		}
	}

	// create 1 row per user id group pair
	if permission.UserIdGroupPairs != nil {
		for _, r := range permission.UserIdGroupPairs {
			descArr = append(descArr, model.EC2SecurityGroupRuleDescription{
				Group:           group,
				Permission:      permission,
				IPRange:         nil,
				Ipv6Range:       nil,
				UserIDGroupPair: &r,
				PrefixListId:    nil,
				Type:            groupType,
			})
		}
	}

	return descArr
}

func EC2SecurityGroupRule(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	groups, err := EC2SecurityGroup(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	var values []Resource
	descArr := make([]model.EC2SecurityGroupRuleDescription, 0, 128)
	for _, groupWrapper := range groups {
		group := groupWrapper.Description.(model.EC2SecurityGroupDescription).SecurityGroup
		if group.IpPermissions != nil {
			for _, permission := range group.IpPermissions {
				descArr = append(descArr, getEC2SecurityGroupRuleDescriptionFromIPPermission(group, permission, "ingress")...)
			}
		}
		if group.IpPermissionsEgress != nil {
			for _, permission := range group.IpPermissionsEgress {
				descArr = append(descArr, getEC2SecurityGroupRuleDescriptionFromIPPermission(group, permission, "egress")...)
			}
		}
	}
	for _, desc := range descArr {
		hashCode := desc.Type + "_" + *desc.Permission.IpProtocol
		if desc.Permission.FromPort != nil {
			hashCode = hashCode + "_" + fmt.Sprint(desc.Permission.FromPort) + "_" + fmt.Sprint(desc.Permission.ToPort)
		}

		if desc.IPRange != nil && desc.IPRange.CidrIp != nil {
			hashCode = hashCode + "_" + *desc.IPRange.CidrIp
		} else if desc.Ipv6Range != nil && desc.Ipv6Range.CidrIpv6 != nil {
			hashCode = hashCode + "_" + *desc.Ipv6Range.CidrIpv6
		} else if desc.UserIDGroupPair != nil && *desc.UserIDGroupPair.GroupId == *desc.Group.GroupId {
			hashCode = hashCode + "_" + *desc.Group.GroupId
		} else if desc.PrefixListId != nil && desc.PrefixListId.PrefixListId != nil {
			hashCode = hashCode + "_" + *desc.PrefixListId.PrefixListId
		}

		arn := fmt.Sprintf("arn:%s:ec2:%s:%s:security-group/%s:%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *desc.Group.GroupId, hashCode)
		values = append(values, Resource{
			Region:      describeCtx.Region,
			ARN:         arn,
			Name:        fmt.Sprintf("%s_%s", *desc.Group.GroupId, hashCode),
			Description: desc,
		})
	}

	return values, nil
}

func EC2SpotFleet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeSpotFleetRequestsPaginator(client, &ec2.DescribeSpotFleetRequestsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.SpotFleetRequestConfigs {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.SpotFleetRequestId,
				Name:        *v.SpotFleetRequestId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2Subnet(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeSubnetsPaginator(client, &ec2.DescribeSubnetsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Subnets {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.SubnetArn,
				Name:   *v.SubnetId,
				Description: model.EC2SubnetDescription{
					Subnet: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GetEC2Subnet(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	subnetId := fields["id"]
	client := ec2.NewFromConfig(cfg)
	out, err := client.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		SubnetIds: []string{subnetId},
	})
	if err != nil {
		return nil, err
	}

	var values []Resource

	for _, v := range out.Subnets {
		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    *v.SubnetArn,
			Name:   *v.SubnetId,
			Description: model.EC2SubnetDescription{
				Subnet: v,
			},
		})
	}

	return values, nil
}

func EC2TrafficMirrorFilter(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTrafficMirrorFiltersPaginator(client, &ec2.DescribeTrafficMirrorFiltersInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.TrafficMirrorFilters {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.TrafficMirrorFilterId,
				Name:        *v.TrafficMirrorFilterId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2TrafficMirrorSession(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTrafficMirrorSessionsPaginator(client, &ec2.DescribeTrafficMirrorSessionsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.TrafficMirrorSessions {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.TrafficMirrorSessionId,
				Name:        *v.TrafficMirrorFilterId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2TrafficMirrorTarget(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTrafficMirrorTargetsPaginator(client, &ec2.DescribeTrafficMirrorTargetsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.TrafficMirrorTargets {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.TrafficMirrorTargetId,
				Name:        *v.TrafficMirrorTargetId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2TransitGateway(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTransitGatewaysPaginator(client, &ec2.DescribeTransitGatewaysInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			if !isErr(err, "InvalidTransitGatewayID.NotFound") && !isErr(err, "InvalidTransitGatewayID.Unavailable") && !isErr(err, "InvalidTransitGatewayID.Malformed") {
				return nil, err
			}
			continue
		}

		for _, v := range page.TransitGateways {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.TransitGatewayArn,
				Name:   *v.TransitGatewayId,
				Description: model.EC2TransitGatewayDescription{
					TransitGateway: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayConnect(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTransitGatewayConnectsPaginator(client, &ec2.DescribeTransitGatewayConnectsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.TransitGatewayConnects {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.TransitGatewayAttachmentId,
				Name:        *v.TransitGatewayAttachmentId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayMulticastDomain(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTransitGatewayMulticastDomainsPaginator(client, &ec2.DescribeTransitGatewayMulticastDomainsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.TransitGatewayMulticastDomains {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.TransitGatewayMulticastDomainArn,
				Name:        *v.TransitGatewayMulticastDomainArn,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayMulticastDomainAssociation(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	domains, err := EC2TransitGatewayMulticastDomain(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)
	var values []Resource
	for _, domain := range domains {
		paginator := ec2.NewGetTransitGatewayMulticastDomainAssociationsPaginator(client, &ec2.GetTransitGatewayMulticastDomainAssociationsInput{
			TransitGatewayMulticastDomainId: domain.Description.(types.TransitGatewayMulticastDomain).TransitGatewayMulticastDomainId,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.MulticastDomainAssociations {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          *v.TransitGatewayAttachmentId,
					Name:        *v.TransitGatewayAttachmentId,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayMulticastGroupMember(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	domains, err := EC2TransitGatewayMulticastDomain(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)
	var values []Resource
	for _, domain := range domains {
		tgmdID := domain.Description.(types.TransitGatewayMulticastDomain).TransitGatewayMulticastDomainId
		paginator := ec2.NewSearchTransitGatewayMulticastGroupsPaginator(client, &ec2.SearchTransitGatewayMulticastGroupsInput{
			TransitGatewayMulticastDomainId: tgmdID,
			Filters: []types.Filter{
				{
					Name:   aws.String("is-group-member"),
					Values: []string{"true"},
				},
			},
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.MulticastGroups {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*tgmdID, *v.GroupIpAddress),
					Name:        *v.GroupIpAddress,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayMulticastGroupSource(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	domains, err := EC2TransitGatewayMulticastDomain(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)
	var values []Resource
	for _, domain := range domains {
		tgmdID := domain.Description.(types.TransitGatewayMulticastDomain).TransitGatewayMulticastDomainId
		paginator := ec2.NewSearchTransitGatewayMulticastGroupsPaginator(client, &ec2.SearchTransitGatewayMulticastGroupsInput{
			TransitGatewayMulticastDomainId: tgmdID,
			Filters: []types.Filter{
				{
					Name:   aws.String("is-group-source"),
					Values: []string{"true"},
				},
			},
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, v := range page.MulticastGroups {
				resource := Resource{
					Region:      describeCtx.Region,
					ID:          CompositeID(*tgmdID, *v.GroupIpAddress),
					Name:        *v.GroupIpAddress,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayPeeringAttachment(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTransitGatewayPeeringAttachmentsPaginator(client, &ec2.DescribeTransitGatewayPeeringAttachmentsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.TransitGatewayPeeringAttachments {
			resource := Resource{
				Region:      describeCtx.Region,
				ID:          *v.TransitGatewayAttachmentId,
				Name:        *v.TransitGatewayAttachmentId,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2VPC(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVpcsPaginator(client, &ec2.DescribeVpcsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Vpcs {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":vpc/" + *v.VpcId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.VpcId,
				Description: model.EC2VpcDescription{
					Vpc: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func GetEC2VPC(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)

	vpcID := fields["id"]

	out, err := client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		VpcIds: []string{vpcID},
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range out.Vpcs {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":vpc/" + *v.VpcId
		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.VpcId,
			Description: model.EC2VpcDescription{
				Vpc: v,
			},
		})
	}

	return values, nil
}

func EC2VPCEndpoint(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVpcEndpointsPaginator(client, &ec2.DescribeVpcEndpointsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.VpcEndpoints {
			arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":vpc-endpoint/" + *v.VpcEndpointId
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				ID:     *v.VpcEndpointId,
				Name:   *v.VpcEndpointId,
				Description: model.EC2VPCEndpointDescription{
					VpcEndpoint: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2VPCEndpointConnectionNotification(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVpcEndpointConnectionNotificationsPaginator(client, &ec2.DescribeVpcEndpointConnectionNotificationsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ConnectionNotificationSet {
			resource := Resource{
				Region:      describeCtx.Region,
				ARN:         *v.ConnectionNotificationArn,
				Name:        *v.ConnectionNotificationArn,
				Description: v,
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2VPCEndpointService(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)

	var values []Resource
	err := PaginateRetrieveAll(func(prevToken *string) (nextToken *string, err error) {
		output, err := client.DescribeVpcEndpointServices(ctx, &ec2.DescribeVpcEndpointServicesInput{
			NextToken: prevToken,
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.ServiceDetails {
			splitServiceName := strings.Split(*v.ServiceName, ".")
			arn := fmt.Sprintf("arn:%s:ec2:%s:%s:vpc-endpoint-service/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, splitServiceName[len(splitServiceName)-1])

			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.ServiceName,
				Description: model.EC2VPCEndpointServiceDescription{
					VpcEndpointService: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}

		return output.NextToken, nil
	})
	if err != nil {
		return nil, err
	}

	return values, nil
}

func EC2VPCEndpointServicePermissions(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	services, err := EC2VPCEndpointService(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}

	client := ec2.NewFromConfig(cfg)

	var values []Resource
	for _, s := range services {
		service := s.Description.(model.EC2VPCEndpointServiceDescription).VpcEndpointService

		paginator := ec2.NewDescribeVpcEndpointServicePermissionsPaginator(client, &ec2.DescribeVpcEndpointServicePermissionsInput{
			ServiceId: service.ServiceId,
		})
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx)
			if err != nil {
				var ae smithy.APIError
				if errors.As(err, &ae) && ae.ErrorCode() == "InvalidVpcEndpointServiceId.NotFound" {
					// VpcEndpoint doesn't have permissions set. Move on!
					break
				}
				return nil, err
			}

			for _, v := range page.AllowedPrincipals {
				resource := Resource{
					Region:      describeCtx.Region,
					ARN:         *v.Principal,
					Name:        *v.Principal,
					Description: v,
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2VPCPeeringConnection(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVpcPeeringConnectionsPaginator(client, &ec2.DescribeVpcPeeringConnectionsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.VpcPeeringConnections {
			arn := fmt.Sprintf("arn:%s:ec2:%s:%s:vpc-peering-connection/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.VpcPeeringConnectionId)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.VpcPeeringConnectionId,
				Description: model.EC2VpcPeeringConnectionDescription{
					VpcPeeringConnection: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2VPNConnection(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeVpnConnections(ctx, &ec2.DescribeVpnConnectionsInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.VpnConnections {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":vpn-connection/" + *v.VpnConnectionId
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.VpnConnectionId,
			Description: model.EC2VPNConnectionDescription{
				VpnConnection: v,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func EC2VPNGateway(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeVpnGateways(ctx, &ec2.DescribeVpnGatewaysInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.VpnGateways {
		resource := Resource{
			Region: describeCtx.Region,
			ID:     *v.VpnGatewayId,
			Name:   *v.VpnGatewayId,
			Description: model.EC2VPNGatewayDescription{
				VPNGateway: v,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func EC2Region(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.Regions {
		arn := "arn:" + describeCtx.Partition + "::" + *v.RegionName + ":" + describeCtx.AccountID
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.RegionName,
			Description: model.EC2RegionDescription{
				Region: v,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func EC2AvailabilityZone(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)

	regionsOutput, err := client.DescribeRegions(ctx, &ec2.DescribeRegionsInput{
		AllRegions: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, region := range regionsOutput.Regions {
		if region.OptInStatus != nil && *region.OptInStatus != "not-opted-in" {
			continue
		}
		output, err := client.DescribeAvailabilityZones(ctx, &ec2.DescribeAvailabilityZonesInput{
			AllAvailabilityZones: aws.Bool(true),
			Filters: []types.Filter{
				{
					Name:   aws.String("region-name"),
					Values: []string{*region.RegionName},
				},
			},
		})
		if err != nil {
			return nil, err
		}

		for _, v := range output.AvailabilityZones {
			arn := fmt.Sprintf("arn:%s::%s::availability-zone/%s", describeCtx.Partition, *region.RegionName, *v.ZoneName)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   *v.RegionName,
				Description: model.EC2AvailabilityZoneDescription{
					AvailabilityZone: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}
	return values, nil
}

func EC2KeyPair(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.KeyPairs {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":key-pair/" + *v.KeyName
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.KeyName,
			Description: model.EC2KeyPairDescription{
				KeyPair: v,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func GetEC2KeyPair(ctx context.Context, cfg aws.Config, fields map[string]string) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	keyPairID := fields["id"]

	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
		KeyPairIds: []string{keyPairID},
	})
	if err != nil {
		return nil, err
	}

	var values []Resource
	for _, v := range output.KeyPairs {
		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":key-pair/" + *v.KeyName
		values = append(values, Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.KeyName,
			Description: model.EC2KeyPairDescription{
				KeyPair: v,
			},
		})
	}

	return values, nil
}

func EC2AMI(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Owners: []string{"self"},
	})
	if err != nil {
		if isErr(err, "InvalidAMIID.NotFound") || isErr(err, "InvalidAMIID.Unavailable") || isErr(err, "InvalidAMIID.Malformed") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource

	for _, v := range output.Images {
		imageAttribute, err := client.DescribeImageAttribute(ctx, &ec2.DescribeImageAttributeInput{
			Attribute: types.ImageAttributeNameLaunchPermission,
			ImageId:   v.ImageId,
		})
		if err != nil {
			if isErr(err, "InvalidAMIID.NotFound") || isErr(err, "InvalidAMIID.Unavailable") || isErr(err, "InvalidAMIID.Malformed") {
				continue
			}
			return nil, err
		}

		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":image/" + *v.ImageId
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.ImageId,
			Description: model.EC2AMIDescription{
				AMI:               v,
				LaunchPermissions: *imageAttribute,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func EC2ReservedInstances(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	output, err := client.DescribeReservedInstances(ctx, &ec2.DescribeReservedInstancesInput{})
	if err != nil {
		if isErr(err, "InvalidParameterValue") || isErr(err, "InvalidInstanceID.Unavailable") || isErr(err, "InvalidInstanceID.Malformed") {
			return nil, nil
		}
		return nil, err
	}

	var values []Resource

	filterName := "reserved-instances-id"
	for _, v := range output.ReservedInstances {
		var modifications []types.ReservedInstancesModification
		modificationPaginator := ec2.NewDescribeReservedInstancesModificationsPaginator(client, &ec2.DescribeReservedInstancesModificationsInput{
			Filters: []types.Filter{
				{
					Name:   &filterName,
					Values: []string{*v.ReservedInstancesId},
				},
			},
		})
		for modificationPaginator.HasMorePages() {
			page, err := modificationPaginator.NextPage(ctx)
			if err != nil {
				if isErr(err, "InvalidParameterValue") || isErr(err, "InvalidInstanceID.Unavailable") || isErr(err, "InvalidInstanceID.Malformed") {
					continue
				}
				return nil, err
			}

			modifications = append(modifications, page.ReservedInstancesModifications...)
		}

		arn := "arn:" + describeCtx.Partition + ":ec2:" + describeCtx.Region + ":" + describeCtx.AccountID + ":reserved-instances/" + *v.ReservedInstancesId
		resource := Resource{
			Region: describeCtx.KaytuRegion,
			ARN:    arn,
			Name:   *v.ReservedInstancesId,
			Description: model.EC2ReservedInstancesDescription{
				ReservedInstances:   v,
				ModificationDetails: modifications,
			},
		}
		if stream != nil {
			if err := (*stream)(resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, resource)
		}
	}

	return values, nil
}

func EC2IpamPool(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeIpamPoolsPaginator(client, &ec2.DescribeIpamPoolsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.IpamPools {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.IpamPoolArn,
				Name:   *v.IpamPoolId,
				Description: model.EC2IpamPoolDescription{
					IpamPool: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2Ipam(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeIpamsPaginator(client, &ec2.DescribeIpamsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Ipams {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.IpamArn,
				Name:   *v.IpamId,
				Description: model.EC2IpamDescription{
					Ipam: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2InstanceAvailability(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeInstanceTypeOfferingsPaginator(client, &ec2.DescribeInstanceTypeOfferingsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.InstanceTypeOfferings {
			arn := fmt.Sprintf("arn:%s:ec2:%s::instance-type/%s", describeCtx.Partition, *v.Location, v.InstanceType)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   fmt.Sprintf("%s (%s)", v.InstanceType, *v.Location),
				Description: model.EC2InstanceAvailabilityDescription{
					InstanceAvailability: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
	}

	return values, nil
}

func EC2InstanceType(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeInstanceTypesPaginator(client, &ec2.DescribeInstanceTypesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.InstanceTypes {
			arn := fmt.Sprintf("arn:%s:ec2:::instance-type/%s", describeCtx.Partition, v.InstanceType)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    arn,
				Name:   string(v.InstanceType),
				Description: model.EC2InstanceTypeDescription{
					InstanceType: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2ManagedPrefixList(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeManagedPrefixListsPaginator(client, &ec2.DescribeManagedPrefixListsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.PrefixLists {
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ARN:    *v.PrefixListArn,
				Name:   *v.PrefixListName,
				Description: model.EC2ManagedPrefixListDescription{
					ManagedPrefixList: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EC2SpotPrice(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)

	endTime := time.Now()
	startTime := endTime.AddDate(0, 0, -1)
	paginator := ec2.NewDescribeSpotPriceHistoryPaginator(client, &ec2.DescribeSpotPriceHistoryInput{
		StartTime: &startTime,
		EndTime:   &endTime,
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.SpotPriceHistory {
			if v.SpotPrice == nil {
				continue
			}
			avZone := ""
			if v.AvailabilityZone != nil {
				avZone = *v.AvailabilityZone
			}
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				Name:   fmt.Sprintf("%s-%s (%s)", v.InstanceType, *v.SpotPrice, avZone),
				Description: model.EC2SpotPriceDescription{
					SpotPrice: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}

		}
	}

	return values, nil
}

func EC2TransitGatewayRoute(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTransitGatewayRouteTablesPaginator(client, &ec2.DescribeTransitGatewayRouteTablesInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, transitGatewayRouteTable := range page.TransitGatewayRouteTables {
			routes, err := client.SearchTransitGatewayRoutes(ctx, &ec2.SearchTransitGatewayRoutesInput{
				Filters: []types.Filter{
					{
						Name:   aws.String("state"),
						Values: []string{"active", "blackhole", "pending"},
					},
				},
				TransitGatewayRouteTableId: transitGatewayRouteTable.TransitGatewayRouteTableId,
			})
			if err != nil {
				return nil, err
			}
			for _, route := range routes.Routes {
				arn := fmt.Sprintf("arn:%s:ec2:%s:%s:transit-gateway-route-table/%s:%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *transitGatewayRouteTable.TransitGatewayRouteTableId, *route.DestinationCidrBlock)
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ARN:    arn,
					Name:   *route.DestinationCidrBlock,
					Description: model.EC2TransitGatewayRouteDescription{
						TransitGatewayRoute:        route,
						TransitGatewayRouteTableId: *transitGatewayRouteTable.TransitGatewayRouteTableId,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}

	return values, nil
}

func EC2TransitGatewayAttachment(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)

	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeTransitGatewayAttachmentsPaginator(client, &ec2.DescribeTransitGatewayAttachmentsInput{})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.TransitGatewayAttachments {
			arn := fmt.Sprintf("arn:%s:ec2:%s:%s:transit-gateway-attachment/%s", describeCtx.Partition, describeCtx.Region, describeCtx.AccountID, *v.TransitGatewayAttachmentId)
			resource := Resource{
				Region: describeCtx.KaytuRegion,
				ID:     *v.TransitGatewayAttachmentId,
				ARN:    arn,
				Description: model.EC2TransitGatewayAttachmentDescription{
					TransitGatewayAttachment: v,
				},
			}
			if stream != nil {
				if err := (*stream)(resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, resource)
			}
		}
	}

	return values, nil
}

func EbsVolumeMetricReadOps(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVolumesPaginator(client, &ec2.DescribeVolumesInput{
		MaxResults: aws.Int32(500),
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Volumes {
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "5_MIN", "AWS/EBS", "VolumeReadOps", "VolumeId", *v.VolumeId)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ID:     fmt.Sprintf("%s:%s:%s:%s", *v.VolumeId, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.EbsVolumeMetricReadOpsDescription{
						CloudWatchMetricRow: metric,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
		}
	}

	return values, nil
}

func EbsVolumeMetricReadOpsDaily(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVolumesPaginator(client, &ec2.DescribeVolumesInput{
		MaxResults: aws.Int32(500),
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Volumes {
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "DAILY", "AWS/EBS", "VolumeReadOps", "VolumeId", *v.VolumeId)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ID:     fmt.Sprintf("%s:%s:%s:%s-daily", *v.VolumeId, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.EbsVolumeMetricReadOpsDailyDescription{
						CloudWatchMetricRow: metric,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
		}
	}

	return values, nil
}

func EbsVolumeMetricReadOpsHourly(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVolumesPaginator(client, &ec2.DescribeVolumesInput{
		MaxResults: aws.Int32(500),
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Volumes {
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "HOURLY", "AWS/EBS", "VolumeReadOps", "VolumeId", *v.VolumeId)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ID:     fmt.Sprintf("%s:%s:%s:%s-hourly", *v.VolumeId, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.EbsVolumeMetricReadOpsHourlyDescription{
						CloudWatchMetricRow: metric,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
		}
	}

	return values, nil
}

func EbsVolumeMetricWriteOps(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVolumesPaginator(client, &ec2.DescribeVolumesInput{
		MaxResults: aws.Int32(500),
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Volumes {
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "5_MIN", "AWS/EBS", "VolumeWriteOps", "VolumeId", *v.VolumeId)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ID:     fmt.Sprintf("%s:%s:%s:%s", *v.VolumeId, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.EbsVolumeMetricWriteOpsDescription{
						CloudWatchMetricRow: metric,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
		}
	}

	return values, nil
}

func EbsVolumeMetricWriteOpsDaily(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVolumesPaginator(client, &ec2.DescribeVolumesInput{
		MaxResults: aws.Int32(500),
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Volumes {
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "DAILY", "AWS/EBS", "VolumeWriteOps", "VolumeId", *v.VolumeId)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ID:     fmt.Sprintf("%s:%s:%s:%s-daily", *v.VolumeId, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.EbsVolumeMetricWriteOpsDailyDescription{
						CloudWatchMetricRow: metric,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
		}
	}

	return values, nil
}

func EbsVolumeMetricWriteOpsHourly(ctx context.Context, cfg aws.Config, stream *StreamSender) ([]Resource, error) {
	describeCtx := GetDescribeContext(ctx)
	client := ec2.NewFromConfig(cfg)
	paginator := ec2.NewDescribeVolumesPaginator(client, &ec2.DescribeVolumesInput{
		MaxResults: aws.Int32(500),
	})

	var values []Resource
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.Volumes {
			metrics, err := listCloudWatchMetricStatistics(ctx, cfg, "HOURLY", "AWS/EBS", "VolumeWriteOps", "VolumeId", *v.VolumeId)
			if err != nil {
				return nil, err
			}
			for _, metric := range metrics {
				resource := Resource{
					Region: describeCtx.KaytuRegion,
					ID:     fmt.Sprintf("%s:%s:%s:%s-hourly", *v.VolumeId, metric.Timestamp.Format(time.RFC3339), *metric.DimensionName, *metric.DimensionValue),
					Description: model.EbsVolumeMetricWriteOpsHourlyDescription{
						CloudWatchMetricRow: metric,
					},
				}
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}

			}
		}
	}

	return values, nil
}
