// weather.go
package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	// Import Logrus for structured logging
	"github.com/PunitNaran/weather_app/configs"
	"github.com/PunitNaran/weather_app/internal/forecast"
	"github.com/labstack/echo/v4"
)

type CityForecast struct {
	Name   string                   `json:"name"`
	Detail []configs.ForecastPeriod `json:"detail"`
}

type ForecastData struct {
	Forecast  []CityForecast `json:"forecast"`
	ErrorResp []string       `json:"errors,omitempty"`
}

func (h *Handler) Weather(c echo.Context) error {
	cities := c.QueryParam("city")
	cityList := strings.Split(cities, ",")
	var errs []error
	var cForcasts []CityForecast

	for _, city := range cityList {
		city = strings.ToLower(strings.TrimSpace(city))
		forecastInfo, err := h.store.Get(city)
		if err != nil {
			h.l.Debug(err)
		}

		if len(forecastInfo) != 0 {
			sT := forecastInfo[0].StartTime.UTC()
			cTime := time.Now()
			if isToday := sT.Year() == cTime.Year() && sT.YearDay() == cTime.YearDay(); !isToday {
				forecastInfo, err = h.getForecastData(city)
				if err != nil {
					errs = append(errs, err)
					continue
				}
			}
		} else {
			forecastInfo, err = h.getForecastData(city)
			if err != nil {
				errs = append(errs, err)
				continue
			}
		}

		cForcasts = append(cForcasts, CityForecast{
			Name:   city,
			Detail: forecastInfo,
		})
	}

	code := http.StatusOK
	forcastData := &ForecastData{
		Forecast: cForcasts,
	}

	if len(errs) > 0 {
		if len(errs) != len(cityList) {
			code = http.StatusPartialContent
		} else {
			code = http.StatusBadRequest
		}
		forcastData.ErrorResp = errorsToStrings(errs)
	}

	return c.JSON(code, forcastData)
}

func errorsToStrings(errs []error) []string {
	result := make([]string, len(errs))
	for i, err := range errs {
		if err != nil {
			result[i] = err.Error()
		}
	}
	return result
}

func (h *Handler) getForecastData(city string) ([]configs.ForecastPeriod, error) {
	link, err := forecast.CreateOpenStreetMapLink(city)
	if err != nil {
		return nil, err
	}
	u, err := h.getCoordinates(link)
	if err != nil {
		return nil, err
	}
	u, err = h.getForecast(u.String())
	if err != nil {
		return nil, err
	}

	forecastInfo, err := h.getForecastPeriodInfo(u.String())
	if err != nil {
		return nil, err
	}

	h.store.Set(city, forecastInfo)
	return forecastInfo, nil
}

func (h *Handler) getResponse(link string) ([]byte, error) {
	link = strings.ReplaceAll(link, " ", "%20")
	resp, err := h.c.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	return b, err
}

func (h *Handler) getCoordinates(link string) (*url.URL, error) {
	placeInfo := []configs.ForecastCoordinates{}

	b, err := h.getResponse(link)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &placeInfo); err != nil {
		h.l.Errorf("Error unmarshaling JSON: %v", err)
		return nil, err
	}

	if len(placeInfo) == 0 {
		return nil, errors.New("unable to get coordinates for a given city")
	}

	u, err := placeInfo[0].GetForecastCoordinatesLink()
	return u, err
}

func (h *Handler) getForecast(link string) (*url.URL, error) {
	b, err := h.getResponse(link)
	if err != nil {
		return nil, err
	}

	var forecastInfo configs.PropertiesInfo
	if err := json.Unmarshal([]byte(b), &forecastInfo); err != nil {
		h.l.Errorf("Error unmarshaling JSON: %v", err)
		return nil, err
	}

	u, err := url.Parse(forecastInfo.Properties.ForecastURL)
	return u, err
}

func (h *Handler) getForecastPeriodInfo(link string) ([]configs.ForecastPeriod, error) {
	var forecastPeriodsInfo configs.PropertiesForecastInfo

	b, err := h.getResponse(link)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &forecastPeriodsInfo); err != nil {
		h.l.Errorf("Error unmarshaling JSON: %v", err)
		return nil, err
	}

	now := time.Now()
	forecastInfo := make([]configs.ForecastPeriod, 0)

	if len(forecastPeriodsInfo.Periods.Periods) > 0 {
		for _, period := range forecastPeriodsInfo.Periods.Periods {
			if period.EndTime.After(now) && period.StartTime.Before(now.Add(72*time.Hour)) {
				forecastInfo = append(forecastInfo, period)
				continue
			}
			if len(forecastInfo) != 0 {
				break
			}
		}
		return forecastInfo, nil
	}

	h.l.Info("No forecast data available.")
	return nil, errors.New("")
}
