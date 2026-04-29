package gui

import (
	guisettings "github.com/Den4ik-droid/weather-widget/internal/domain/gui_settings"
	"github.com/Den4ik-droid/weather-widget/pkg/config"
)

type Logger interface {
	Info(string)
	Debug(string)
	Error(string, error)
}

type WeatherInfo interface {
	GetTemperature(float64, float64) (float32, error)
}

type AppRunnerProvider interface {
	GetAppRunner() guisettings.AppRunner
}

type WindowProvider interface {
	CreateWindow(name string, size guisettings.WindowSize) (guisettings.Window, error)
	GetTextWidget(text string) guisettings.TextWidget
}

type guiApp struct {
	l          Logger
	wp         WindowProvider
	wi         WeatherInfo
	conf       config.Config
	appRunner  guisettings.AppRunner
	window     guisettings.Window
	tempWidget guisettings.TextWidget
}

func New(l Logger, wp WindowProvider, wi WeatherInfo, conf config.Config) *guiApp {
	return &guiApp{
		l:    l,
		wp:   wp,
		wi:   wi,
		conf: conf,
	}
}

func (g *guiApp) Run() error {
	g.l.Info("Запуск GUI приложения")

	// Создаём окно
	windowSize := guisettings.NewWS(400, 200)
	window, err := g.wp.CreateWindow("Информер погоды", windowSize)
	if err != nil {
		g.l.Error("Ошибка создания окна", err)
		return err
	}
	g.window = window

	// Создаём виджет для температуры
	tempWidget := g.wp.GetTextWidget("Загрузка...")
	g.tempWidget = tempWidget

	// Устанавливаем виджет в окно
	err = g.window.SetTemperatureWidget(tempWidget)
	if err != nil {
		g.l.Error("Ошибка установки виджета", err)
		return err
	}

	// Получаем температуру
	temp, err := g.wi.GetTemperature(g.conf.L.Lat, g.conf.L.Long)
	if err != nil {
		g.l.Error("Ошибка получения температуры", err)
		temp = 0
	}
	g.window.UpdateTemperature(temp)

	// Рендерим окно
	g.window.Render()

	// Запускаем приложение
	g.wp.GetAppRunner().Run()

	return nil
}