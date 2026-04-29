package fyne

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	guisettings "github.com/Den4ik-droid/weather-widget/internal/domain/gui_settings"
)

type window struct {
	w       fyne.Window
	tempWid guisettings.TextWidget
}

func NewW(w fyne.Window) *window {
	return &window{w: w}
}

func (win *window) Resize(ws guisettings.WindowSize) error {
	if ws.IsFull() {
		win.w.Resize(fyne.NewSize(float32(ws.Width()), float32(ws.Height())))
	}
	win.w.Resize(fyne.NewSize(float32(ws.Width()), float32(ws.Height())))
	return nil
}

func (win *window) UpdateTemperature(t float32) error {
	if win.tempWid == nil {
		return fmt.Errorf("temperature widget is not set")
	}
	win.tempWid.SetText(fmt.Sprintf("Температура: %.1f°C", t))
	return nil
}

func (win *window) SetTemperatureWidget(tw guisettings.TextWidget) error {
	win.tempWid = tw
	return nil
}

func (win *window) Render() error {
	win.w.CenterOnScreen()
	return nil
}