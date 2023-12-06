package aws

import (
	"context"
	"fmt"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWellArchitectedShareInvitation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wellarchitected_share_invitation",
		Description: "AWS Well-Architected Share Invitation",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWellArchitectedShareInvitation,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "workload_name", Require: plugin.Optional},
				{Name: "lens_name", Require: plugin.Optional},
				{Name: "share_resource_type", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "lens_arn",
				Description: "The ARN for the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.LensArn"),
			},
			{
				Name:        "lens_name",
				Description: "The full name of the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.LensName"),
			},
			{
				Name:        "permission_type",
				Description: "Permission granted on a workload share.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.PermissionType"),
			},
			{
				Name:        "share_invitation_id",
				Description: "The ID assigned to the share invitation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.ShareInvitationId"),
			},
			{
				Name:        "share_resource_type",
				Description: "The resource type of the share invitation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(shareResourceType),
			},
			{
				Name:        "shared_by",
				Description: "An Amazon Web Services account ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.SharedBy"),
			},
			{
				Name:        "shared_with",
				Description: "The Amazon Web Services account ID, IAM role, organization ID, or organizational unit (OU) ID with which the workload is shared.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.SharedWith"),
			},
			{
				Name:        "workload_name",
				Description: "The name of the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.WorkloadName"),
			},
			{
				Name:        "workload_id",
				Description: "The ID assigned to the workload.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.WorkloadId"),
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ShareInvitation.ShareInvitationId"),
			},
		}),
	}
}

//// LIST FUNCTION

func listWellArchitectedShareInvitations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	// Create client
	svc, err := WellArchitectedClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_wellarchitected_share_invitation.listWellArchitectedShareInvitations", "client_error", err)
		return nil, err
	}

	// Unsupported region, return no data
	if svc == nil {
		return nil, nil
	}

	// Limiting the results
	maxLimit := int32(50)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	input := &wellarchitected.ListShareInvitationsInput{
		MaxResults: maxLimit,
	}

	if d.EqualsQualString("workload_name") != "" && d.EqualsQualString("lens_name") != "" {
		return nil, fmt.Errorf("you can either pass workload_name or lens_name as an optional qual but not both")
	}

	if d.EqualsQualString("workload_name") != "" {
		input.WorkloadNamePrefix = aws.String(d.EqualsQualString("workload_name"))
	}

	if d.EqualsQualString("lens_name") != "" {
		input.LensNamePrefix = aws.String(d.EqualsQualString("lens_name"))
		input.ShareResourceType = types.ShareResourceTypeLens
	}

	if d.EqualsQualString("share_resource_type") != "" {
		input.ShareResourceType = types.ShareResourceType(d.EqualsQualString("share_resource_type"))
	}

	paginator := wellarchitected.NewListShareInvitationsPaginator(svc, input, func(o *wellarchitected.ListShareInvitationsPaginatorOptions) {
		o.Limit = maxLimit
		o.StopOnDuplicateToken = true
	})

	// List call
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("aws_wellarchitected_share_invitation.listWellArchitectedShareInvitations", "api_error", err)
			return nil, err
		}

		for _, item := range output.ShareInvitationSummaries {
			d.StreamListItem(ctx, item)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func shareResourceType(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	shareInvitation := d.HydrateItem.(types.ShareInvitationSummary)

	if shareInvitation.LensArn != nil && *shareInvitation.LensArn != "" {
		return types.ShareResourceTypeLens, nil
	} else {
		return types.ShareResourceTypeWorkload, nil
	}
}
