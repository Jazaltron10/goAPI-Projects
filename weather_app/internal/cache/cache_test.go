package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/PunitNaran/weather_app/configs"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -destination=./mocks/mock_cache.go -package=mocks github.com/PunitNaran/weather_app/internal/cache Cache

// The following mock will be used by other packages

func TestGetFileCache(t *testing.T) {
	folder := createTempFolder(t)
	defer os.RemoveAll(folder)

	f := &FileCache{
		filePath: folder,
		l:        logrus.New(),
	}
	// Write to file - "The setup" this can be improved and refactored
	b, err := json.Marshal([]configs.ForecastPeriod{
		{
			DetailedForecast: "test1",
			StartTime:        time.Now().Add(-time.Minute),
			EndTime:          time.Now(),
		},
	})
	assert.NoError(t, err)
	err = os.WriteFile(folder+"aCity", b, 0644)
	assert.NoError(t, err)

	tests := map[string]struct {
		key            string
		forcastDetails string
		errResp        error
	}{
		"success - got data from file cache": {
			key:            "aCity",
			forcastDetails: "test1",
			errResp:        nil,
		},
		"fail - unable to get from file cache": {
			key:            "aCity2",
			forcastDetails: "test1",
			errResp:        fmt.Errorf(""), // this can be something else, or a bool
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			forecastInfo, err := f.Get(tc.key)
			if tc.errResp != nil {
				assert.True(t, os.IsNotExist(err))
				// assert.ErrorIs(t, err, tc.errResp) // if we needed to check an error
				return
			}

			assert.NoError(t, err)
			// for now due to the time of the exersize we will just check the first value if it has it
			assert.Equal(t, tc.forcastDetails, forecastInfo[0].DetailedForecast)
		})
	}
}

func TestSetFileCache(t *testing.T) {
	folder := createTempFolder(t)
	defer os.RemoveAll(folder)
	log := logrus.New()
	tests := map[string]struct {
		f              *FileCache
		key            string
		forcastDetails []configs.ForecastPeriod
		errResp        error
	}{
		"success - set forcast periods in a file cache": {
			f: &FileCache{
				filePath: folder,
				l:        log,
			},
			key: "aCity",
			forcastDetails: []configs.ForecastPeriod{
				{
					DetailedForecast: "test1",
					StartTime:        time.Now().Add(-time.Minute),
					EndTime:          time.Now(),
				},
			},
			errResp: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := tc.f.Set(tc.key, tc.forcastDetails)
			assert.NoError(t, err)
		})
	}
}

func TestGetMemoryCache(t *testing.T) {
	mc := &MemoryCache{
		cache: make(map[string][]configs.ForecastPeriod),
		l:     logrus.New(),
	}
	mc.cache["aCity"] = []configs.ForecastPeriod{
		{
			DetailedForecast: "test1",
			StartTime:        time.Now().Add(-time.Minute),
			EndTime:          time.Now(),
		},
	}
	tests := map[string]struct {
		key            string
		forcastDetails string
		errResp        error
	}{
		"success - got data from memory cache": {
			key:            "aCity",
			forcastDetails: "test1",
			errResp:        nil,
		},
		"fail - unable to get from memory cache": {
			key:            "aCity2",
			forcastDetails: "test1",
			errResp:        fmt.Errorf("key not found"), // this can be something else, or a bool
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			forecastInfo, err := mc.Get(tc.key)
			if tc.errResp != nil {
				assert.ErrorContains(t, err, tc.errResp.Error())
				return
			}

			assert.NoError(t, err)
			// for now due to the time of the exersize we will just check the first value if it has it
			assert.Equal(t, tc.forcastDetails, forecastInfo[0].DetailedForecast)
		})
	}
}

func TestSetMemoryCache(t *testing.T) {

	log := logrus.New()
	tests := map[string]struct {
		mc             *MemoryCache
		key            string
		forcastDetails []configs.ForecastPeriod
		errResp        error
	}{
		"success - set forcast periods in a memory cache": {
			mc: &MemoryCache{
				cache: make(map[string][]configs.ForecastPeriod),
				l:     log,
			},
			key: "aCity",
			forcastDetails: []configs.ForecastPeriod{
				{
					DetailedForecast: "test1",
					StartTime:        time.Now().Add(-time.Minute),
					EndTime:          time.Now(),
				},
			},
			errResp: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := tc.mc.Set(tc.key, tc.forcastDetails)
			assert.NoError(t, err)
		})
	}
}

func createTempFolder(t *testing.T) string {
	t.Helper()
	folder, err := os.MkdirTemp("", "cache")
	assert.NoError(t, err)
	return folder
}
