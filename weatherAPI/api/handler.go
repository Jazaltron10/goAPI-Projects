// api/handler.go
package api

import (
	"net/http"
	"strings"

	"github.com/jazaltron10/goAPI/weatherAPI/internal/weather"
	"github.com/labstack/echo/v4"
)


type Handler struct {
	c     *http.Client
	store cache.Cache
	l     *logrus.Logger
}


func (h *Handler) CreateClient(store cache.Cache, l *logrus.Logger) {
	h.c = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	h.l = l
	h.store = store
}


// internal/handler/handler.go
// -> type Handler struct {... Objects -> client, store, logger}
// instead of func GetWeatherForecastHandler -> func (h *Handler) GetWeatherForecastHandler

// GetWeatherForecastHandler is the handler for the weather forecast endpoint.
func GetWeatherForecastHandler(c echo.Context) error {
	// Parse the list of cities from the query parameter.
	cityList := c.QueryParam("city")
	cities := strings.Split(cityList, ",")

	// Initialize the forecast slice.
	var forecasts []weather.Forecast

	// Process each city.
	for _, city := range cities {
		// Get the coordinates for the city using the geocoding service.
		coordinates, err := weather.GetCoordinates(city)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Get the forecast for the coordinates using the weather API.
		forecast, err := weather.GetWeatherForecastForCoordinates(coordinates)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// Append the forecast to the slice.
		forecasts = append(forecasts, forecast)
	}

	// Create the response JSON.
	responseJSON := map[string][]weather.Forecast{"forecast": forecasts}

	// Return the response JSON.
	return c.JSON(http.StatusOK, responseJSON)
}
