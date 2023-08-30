package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsConfigConfigurationRecorder(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_config_configuration_recorder",
		Description: "AWS Config Configuration Recorder",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"NoSuchConfigurationRecorderException"}),
			},
			Hydrate: kaytu.GetConfigConfigurationRecorder,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListConfigConfigurationRecorder,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the recorder. By default, AWS Config automatically assigns the name default when creating the configuration recorder.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationRecorder.Name"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the configuration recorder.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getAwsConfigurationRecorderARN,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "recording_group",
				Description: "Specifies the types of AWS resources for which AWS Config records configuration changes.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConfigurationRecorder.RecordingGroup"),
			},
			{
				Name:        "role_arn",
				Description: "Amazon Resource Name (ARN) of the IAM role used to describe the AWS resources associated with the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationRecorder.RoleARN"),
			},
			{
				Name:        "status_recording",
				Description: "Specifies whether or not the recorder is currently recording.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ConfigurationRecordersStatus.Recording"),
			},
			{
				Name:        "status",
				Description: "The current status of the configuration recorder.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ConfigurationRecordersStatus"),
			},
			// Standard columns
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Hydrate:     getAwsConfigurationRecorderARN,
				Transform:   transform.FromValue().Transform(transform.EnsureStringArray),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ConfigurationRecorder.Name"),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func getAwsConfigurationRecorderARN(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	region := d.EqualsQualString(matrixKeyRegion)

	configurationRecorder := h.Item.(kaytu.ConfigConfigurationRecorder).Description.ConfigurationRecorder
	metadata := h.Item.(kaytu.ConfigConfigurationRecorder).Metadata

	arn := "arn:" + metadata.Partition + ":config:" + region + ":" + metadata.AccountID + ":config-recorder" + "/" + *configurationRecorder.Name

	return arn, nil
}
