//go:generate go run ./table_aws_service_template.go

package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var table = Table{
	Service: "EC2",
	Name:    "LaunchTemplate",
}

const awsTableServiceTemplate = `
package aws

import (
	"context"
	"github.com/kaytu-io/kaytu-aws-describer/pkg/kaytu-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAws{{.Service}}{{.Name}}(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_{{.ServiceLowerCase}}_{{.NameLowerCase}}",
		Description: "AWS {{.Service}} {{.Name}}",
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("arn"), //TODO: change this to the primary key columns in model.go
			Hydrate:           kaytu.Get{{.Service}}{{.Name}},
		},
		List: &plugin.ListConfig{
			Hydrate: kaytu.List{{.Service}}{{.Name}},
		},
		Columns: awskaytuRegionalColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the {{.NameLowerCase}}.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.{{.Name}}.Id")},
			{
				Name:        "name",
				Description: "The name of the {{.NameLowerCase}}.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.{{.Name}}.Name")},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the {{.NameLowerCase}}",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ARN")},
			{
				Name:        "title",
				Description: resourceInterfaceDescription("title"),
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.{{.Name}}.Name")},
			{
				Name:        "tags",
				Description: resourceInterfaceDescription("tags"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.{{.Name}}.Tags"), // probably needs a transform function
			},
			{
				Name:        "akas",
				Description: resourceInterfaceDescription("akas"),
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("ARN").Transform(arnToAkas),
			},
		}),
	}
}

//// TRANSFORM FUNCTIONS

func get{{.Service}}{{.Name}}TurbotTags(_ context.Context, d *transform.TransformData) (interface{}, error) {
	tags := d.HydrateItem.(kaytu.{{.Service}}{{.Name}}).Description.{{.Name}}.Tags
	return ec2V2TagsToMap(tags)
}
`

type Table struct {
	Service, Name, ServiceLowerCase, NameLowerCase string
}

func main() {
	tmpl := template.New("aws_table")
	tmpl, err := tmpl.Parse(awsTableServiceTemplate)
	if err != nil {
		panic(err)
	}

	table.ServiceLowerCase = strings.ToLower(table.Service)
	table.NameLowerCase = strings.ToLower(table.Name)

	fileName := fmt.Sprintf("../table_aws_%s_%s.go", table.ServiceLowerCase, table.NameLowerCase)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, table)
	if err != nil {
		panic(err)
	}
}
