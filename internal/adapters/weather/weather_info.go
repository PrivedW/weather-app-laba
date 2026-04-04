package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/PrivedW/weather-app-laba_info/internal/domain/models"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
)

const url = "https://api.open-meteo.com/v1/forecast"

type response struct {
	Curr models.TempInfo `json:"current"`
}
type weatherInfo struct {
	c        models.TempInfo
	l        cli.Logger
	isLoaded bool
}

func New(l cli.Logger) cli.WeatherInfo {
	return &weatherInfo{
		l: l,
	}
}
func (wi *weatherInfo) getWeatherInfo(lat, long float64) error {
	if wi.isLoaded {
		return nil
	}

	var response response
	params := fmt.Sprintf(
		"latitude=%f&longitude=%f&current=temperature_2m",
		lat,
		long,
	)
	url := fmt.Sprintf("%s?%s", url, params)
	wi.l.Debug(fmt.Sprintf("url was generated success- %s", url))
	resp, err := http.Get(url)
	if err != nil {
		wi.l.Error("can`t get weather data", err)
		customErr := errors.New("can`t get weather data from openmeteo")
		return errors.Join(customErr, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			wi.l.Error("can`t close body", err)
		}
	}()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		wi.l.Error("can`t read data from body", err)
		customErr := errors.New("can`t read data from response")
		return errors.Join(customErr, err)
	}
	wi.l.Debug(fmt.Sprintf("data was readed successfuly size- %d",
		len(data)),
	)
	if err := json.Unmarshal(data, &response); err != nil {
		wi.l.Error("can`t unmarshal json data", err)
		customErr := errors.New("can`t unmarshal data from response")
		return errors.Join(customErr, err)
	}
	wi.c = response.Curr
	wi.isLoaded = true
	return nil
}
func (wi *weatherInfo) GetTemperature(lat, long float64) (models.TempInfo, error) {
	err := wi.getWeatherInfo(lat, long)
	return models.TempInfo{
		Temp: wi.c.Temp,
	}, err
}
