package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamVirtualMfaDevice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_virtual_mfa_device",
		Description: "AWS IAM Virtual MFA device",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListIAMVirtualMFADevice,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "assignment_status", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "serial_number",
				Description: "The serial number associated with VirtualMFADevice.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMFADevice.SerialNumber"),
			},
			{
				Name:        "enable_date",
				Description: "The date and time on which the virtual MFA device was enabled.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VirtualMFADevice.EnableDate"),
			},
			{
				Name:        "assignment_status",
				Description: "The status (Unassigned or Assigned) of the device.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getAssignmentStatus),
			},
			//we don't have this field in struct
			{
				Name:        "user_id",
				Description: "The user id of the user associated with this virtual MFA device.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMFADevice.User.UserId"),
			},
			{
				Name:        "user_name",
				Description: "The friendly name of the user associated with this virtual MFA device.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.VirtualMFADevice.User.UserName"),
			},
			{
				Name:        "user",
				Description: "Details of the IAM user associated with this virtual MFA device.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMFADevice.User"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags attached with the MFA device.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Tags"),
			},

			// {
			// 	Name:        "base_32_string",
			// 	Description: "The friendly name identifying the user",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("Base32StringSeed"),
			// },
			// {
			// 	Name:        "QRCodePNG",
			// 	Description: "A QR code PNG image that encodes otpauth://totp/$virtualMFADeviceName@$AccountName?secret=$Base32String where $virtualMFADeviceName is one of the create call arguments. AccountName is the user name if set (otherwise, the account ID otherwise), and Base32String is the seed in base32 format. The Base32String value is base64-encoded.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("QRCodePNG"),
			// },

			// Standard columns for all tables
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(virtualMfaDeviceTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.VirtualMFADevice.SerialNumber"),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.VirtualMFADevice.SerialNumber").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func virtualMfaDeviceTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.IAMVirtualMFADevice).Description.Tags
	var turbotTagsMap map[string]string
	if tags == nil {
		return nil, nil
	}

	turbotTagsMap = map[string]string{}
	for _, i := range tags {
		turbotTagsMap[*i.Key] = *i.Value
	}

	return &turbotTagsMap, nil
}

func getAssignmentStatus(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(kaytu.IAMVirtualMFADevice).Description.VirtualMFADevice
	if data.User != nil {
		return "Assigned", nil
	}
	return "Unassigned", nil
}
