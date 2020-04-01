package interpreter

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/types"
	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
)

type Interpret interface {
	ToFile(output string) ([]string, error)
}

type interpret struct {
	root           *parser.Parse
	templateStruct Template
}

func NewInterpret(root *parser.Parse, templateStruct Template) Interpret {
	return &interpret{
		root:           root,
		templateStruct: templateStruct,
	}
}

func NewInterpretDefaults(root *parser.Parse) (Interpret, error) {
	templates, err := types.Template()
	if err != nil {
		return nil, err
	}
	return NewInterpret(root, templates), nil
}

func (s *interpret) ToFile(output string) ([]string, error) {
	files := []string{}
	green := color.FgCyan.Render
	red := color.FgRed.Render

	for _, obj := range s.root.Structs {
		filename := path.Join(output, obj.Filename+".go")

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

	fmt.Printf(green("✓")+" Create %v files\n", len(s.root.Structs))

	return files, nil
}
