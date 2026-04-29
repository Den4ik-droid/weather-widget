package main

import (
	"os"

	"github.com/Den4ik-droid/weather-widget/internal/adapters/weather"
	"github.com/Den4ik-droid/weather-widget/internal/pkg/app/cli"
	"github.com/Den4ik-droid/weather-widget/pkg/logger"
)

func main() {
	l := logger.New()
	wi := weather.New(l)
	app := cli.New(l, wi)

	err := app.Run()
	if err != nil {
		l.Error("Some error", err)
		os.Exit(1)
	}
	os.Exit(0)
}