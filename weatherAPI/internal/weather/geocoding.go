package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetCoordinates retrieves the coordinates for a given city using the geocoding service.
func GetCoordinates(city string) (Coordinates, error) {
	// Build the URL for the geocoding service.
	geocodingURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s,usa&format=json", city)

	// in server.go -> c := http.Client{...}

	// Make an HTTP GET request to the geocoding service.
	resp, err := http.Get(geocodingURL)
	if err != nil {
		return Coordinates{}, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Coordinates{}, err
	}

	// Unmarshal the JSON response to get the coordinates.
	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return Coordinates{}, err
	}

	// Extract the coordinates from the response.
	lat, _ := result[0]["lat"].(float64)
	lon, _ := result[0]["lon"].(float64)

	return Coordinates{Latitude: lat, Longitude: lon}, nil
}
