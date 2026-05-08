package flags

import "flag"

const DefaultConfigPath = "./config/config.yaml"

type flags struct {
	Path string
}

func Parse() *flags {
	config := flag.String("config", DefaultConfigPath, "path to config")

	flag.Parse()
	return &flags{
		Path: *config,
	}
}
