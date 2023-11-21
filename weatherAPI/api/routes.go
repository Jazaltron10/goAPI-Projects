// api/routes.go
package api

import (
	"github.com/jazaltron10/goAPI/weatherAPI/internal/weather"
	"github.com/labstack/echo/v4"
)

// InitializeRoutes sets up the routes for the application.
func InitializeRoutes(e *echo.Echo) {
	// Create a group for API routes
	apiGroup := e.Group("/api")

	// Define the endpoint for the weather forecast.
	apiGroup.GET("/weather", weather.GetWeatherForecastHandler)
}
