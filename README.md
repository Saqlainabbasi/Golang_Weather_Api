# Golang_Weather_Api

In this Golang serve project we created a endpoint which will return a weather data by city

## How to run

1. Clone the repo
2. make api config file in root directory which holds the openweathermap api key
3. Run `go run main.go`
4. Open "http://localhost:8080/weather/London" in your browser

## How to make api config file

{
"OpenWeatherMapApi":"your api key"
}

## Resource

- `https://openweathermap.org/api`

## Endpoints

1. "http://localhost:8080/hello"
2. "http://localhost:8080/weather/London"

## Just For Info ( live reload of server)

nodemon --exec go run main.go --signal SIGTERM
