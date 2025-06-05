package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	apiKey = "814334aca4b1e933fa41744a83ba32f3"
	lat    = "-2.9635687"
	lon    = "104.7953744"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}

func getWeather(c *fiber.Ctx) error {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&units=metric&appid=%s", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal ambil data"})
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var weather WeatherResponse
	json.Unmarshal(body, &weather)

	return c.JSON(fiber.Map{
		"city":        weather.Name,
		"temp_c":      weather.Main.Temp,
		"description": weather.Weather[0].Description,
	})
}

func main() {
	app := fiber.New()

	app.Get("/weather", getWeather)

	log.Fatal(app.Listen(":3000"))
}

