package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/RossMerr/jsonschema"
)

var (
	packagename   = flag.String("package", "main", "Go package")
	output        = flag.String("output", "", "Output folder")
	schemaversion = flag.String("schema", "", "If you need to declare a schema")
)

func main() {
	flag.Parse()

	config := jsonschema.Config{
		Packagename:   *packagename,
		Output:        *output,
		Schemaversion: *schemaversion,
	}

	schema := jsonschema.NewSchemaReferences(config)

	var files []string
	for _, file := range flag.Args() {
		fileInfo, err := os.Stat(file)

		if err != nil {
			if os.IsNotExist(err) {
				panic(fmt.Errorf("File does not exist."))
			}

		}

		if fileInfo.IsDir() {
			err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
				if filepath.Ext(path) == ".json" {
					files = append(files, path)
				}
				return nil
			})
			if err != nil {
				panic(err)
			}
		} else {
			if filepath.Ext(file) == ".json" {
				files = append(files, file)
			}
		}
	}

	err := schema.Parse(files)
	if err != nil {
		panic(err)
	}

	err = schema.Generate()
	if err != nil {
		panic(err)
	}
}
