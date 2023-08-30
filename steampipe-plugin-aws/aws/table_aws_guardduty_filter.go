package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsGuardDutyFilter(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_guardduty_filter",
		Description: "AWS GuardDuty Filter",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"detector_id", "name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException", "NoSuchEntityException", "BadRequestException"}),
			},
			Hydrate: kaytu.GetGuardDutyFilter,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGuardDutyFilter,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "detector_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name for the filter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Filter.Name")},
			{
				Name:        "detector_id",
				Description: "The ID of the detector.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DetectorId")},
			{
				Name:        "action",
				Description: "Specifies the action that is to be applied to the findings that match the filter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Filter.Action")},
			{
				Name:        "description",
				Description: "The description of the filter.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Filter.Description")},
			{
				Name:        "rank",
				Description: "Specifies the position of the filter in the list of current filters. Also specifies the order in which this filter is applied to the findings.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Filter.Rank")},
			{
				Name:        "finding_criteria",
				Description: "Represents the criteria to be used in the filter for querying findings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Filter.FindingCriteria")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Filter.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Filter.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
