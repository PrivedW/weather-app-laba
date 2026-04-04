package main

import (
	"os"

	pogodaby "github.com/PrivedW/weather-app-laba_info/internal/adapters/pogoda_by"
	"github.com/PrivedW/weather-app-laba_info/internal/adapters/weather"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/flags"
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

	logger := cli.NewSimpleLogger()
	wi := getProvider(c, logger)
	app := cli.New(logger, wi, c)
	err = app.Run()
	if err != nil {
		logger.Error("Some error", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func getProvider(c config.Config, l cli.Logger) cli.WeatherInfo {
	var wi cli.WeatherInfo
	switch c.P.Type {
	case "open-meteo":
		wi = weather.New(l)
	case "pogoda":
		wi = pogodaby.New(l)
	default:
		wi = weather.New(l)
	}
	return wi
}
