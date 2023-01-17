package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// Linestring is LINESTRING in database.
type Linestring []Point

// NullLinestring represents a Linestring that may be null.
// NullLinestring implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var line gogis.NullLinestring
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&line)
//	...
//	if line.Valid {
//	   // use line.Linestring
//	} else {
//	   // NULL value
//	}
type NullLinestring struct {
	Linestring Linestring
	Valid      bool
}

// Scan implements the SQL driver.Scanner interface.
func (l *NullLinestring) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	linestring := ewkb.Linestring{}

	if err := ewkb.Unmarshal(&linestring, value); err != nil {
		return err
	}

	pointSet := make([]Point, len(linestring.CoordinateSet))

	for idx, pnt := range linestring.CoordinateSet {
		pointSet[idx].Coordinate = pnt
	}

	l.Linestring = Linestring(pointSet)
	l.Valid = true

	return nil
}

// Scan implements the SQL driver.Scanner interface.
func (l *Linestring) Scan(value interface{}) error {
	linestring := ewkb.Linestring{}

	if err := ewkb.Unmarshal(&linestring, value); err != nil {
		return err
	}

	pointSet := make([]Point, len(linestring.CoordinateSet))

	for idx, pnt := range linestring.CoordinateSet {
		pointSet[idx].Coordinate = pnt
	}

	*l = Linestring(pointSet)

	return nil
}

// Value implements the driver.Valuer interface.
func (l Linestring) Value() (driver.Value, error) {
	linestring := ewkb.Linestring{
		CoordinateSet: make(ewkb.CoordinateSet, len(l)),
	}

	if len(l) > 0 {
		linestring.SRID = l[0].SRID
	}

	for idx, pnt := range l {
		linestring.CoordinateSet[idx] = pnt.Coordinate
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
