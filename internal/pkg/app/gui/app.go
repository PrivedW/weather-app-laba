package gui

import (
	"github.com/PrivedW/weather-app-laba_info/internal/domain/gui_settings"
	"github.com/PrivedW/weather-app-laba_info/internal/pkg/app/cli"
	"github.com/PrivedW/weather-app-laba_info/pkg/config"
)

type guiApp struct {
	l        cli.Logger
	p        guisettings.Provider
	provider cli.WeatherInfo
	c        config.Config
}

func New(l cli.Logger, p guisettings.Provider, provider cli.WeatherInfo, c config.Config) *guiApp {
	return &guiApp{
		l:        l,
		p:        p,
		provider: provider,
		c:        c,
	}
}

func (g *guiApp) Run() error {
	tempInfo, err := g.provider.GetTemperature(g.c.L.Lat, g.c.L.Long)
	if err != nil {
		g.l.Error("can`t get temp info", err)
		return err
	}

	w, err := g.p.CreateWindow("Информер погоды", guisettings.NewWS(400, 200))
	if err != nil {
		g.l.Error("can`t create window", err)
		return err
	}

	tw := g.p.GetTextWidget("")
	if err := w.SetTemperatureWidget(tw); err != nil {
		g.l.Error("can`t set temperature widget", err)
		return err
	}

	if err := w.UpdateTemperature(tempInfo.Temp); err != nil {
		g.l.Error("can`t update temperature", err)
		return err
	}

	if err := w.Render(); err != nil {
		g.l.Error("can`t render window", err)
		return err
	}

	ar := g.p.GetAppRunner()
	ar.Run()
	return nil
}
