package main

import (
	"log"

	"github.com/Den4ik-droid/weather-widget/internal/pkg/app/cli"
	"github.com/Den4ik-droid/weather-widget/pkg/logger"
)	

func main() {
	logger := logger.NewSimpleLogger()
	app := cli.New(logger)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}