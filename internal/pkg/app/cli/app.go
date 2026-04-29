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
	GetTemperature(float64, float64) models.TempInfo
}

type cliApp struct {
	l  Logger
	wi WeatherInfo
	cfg config.Config
}

func New(l Logger, wi WeatherInfo, cfg config.Config) *cliApp {
	return &cliApp{
		l:   l,
		wi:  wi,
		cfg: cfg,
	}
}

func (c *cliApp) Run() error {
	c.l.Info("Запуск приложения")
	c.l.Debug(fmt.Sprintf("Используется провайдер: %s", c.cfg.P.Type))
	c.l.Debug(fmt.Sprintf("Координаты: lat=%.4f, long=%.4f", c.cfg.L.Lat, c.cfg.L.Long))

	fmt.Printf(
		"Температура воздуха - %.2f градусов цельсия\n",
		c.wi.GetTemperature(c.cfg.L.Lat, c.cfg.L.Long).Temp,
	)
	return nil
}