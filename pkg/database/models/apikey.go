package models

import (
	"errors"
	"time"

	"github.com/vatusa/api2/pkg/database"
	"gorm.io/gorm/clause"
)

type APIKey struct {
	ID          uint64    `xml:"id" json:"id" yaml:"id"`
	FacilityID  uint64    `xml:"-" json:"-" yaml:"-" gorm:"index:facility"`
	Facility    Facility  `xml:"facility" json:"facility" yaml:"facility"`
	Name        string    `xml:"name" json:"name" yaml:"name"`
	Token       string    `xml:"-" json:"-" yaml:"-" gorm:"type:varchar(48);index:token"`
	NeverExpire bool      `xml:"never_expire" json:"never_expire" yaml:"never_expire"`
	ExpiresAt   time.Time `xml:"expires_at" json:"expires_at" yaml:"expires_at"`
	CreatedAt   time.Time `xml:"created_at" json:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `xml:"updated_at" json:"updated_at" yaml:"updated_at"`
}

func FindFacilityByAPIKey(token string) (*Facility, error) {
	var apiKey APIKey
	if err := database.DB.Preload(clause.Associations).Where("token = ?", token).First(&apiKey).Error; err != nil {
		return nil, err
	}

	if !apiKey.NeverExpire && apiKey.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("api key expired")
	}

	return &apiKey.Facility, nil
}
