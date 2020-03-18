package interpreter

import (
	"os"
	"os/exec"

	"github.com/RossMerr/jsonschema/parser"
)

type Interpret interface {
	ToFile(output string) error
}

type interpret struct {
	root     *parser.Parse
	templateStruct Template
}

func NewInterpret(root *parser.Parse, templateStruct Template) Interpret {
	return &interpret{
		root:     root,
		templateStruct: templateStruct,
	}
}

func NewInterpretDefaults(root *parser.Parse) Interpret {
	return NewInterpret(root, parser.TemplateStruct())
}

func (s *interpret) ToFile(output string) error {
	for _, obj := range s.root.Structs {
		filename := output + obj.ID().Filename() + ".go"
		_, err := os.Stat(filename)
		if !os.IsNotExist(err) {
			err = os.Remove(filename)
			if err != nil {
				return err
			}
		}

		file, err := os.Create(filename)
		if err != nil {
			return err
		}

		err = s.templateStruct.Execute(file, obj)
		if err != nil {
			return err
		}

		cmd := exec.Command("gofmt", "-w", filename)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
