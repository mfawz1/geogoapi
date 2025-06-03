package database

import (
	"github.com/nferruzzi/gormgis"
	"gorm.io/gorm"
)

type GeoEntity struct {
	gorm.Model
	Entity
	//srid4326 point
	gormGIS.GeoPoint `gorm:"type:geometry"`
}
