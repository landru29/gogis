package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// CircularString is CIRCULARSTRING in database.
type CircularString []Point

// NullCircularString represents a CircularString that may be null.
// NullCircularString implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var triangle gogis.NullCircularString
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&triangle)
//	...
//	if triangle.Valid {
//	   // use triangle.CircularString
//	} else {
//	   // NULL value
//	}
type NullCircularString struct {
	CircularString CircularString
	Valid          bool
}

// Scan implements the SQL driver.Scanner interface.
func (c *NullCircularString) Scan(value interface{}) error {
	if dataBytes, ok := value.([]byte); ok && dataBytes == nil {
		return nil
	}

	triangle := ewkb.CircularString{}

	if err := ewkb.Unmarshal(&triangle, value); err != nil {
		return err
	}

	c.CircularString = CircularStringFromEWKB(triangle)
	c.Valid = true

	return nil
}

// Scan implements the SQL driver.Scanner interface.
func (c *CircularString) Scan(value interface{}) error {
	triangle := ewkb.CircularString{}

	if err := ewkb.Unmarshal(&triangle, value); err != nil {
		return err
	}

	*c = CircularStringFromEWKB(triangle)

	return nil
}

// Value implements the driver.Valuer interface.
func (c CircularString) Value() (driver.Value, error) {
	var srid *ewkb.SystemReferenceID

	triangle := ewkb.CircularString{
		CoordinateSet: make(ewkb.CoordinateSet, len(c)),
	}

	for idx, pnt := range c {
		triangle.CoordinateSet[idx] = pnt.Coordinate
	}

	triangle.SRID = srid

	return ewkb.Marshal(triangle)
}

// Value implements the driver.Valuer interface.
func (c NullCircularString) Value() (driver.Value, error) {
	if !c.Valid {
		return nil, nil
	}

	return c.CircularString.Value()
}

// CircularStringFromEWKB converts EWKB to CircularString.
func CircularStringFromEWKB(circular ewkb.CircularString) CircularString {
	poly := make([]Point, len(circular.CoordinateSet))

	for idx, pnt := range circular.CoordinateSet {
		poly[idx].Coordinate = pnt
	}

	return CircularString(poly)
}
