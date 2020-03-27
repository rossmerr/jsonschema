package main

import (
	"go/scanner"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/RossMerr/jsonschema/interpreter"
)

var (
	packagename = flag.StringP("package", "p", "main", "Go package name to use")
	output      = flag.StringP("output", "o", ".", "Output folder")
)

var (
	exitCode = 0
	files    []string
)

func report(err error) {
	scanner.PrintError(os.Stderr, err)
	exitCode = 2
}

func isJsonFile(f os.FileInfo) bool {
	// ignore non-JSON files
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".json")
}

func visitFile(path string, f os.FileInfo, err error) error {
	if err == nil && isJsonFile(f) {
		files = append(files, path)
	}
	if err != nil && !os.IsNotExist(err) {
		report(err)
	}
	return nil
}

func walkDir(path string) {
	filepath.Walk(path, visitFile)
}

func main() {
	flag.Parse()

	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		switch dir, err := os.Stat(path); {
		case err != nil:
			report(err)
		case dir.IsDir():
			walkDir(path)
		default:
			files = append(files, path)
		}
	}

	interpreter := interpreter.NewInterpreterDefaults(*packagename)
	interpret, err := interpreter.Interpret(files)
	if err != nil {
		report(err)
	}
	err = interpret.ToFile(*output)
	if err != nil {
		report(err)
	}
	os.Exit(exitCode)
}
