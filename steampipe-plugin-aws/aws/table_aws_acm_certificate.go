package aws

import (
	"context"
	"strings"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsAcmCertificate(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_acm_certificate",
		Description: "AWS ACM Certificate",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("certificate_arn"),
			Hydrate:    kaytu.GetCertificateManagerCertificate,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListCertificateManagerCertificate,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "status",
					Require: plugin.Optional,
				},
				{
					Name:    "key_algorithm",
					Require: plugin.Optional,
				},
			},
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "certificate_arn",
				Description: "Amazon Resource Name (ARN) of the certificate. This is of the form: arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.CertificateArn"),
			},
			{
				Name:        "certificate",
				Description: "The ACM-issued certificate corresponding to the ARN specified as input",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.Certificate"),
			},
			{
				Name:        "certificate_chain",
				Description: "The ACM-issued certificate corresponding to the ARN specified as input",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Attributes.CertificateChain"),
			},
			{
				Name:        "domain_name",
				Description: "Fully qualified domain name (FQDN), such as www.example.com or example.com, for the certificate",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.DomainName"),
			},
			{
				Name:        "certificate_transparency_logging_preference",
				Description: "Indicates whether to opt in to or out of certificate transparency logging. Certificates that are not logged typically generate a browser error. Transparency makes it possible for you to detect SSL/TLS certificates that have been mistakenly or maliciously issued for your domain.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.Options.CertificateTransparencyLoggingPreference"),
			},
			{
				Name:        "created_at",
				Description: "The time at which the certificate was requested. This value exists only when the certificate type is AMAZON_ISSUED",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Certificate.CreatedAt"),
			},
			{
				Name:        "subject",
				Description: "The name of the entity that is associated with the public key contained in the certificate",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.Subject"),
			},
			{
				Name:        "imported_at",
				Description: "The name of the certificate authority that issued and signed the certificate",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Certificate.ImportedAt"),
			},
			{
				Name:        "issuer",
				Description: "The name of the certificate authority that issued and signed the certificate",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.Issuer"),
			},
			{
				Name:        "signature_algorithm",
				Description: "The algorithm that was used to sign the certificate",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.SignatureAlgorithm"),
			},
			{
				Name:        "extended_key_usages",
				Description: "Specify one or more ExtendedKeyUsage extension values.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Certificate.ExtendedKeyUsages"),
			},
			{
				Name:        "failure_reason",
				Description: "The reason the certificate request failed. This value exists only when the certificate status is FAILED",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.FailureReason"),
			},
			{
				Name:        "issued_at",
				Description: "A list of ARNs for the AWS resources that are using the certificate. A certificate can be used by multiple AWS resources",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Certificate.IssuedAt"),
			},
			{
				Name:        "status",
				Description: "The status of the certificate",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.Status"),
			},
			{
				Name:        "key_algorithm",
				Description: "The algorithm that was used to generate the public-private key pair",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.KeyAlgorithm"),
			},
			{
				Name:        "not_after",
				Description: "The time after which the certificate is not valid",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Certificate.NotAfter"),
			},
			{
				Name:        "not_before",
				Description: "The time before which the certificate is not valid",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Certificate.NotBefore"),
			},
			{
				Name:        "renewal_eligibility",
				Description: "Specifies whether the certificate is eligible for renewal.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.RenewalEligibility"),
			},
			{
				Name:        "revocation_reason",
				Description: "The reason the certificate was revoked. This value exists only when the certificate status is REVOKED",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.RevocationReason"),
			},
			{
				Name:        "revoked_at",
				Description: "The time at which the certificate was revoked. This value exists only when the certificate status is REVOKED",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Certificate.RevokedAt"),
			},
			{
				Name:        "serial",
				Description: "The serial number of the certificate",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.Serial"),
			},
			{
				Name:        "type",
				Description: "The source of the certificate. For certificates provided by ACM, this value is AMAZON_ISSUED.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.Type"),
			},
			{
				Name:        "domain_validation_options",
				Description: "Contains information about the initial validation of each domain name that occurs as a result of the RequestCertificate request. This field exists only when the certificate type is AMAZON_ISSUED",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Certificate.DomainValidationOptions"),
			},
			{
				Name:        "in_use_by",
				Description: "A list of ARNs for the AWS resources that are using the certificate",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Certificate.InUseBy"),
			},
			{
				Name:        "subject_alternative_names",
				Description: "One or more domain names (subject alternative names) included in the certificate. This list contains the domain names that are bound to the public key that is contained in the certificate. The subject alternative names include the canonical domain name (CN) of the certificate and additional domain names that can be used to connect to the website",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Certificate.SubjectAlternativeNames"),
			},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with certificate",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Certificate.CertificateArn").Transform(certificateArnToTitle),
			},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.From(certificateTurbotTags),
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Certificate.CertificateArn").Transform(arnToAkas),
			},
		}),
	}
}

//// LIST FUNCTION

//// HYDRATE FUNCTIONS

//// TRANSFORM FUNCTIONS

func certificateTurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.CertificateManagerCertificate).Description.Tags
	var turbotTagsMap map[string]string
	if tags != nil {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}
	return turbotTagsMap, nil
}

func certificateArnToTitle(_ context.Context, d *transform.TransformData) (interface{}, error) {
	item := *d.Value.(*string)

	// Get the resource title
	title := item[strings.LastIndex(item, "/")+1:]

	return title, nil
}
