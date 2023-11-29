package aws

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateDocs(t *testing.T) {
	plg := Plugin(context.Background())

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	pathPrefix := "./docs/table_def"
	if strings.HasSuffix(currentDir, "aws") {
		pathPrefix = "." + pathPrefix
	}

	for _, table := range plg.TableMap {
		doc := `# Columns  

<table>
	<tr><td>Column Name</td><td>Description</td></tr>
`
		for _, column := range table.Columns {
			desc := column.Description
			desc = html.EscapeString(desc)
			doc += fmt.Sprintf(`	<tr><td>%s</td><td>%s</td></tr>
`, column.Name, desc)
		}

		doc += "</table>"

		err := os.WriteFile(pathPrefix+"/"+table.Name+".md", []byte(doc), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func TestGenerateTableList(t *testing.T) {
	var tablesFiles []string
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.HasPrefix(info.Name(), "table_") {
			name := info.Name()
			name = name[6:]
			name = strings.ReplaceAll(name, ".go", "")
			if name == "aws_api_gateway_api_authorizer" {
				name = "aws_api_gateway_authorizer"
			}
			tablesFiles = append(tablesFiles, name)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	plg := Plugin(context.Background())

	awsResourceType, err := os.ReadFile("../../../kaytu-deploy/kaytu/inventory-data/aws-resource-types.json")
	if err != nil {
		panic(err)
	}

	var rts []ResourceType
	err = json.Unmarshal(awsResourceType, &rts)
	if err != nil {
		panic(err)
	}

	for _, rt := range rts {
		exists := false
		for idx, _ := range plg.TableMap {
			if rt.SteampipeTable == idx {
				exists = true
			}
		}
		if !exists {
			panic("rt " + rt.SteampipeTable + " does not exists")
		}
	}

	for idx, _ := range plg.TableMap {
		exists := false
		for _, rt := range rts {
			if rt.SteampipeTable == idx {
				exists = true
			}
		}
		if !exists {
			fmt.Println(idx + " not supported")
		}
	}

	var cv [][]string
	for _, t := range tablesToPopulate {
		resourceName := ""
		status := ""
		for _, rt := range rts {
			if rt.SteampipeTable == t {
				resourceName = rt.ResourceName
				status = rt.Discovery
			}
		}
		cv = append(cv, []string{t, resourceName, status})
	}
	csvfiler, err := os.Create("tableInformation-part2.csv")
	if err != nil {
		panic(err)
	}
	defer csvfiler.Close()

	csvWriter := csv.NewWriter(csvfiler)
	err = csvWriter.WriteAll(cv)
	if err != nil {
		panic(err)
	}
	csvWriter.Flush()

}

type ResourceType struct {
	ResourceName         string
	ResourceLabel        string
	Category             []string
	Tags                 map[string][]string
	TagsString           string `json:"-"`
	ServiceName          string
	ListDescriber        string
	GetDescriber         string
	TerraformName        []string
	TerraformNameString  string `json:"-"`
	TerraformServiceName string
	Discovery            string
	IgnoreSummarize      bool
	SteampipeTable       string
	Model                string
}

var tablesToPopulate = []string{
	"aws_ec2_localgateway",
}
