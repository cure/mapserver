package params

import (
	"flag"
)

type ParamsType struct {
	Config       string
	World        string
	Help         bool
	Version      bool
	Debug        bool
	CreateConfig bool
}

func Parse() ParamsType {
	params := ParamsType{}

	flag.StringVar(&(params.Config), "config", "", "Load configuration from specified file")
	flag.StringVar(&(params.World), "world", "", "Set world path")
	flag.BoolVar(&(params.Help), "help", false, "Show help")
	flag.BoolVar(&(params.Version), "version", false, "Show version")
	flag.BoolVar(&(params.Debug), "debug", false, "Enable debug log")
	flag.BoolVar(&(params.CreateConfig), "createconfig", false, "Creates a config and exits")
	flag.Parse()

	return params
}

func PrintHelp() {
	flag.PrintDefaults()
}
