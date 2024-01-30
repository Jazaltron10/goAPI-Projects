package cache

import(
	"github.com/jazaltron10/forecastApp/configs"
	
)


//caches is the interface for different types of caches 
type cache interface{
	Get(key string) ([]configs.ForecastPeriod, error)
	Set(key string, forecastData []configs.ForecastPeriod) (error)

}