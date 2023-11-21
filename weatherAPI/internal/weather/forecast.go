package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Period represents a time period in the weather forecast.
type Period struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"shortForecast"`
}

// Forecast represents the forecast for a city.
type Forecast struct {
	Name   string   `json:"name"`
	Detail []Period `json:"detail"`
}

// WeatherResponse represents the JSON structure of the weather API response.
type WeatherResponse struct {
	Properties struct {
		Forecast struct {
			Periods []Period `json:"periods"`
		} `json:"forecast"`
	} `json:"properties"`
}

// GetWeatherForecastForCoordinates retrieves the weather forecast for a given set of coordinates.
func GetWeatherForecastForCoordinates(coordinates Coordinates) (Forecast, error) {
	// Build the URL for the weather forecast service.
	weatherURL := fmt.Sprintf("https://api.weather.gov/points/%f,%f", coordinates.Latitude, coordinates.Longitude)

	// Make an HTTP GET request to the weather forecast service.
	resp, err := http.Get(weatherURL)
	if err != nil {
		return Forecast{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Forecast{}, err
	}

	// Unmarshal the JSON response to get the forecast.
	var weatherResponse WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return Forecast{}, err
	}

	// Filter and format the forecast periods.
	var periods []Period
	for _, p := range weatherResponse.Properties.Forecast.Periods {
		if p.StartTime.After(time.Now()) && p.StartTime.Before(time.Now().Add(72*time.Hour)) {
			periods = append(periods, p)
		}
	}

	// Create the forecast object.
	forecast := Forecast{
		Name:   fmt.Sprintf("%s, %s", resp.Request.URL.Query().Get("city"), "USA"),
		Detail: periods,
	}

	return forecast, nil
}
