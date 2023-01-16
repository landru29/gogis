package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// Linestring is LINESTRING in database.
type Linestring []Point

// NullLinestring is a nullable value of Linestring.
type NullLinestring struct {
	Linestring Linestring
	Valid      bool
}

// Scan is the implementation of sql driver.
func (l *NullLinestring) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	linestring := ewkb.Linestring{}

	if err := ewkb.Unmarshal(&linestring, value); err != nil {
		return err
	}

	pointSet := make([]Point, len(linestring.Points))

	for idx, pnt := range linestring.Points {
		pointSet[idx] = Point(pnt)
	}

	l.Linestring = Linestring(pointSet)
	l.Valid = true

	return nil
}

// Scan is the implementation of sql driver.
func (l *Linestring) Scan(value interface{}) error {
	linestring := ewkb.Linestring{}

	if err := ewkb.Unmarshal(&linestring, value); err != nil {
		return err
	}

	pointSet := make([]Point, len(linestring.Points))

	for idx, pnt := range linestring.Points {
		pointSet[idx] = Point(pnt)
	}

	*l = Linestring(pointSet)

	return nil
}

// Value implements the driver.Valuer interface.
func (l Linestring) Value() (driver.Value, error) {
	linestring := ewkb.Linestring{
		Points: make([]ewkb.Point, len(l)),
	}

	if len(l) > 0 {
		linestring.SRID = l[0].SRID
	}

	for idx, pnt := range l {
		linestring.Points[idx] = ewkb.Point(pnt)
	}

	return ewkb.Marshal(linestring)
}

// Value implements the driver.Valuer interface.
func (l NullLinestring) Value() (driver.Value, error) {
	if !l.Valid {
		return nil, nil
	}

	return l.Linestring.Value()
}
