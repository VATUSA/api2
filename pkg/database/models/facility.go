package models

type Facility struct {
	ID     uint64 `json:"-" yaml:"-" xml:"-"`
	IATA   string `json:"iata" yaml:"iata" xml:"iata" gorm:"type:varchar(3);unique_index"`
	Name   string `json:"name" yaml:"name" xml:"name" gorm:"type:varchar(255)"`
	Url    string `json:"url" yaml:"url" xml:"url" gorm:"type:varchar(255)"`
	Active bool   `json:"active" yaml:"active" xml:"active" gorm:"type:boolean"`
}
