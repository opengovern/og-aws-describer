package aws

import (
	"context"
	"github.com/opengovern/og-aws-describer/pkg/opengovernance-es-sdk"

	ecrv1 "github.com/aws/aws-sdk-go/service/ecr"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAwsEcrRegistryScanningConfiguration(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_ecr_registry_scanning_configuration",
		Description: "AWS ECR Registry Scanning Configuration",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.GetECRRegistryScanningConfiguration,
		},
		GetMatrixItemFunc: SupportedRegionMatrix(ecrv1.EndpointsID),
		Columns: awsRegionalColumns([]*plugin.Column{
			{
				Name:        "registry_id",
				Description: "The ID of the registry.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "scanning_configuration",
				Description: "The scanning configuration for the registry.",
				Type:        proto.ColumnType_JSON,
			},
			// Steampipe standard columns
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("RegistryId"),
			},
		}),
	}
}
