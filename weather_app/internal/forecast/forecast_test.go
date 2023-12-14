package forecast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOpenStreenMapLink(t *testing.T) {

	// Call the CreateOpenStreenMapLink function
	link, err := CreateOpenStreetMapLink("london")

	assert.NoError(t, err)
	assert.Equal(t, "https://nominatim.openstreetmap.org/search?q=london,usa&format=json", link)
}
