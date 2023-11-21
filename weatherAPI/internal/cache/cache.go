package cache

import (
	"github.com/jazaltron10/goAPI/weatherAPI/configs"
)

// Cache is the interface for different types of caches.
type Cache interface {
	Get(key string) ([]configs.ForecastPeriod, error)
	Set(key string, forecastData []configs.ForecastPeriod) error
}
//create an internal cache database 