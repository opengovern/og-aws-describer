package aws

import (
	"context"

	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcrImageScanFinding(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecr_image_scan_finding",
		Description: "AWS ECR Image Scan Finding",
		List: &plugin.ListConfig{
			Hydrate: kaytu.ListECRImage,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"RepositoryNotFoundException", "ImageNotFoundException", "ScanNotFoundException"}),
			},
			// Ideally image_tag and image_digest could both be used as optional
			// key columns, but the query planner only works with required key
			// columns when there are multiple. We chose image_tag instead of
			// image_digest as it's more common/friendly to use.
			KeyColumns: []*plugin.KeyColumn{
				{Name: "repository_name", Require: plugin.Required},
				{Name: "image_tag", Require: plugin.Required},
			},
		},

		Columns: awsKaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "repository_name",
				Description: "The name of the repository.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Image.RepositoryName"),
			},
			{
				Name:        "image_tag",
				Description: "The image tag",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ImageTag")},
			{
				Name:        "image_digest",
				Description: "The image digest",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ImageDigest")},
			{
				Name:        "name",
				Description: "The name associated with the finding, usually a CVE number.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ImageScanFinding.Name")},
			{
				Name:        "severity",
				Description: "The finding severity.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ImageScanFinding.Severity")},
			{
				Name:        "attributes",
				Description: "A collection of attributes of the host from which the finding is generated.",
				Type:        proto.ColumnType_JSON,

				Transform: transform.FromField("Description.ImageScanFinding.Attributes")},
			{
				Name:        "description",
				Description: "The description of the finding.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ImageScanFinding.Description")},
			{
				Name:        "uri",
				Description: "A link containing additional details about the security vulnerability.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ImageScanFinding.Uri")},
			{
				Name:        "image_scan_status",
				Description: "The current state of the scan",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ImageScanStatus.Status")},
			{
				Name:        "image_scan_status_description",
				Description: "The description of the image scan status.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ImageScanStatus.Description")},
			{
				Name:        "image_scan_completed_at",
				Description: "The date and time, in JavaScript date format, when the repository was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.ImageScanCompletedAt")},
			{
				Name:        "vulnerability_source_updated_at",
				Description: "The date and time, in JavaScript date format, when the repository was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.VulnerabilitySourceUpdatedAt")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ImageScanFinding.Name")},
		}),
	}
}
