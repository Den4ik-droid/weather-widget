package main

import (
	"os"

	"github.com/Den4ik-droid/weather-widget/internal/adapters/weather"
	"github.com/Den4ik-droid/weather-widget/internal/pkg/app/cli"
	"github.com/Den4ik-droid/weather-widget/internal/pkg/flags"
	"github.com/Den4ik-droid/weather-widget/pkg/config"
	"github.com/Den4ik-droid/weather-widget/pkg/logger"
)

func main() {
	// Парсим аргументы командной строки
	arguments := flags.Parse()

	// Открываем файл конфигурации
	r, err := os.Open(arguments.Path)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	// Парсим конфигурацию
	cfg, err := config.Parse(r)
	if err != nil {
		panic(err)
	}

	// Инициализируем логгер
	l := logger.New()

	// Получаем провайдера погоды
	wi := getProvider(cfg, l)

	// Создаём и запускаем приложение
	app := cli.New(l, wi, cfg)
	err = app.Run()
	if err != nil {
		l.Error("Some error", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func getProvider(cfg config.Config, l cli.Logger) cli.WeatherInfo {
	var wi cli.WeatherInfo
	switch cfg.P.Type {
	case "open-meteo":
		wi = weather.New(l)
	default:
		wi = weather.New(l)
	}
	return wi
}