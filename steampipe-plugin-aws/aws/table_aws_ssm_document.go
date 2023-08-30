package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAwsSSMDocument(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ssm_document",
		Description: "AWS SSM Document",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"ValidationException", "InvalidDocument"}),
			},
			Hydrate: kaytu.GetSSMDocument,
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListSSMDocument,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "document_type", Require: plugin.Optional},
				{Name: "owner_type", Require: plugin.Optional},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the Systems Manager document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.Name")},
			{
				Name:        "account_ids",
				Description: "The account IDs that have permission to use this document.The ID can be either an AWS account or All.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Permissions.AccountIds")},
			{
				Name:        "account_sharing_info_list",
				Description: "A list of AWS accounts where the current document is shared and the version shared with each account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Permissions.AccountSharingInfoList")},
			{
				//error
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.From(getAwsSSMDocumentArn),
			},
			{
				Name:        "approved_version",
				Description: "The version of the document currently approved for use in the organization.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.ApprovedVersion")},
			{
				Name:        "attachments_information",
				Description: "Details about the document attachments, including names, locations, sizes,and so on.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Document.Document.AttachmentsInformation")},
			{
				Name:        "author",
				Description: "The user in your organization who created the document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.Author")},
			{
				Name:        "created_date",
				Description: "The date when the document was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Document.Document.CreatedDate")},
			{
				Name:        "default_version",
				Description: "The default version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.DefaultVersion")},
			{
				Name:        "description",
				Description: "A description of the document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.Description")},
			{
				Name:        "document_format",
				Description: "The document format, either JSON or YAML.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DocumentIdentifier.DocumentFormat")},
			{
				Name:        "document_type",
				Description: "The type of document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.DocumentType")},
			{
				Name:        "document_version",
				Description: "The document version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DocumentIdentifier.DocumentVersion")},
			{
				Name:        "hash",
				Description: "The Sha256 or Sha1 hash created by the system when the document was created.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.Hash")},
			{
				Name:        "hash_type",
				Description: "The hash type of the document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.HashType")},
			{
				Name:        "latest_version",
				Description: "The latest version of the document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.LatestVersion")},
			{
				Name:        "owner_type",
				Description: "The AWS user account type to filter the documents. Possible values: Self, Amazon, Public, Private, ThirdParty, All, Default.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("owner_type"),
			},
			{
				Name:        "owner",
				Description: "The AWS user account that created the document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.Owner")},
			{
				Name:        "parameters",
				Description: "A description of the parameters for a document.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Document.Document.Parameters")},
			{
				Name:        "pending_review_version",
				Description: "The version of the document that is currently under review.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.PendingReviewVersion")},
			{
				Name:        "platform_types",
				Description: "The operating system platform.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Document.Document.PlatformTypes")},
			{
				Name:        "requires",
				Description: "A list of SSM documents required by a document.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DocumentIdentifier.Requires")},
			{
				Name:        "review_information",
				Description: "Details about the review of a document.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Document.Document.ReviewInformation")},
			{
				Name:        "review_status",
				Description: "The current status of the review.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.DocumentIdentifier.ReviewStatus")},
			{
				Name:        "schema_version",
				Description: "The schema version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.SchemaVersion")},
			{
				Name:        "sha1",
				Description: "The SHA1 hash of the document, which you can use for verification.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.Sha1")},
			{
				Name:        "status",
				Description: "The user in your organization who created the document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.Status")},
			{
				Name:        "status_information",
				Description: "A message returned by AWS Systems Manager that explains the Status value.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.StatusInformation")},
			{
				Name:        "tags_src",
				Description: "A list of tags associated with document",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Document.Document.Tags")},
			{
				Name:        "target_type",
				Description: "The target type which defines the kinds of resources the document can run on.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.TargetType")},
			{
				Name:        "version_name",
				Description: "The version of the artifact associated with the document.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Document.Document.VersionName")},

			// Standard columns

			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Document.Document.Parameters")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.Document.Document.Tags").Transform(ssmDocumentTagListToTurbotTags),
			},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Document.Document.Name")},
		}),
	}
}

func ssmDocumentTagListToTurbotTags(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.HydrateItem == nil {
		return nil, nil
	}

	var tags []types.Tag
	switch item := d.HydrateItem.(type) {
	case *types.DocumentDescription:
		tags = item.Tags
	case types.DocumentIdentifier:
		tags = item.Tags
	}

	// Mapping the resource tags inside turbotTags
	var turbotTagsMap map[string]string
	if len(tags) > 0 {
		turbotTagsMap = map[string]string{}
		for _, i := range tags {
			turbotTagsMap[*i.Key] = *i.Value
		}
	}

	return turbotTagsMap, nil
}

func getAwsSSMDocumentArn(_ context.Context, d *transform.TransformData) (interface{}, error) {
	document := d.HydrateItem.(kaytu.SSMDocument).Description.Document
	metadata := d.HydrateItem.(kaytu.SESConfigurationSet).Metadata

	arn := "arn:" + metadata.Partition + ":ssm:" + metadata.Region + ":" + metadata.AccountID + ":document"

	if strings.HasPrefix(*document.Document.Name, "/") {
		arn = arn + *document.Document.Name
	} else {
		arn = arn + "/" + *document.Document.Name
	}

	return arn, nil
}
