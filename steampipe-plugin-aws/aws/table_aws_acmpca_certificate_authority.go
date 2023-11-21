package aws

import (
	"context"
	"strings"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsACMPCACertificateAuthority(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_acmpca_certificate_authority",
		Description: "AWS ACMPCA CertificateAuthority",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("arn"),
			Hydrate:    kaytu.GetACMPCACertificateAuthority,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListACMPCACertificateAuthority,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "arn",
				Description: "Amazon Resource Name (ARN) for your private certificate authority (CA). The format is 12345678-1234-1234-1234-123456789012.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.Arn"),
			},
			{
				Name:        "created_at",
				Description: "Date and time at which your private CA was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CertificateAuthority.CreatedAt"),
			},
			{
				Name:        "failure_reason",
				Description: "Reason the request to create your private CA failed.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.FailureReason"),
			},
			{
				Name:        "key_storage_security_standard",
				Description: "Defines a cryptographic key management compliance standard used for handling CA keys. Default: FIPS_140_2_LEVEL_3_OR_HIGHER Note: Amazon Web Services Region ap-northeast-3 supports only FIPS_140_2_LEVEL_2_OR_HIGHER. You must explicitly specify this parameter and value when creating a CA in that Region. Specifying a different value (or no value) results in an InvalidArgsException with the message 'A certificate authority cannot be created in this region with the specified security standard.'",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.KeyStorageSecurityStandard"),
			},
			{
				Name:        "last_state_change_at",
				Description: "Date and time at which your private CA was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CertificateAuthority.LastStateChangeAt"),
			},
			{
				Name:        "not_after",
				Description: "Date and time after which your private CA certificate is not valid.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CertificateAuthority.NotAfter"),
			},
			{
				Name:        "not_before",
				Description: "Date and time before which your private CA certificate is not valid.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CertificateAuthority.NotBefore"),
			},
			{
				Name:        "owner_account",
				Description: "The Amazon Web Services account ID that owns the certificate authority.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.OwnerAccount"),
			},
			{
				Name:        "restorable_until",
				Description: "The period during which a deleted CA can be restored. For more information, see the PermanentDeletionTimeInDays parameter of the DeleteCertificateAuthorityRequest action.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CertificateAuthority.RestorableUntil"),
			},
			{
				Name:        "serial",
				Description: "Serial number of your private CA.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.Serial"),
			},
			{
				Name:        "status",
				Description: "Status of your private CA.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.Status"),
			},
			{
				Name:        "type",
				Description: "Type of your private CA.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.Type"),
			},
			{
				Name:        "usage_mode",
				Description: "Specifies whether the CA issues general-purpose certificates that typically require a revocation mechanism, or short-lived certificates that may optionally omit revocation because they expire quickly. Short-lived certificate validity is limited to seven days. The default value is GENERAL_PURPOSE.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.UsageMode"),
			},
			{
				Name:        "certificate_authority_configuration",
				Description: "Your private CA configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CertificateAuthority.CertificateAuthorityConfiguration"),
			},
			{
				Name:        "revocation_configuration",
				Description: "Information about the Online Certificate Status Protocol (OCSP) configuration or certificate revocation list (CRL) created and maintained by your private CA.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.CertificateAuthority.RevocationConfiguration"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with private certificate authority (CA).",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Steampipe Standard Columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CertificateAuthority.Arn").Transform(authorityArnToTitle),
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
				Transform:   transform.FromField("Description.CertificateAuthority.Arn").Transform(transform.EnsureStringArray),
			},
		}),
	}
}

func authorityArnToTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := *d.Value.(*string)

	// Get the resource title
	title := item[strings.LastIndex(item, "/")+1:]

	return title, nil
}
