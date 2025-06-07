package database

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

// GeoPoint implementation copied from and modified to Scan for string not uint8 "github.com/nferruzzi/gormgis"
type GeoPoint struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p.Lng, p.Lat)
}

func (p *GeoPoint) Scan(val any) error {
	valString, ok := val.(string)
	if !ok {
		log.Fatal("Scanning GeoPoint failed")
	}
	b, err := hex.DecodeString(valString)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %d", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}

	return nil
}

func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}

type GeoEntity struct {
	ID        uint      `gorm:"primarykey" uri:"id" fake:"-"`
	CreatedAt time.Time `fake:"-"`
	UpdatedAt time.Time `fake:"-"`
	Entity
	//srid4326 point
	GeoPoint `gorm:"type:geometry"`
}
