package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// declare a struct for api config file
type ApiKey struct {
	OpenWeatherMapApi string `json:"OpenWeatherMapApi"`
}

// declare a struct for weather data
// this struct there are nexted struct
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int32   `json:"pressure"`
		Humidity int32   `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   int16   `json:"deg"`
	} `json:"wind"`
}

// func to load api config
func loadApiConfig(filename string) (ApiKey, error) {
	//read data form the file
	// we use the ioutil pkg
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return ApiKey{}, err
	}
	var key ApiKey
	err = json.Unmarshal(bytes, &key)
	if err != nil {
		return ApiKey{}, err
	}
	return key, nil

}

// func
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Form Go!....\n"))
}

// func to hadle the get weather data....
func getWeather(w http.ResponseWriter, r *http.Request) {
	//get the city name form the url params....
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//get the data form api...
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiConfig.OpenWeatherMapApi + "&units=imperial")
	if err != nil {
		http.Error(w, "Error getting data form weather api", http.StatusInternalServerError)
		return
	}
	fmt.Println(resp.Body)
	// decode the resp into json data
	var d WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		http.Error(w, "Data format not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// convert the data back to json foramat
	json.NewEncoder(w).Encode(d)
}

func main() {
	// route to check server is up or not
	http.HandleFunc("/hello", hello)

	// route to handle request for weather data.....
	http.HandleFunc("/weather/", getWeather)

	//start the http server....
	fmt.Println("Server listing on port 8080")
	http.ListenAndServe(":8080", nil)
	//log which port the server is lisning on....
}
