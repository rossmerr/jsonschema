package interpreter

import (
	"fmt"
	"go/token"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
	log "github.com/sirupsen/logrus"
)

type Interpret interface {
	ToFile(output string) ([]string, error)
}

var _ Interpret = (*interpret)(nil)

type interpret struct {
	documents      map[jsonschema.ID]parser.Component
	templateStruct Template
	packagename    string
}

func NewInterpret(documents map[jsonschema.ID]parser.Component, templateStruct Template) Interpret {
	return &interpret{
		documents:      documents,
		templateStruct: templateStruct,
	}
}

func NewInterpretDefaults(documents map[jsonschema.ID]parser.Component) (Interpret, error) {
	templates, err := templates.DefaultSchemaTemplate()
	if err != nil {
		return nil, err
	}
	return NewInterpret(documents, templates), nil
}

func (s *interpret) ToFile(output string) ([]string, error) {
	files := []string{}

	for id, obj := range s.documents {
		filename := path.Join(output, toFilename(id.String())+".go")

		err := s.removeFile(filename)
		if err != nil {
			return files, err
		}

		err = s.makeDir(filename)
		if err != nil {
			return files, err
		}

		file, err := s.createFile(filename)
		if err != nil {
			return files, err
		}
		files = append(files, filename)

		err = s.executeTemplate(file, obj, filename)
		if err != nil {
			return files, err
		}

		err = s.format(filename)
		if err != nil {
			return files, err
		}
	}

	log.Printf(green("✓")+" Create %v files\n", len(s.documents))
	//fmt.Printf(green("✓")+" Create %v files\n", len(s.documents))

	return files, nil
}

func (s *interpret) format(filename string) error {
	cmd := exec.Command("gofmt", "-w", filename)
	data, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf(red("🗴")+"Go Fmt %v\n", filename)
		fmt.Printf("%v\n", string(data))
	}
	return err
}

func (s *interpret) executeTemplate(file *os.File, obj parser.Component, filename string) error {
	err := s.templateStruct.Execute(file, obj)
	if err != nil {
		fmt.Printf(red("🗴")+"Execute template for %v\n", filename)
	}

	log.Infof("Create file %v", filename)
	return err
}

func (s *interpret) createFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf(red("🗴")+"Create %v\n", filename)
	}

	return file, err
}

func (s *interpret) makeDir(filename string) error {
	path := filepath.Dir(filename)
	if path == "." {
		return nil
	}

	err := os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Printf(red("🗴")+"Making %v\n", path)
	}
	return err
}

func (s *interpret) removeFile(filename string) error {
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		err = os.Remove(filename)
		if err != nil {
			fmt.Printf(red("🗴")+"Removing old file %v\n", filename)
		}
		return err
	}
	return nil
}

// toFilename returns the file name from the ID.
func toFilename(s string) string {
	if len(s) < 1 {
		return "."
	}

	index := strings.Index(s, "#")
	var basename string
	if index < 0 {
		basename = filepath.Base(s)
	} else {
		basename = filepath.Base(s[:index])
	}

	name := strings.TrimSuffix(basename, filepath.Ext(basename))

	// Valid field names must start with a unicode letter
	if !unicode.IsLetter(rune(name[0])) {
		name = "No " + name
	}

	// Valid field names must not be a reserved word
	if token.IsKeyword(name) {
		name = "Key " + name
	}

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	filename := reg.ReplaceAllString(strings.Title(clean), "")

	if len(filename) > 0 {
		return string(unicode.ToLower(rune(filename[0]))) + filename[1:]
	}
	return filename
}
