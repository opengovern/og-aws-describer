package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	directoryservicev1 "github.com/aws/aws-sdk-go/service/directoryservice"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsDirectoryServiceCertificate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_directory_service_certificate",
		Description: "AWS Directory Service Certificate",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"directory_id", "certificate_id"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"CertificateDoesNotExistException", "DirectoryDoesNotExistException", "InvalidParameterException"}),
			},
			Hydrate: kaytu.GetDirectoryServiceCertificate,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListDirectoryServiceCertificate,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"DirectoryDoesNotExistException"}),
			},
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "directory_id",
					Require: plugin.Optional,
				},
			},
		},
		GetMatrixItemFunc: SupportedRegionMatrix(directoryservicev1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "directory_id",
				Description: "The directory identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DirectoryId"),
			},
			{
				Name:        "certificate_id",
				Description: "The identifier of the certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.CertificateId"),
			},
			{
				Name:        "common_name",
				Description: "The common name for the certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.CommonName"),
			},
			{
				Name:        "type",
				Description: "The function that the registered certificate performs. Valid values include ClientLDAPS or ClientCertAuth. The default value is ClientLDAPS.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.Type"),
			},
			{
				Name:        "state",
				Description: "The state of the certificate. Valid values: Registering | Registered | RegisterFailed | Deregistering | Deregistered | DeregisterFailed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.State"),
			},
			{
				Name:        "expiry_date_time",
				Description: "The date and time when the certificate will expire.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Certificate.ExpiryDateTime"),
			},
			{
				Name:        "registered_date_time",
				Description: "The date and time that the certificate was registered.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Certificate.RegisteredDateTime"),
			},
			{
				Name:        "state_reason",
				Description: "Describes a state change for the certificate.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.StateReason"),
			},
			{
				Name:        "client_cert_auth_settings",
				Description: "A ClientCertAuthSettings object that contains client certificate authentication settings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Certificate.ClientCertAuthSettings"),
			},

			// Steampipe Standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.CommonName"),
			},
		}),
	}
}

type CertInfo struct {
	DirectoryId            *string
	CertificateId          *string
	ClientCertAuthSettings *types.ClientCertAuthSettings
	CommonName             *string
	ExpiryDateTime         *time.Time
	RegisteredDateTime     *time.Time
	State                  types.CertificateState
	StateReason            *string
	Type                   types.CertificateType
}
