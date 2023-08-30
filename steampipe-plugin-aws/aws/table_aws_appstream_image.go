package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	appstreamv1 "github.com/aws/aws-sdk-go/service/appstream"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAppStreamImage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_appstream_image",
		Description: "AWS AppStream Image",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListAppStreamImage,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "arn",
					Require: plugin.Optional,
				},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(appstreamv1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Name"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Arn"),
			},
			{
				Name:        "appstream_agent_version",
				Description: "The version of the AppStream 2.0 agent to use for instances that are launched from this image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.AppstreamAgentVersion"),
			},
			{
				Name:        "base_image_arn",
				Description: "The ARN of the image from which this image was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.BaseImageArn"),
			},
			{
				Name:        "created_time",
				Description: "The time the image was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Image.CreatedTime"),
			},
			{
				Name:        "description",
				Description: "The description to display.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Description"),
			},
			{
				Name:        "display_name",
				Description: "The image name to display.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.DisplayName"),
			},
			{
				Name:        "image_builder_name",
				Description: "The name of the image builder that was used to create the private image. If the image is shared, this value is null.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.ImageBuilderName"),
			},
			{
				Name:        "image_builder_supported",
				Description: "Indicates whether an image builder can be launched from this image.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Image.ImageBuilderSupported"),
			},
			{
				Name:        "platform",
				Description: "The operating system platform of the image.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Platform"),
			},
			{
				Name:        "public_base_image_released_date",
				Description: "The release date of the public base image. For private images, this date is the release date of the base image from which the image was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Image.PublicBaseImageReleasedDate"),
			},
			{
				Name:        "state",
				Description: "The image starts in the PENDING state. If image creation succeeds, the state is AVAILABLE. If image creation fails, the state is FAILED.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.State"),
			},
			{
				Name:        "visibility",
				Description: "Indicates whether the image is public or private.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Visibility"),
			},
			{
				Name:        "applications",
				Description: "The applications associated with the image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.Applications"),
			},
			{
				Name:        "image_errors",
				Description: "Describes the errors that are returned when a new image can't be created.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.ImageErrors"),
			},
			{
				Name:        "image_permissions",
				Description: "The permissions to provide to the destination AWS account for the specified image.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.ImagePermissions"),
			},
			{
				Name:        "state_change_reason",
				Description: "The reason why the last state change occurred.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.StateChangeReason"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.Name"),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Image.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}
