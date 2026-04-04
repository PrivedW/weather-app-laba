package cli

import (
	"fmt"

	"github.com/PrivedW/weather-app-laba_info/internal/domain/models"
	"github.com/PrivedW/weather-app-laba_info/pkg/config"
)

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string, err error)
}

type WeatherInfo interface {
	GetTemperature(float64, float64) (models.TempInfo, error)
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
	tempInfo, err := c.wi.GetTemperature(c.c.L.Lat, c.c.L.Long)
	if err != nil {
		c.logger.Error("can`t get temp info", err)
		return err
	}

	fmt.Printf(
		"Температура воздуха- %.2f градусов цельсия\n",
		tempInfo.Temp,
	)

	return nil
}
