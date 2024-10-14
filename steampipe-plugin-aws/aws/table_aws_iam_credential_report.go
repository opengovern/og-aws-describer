package aws

import (
	"context"

	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsIamCredentialReport(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_iam_credential_report",
		Description: "AWS IAM Credential Report",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIAMCredentialReport,
		},
		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "user_name",
				Description: "The friendly name of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CredentialReport.UserName"),
			},
			{
				Name:        "user_arn",
				Description: "The Amazon Resource Name (ARN) of the user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CredentialReport.UserArn"),
			},
			{
				Name:        "user_creation_time",
				Description: "The date and time when the user was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.UserCreationTime"),
			},
			{
				Name:        "generated_time",
				Description: "The date and time when the credential report was created, in ISO 8601 date-time format (http://www.iso.org/iso/iso8601).",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.GeneratedTime"),
			},
			{
				Name:        "access_key_1_active",
				Description: "Does the user have an access key and is the access key's status Active.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey1Active"),
			},
			{
				Name:        "access_key_1_last_rotated",
				Description: "The date and time when the user's access key was created or last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey1LastRotated").NullIfEqual("N/A"),
			},
			{
				Name:        "access_key_1_last_used_date",
				Description: "The date and time when the user's access key was most recently used to sign an AWS API request. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey1LastUsedDate").NullIfEqual("N/A"),
			},
			{
				Name:        "access_key_1_last_used_region",
				Description: "The AWS Region in which the access key was most recently used. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey1LastUsedRegion").NullIfEqual("N/A"),
			},
			{
				Name:        "access_key_1_last_used_service",
				Description: "The AWS service that was most recently accessed with the access key. The value in this field uses the service's namespace—for example, s3 for Amazon S3 and ec2 for Amazon EC2. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey1LastUsedService").NullIfEqual("N/A"),
			},
			{
				Name:        "access_key_2_active",
				Description: "Does the user have a second access key and is the access key's status Active.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey2Active"),
			},
			{
				Name:        "access_key_2_last_rotated",
				Description: "The date and time when the user's second access key was created or last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey2LastRotated").NullIfEqual("N/A"),
			},
			{
				Name:        "access_key_2_last_used_date",
				Description: "The date and time when the user's second access key was most recently used to sign an AWS API request. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey2LastUsedDate").NullIfEqual("N/A"),
			},
			{
				Name:        "access_key_2_last_used_region",
				Description: "The AWS Region in which the user's second access key was most recently used. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey2LastUsedRegion").NullIfEqual("N/A"),
			},
			{
				Name:        "access_key_2_last_used_service",
				Description: "The AWS service that was most recently accessed with the user's second access key. The value in this field uses the service's namespace—for example, s3 for Amazon S3 and ec2 for Amazon EC2. When an access key is used more than once in a 15-minute span, only the first use is recorded in this field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CredentialReport.AccessKey2LastUsedService").NullIfEqual("N/A"),
			},
			{
				Name:        "cert_1_active",
				Description: "Does the user have an X.509 signing certificate and is that certificate's status Active.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CredentialReport.Cert1Active"),
			},
			{
				Name:        "cert_1_last_rotated",
				Description: "The date and time when the user's signing certificate was created or last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.Cert1LastRotated").NullIfEqual("N/A"),
			},
			{
				Name:        "cert_2_active",
				Description: "Does the user have a second X.509 signing certificate and is that certificate's status Active.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CredentialReport.Cert2Active"),
			},
			{
				Name:        "cert_2_last_rotated",
				Description: "The date and time when the user's second signing certificate was created or last changed.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.Cert2LastRotated").NullIfEqual("N/A"),
			},
			{
				Name:        "mfa_active",
				Description: "Whether a multi-factor authentication (MFA) device has been enabled for the user.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CredentialReport.MFAActive"),
			},
			{
				Name:        "password_enabled",
				Description: "When the user has a password, this value is true. Otherwise it is false. The value for the AWS account root user is always false.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.CredentialReport.PasswordEnabled").Transform(passwordEnabledToBool),
			},
			{
				Name:        "password_last_changed",
				Description: "The date and time when the user's password was last set.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.PasswordLastChanged").NullIfEqual("N/A").NullIfEqual("not_supported"),
			},
			{
				Name:        "password_last_used",
				Description: "The date and time when the AWS account root user or IAM user's password was last used to sign in to an AWS website.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.PasswordLastUsed").NullIfEqual("N/A").NullIfEqual("no_information"),
			},
			{
				Name:        "password_status",
				Description: "The status of an user password. Password status can be one of used, never_used and not_set.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CredentialReport.PasswordLastUsed").Transform(passwordStatus),
			},
			{
				Name:        "password_next_rotation",
				Description: "When the account has a password policy that requires password rotation, this field contains the date and time.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.CredentialReport.PasswordNextRotation").NullIfEqual("N/A").NullIfEqual("not_supported"),
			},
		}),
	}
}

//// LIST FUNCTION

//// TRANSFORM FUNCTIONS

func passwordEnabledToBool(_ context.Context, d *transform.TransformData) (interface{}, error) {
	enabled := types.SafeString(d.Value)
	// The value for the AWS root account <root_account> is always returned as not_supported for password_enabled by AWS API. The root password can not be disabled.
	if enabled == "not_supported" {
		return true, nil
	}
	if enabled == "true" {
		return true, nil
	}
	return false, nil
}

func passwordStatus(_ context.Context, d *transform.TransformData) (interface{}, error) {
	used := types.SafeString(d.Value)
	pwdStatus := "used"
	if used == "no_information" {
		pwdStatus = "never_used"
	} else if used == "N/A" {
		pwdStatus = "not_set"
	}
	return pwdStatus, nil
}
