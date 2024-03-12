package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// OpenWeather APIからのレスポンスを格納するための構造体
type WeatherResponse struct {
	Main struct {
		Pressure float64 `json:"pressure"` // 気圧
	} `json:"main"`
}

// OpenWeather APIキー（自分のAPIキーに置き換えてください）
const apiKey = "fed2a2a92489ca118efce26b45df2c64"

// 指定した都市の現在の気圧を基にブール値を返す関数
func CheckPressure(city string) bool {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching weather data: %s", err)
	}
	defer resp.Body.Close()


	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Fatalf("Error decoding weather data: %s", err)
	}
	

	currentPressure := weather.Main.Pressure
	averagePressure := 1013.0 // 平均気圧

	// 気圧が平均から6〜10ヘクトパスカル下がっているか判定
	if currentPressure >= averagePressure-10 && currentPressure <= averagePressure-6 {
		return true
	}

	return false
}

func Service() {
	city := "Tokyo"
	if CheckPressure(city) {
		fmt.Printf("The pressure in %s has dropped below the average threshold.\n", city)
	} else {
		fmt.Printf("The pressure in %s is within the normal range.\n", city)
	}

}