package seed

import "github.com/vatusa/api2/pkg/database/models"

func BuildFacilities(values []map[string]interface{}) []models.Facility {
	facilities := make([]models.Facility, len(values))

	for i, value := range values {
		facilities[i] = BuildFacilityFromSeed(value)
	}

	return facilities
}

func BuildFacilityFromSeed(value map[string]interface{}) models.Facility {
	facility := models.Facility{}
	facility.IATA = value["iata"].(string)
	facility.Name = value["name"].(string)
	facility.Url = value["url"].(string)
	facility.Active = value["active"].(bool)

	return facility
}
