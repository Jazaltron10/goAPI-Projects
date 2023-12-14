package cache

import (
	"fmt"
	"sync"

	"github.com/PunitNaran/weather_app/configs"
	"github.com/sirupsen/logrus"
)

// MemoryCache is an in-memory cache.
type MemoryCache struct {
	cache map[string][]configs.ForecastPeriod
	mu    sync.RWMutex
	l     *logrus.Logger
}

func NewMemoryCache(l *logrus.Logger) *MemoryCache {
	return &MemoryCache{
		cache: make(map[string][]configs.ForecastPeriod),
		l:     l,
	}
}

func (c *MemoryCache) Get(key string) ([]configs.ForecastPeriod, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, ok := c.cache[key]
	if !ok {
		err := fmt.Errorf("key not found: %s", key)
		c.l.Error(err)
		return nil, err
	}
	return data, nil
}

func (c *MemoryCache) Set(key string, data []configs.ForecastPeriod) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = data
	return nil
}
