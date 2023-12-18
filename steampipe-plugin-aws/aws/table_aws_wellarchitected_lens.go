package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsWellArchitectedLens(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_wellarchitected_lens",
		Description: "AWS Well-Architected Lens",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListWellArchitectedLens,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ResourceNotFoundException", "ValidationException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "lens_name", Require: plugin.Optional},
				{Name: "lens_status", Require: plugin.Optional},
				{Name: "lens_type", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "lens_name",
				Description: "The full name of the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LensSummary.LensName"),
			},
			{
				Name:        "lens_alias",
				Description: "The alias of the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LensSummary.LensAlias"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LensSummary.LensArn"),
			},
			{
				Name:        "created_at",
				Description: "The date and time when the lens was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LensSummary.CreatedAt"),
			},
			{
				Name:        "updated_at",
				Description: "The date and time when the lens was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LensSummary.UpdatedAt"),
			},
			{
				Name:        "description",
				Description: "The description of the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LensSummary.Description"),
			},
			{
				Name:        "lens_status",
				Description: "The status of the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LensSummary.LensStatus"),
			},
			{
				Name:        "lens_type",
				Description: "The type of the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LensSummary.LensType").Transform(transform.ToString),
			},
			{
				Name:        "lens_version",
				Description: "The version of the lens.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LensSummary.LensVersion"),
			},
			{
				Name:        "owner",
				Description: "An Amazon Web Services account ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LensSummary.Owner"),
			},
			{
				Name:        "share_invitation_id",
				Description: "The ID assigned to the shared invitation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lens.ShareInvitationId"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Lens.Name"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Lens.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Lens.LensArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

type LensInfo struct {
	CreatedAt         *time.Time
	Description       *string
	LensAlias         *string
	LensArn           *string
	LensName          *string
	LensStatus        types.LensStatus
	LensType          types.LensType
	LensVersion       *string
	Owner             *string
	UpdatedAt         *time.Time
	ShareInvitationId *string
	Tags              map[string]string
}

//// LIST FUNCTION

func listWellArchitectedLenses(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create session
	svc, err := WellArchitectedClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aws_wellarchitected_lens.listWellArchitectedLenses", "connection_error", err)
		return nil, err
	}
	if svc == nil {
		// Unsupported region, return no data
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

	input := &wellarchitected.ListLensesInput{
		MaxResults: aws.Int32(maxLimit),
	}

	if d.EqualsQualString("lens_status") != "" {
		input.LensStatus = types.LensStatusType(d.EqualsQualString("lens_status"))
	}
	if d.EqualsQualString("lens_name") != "" {
		input.LensName = aws.String(d.EqualsQualString("lens_name"))
	}
	if d.EqualsQualString("lens_type") != "" {
		input.LensType = types.LensType(d.EqualsQualString("lens_type"))
	}

	paginator := wellarchitected.NewListLensesPaginator(svc, input, func(o *wellarchitected.ListLensesPaginatorOptions) {
		o.Limit = maxLimit
		o.StopOnDuplicateToken = true
	})

	// List call
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("aws_wellarchitected_lens.listWellArchitectedLenses", "api_error", err)
			return nil, err
		}

		for _, item := range output.LensSummaries {
			// As per the doc(https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_LensSummary.html)
			// The lens alias is same as lens arn if it is a custom lens.
			// API returns nil value for lens arn in the case of custom lenses, so we are replacing the nil value with lens arn.
			if item.LensAlias == nil {
				item.LensAlias = item.LensArn
			}

			d.StreamListItem(ctx, LensInfo{
				CreatedAt:   item.CreatedAt,
				Description: item.Description,
				LensAlias:   item.LensAlias,
				LensArn:     item.LensArn,
				LensName:    item.LensName,
				LensStatus:  item.LensStatus,
				LensType:    item.LensType,
				LensVersion: item.LensVersion,
				Owner:       item.Owner,
				UpdatedAt:   item.UpdatedAt,
			})

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}
