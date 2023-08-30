package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDirectoryServiceDirectory(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_directory_service_directory",
		Description: "AWS Directory Service Directory",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("directory_id"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"InvalidParameterValueException", "ResourceNotFoundFault", "EntityDoesNotExistException"}),
			},
			Hydrate: kaytu.GetDirectoryServiceDirectory,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDirectoryServiceDirectory,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "directory_id",
					Require: plugin.Optional,
				},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The fully qualified name of the directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.Name")},
			{
				Name:        "directory_id",
				Description: "The directory identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.DirectoryId")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that uniquely identifies the directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN"),
			},
			{
				Name:        "stage",
				Description: "The current stage of the directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.Stage")},
			{
				Name:        "type",
				Description: "The directory type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.Type")},
			{
				Name:        "access_url",
				Description: "The access URL for the directory, such as http://<alias>.awsapps.com.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.AccessUrl")},
			{
				Name:        "alias",
				Description: "The alias for the directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.Alias")},
			{
				Name:        "description",
				Description: "The description for the directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.Description")},
			{
				Name:        "desired_number_of_domain_controllers",
				Description: "The desired number of domain controllers in the directory if the directory is Microsoft AD.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Directory.DesiredNumberOfDomainControllers")},
			{
				Name:        "edition",
				Description: "The edition associated with this directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.Edition")},
			{
				Name:        "launch_time",
				Description: "Specifies when the directory was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Directory.LaunchTime")},
			{
				Name:        "radius_status",
				Description: "The status of the RADIUS MFA server connection.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.RadiusStatus")},
			{
				Name:        "share_method",
				Description: "The method used when sharing a directory to determine whether the directory should be shared within your AWS organization (ORGANIZATIONS) or with any AWS account by sending a shared directory request (HANDSHAKE).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.ShareMethod")},
			{
				Name:        "share_notes",
				Description: "A directory share request that is sent by the directory owner to the directory consumer.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.ShareNotes")},
			{
				Name:        "share_status",
				Description: "Current directory status of the shared AWS Managed Microsoft AD directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.ShareStatus")},
			{
				Name:        "short_name",
				Description: "The short name of the directory.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.ShortName")},
			{
				Name:        "size",
				Description: "The directory size.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.Size")},
			{
				Name:        "sso_enabled",
				Description: "Indicates if single sign-on is enabled for the directory. For more information, see EnableSso and DisableSso.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Directory.SsoEnabled")},
			{
				Name:        "stage_last_updated_date_time",
				Description: "The date and time that the stage was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Directory.StageLastUpdatedDateTime")},
			{
				Name:        "stage_reason",
				Description: "Additional information about the directory stage.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.StageReason")},
			{
				Name:        "connect_settings",
				Description: "A DirectoryConnectSettingsDescription object that contains additional information about an AD Connector directory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Directory.ConnectSettings")},
			{
				Name:        "dns_ip_addrs",
				Description: "The IP addresses of the DNS servers for the directory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Directory.DnsIpAddrs")},
			{
				Name:        "event_topics",
				Description: "Amazon SNS topic names that receive status messages from the specified Directory ID.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.EventTopics"),
			},
			{
				Name:        "snapshot_limit",
				Description: "Obtains the manual snapshot limits for a directory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Snapshot.ManualSnapshotsLimit")},
			{
				Name:        "owner_directory_description",
				Description: "Describes the AWS Managed Microsoft AD directory in the directory owner account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Directory.OwnerDirectoryDescription")},
			{
				Name:        "radius_settings",
				Description: "A RadiusSettings object that contains information about the RADIUS server.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Directory.RadiusSettings")},
			{
				Name:        "regions_info",
				Description: "Lists the Regions where the directory has replicated.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Directory.RegionsInfo")},
			{
				Name:        "shared_directories",
				Description: "Details about the shared directory in the directory owner account for which the share request in the directory consumer account has been accepted.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.SharedDirectory"),
			},
			{
				Name:        "vpc_settings",
				Description: "A DirectoryVpcSettingsDescription object that contains additional information about a directory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Directory.VpcSettings")},
			{
				Name:        "tags_src",
				Description: "A list of tags currently associated with the Directory Service Directory.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Directory.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(directoryServiceDirectoryTurbotData),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func directoryServiceDirectoryTurbotData(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.DirectoryServiceDirectory).Description.Tags

	// Mapping the resource tags inside turbotTags
	if tags != nil {
		turbotTagsMap := map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
		return turbotTagsMap, nil
	}
	return nil, nil
}
