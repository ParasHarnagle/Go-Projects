package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey,omitempty"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := query(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			//json.NewEncoder(w).Encode(data)
			renderWeather(w, data)
		})

	http.ListenAndServe(":9000", nil)
}
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world\n"))
}
func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename) // Corrected this line
	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig("apiConfig")
	if err != nil {
		return weatherData{}, err

	}
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()
	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}

func renderWeather(w http.ResponseWriter, data weatherData) {
	html := `
    <html>
        <head>
            <title>Weather Report</title>
            <style>
                body { font-family: Arial, sans-serif; text-align: center; }
                .weather { margin-top: 20px; }
            </style>
        </head>
        <body>
            <h1>Weather Report</h1>
            <div class="weather">
                <h2>%s</h2>
                <p>Temperature: %.2f Kelvin</p>
            </div>
        </body>
    </html>`
	fmt.Fprintf(w, html, data.Name, data.Main.Kelvin)
}
