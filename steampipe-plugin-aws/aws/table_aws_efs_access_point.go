package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEfsAccessPoint(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_efs_access_point",
		Description: "AWS EFS Access Point",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("access_point_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"AccessPointNotFound"}),
			},
			Hydrate: opengovernance.GetEFSAccessPoint,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEFSAccessPoint,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"FileSystemNotFound"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{Name: "file_system_id", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the access point. This is the value of the Name tag.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.Name")},
			{
				Name:        "access_point_id",
				Description: "The ID of the access point, assigned by Amazon EFS.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.AccessPointId")},
			{
				Name:        "access_point_arn",
				Description: "The unique Amazon Resource Name (ARN) associated with the access point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.AccessPointArn")},
			{
				Name:        "life_cycle_state",
				Description: "Identifies the lifecycle phase of the access point.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.LifeCycleState")},
			{
				Name:        "file_system_id",
				Description: "The ID of the EFS file system that the access point applies to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.FileSystemId")},
			{
				Name:        "client_token",
				Description: "The opaque string specified in the request to ensure idempotent creation.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.ClientToken")},
			{
				Name:        "owner_id",
				Description: "Identified the AWS account that owns the access point resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccessPoint.OwnerId")},
			{
				Name:        "posix_user",
				Description: "The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by NFS clients using the access point.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AccessPoint.PosixUser")},
			{
				Name:        "root_directory",
				Description: "The directory on the Amazon EFS file system that the access point exposes as the root directory to NFS clients using the access point.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AccessPoint.RootDirectory")},
			{
				Name:        "tags_src",
				Description: "The tags associated with the access point, presented as an array of Tag objects.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AccessPoint.Tags")},

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(efsAccessPointTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(efsAccessPointTitle),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.AccessPoint.AccessPointArn").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func efsAccessPointTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	tagList := d.HydrateItem.(opengovernance.EFSAccessPoint).Description.AccessPoint

	if tagList.Tags == nil {
		return nil, nil
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string

	if tagList.Tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tagList.Tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

// Generate title for the resource
func efsAccessPointTitle(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(opengovernance.EFSAccessPoint).Description.AccessPoint

	// If name is available, then setting name as title, else setting Access Point ID as title
	if data.Name != nil {
		return data.Name, nil
	}

	return data.AccessPointId, nil
}
