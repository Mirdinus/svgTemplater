package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func loadForecastData() (*ForecastWeatherData, error) {
	url := fmt.Sprintf("%s/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", Config.Weather.BaseUrl, Config.Weather.ApiKey, Config.Weather.City)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData ForecastWeatherData
	err = json.Unmarshal(data, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func ProcessWeather(template string) string {
	forecastData, err := loadForecastData()
	if err != nil {
		log.Fatalf("Error fetching weather: %v", err)
	}

	template = strings.ReplaceAll(template, "{weather_temperature_min}", fmt.Sprintf("%.0f", forecastData.Forecast.Forecastday[0].Day.MintempC))
	template = strings.ReplaceAll(template, "{weather_temperature_max}", fmt.Sprintf("%.0f", forecastData.Forecast.Forecastday[0].Day.MaxtempC))
	template = strings.ReplaceAll(template, "{weather_temperature}", fmt.Sprintf("%.0f", forecastData.Current.TempC))
	template = strings.ReplaceAll(template, "{weather_temperature_condition}", forecastData.Current.Condition.Text)

	if strings.Contains(forecastData.Current.Condition.Text, "rain") {
		data, err := os.ReadFile("icons/rain.svg")
		if err != nil {
			log.Fatalf("Error reading rain icon: %v", err)
		}
		template = strings.ReplaceAll(template, "{weather_icon}", string(data))
	} else if strings.Contains(forecastData.Current.Condition.Text, "cloud") {
		data, err := os.ReadFile("icons/cloud.svg")
		if err != nil {
			log.Fatalf("Error reading rain icon: %v", err)
		}
		template = strings.ReplaceAll(template, "{weather_icon}", string(data))
	} else {
		data, err := os.ReadFile("icons/clear.svg")
		if err != nil {
			log.Fatalf("Error reading rain icon: %v", err)
		}
		template = strings.ReplaceAll(template, "{weather_icon}", string(data))
	}

	return template
}
