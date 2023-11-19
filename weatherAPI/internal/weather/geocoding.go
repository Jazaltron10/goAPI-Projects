package weather

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Coordinates represents latitude and longitude of a location.
type Coordinates struct {
    Latitude  float64 `json:"lat"`
    Longitude float64 `json:"lon"`
}

// GetCoordinates retrieves the coordinates for a given city using the geocoding service.
func GetCoordinates(city string) (Coordinates, error) {
    // Build the URL for the geocoding service.
    geocodingURL := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s,usa&format=json", city)

    // Make an HTTP GET request to the geocoding service.
    resp, err := http.Get(geocodingURL)
    if err != nil {
        return Coordinates{}, err
    }
    defer resp.Body.Close()

    // Read the response body.
    var result []map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return Coordinates{}, err
    }

    // Extract the coordinates from the response.
    lat, _ := result[0]["lat"].(float64)
    lon, _ := result[0]["lon"].(float64)

    return Coordinates{Latitude: lat, Longitude: lon}, nil
}
