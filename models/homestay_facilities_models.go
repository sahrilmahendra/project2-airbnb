package models

import "gorm.io/gorm"

// struktur homestay facilities
type Homestay_Facility struct {
	gorm.Model
	HomestayID uint `json:"homestay_id" form:"homestay_id"`
	FacilityID uint `json:"facility_id" form:"facility_id"`
}

type Get_Homestay_Facility struct {
	ID            uint
	Name_Homestay string
	Name_Facility string
}
