package seed

import (
	"github.com/vatusa/api2/pkg/database/models"
	"github.com/vatusa/api2/pkg/utils"
)

func BuildRatings(values []map[string]interface{}) []models.Rating {
	ratings := make([]models.Rating, len(values))

	for i, value := range values {
		ratings[i] = BuildRatingFromSeed(value)
	}

	return ratings
}

func BuildRatingFromSeed(value map[string]interface{}) models.Rating {
	rating := models.Rating{}
	rating.ID = utils.GetInt(value["id"])
	rating.Long = value["long"].(string)
	rating.Short = value["short"].(string)

	return rating
}
