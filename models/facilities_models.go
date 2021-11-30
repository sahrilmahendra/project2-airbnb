package models

import "gorm.io/gorm"

type Facility struct {
	gorm.Model
	Name_Facility       string `json:"name_facility" form:"name_facility"`
	Homestay_Facilities []Homestay_Facility
}

type GetFacility struct {
	ID            uint
	Name_Facility string
}
