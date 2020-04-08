package interpreter

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"unicode"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
)

type Interpret interface {
	ToFile(output string) ([]string, error)
}

type interpret struct {
	documents      map[jsonschema.ID]*parser.Document
	templateStruct Template
	packagename string
}

func NewInterpret(documents map[jsonschema.ID]*parser.Document, templateStruct Template) Interpret {
	return &interpret{
		documents:      documents,
		templateStruct: templateStruct,
	}
}

func NewInterpretDefaults(documents map[jsonschema.ID]*parser.Document) (Interpret, error) {
	templates, err := templates.DefaultSchemaTemplate()
	if err != nil {
		return nil, err
	}
	return NewInterpret(documents, templates), nil
}

func (s *interpret) ToFile(output string) ([]string, error) {
	files := []string{}
	green := color.FgCyan.Render
	red := color.FgRed.Render

	for _, obj := range s.documents {
		obj.WithPackageName(s.packagename)

		filename := path.Join(output, toFilename(obj.ID)+".go")

		_, err := os.Stat(filename)
		if !os.IsNotExist(err) {
			err = os.Remove(filename)
			if err != nil {
				fmt.Printf(red("🗴")+"Removing old file %v\n", filename)
				return files, err
			}
		}

		path := filepath.Dir(filename)
		if path != "." {
			err = os.MkdirAll(path, 0777)
		}
		if err != nil {
			fmt.Printf(red("🗴")+"Making %v\n", path)
			return files, err
		}

		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf(red("🗴")+"Create %v\n", filename)
			return files, err
		}
		files = append(files, filename)

		err = s.templateStruct.Execute(file, obj)
		if err != nil {
			fmt.Printf(red("🗴")+"Execute template for %v\n", filename)
			return files, err
		}

		log.Infof("Create file %v", filename)

		cmd := exec.Command("gofmt", "-w", filename)
		data, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf(red("🗴")+"Go Fmt %v\n", filename)
			fmt.Printf("%v\n", string(data))
			return files, err
		}
	}

	fmt.Printf(green("✓")+" Create %v files\n", len(s.documents))

	return files, nil
}

// toFilename returns the file name from the ID.
func toFilename(s string) string {
	id := jsonschema.ID(s)
	name := id.ToTypename()

	if len(name) > 0 {
		return string(unicode.ToLower(rune(name[0]))) + name[1:]
	}
	return name
}
