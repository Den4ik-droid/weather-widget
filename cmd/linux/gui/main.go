package main

import (
	"os"

	"github.com/Den4ik-droid/weather-widget/internal/pkg/app/gui"
	"github.com/Den4ik-droid/weather-widget/internal/pkg/config"
	"github.com/Den4ik-droid/weather-widget/internal/pkg/flags"
	"github.com/Den4ik-droid/weather-widget/internal/pkg/gui/fyne"
	"github.com/Den4ik-droid/weather-widget/internal/pkg/providers"
	"github.com/Den4ik-droid/weather-widget/pkg/logger"
)

func main() {
	arguments := flags.Parse()

	r, err := os.Open(arguments.Path)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	cfg, err := config.Parse(r)
	if err != nil {
		panic(err)
	}

	l := logger.New()
	provider := providers.GetProvider(cfg, l)
	p := fyne.NewP()
	g := gui.New(l, p, provider, cfg)

	err = g.Run()
	if err != nil {
		panic(err)
	}
}