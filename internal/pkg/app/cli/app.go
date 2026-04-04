package cli

import (
	"fmt"

	"github.com/PrivedW/weather-app-laba_info/pkg/config"
)

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string, err error)
}

type Current struct {
	Temp float32 `json:"temperature_2m"`
}

type WeatherInfo interface {
	GetTemperature(float64, float64) Current
}

type cliApp struct {
	logger Logger
	wi     WeatherInfo
	c      config.Config
}

func New(logger Logger, wi WeatherInfo, c config.Config) *cliApp {
	return &cliApp{
		logger: logger,
		wi:     wi,
		c:      c,
	}
}

func (c *cliApp) Run() error {
	fmt.Printf(
		"Температура воздуха- %.2f градусов цельсия\n",
		c.wi.GetTemperature(c.c.L.Lat, c.c.L.Long).Temp,
	)

	return nil
}
