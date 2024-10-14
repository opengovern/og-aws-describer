package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsIamUserServiceSpecificCredential(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_service_specific_credential",
		Description: "AWS IAM User Service Specific Credential",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIAMServiceSpecificCredential,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "service_name", Require: plugin.Optional},
				{Name: "user_name", Require: plugin.Optional},
			},
		},
		Columns: awsKaytuColumns([]*plugin.Column{
			{
				Name:        "service_name",
				Description: "The name of the service associated with the service-specific credential.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceSpecificCredential.ServiceName")},
			{
				Name:        "service_specific_credential_id",
				Description: "The unique identifier for the service-specific credential.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceSpecificCredential.ServiceSpecificCredentialId")},
			{
				Name:        "create_date",
				Description: "The date and time, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601), when the service-specific credential were created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ServiceSpecificCredential.CreateDate")},
			{
				Name:        "service_user_name",
				Description: "The generated user name for the service-specific credential.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceSpecificCredential.ServiceUserName")},
			{
				Name:        "status",
				Description: "The status of the service-specific credential. Active means that the key is valid for API calls, while Inactive means it is not.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceSpecificCredential.Status")},
			{
				Name:        "user_name",
				Description: "The name of the IAM user associated with the service-specific credential.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceSpecificCredential.UserName")},

			// Standard columns for all tables
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ServiceSpecificCredential.ServiceName")},
		}),
	}
}
