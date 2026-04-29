package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Den4ik-droid/weather-widget/pkg/logger"
)

type cliApp struct {
	log logger.Logger
}

func New(log logger.Logger) *cliApp {
	return &cliApp{log: log}
}

func (c *cliApp) Run() error {
	c.log.Info("Запуск получения погоды")

	type Current struct {
		Temp float32 `json:"temperature_2m"`
	}
	type Response struct {
		Curr Current `json:"current"`
	}

	var response Response

	params := fmt.Sprintf(
		"latitude=%f&longitude=%f&current=temperature_2m",
		53.6688,
		23.8223,
	)
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)

	c.log.Debug("Формируем запрос: " + url)

	resp, err := http.Get(url)
	if err != nil {
		c.log.Error("Ошибка HTTP-запроса")
		customErr := errors.New("can`t get weather data from openmeteo")
		return errors.Join(customErr, err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.log.Error("Ошибка закрытия тела ответа")
			fmt.Printf("can`t close body err - %s\n", err.Error())
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.log.Error("Ошибка чтения данных")
		customErr := errors.New("can`t read data from response")
		return errors.Join(customErr, err)
	}

	if err := json.Unmarshal(data, &response); err != nil {
		c.log.Error("Ошибка парсинга JSON")
		customErr := errors.New("can`t unmarshal data from response")
		return errors.Join(customErr, err)
	}

	c.log.Info("Погода успешно получена")
	fmt.Printf(
		"Температура воздуха - %.2f градусов цельсия\n",
		response.Curr.Temp,
	)
	return nil
}