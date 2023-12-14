package forecast

import (
	"os"

	"github.com/PunitNaran/weather_app/configs"
	"github.com/sirupsen/logrus"
)

// Import Logrus for structured logging

const (
	country = "usa"
	format  = "json"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.InfoLevel)
}

func CreateOpenStreetMapLink(city string) (string, error) {
	c := configs.CityCountryEndpoint{
		City:    city,
		Country: country,
		Format:  format,
	}

	u, err := c.GetOpenStreetMapLink()
	if err != nil {
		log.Errorf("Error creating OpenStreetMap link: %v", err)
		return "", err
	}

	return u.String(), nil
}
