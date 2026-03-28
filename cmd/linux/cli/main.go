package main

import (
	"fmt"
	"os"

	"github.com/PrivedW/weather-app-laba_info/internal/adapters/weather"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
)

func main() {
	logger := cli.NewSimpleLogger()
	app := cli.New(logger, weather.New(logger))
	err := app.Run()
	if err != nil {
		fmt.Printf("Some error- %s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
