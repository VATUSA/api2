package models

import (
	"time"

	"github.com/vatusa/api2/pkg/database"
	"gorm.io/gorm/clause"
)

type APIKey struct {
	ID          uint64     `xml:"id" json:"id"`
	FacilityID  uint64     `xml:"-" json:"-" gorm:"index:facility"`
	Facility    Facility   `xml:"facility" json:"facility"`
	Name        string     `xml:"name" json:"name"`
	Token       string     `xml:"-" json:"-" gorm:"type:varchar(48);index:token"`
	NeverExpire bool       `xml:"never_expire" json:"never_expire"`
	ExpiresAt   *time.Time `xml:"expires_at" json:"expires_at"`
	CreatedAt   time.Time  `xml:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `xml:"updated_at" json:"updated_at"`
}

func FindFacilityByAPIKey(token string) (*Facility, error) {
	var apiKey APIKey
	if err := database.DB.Preload(clause.Associations).Where("token = ?", token).First(&apiKey).Error; err != nil {
		return nil, err
	}
	return &apiKey.Facility, nil
}

func IsAPIKeyValid()
