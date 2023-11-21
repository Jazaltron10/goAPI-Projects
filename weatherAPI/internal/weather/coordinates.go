// internal/weather/coordinates.go
package weather

// Coordinates represents latitude and longitude of a location.
type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
