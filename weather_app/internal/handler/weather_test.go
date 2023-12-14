package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PunitNaran/weather_app/configs"
	"github.com/PunitNaran/weather_app/internal/cache/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func TestWeatherHandler(t *testing.T) {
	// Given a new Echo instance to simulate an HTTP server
	e := echo.New()

	tests := map[string]struct {
		cityGetCalls     map[string]int
		citySetCalls     map[string]int
		expectedErr      error
		httpResponseCode int
		invalidCities    []string
	}{
		"success": {
			cityGetCalls: map[string]int{
				"newyork":    1,
				"losangeles": 1,
			},
			citySetCalls: map[string]int{
				"newyork":    1,
				"losangeles": 1,
			},
			expectedErr:      nil,
			httpResponseCode: 200,
			invalidCities:    []string{},
		},
		// if would be nice to also include more expected failures
		"partial failure": {
			cityGetCalls: map[string]int{
				"choclates!": 1,
				"losangeles": 1,
			},
			citySetCalls: map[string]int{
				"losangeles": 1,
			},
			expectedErr:      fmt.Errorf("unable to get coordinates for a given city"),
			httpResponseCode: 206,
			invalidCities:    []string{"choclates!"},
		},
		"absolute failure": {
			cityGetCalls: map[string]int{
				"choclates!": 1,
				"mmmmmmm":    1,
			},
			citySetCalls:     map[string]int{},
			expectedErr:      fmt.Errorf("unable to get coordinates for a given city"),
			httpResponseCode: 400,
			invalidCities:    []string{"choclates!", "mmmmmmm"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			// Given a mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			cMock := mocks.NewMockCache(mockCtrl)

			cities := ""
			for city, calls := range tc.cityGetCalls {
				cities += city + ","
				// Monitor the calls and expected result
				cMock.EXPECT().Get(city).Return([]configs.ForecastPeriod{}, nil).Times(calls)
			}

			// Remove the trailing comma
			if len(cities) > 0 {
				cities = cities[:len(cities)-1]
			}

			for city, calls := range tc.citySetCalls {
				// Monitor the store in a cache and number of calls
				cMock.EXPECT().Set(city, gomock.Any()).Return(nil).Times(calls)
			}

			// Then create a request with city query parameters
			path := fmt.Sprintf("/weather?city=%s", cities)
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()

			// Bind the request to an Echo context
			c := e.NewContext(req, rec)

			// And create a new handler
			h := &Handler{
				c:     &http.Client{}, // I would create a mock for this. But for constrained time I'll just use the http client
				store: cMock,
			}

			// And call the Weather handler function
			err := h.Weather(c)

			// Then check the expected response
			assert.NoError(t, err)
			assert.Equal(t, tc.httpResponseCode, rec.Code)

			forecastData := ForecastData{}

			err = json.Unmarshal(rec.Body.Bytes(), &forecastData)
			assert.NoError(t, err)

			if tc.expectedErr != nil {
				// we should expect that there is a error response
				assert.NotEqual(t, 0, len(forecastData.ErrorResp))
				for _, err := range forecastData.ErrorResp {
					assert.Equal(t, err, tc.expectedErr.Error())
				}
			}
			// And check if the cities requested are correct
			// If i had mocked the http.client I could also check the details of the response even further
			for _, cityForcast := range forecastData.Forecast {
				_, ok := tc.cityGetCalls[cityForcast.Name]
				if stringExists(t, cityForcast.Name, tc.invalidCities) {
					// we dont need to assert the fail since the invalid city has not been stored in the cache
					continue
				}
				if !ok {
					assert.Failf(t, "City not found: %s", cityForcast.Name)
				}
			}
		})
	}

}

func stringExists(t *testing.T, s string, slice []string) bool {
	t.Helper()
	for _, element := range slice {
		if element == s {
			return true
		}
	}
	return false
}
