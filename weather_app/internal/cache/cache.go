package cache

import (
	"github.com/PunitNaran/weather_app/configs"
	// Import Logrus for structured logging
)

// Cache is the interface for different types of caches.
type Cache interface {
	Get(key string) ([]configs.ForecastPeriod, error)
	Set(key string, forecastData []configs.ForecastPeriod) error
}
