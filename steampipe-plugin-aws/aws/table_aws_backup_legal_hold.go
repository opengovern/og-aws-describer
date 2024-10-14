package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsBackupLegalHold(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_backup_legal_hold",
		Description: "AWS Backup Legal Hold",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("legal_hold_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValueException"}),
			},
			Hydrate: opengovernance.GetBackupLegalHold,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBackupLegalHold,
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "legal_hold_id",
				Description: "ID of specific legal hold on one or more recovery points.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LegalHold.LegalHoldId")},
			{
				Name:        "arn",
				Description: "This is an Amazon Resource Number (ARN) that uniquely identifies the legal hold.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LegalHold.LegalHoldArn"),
			},
			{
				Name:        "creation_date",
				Description: "This is the time in number format when legal hold was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LegalHold.CreationDate")},
			{
				Name:        "status",
				Description: "This is the status of the legal hold. Statuses can be ACTIVE, CREATING, CANCELED, and CANCELING.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LegalHold.Status")},
			{
				Name:        "cancellation_date",
				Description: "This is the time in number format when legal hold was cancelled.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LegalHold.CancellationDate")},
			{
				Name:        "description",
				Description: "This is the description of a legal hold.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LegalHold.Description")},
			{
				Name:        "retain_record_until",
				Description: "This is the date and time until which the legal hold record will be retained.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.LegalHold.RetainRecordUntil")},
			{
				Name:        "recovery_point_selection",
				Description: "This specifies criteria to assign a set of resources, such as resource types or backup vaults.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LegalHold.RecoveryPointSelection")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.LegalHold.Title")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.LegalHold.LegalHoldArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
