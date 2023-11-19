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
	Description string    `json:"description"`
}

// Forecast represents the forecast for a city.
type Forecast struct {
	Name   string   `json:"name"`
	Detail []Period `json:"detail"`
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

	// Unmarshal the JSON response to get the forecast URL.
	var forecastURL string
	if err := json.Unmarshal(body, &forecastURL); err != nil {
		return Forecast{}, err
	}

	// Make another HTTP GET request to the forecast URL.
	resp, err = http.Get(forecastURL)
	if err != nil {
		return Forecast{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return Forecast{}, err
	}

	// Unmarshal the JSON response to get the forecast periods.
	var periods []Period
	if err := json.Unmarshal(body, &periods); err != nil {
		return Forecast{}, err
	}

	// Filter and format the forecast periods.
	var filteredPeriods []Period
	for _, p := range periods {
		if p.StartTime.After(time.Now()) && p.StartTime.Before(time.Now().Add(72*time.Hour)) {
			filteredPeriods = append(filteredPeriods, p)
		}
	}

	// Create the forecast object.
	forecast := Forecast{
		Name:   fmt.Sprintf("%s, %s", resp.Request.URL.Query().Get("city"), "USA"),
		Detail: filteredPeriods,
	}

	return forecast, nil
}
