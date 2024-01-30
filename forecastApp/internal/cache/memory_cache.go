package cache


import(
	"fmt"
	"sync"

	"github.com/jazaltron10/forecastApp/configs"
	"github.com/sirupsen/logrus"

)

type MemoryCache struct{
	cache map[string][]configs.ForecastPeriod
	mu	sync.RWMutex
	l	*logrus.Logger
}


func NewMemoryCache(l *logrus.Logger) *MemoryCache{
	return &MemoryCache{
		cache: make(map[string][]configs.ForecastPeriod),
		l: l,
	}
}
