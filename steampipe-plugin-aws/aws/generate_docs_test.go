package aws

import (
	"context"
	"fmt"
	"os"
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
		doc := `<table>
	<tr><td>Column Name</td><td>Description</td></tr>
`
		for _, column := range table.Columns {
			doc += fmt.Sprintf(`	<tr><td>%s</td><td>%s</td></tr>
`, column.Name, column.Description)
		}

		doc += "</table>"

		err := os.WriteFile(pathPrefix+"/"+table.Name+".md", []byte(doc), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
