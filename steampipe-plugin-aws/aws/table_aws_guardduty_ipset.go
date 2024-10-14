package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsGuardDutyIPSet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_guardduty_ipset",
		Description: "AWS GuardDuty IPSet",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"detector_id", "ipset_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException", "NoSuchEntityException", "BadRequestException"}),
			},
			Hydrate: opengovernance.GetGuardDutyIPSet,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGuardDutyIPSet,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "detector_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name for the IPSet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSet.Name")},
			{
				Name:        "detector_id",
				Description: "The ID of the detector.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DetectorId")},
			{
				Name:        "ipset_id",
				Description: "The ID of the IPSet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSetId")},
			{
				Name:        "format",
				Description: "The format of the file that contains the IPSet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSet.Format")},
			{
				Name:        "status",
				Description: "The status of IPSet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSet.Status")},
			{
				Name:        "location",
				Description: "The URI of the file that contains the IPSet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSet.Location")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.IPSet.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.IPSet.Tags")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas)},
		}),
	}
}
