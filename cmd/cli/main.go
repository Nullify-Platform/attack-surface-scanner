package main

import (
	"os"

	"github.com/nullify-platform/attack-surface-scanner/internal/scan"
	"github.com/nullify-platform/logger/pkg/logger"

	"github.com/alexflint/go-arg"
)

type Scan struct {
	SpecPath   string `arg:"--spec-path" help:"The file path to the OpenAPI file (both yaml and json are supported) e.g. ./openapi.yaml"`
	TargetHost string `arg:"--target-host" help:"The base URL of the API to be scanned e.g. https://api.nullify.ai"`
}

type args struct {
	Scan *Scan `arg:"subcommand:scan" help:"test the given app for vulnerabilities"`

	Verbose bool `arg:"-v" help:"enable verbose logging"`
	Debug   bool `arg:"-d" help:"enable debug logging"`
}

func (args) Version() string {
	return logger.Version
}

func main() {
	var args args
	p := arg.MustParse(&args)

	logLevel := "warn"
	if args.Verbose {
		logLevel = "info"
	}
	if args.Debug {
		logLevel = "debug"
	}
	log, err := logger.ConfigureDevelopmentLogger(logLevel)
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	switch {
	case args.Scan != nil:
		err := scan.Scan(args.Scan.SpecPath, args.Scan.TargetHost)
		if err != nil {
			logger.Error(
				"failed to run scan",
				logger.Err(err),
			)
			os.Exit(1)
		}
	default:
		p.WriteHelp(os.Stdout)
	}
}
