package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A simple example of a BDD approch - Other tests are TDD

func TestCityCountryEndpoint_GetOpenStreetMapLink_ValidFormat(t *testing.T) {
	// Given a CityCountryEndpoint with a valid format
	c := CityCountryEndpoint{
		City:    "new York",
		Country: "usa",
		Format:  "json",
	}

	// When we call the GetOpenStreetMapLink method
	link, err := c.GetOpenStreetMapLink()

	// Then Expect that there's no error and the link is valid
	assert.NoError(t, err)
	assert.NotNil(t, link)
	assert.Contains(t, link.String(), "q=new york,usa&format=json")
}

func TestCityCountryEndpoint_GetOpenStreetMapLink_InvalidFormat(t *testing.T) {
	// Given a CityCountryEndpoint with an invalid format
	c := CityCountryEndpoint{
		City:    "new York",
		Country: "usa",
		Format:  "invalid",
	}

	// When we call the GetOpenStreetMapLink method
	link, err := c.GetOpenStreetMapLink()

	// Then Expect that an error is returned
	assert.Error(t, err)
	assert.Nil(t, link)
}

func TestForecastCoordinates_GetForecastCoordinatesLink(t *testing.T) {
	// Given a ForecastCoordinates with valid coordinates
	f := ForecastCoordinates{
		Longitude: "10.123",
		Latitude:  "20.456",
	}

	// When we call the GetForecastCoordinatesLink method
	link, err := f.GetForecastCoordinatesLink()

	// Then Expect that there's no error and the link is valid
	assert.NoError(t, err)
	assert.NotNil(t, link)
	assert.Contains(t, link.String(), "/points/20.456,10.123")
}

func TestCityCountryEndpoint_FormatIsValid_ValidFormat(t *testing.T) {
	// Given a CityCountryEndpoint with a valid format
	c := CityCountryEndpoint{
		City:    "new york",
		Country: "usa",
		Format:  "json",
	}

	// When we call the FormatIsValid method
	isValid := c.formatIsValid()

	// Then expect that the format is valid
	assert.True(t, isValid)
}

func TestCityCountryEndpoint_FormatIsValid_InvalidFormat(t *testing.T) {
	// Given a CityCountryEndpoint with an invalid format
	c := CityCountryEndpoint{
		City:    "new york",
		Country: "usa",
		Format:  "invalid",
	}

	// When we the FormatIsValid method
	isValid := c.formatIsValid()

	// Then expect that the format is invalid
	assert.False(t, isValid)
}