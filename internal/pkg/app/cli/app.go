package cli

import (
	"fmt"
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
}

func New(logger Logger, wi WeatherInfo) *cliApp {
	return &cliApp{
		logger: logger,
		wi:     wi,
	}
}

func (c *cliApp) Run() error {
	fmt.Printf(
		"Температура воздуха- %.2f градусов цельсия\n",
		c.wi.GetTemperature(53.6688, 23.8223).Temp,
	)

	return nil
}
