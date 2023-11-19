package api

import (
    "github.com/labstack/echo/v4"
)

// StartServer initializes and starts the Echo server.
func StartServer() {
    e := echo.New()

    // Define the endpoint for the weather forecast.
    e.GET("/weather", GetWeatherForecast)

    // Start the server on port 1323.
    e.Logger.Fatal(e.Start(":1323"))
}
