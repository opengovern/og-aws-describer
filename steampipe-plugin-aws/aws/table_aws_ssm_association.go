package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsSSMAssociation(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_association",
		Description: "AWS SSM Association",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("association_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"AssociationDoesNotExist", "ValidationException"}),
			},
			Hydrate: kaytu.GetSSMAssociation,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSSMAssociation,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "association_name", Require: plugin.Optional},
				{Name: "instance_id", Require: plugin.Optional},
				{Name: "status", Require: plugin.Optional},
				{Name: "last_execution_date", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "association_id",
				Description: "The ID created by the system when you create an association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssociationItem.AssociationId")},
			{
				Name:        "association_name",
				Description: "The Name of association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssociationItem.AssociationName")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) specifying the association.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getSSMAssociationARN,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "document_name",
				Description: "The name of the Systems Manager document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssociationItem.Name")},
			{
				Name:        "date",
				Description: "The date when the association was made.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssociationItem.LastExecutionDate")},
			{
				Name:        "compliance_severity",
				Description: "A cron expression that specifies a schedule when the association runs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Association.AssociationDescription.ComplianceSeverity")},
			{
				Name:        "apply_only_at_cron_interval",
				Description: "By default, when you create a new associations, the system runs it immediately after it is created and then according to the schedule you specified. Specify this option if you don't want an association to run immediately after you create it. This parameter is not supported for rate expressions.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Association.AssociationDescription.ApplyOnlyAtCronInterval")},
			{
				Name:        "association_version",
				Description: "The association version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssociationItem.AssociationVersion")},
			{
				Name:        "automation_target_parameter_name",
				Description: "Specify the target for the association. This target is required for associations that use an Automation document and target resources by using rate controls.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Association.AssociationDescription.AutomationTargetParameterName")},
			{
				Name:        "document_version",
				Description: "The version of the document used in the association.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssociationItem.DocumentVersion")},
			{
				Name:        "instance_id",
				Description: "The ID of the instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssociationItem.InstanceId")},
			{
				Name:        "last_execution_date",
				Description: "The date on which the association was last run.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.AssociationItem.LastExecutionDate")},
			{
				Name:        "last_successful_execution_date",
				Description: "The last date on which the association was successfully run.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Association.AssociationDescription.LastSuccessfulExecutionDate")},
			{
				Name:        "last_update_association_date",
				Description: "The date when the association was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Association.AssociationDescription.LastUpdateAssociationDate")},
			{
				Name:        "schedule_expression",
				Description: "A cron expression that specifies a schedule when the association runs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AssociationItem.ScheduleExpression")},
			{
				Name:        "max_concurrency",
				Description: "The maximum number of targets allowed to run the association at the same time.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Association.AssociationDescription.MaxConcurrency")},
			{
				Name:        "max_errors",
				Description: "The number of errors that are allowed before the system stops sending requests to run the association on additional targets.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Association.AssociationDescription.MaxErrors")},
			{
				Name:        "sync_compliance",
				Description: "The mode for generating association compliance. You can specify AUTO or MANUAL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Association.AssociationDescription.SyncCompliance")},
			{
				Name:        "overview",
				Description: "Information about the association.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssociationItem.Overview")},
			{
				Name:        "output_location",
				Description: "An S3 bucket where you want to store the output details of the request.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Association.AssociationDescription.OutputLocation")},
			{
				Name:        "parameters",
				Description: "A description of the parameters for a document.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Association.AssociationDescription.Parameters")},
			{
				Name:        "status",
				Description: "The status of the association. Status can be: Pending, Success, or Failed.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AssociationItem.Overview.Status")},
			{
				Name:        "targets",
				Description: "A cron expression that specifies a schedule when the association runs.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AssociationItem.Targets")},
			{
				Name:        "target_locations",
				Description: "The combination of AWS Regions and AWS accounts where you want to run the association.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Association.AssociationDescription.TargetLocations")},

			// Steampipe Standard columns

			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.AssociationItem.Name")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getSSMAssociationARN,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func getSSMAssociationARN(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)
	associationData := h.Item.(kaytu.SSMAssociation)

	arn := "arn:" + associationData.Metadata.Partition + ":ssm:" + region + ":" + associationData.Metadata.AccountID + ":association/" + *associationData.Description.Association.AssociationDescription.AssociationId

	return arn, nil
}
