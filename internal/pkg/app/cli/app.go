package cli

import (
	"fmt"

	"github.com/Den4ik-droid/weather-widget/internal/domain/models"
	"github.com/Den4ik-droid/weather-widget/pkg/config"
)

type Logger interface {
	Info(string)
	Debug(string)
	Error(string, error)
}

type WeatherInfo interface {
	GetTemperature(float64, float64) (models.TempInfo, error) // Добавлен error
}

type cliApp struct {
	l    Logger
	wi   WeatherInfo
	conf config.Config
}

func New(l Logger, wi WeatherInfo, c config.Config) *cliApp {
	return &cliApp{
		l:    l,
		wi:   wi,
		conf: c,
	}
}

func (c *cliApp) Run() error {
	tempInfo, err := c.wi.GetTemperature(c.conf.L.Lat, c.conf.L.Long) // Исправлено: Long, не Lat
	if err != nil {
		c.l.Error("can't get temp info", err)
		return err
	}
	fmt.Printf(
		"Температура воздуха - %.2f градусов цельсия\n",
		tempInfo.Temp,
	)
	return nil
}