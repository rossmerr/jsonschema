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
	template Template
}

func NewInterpret(root *parser.Parse, template Template) Interpret {
	return &interpret{
		root:     root,
		template: template,
	}
}

func NewInterpretDefaults(root *parser.Parse) Interpret {
	return NewInterpret(root, parser.Template())
}

func (s *interpret) ToFile(output string) error {
	for _, obj := range s.root.Structs {
		filename := output + obj.ID.Filename() + ".go"
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

		err = s.template.Execute(file, parser.Document{
			obj,
		})
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
