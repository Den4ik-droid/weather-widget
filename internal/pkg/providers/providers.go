package providers

import (
	pogodaby "github.com/Den4ik-droid/weather-widget/internal/adapters/pogoda_by"
	"github.com/Den4ik-droid/weather-widget/internal/adapters/weather"
	"github.com/Den4ik-droid/weather-widget/internal/domain/models"
	"github.com/Den4ik-droid/weather-widget/pkg/config"
)

type Logger interface {
	Info(string)
	Debug(string)
	Error(string, error)
}

type WeatherInfo interface {
	GetTemperature(float64, float64) (models.TempInfo, error)
}

func GetProvider(cfg config.Config, l Logger) WeatherInfo {
	switch cfg.P.Type {
	case "open-meteo":
		return weather.New(l)
	case "pogoda":
		return pogodaby.New(l)
	default:
		return weather.New(l)
	}
}