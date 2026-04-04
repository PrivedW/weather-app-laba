package pogodaby

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PrivedW/weather-app-laba_info/internal/domain/models"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
)

const url = "http://pogoda.by/api/v2/weather-fact?station=26820"

type resp struct {
	Date         string  `json:"d"`
	Temp         float32 `json:"t"`
	Humidity     int     `json:"w"`
	PressureStat float32 `json:"pStation"`
	PressureSea  float32 `json:"pSea"`
	Trend        int     `json:"tend"`
	Code         any     `json:"code"`
	SpeedWind    float32 `json:"speedWind"`
	SpeedWindMax float32 `json:"speedWindMax"`
	Visibility   int     `json:"vis_m"`
	DirWind      int     `json:"dirWind"`
}

type pogoda struct {
	l cli.Logger
}

func New(l cli.Logger) cli.WeatherInfo {
	return &pogoda{l: l}
}

func (p *pogoda) GetTemperature(lat, long float64) (models.TempInfo, error) {
	response, err := http.Get(url)
	if err != nil {
		p.l.Error("can`t get data from pogoda.by", err)
		return models.TempInfo{}, err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			p.l.Error("can`t close response body", err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			p.l.Error("can`t read error response body", err)
			return models.TempInfo{}, err
		}

		err = fmt.Errorf("pogoda.by returned status %d: %s", response.StatusCode, string(body))
		p.l.Error("can`t get valid response from pogoda.by", err)
		return models.TempInfo{}, err
	}

	var r resp
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		p.l.Error("can`t decode JSON", err)
		return models.TempInfo{}, err
	}
	return models.TempInfo{
		Temp: r.Temp,
	}, nil
}
