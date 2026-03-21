package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
}

type cliApp struct {
	logger Logger
}

func New(logger Logger) *cliApp {
	return &cliApp{
		logger: logger,
	}
}

func (c *cliApp) Run() error {
	type Current struct {
		Temp float32 `json:"temperature_2m"`
	}
	type Response struct {
		Curr Current `json:"current"`
	}
	var response Response

	c.logger.Info("Fetching weather data from Open-Meteo API")

	params := fmt.Sprintf(
		"latitude=%f&longitude=%f&current=temperature_2m",
		53.6688,
		23.8223,
	)
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)

	c.logger.Debug("Request URL: " + url)

	resp, err := http.Get(url)
	if err != nil {
		customErr := errors.New("can`t get weather data from openmeteo")
		c.logger.Error("Failed to get weather data: " + err.Error())
		return errors.Join(customErr, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.logger.Error("can`t close body err- " + err.Error())
		}
	}()

	c.logger.Debug("Response status: " + resp.Status)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		customErr := errors.New("can`t read data from response")
		c.logger.Error("Failed to read response body: " + err.Error())
		return errors.Join(customErr, err)
	}

	c.logger.Debug("Received data length: " + string(rune(len(data))))

	if err := json.Unmarshal(data, &response); err != nil {
		customErr := errors.New("can`t unmarshal data from response")
		c.logger.Error("Failed to unmarshal JSON: " + err.Error())
		return errors.Join(customErr, err)
	}

	c.logger.Info("Weather data successfully retrieved")

	fmt.Printf(
		"Температура воздуха- %.2f градусов цельсия\n",
		response.Curr.Temp,
	)

	return nil
}
