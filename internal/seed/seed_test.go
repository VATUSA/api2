package seed

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vatusa/api2/pkg/database/models"
)

func TestIsValidSeedType(t *testing.T) {
	assert.True(t, IsValidSeedType("yaml"))
	assert.True(t, IsValidSeedType("json"))
	assert.False(t, IsValidSeedType(""))
	assert.False(t, IsValidSeedType("foo"))
}

var ratingYAML = `
kind: rating
values:
- id: 1
  long: Testing
  short: T
`

var ratingJSON = `
{
	"kind": "rating",
	"values": [
		{
			"id": 1,
			"long": "Testing",
			"short": "T"
		}
	]
}`

var ratingSeed = Seed{
	Kind: "rating",
	Values: []map[string]interface{}{
		{
			"id":    1,
			"long":  "Testing",
			"short": "T",
		},
	},
}

var ratingModels = []models.Rating{
	{
		ID:    1,
		Long:  "Testing",
		Short: "T",
	},
}

var facilityYAML = `
kind: facility
values:
- iata: TST
  name: Testing
  url: http://www.testing.com
  active: true
`

var facilityJSON = `
{
	"kind": "facility",
	"values": [
		{
			"iata": "TST",
			"name": "Testing",
			"url": "http://www.testing.com",
			"active": true
		}
	]
}`

var facilitySeed = Seed{
	Kind: "facility",
	Values: []map[string]interface{}{
		{
			"iata":   "TST",
			"name":   "Testing",
			"url":    "http://www.testing.com",
			"active": true,
		},
	},
}

var facilityModels = []models.Facility{
	{
		IATA:   "TST",
		Name:   "Testing",
		Url:    "http://www.testing.com",
		Active: true,
	},
}

func TestRatingSeed(t *testing.T) {
	expected := models.Rating{
		ID:    1,
		Long:  "Testing",
		Short: "T",
	}

	assert.Equal(t, expected, BuildRatingFromSeed(ratingSeed.Values[0]))
}

func TestFacilitySeed(t *testing.T) {
	expected := models.Facility{
		IATA:   "TST",
		Name:   "Testing",
		Url:    "http://www.testing.com",
		Active: true,
	}

	assert.Equal(t, expected, BuildFacilityFromSeed(facilitySeed.Values[0]))
}

func TestBuildSeed(t *testing.T) {
	ret, err := BuildSeed("yaml", []byte(ratingYAML))
	assert.NoError(t, err)
	ratings := BuildRatings(ret.Values)
	assert.Equal(t, ratingModels, ratings)

	ret, err = BuildSeed("json", []byte(ratingJSON))
	assert.NoError(t, err)
	ratings = BuildRatings(ret.Values)
	assert.Equal(t, ratingModels, ratings)

	ret, err = BuildSeed("yaml", []byte(facilityYAML))
	assert.NoError(t, err)
	facilities := BuildFacilities(ret.Values)
	assert.Equal(t, facilityModels, facilities)

	ret, err = BuildSeed("json", []byte(facilityJSON))
	assert.NoError(t, err)
	facilities = BuildFacilities(ret.Values)
	assert.Equal(t, facilityModels, facilities)
}
