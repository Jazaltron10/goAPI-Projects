package configs

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus" // Import Logrus for structured logging
)

const (
	openStreetMapWebLink      = "https://nominatim.openstreetmap.org"
	forecastPairOfCoordinates = "https://api.weather.gov/points"
)

var openStreetFormats = []string{"xml", "geojson", "geocodejson", "json", "jsonv2"}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.InfoLevel)
}

type CityCountryEndpoint struct {
	City    string
	Country string
	Format  string
}

type ForecastCoordinates struct {
	Longitude string `json:"lon"`
	Latitude  string `json:"lat"`
}

type PropertiesInfo struct {
	Properties PropertyInfo `json:"properties"`
}

type PropertyInfo struct {
	ForecastURL string `json:"forecast"`
}

type ForecastPeriod struct {
	DetailedForecast string    `json:"detailedForecast"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
}

type PropertiesForecastInfo struct {
	Periods ForecastPeriodsInfo `json:"properties"`
}

type ForecastPeriodsInfo struct {
	Periods []ForecastPeriod `json:"periods"`
}

func (c *CityCountryEndpoint) GetOpenStreetMapLink() (*url.URL, error) {
	c.City = strings.ToLower(strings.TrimSpace(c.City))
	c.Country = strings.ToLower(strings.TrimSpace(c.Country))

	if !c.formatIsValid() {
		err := errors.New("the 'format' must be one of: xml, geojson, geocodejson, json, jsonv2")
		log.Error(err)
		return nil, err
	}

	link := fmt.Sprintf("%s/search?q=%s,%s&format=%s", openStreetMapWebLink, c.City, c.Country, c.Format)
	return getURL(link)
}

func (c *CityCountryEndpoint) formatIsValid() bool {
	c.Format = strings.ToLower(strings.TrimSpace(c.Format))
	for _, format := range openStreetFormats {
		if c.Format == format {
			return true
		}
	}
	log.Error("Invalid format:", c.Format)
	return false
}

func (f *ForecastCoordinates) GetForecastCoordinatesLink() (*url.URL, error) {
	f.Longitude = strings.ToLower(strings.TrimSpace(f.Longitude))
	f.Latitude = strings.ToLower(strings.TrimSpace(f.Latitude))
	link := fmt.Sprintf("%s/%s,%s", forecastPairOfCoordinates, f.Latitude, f.Longitude)
	return getURL(link)
}

func getURL(link string) (*url.URL, error) {
	u, err := url.Parse(link)
	if err != nil {
		log.Error("Error parsing URL:", err)
		return nil, err
	}
	return u, nil
}
