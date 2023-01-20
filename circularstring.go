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
//	var circle gogis.NullCircularString
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&circle)
//	...
//	if circle.Valid {
//	   // use circle.CircularString
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

	circle := ewkb.CircularString{}

	if err := ewkb.Unmarshal(&circle, value); err != nil {
		return err
	}

	c.Valid = true

	return (&c.CircularString).FromEWKB(circle)
}

// Scan implements the SQL driver.Scanner interface.
func (c *CircularString) Scan(value interface{}) error {
	circle := ewkb.CircularString{}

	if err := ewkb.Unmarshal(&circle, value); err != nil {
		return err
	}

	return c.FromEWKB(circle)
}

// Value implements the driver.Valuer interface.
func (c CircularString) Value() (driver.Value, error) {
	return ewkb.Marshal(c.ToEWKB())
}

// Value implements the driver.Valuer interface.
func (c NullCircularString) Value() (driver.Value, error) {
	if !c.Valid {
		return nil, nil
	}

	return c.CircularString.Value()
}

// ToEWKB implements the ModelConverter interface.
func (c CircularString) ToEWKB() ewkb.Marshaler { //nolint: ireturn
	var srid *ewkb.SystemReferenceID

	circle := ewkb.CircularString{
		CoordinateSet: make(ewkb.CoordinateSet, len(c)),
	}

	for idx, pnt := range c {
		circle.CoordinateSet[idx] = pnt.Coordinate
	}

	circle.SRID = srid

	return circle
}

// FromEWKB implements the ModelConverter interface.
func (c *CircularString) FromEWKB(from interface{}) error {
	circular, ok := fromPtr(from).(ewkb.CircularString)
	if !ok {
		return ewkb.ErrWrongGeometryType
	}

	poly := make([]Point, len(circular.CoordinateSet))

	for idx, pnt := range circular.CoordinateSet {
		poly[idx].Coordinate = pnt
	}

	*c = CircularString(poly)

	return nil
}
