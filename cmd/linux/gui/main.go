package main

import (
	"errors"
	"os"

	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/gui"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/flags"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/gui/fyne"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/providers"
	"github.com/PrivedW/weather-app-laba_info/pkg/config"
)

func main() {
	arguments := flags.Parse()

	l := cli.NewSimpleLogger()
	c, err := config.ParsePath(arguments.Path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}

		l.Info("config file was not found, default settings will be used")
		c, err = config.ParseDefault()
		if err != nil {
			panic(err)
		}
	}

	provider := providers.GetProvider(c, l)
	p := fyne.NewP()
	g := gui.New(l, p, provider, c)
	err = g.Run()
	if err != nil {
		panic(err)
	}
}
