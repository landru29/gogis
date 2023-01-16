package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// Point is a lat lng position in database.
type Point ewkb.Point

// NullPoint is a nullable value of Point.
type NullPoint struct {
	Point Point
	Valid bool
}

// Scan is the implementation of sql driver.
func (p *NullPoint) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	point := ewkb.Point{}
	if err := ewkb.Unmarshal(&point, value); err != nil {
		return err
	}

	p.Point = Point(point)
	p.Valid = true

	return nil
}

// Scan is the implementation of sql driver.
func (p *Point) Scan(value interface{}) error {
	point := ewkb.Point{}
	if err := ewkb.Unmarshal(&point, value); err != nil {
		return err
	}

	*p = Point(point)

	return nil
}

// Value implements the driver.Valuer interface.
func (p Point) Value() (driver.Value, error) {
	return ewkb.Marshal(ewkb.Point(p))
}

// Value implements the driver.Valuer interface.
func (p NullPoint) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return p.Point.Value()
}
