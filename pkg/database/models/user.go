package models

type User struct {
	ID             uint64   `json:"id" yaml:"id" xml:"id"`
	CID            string   `json:"cid" yaml:"cid" xml:"cid" gorm:"type:varchar(24);unique_index"`
	FirstName      string   `json:"first_name" yaml:"first_name" xml:"first_name" gorm:"type:varchar(255)"`
	LastName       string   `json:"last_name" yaml:"last_name" xml:"last_name" gorm:"type:varchar(255)"`
	Email          string   `json:"email" yaml:"email" xml:"email" gorm:"type:varchar(255);unique_index"`
	HomeFacilityID string   `json:"-" yaml:"-" xml:"-" gorm:"type:varchar(3)"`
	HomeFacility   Facility `json:"home_facility" yaml:"home_facility" xml:"home_facility"`
	RatingID       int      `json:"-" yaml:"-" xml:"-"`
	Rating         Rating   `json:"rating" yaml:"rating" xml:"rating"`
}
