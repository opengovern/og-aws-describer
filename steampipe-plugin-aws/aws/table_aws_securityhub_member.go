package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

// AWS SDK v1 to v2 migration isusing middleware
// due to https://github.com/aws/aws-sdk-go-v2/issues/1884

//// TABLE DEFINITION

func tableAwsSecurityHubMember(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_securityhub_member",
		Description: "AWS Securityhub Member",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecurityHubMember,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidInputException", "BadRequestException"}),
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "member_account_id",
				Description: "The Amazon Web Services account ID of the member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.AccountId")},
			{
				Name:        "administrator_id",
				Description: "The Amazon Web Services account ID of the Security Hub administrator account associated with this member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.AdministratorId")},
			{
				Name:        "email",
				Description: "The email address of the member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.Email")},
			{
				Name:        "invited_at",
				Description: "A timestamp for the date and time when the invitation was sent to the member account.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Member.InvitedAt")},
			{
				Name:        "master_id",
				Description: "The Amazon Web Services account ID of the Security Hub administrator account associated with this member account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.MasterId")},
			{
				Name:        "member_status",
				Description: "The status of the relationship between the member account and its administrator account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Member.MemberStatus")},
			{
				Name:        "updated_at",
				Description: "The timestamp for the date and time when the member account was updated.",
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
