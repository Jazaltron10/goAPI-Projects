package cache

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/PunitNaran/weather_app/configs"
	"github.com/sirupsen/logrus"
)

// FileCache is a file-based cache.
type FileCache struct {
	filePath string
	mu       sync.RWMutex
	l        *logrus.Logger
}

func NewFileCache(l *logrus.Logger, filePath string) *FileCache {
	return &FileCache{
		filePath: filePath,
	}
}

func (c *FileCache) Get(key string) ([]configs.ForecastPeriod, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, err := c.loadFromFile(key)
	if err != nil {
		c.l.Errorf("Error reading from file: %v", err)
		return nil, err
	}

	return data, nil
}

func (c *FileCache) Set(key string, forecastData []configs.ForecastPeriod) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.saveToFile(key, forecastData)
}

func (c *FileCache) loadFromFile(key string) ([]configs.ForecastPeriod, error) {
	filePath := c.filePath + key
	data, err := os.ReadFile(filePath)
	if err != nil {
		c.l.Errorf("Error reading from file %s: %v", filePath, err)
		return nil, err
	}

	var forecastData []configs.ForecastPeriod
	if err := json.Unmarshal(data, &forecastData); err != nil {
		c.l.Errorf("Error unmarshaling JSON: %v", err)
		return nil, err
	}

	return forecastData, nil
}

func (c *FileCache) saveToFile(key string, forecastData []configs.ForecastPeriod) error {
	filePath := c.filePath + key
	b, err := json.Marshal(forecastData)
	if err != nil {
		c.l.Errorf("Error marshaling JSON: %v", err)
		return err
	}

	err = os.WriteFile(filePath, b, 0644)
	if err != nil {
		c.l.Errorf("Error writing to file %s: %v", filePath, err)
		return err
	}

	return nil
}
