package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsGuardDutyMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_guardduty_member",
		Description: "AWS GuardDuty Member",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"member_account_id", "detector_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException", "BadRequestException"}),
			},
			Hydrate: kaytu.GetGuardDutyMember,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListGuardDutyMember,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "detector_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "member_account_id",
				Description: "The ID of the member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.AccountId")},
			{
				Name:        "detector_id",
				Description: "The detector ID of the member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.DetectorId")},
			{
				Name:        "master_id",
				Description: "The administrator account ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.MasterId")},
			{
				Name:        "email",
				Description: "The email address of the member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.Email")},
			{
				Name:        "invited_at",
				Description: "The timestamp when the invitation was sent.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Member.InvitedAt")},
			{
				Name:        "relationship_status",
				Description: "The status of the relationship between the member and the administrator.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.RelationshipStatus")},
			{
				Name:        "updated_at",
				Description: "The last-updated timestamp of the member.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Member.UpdatedAt")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.AccountId")},
		}),
	}
}
