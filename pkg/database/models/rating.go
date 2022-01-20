package models

type Rating struct {
	ID    int    `json:"id" yaml:"id" xml:"id" gorm:"primaryKey;autoIncrement:false"`
	Short string `json:"short" yaml:"short" xml:"short" gorm:"type:varchar(3);uniqueIndex"`
	Long  string `json:"long" yaml:"long" xml:"long" gorm:"type:varchar(24)"`
}
