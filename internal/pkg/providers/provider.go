package providers

import (
	pogodaby "github.com/PrivedW/weather-app-laba_info/internal/adapters/pogoda_by"
	"github.com/PrivedW/weather-app-laba_info/internal/adapters/weather"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
	"github.com/PrivedW/weather-app-laba_info/pkg/config"
)

func GetProvider(c config.Config, l cli.Logger) cli.WeatherInfo {
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
