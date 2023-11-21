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

func (c *CityCountryEndpoint) GetOpenStreetMapLink() (*url.URL, error) {}

func (c *CityCountryEndpoint) formatIsValid() bool {}

func (f *ForecastCoordinates) GetForecastCoordinatesLink() (*url.URL, error) {}


func getURL(link string) (*url.URL, error) {}

