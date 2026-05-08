package main

import (
	"errors"
	"os"

	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/flags"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/providers"
	"github.com/PrivedW/weather-app-laba_info/pkg/config"
)

func main() {
	arguments := flags.Parse()

	logger := cli.NewSimpleLogger()
	c, err := config.ParsePath(arguments.Path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}

		logger.Info("config file was not found, default settings will be used")
		c, err = config.ParseDefault()
		if err != nil {
			panic(err)
		}
	}

	wi := providers.GetProvider(c, logger)
	app := cli.New(logger, wi, c)
	err = app.Run()
	if err != nil {
		logger.Error("Some error", err)
		os.Exit(1)
	}
	os.Exit(0)
}
