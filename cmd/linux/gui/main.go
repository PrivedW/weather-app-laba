package main

import (
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

	r, err := os.Open(arguments.Path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = r.Close()
	}()

	c, err := config.Parse(r)
	if err != nil {
		panic(err)
	}

	l := cli.NewSimpleLogger()
	provider := providers.GetProvider(c, l)
	p := fyne.NewP()
	g := gui.New(l, p, provider, c)
	err = g.Run()
	if err != nil {
		panic(err)
	}
}
