package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsRDSReservedDBInstance(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_rds_reserved_db_instance",
		Description: "AWS RDS Reserved DB Instance",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("reserved_db_instance_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ReservedDBInstanceNotFound"}),
			},
			Hydrate: kaytu.GetRDSReservedDBInstance,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListRDSReservedDBInstance,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "class", Require: plugin.Optional},
				{Name: "duration", Require: plugin.Optional},
				{Name: "lease_id", Require: plugin.Optional},
				{Name: "multi_az", Require: plugin.Optional},
				{Name: "offering_type", Require: plugin.Optional},
				{Name: "reserved_db_instances_offering_id", Require: plugin.Optional},
			},
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValue"}),
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "reserved_db_instance_id",
				Description: "The unique identifier for the reservation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.ReservedDBInstanceId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the reserved DB Instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.ReservedDBInstanceArn"),
			},
			{
				Name:        "reserved_db_instances_offering_id",
				Description: "The offering identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.ReservedDBInstancesOfferingId")},
			{
				Name:        "state",
				Description: "The state of the reserved DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.State")},
			{
				Name:        "class",
				Description: "The DB instance class for the reserved DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.DBInstanceClass")},
			{
				Name:        "currency_code",
				Description: "The currency code for the reserved DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.CurrencyCode")},
			{
				Name:        "db_instance_count",
				Description: "The number of reserved DB instances.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ReservedDBInstance.DBInstanceCount")},
			{
				Name:        "duration",
				Description: "The duration of the reservation in seconds.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ReservedDBInstance.Duration")},
			{
				Name:        "fixed_price",
				Description: "The fixed price charged for this reserved DB instance.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.ReservedDBInstance.FixedPrice")},
			{
				Name:        "lease_id",
				Description: "The unique identifier for the lease associated with the reserved DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.LeaseId")},
			{
				Name:        "multi_az",
				Description: "Indicates if the reservation applies to Multi-AZ deployments.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ReservedDBInstance.MultiAZ")},
			{
				Name:        "offering_type",
				Description: "The offering type of this reserved DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.OfferingType")},
			{
				Name:        "product_description",
				Description: "The description of the reserved DB instance.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.ProductDescription")},
			{
				Name:        "start_time",
				Description: "The time the reservation started.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ReservedDBInstance.StartTime")},
			{
				Name:        "usage_price",
				Description: "The hourly price charged for this reserved DB instance.",
				Type:        proto.ColumnType_DOUBLE,
				Transform:   transform.FromField("Description.ReservedDBInstance.UsagePrice")},
			{
				Name:        "recurring_charges",
				Description: "The recurring price charged to run this reserved DB instance.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReservedDBInstance.RecurringCharges")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ReservedDBInstance.ReservedDBInstanceId")},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ReservedDBInstance.ReservedDBInstanceArn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
