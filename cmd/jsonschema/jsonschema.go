package main

import (
	"fmt"
	"go/scanner"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/RossMerr/jsonschema/interpreter"
)

var (
	packagename = flag.StringP("package", "p", "main", "go package name to use")
	output      = flag.StringP("output", "o", ".", "output folder")
	loglevel    = flag.StringP("loglevel", "l", "warn", "standard logger level")
)

var (
	exitCode = 0
	files    []string
	red      = color.FgRed.Render
)

func report(err error) {
	scanner.PrintError(os.Stderr, fmt.Errorf(red("✗")+" %v\n", err))
	exitCode = 2
}

func isJsonFile(f os.FileInfo) bool {
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

func expandHome(path string) string {
	if strings.HasPrefix(path, "~/") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		return filepath.Join(dir, path[2:])
	}
	return path
}

func walkDir(path string) {
	filepath.Walk(path, visitFile)
}

func main() {
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	level, err := log.ParseLevel(*loglevel)
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)

	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		path = expandHome(path)
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
		log.Error(err)
		os.Exit(2)
	}
	_, err = interpret.ToFile(*output)
	if err != nil {
		log.Error(err)
		os.Exit(2)
	}
	os.Exit(exitCode)
}
