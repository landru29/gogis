package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// Point is a lat lng position in database.
type Point ewkb.Point

// NullPoint represents a Point that may be null.
// NullPoint implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var pt gogis.NullPoint
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&pt)
//	...
//	if pt.Valid {
//	   // use pt.Point
//	} else {
//	   // NULL value
//	}
type NullPoint struct {
	Point Point
	Valid bool
}

// Scan implements the SQL driver.Scanner interface.
func (p *NullPoint) Scan(value interface{}) error {
	if dataBytes, ok := value.([]byte); ok && dataBytes == nil {
		return nil
	}

	point := ewkb.Point{}
	if err := ewkb.Unmarshal(&point, value); err != nil {
		return err
	}

	p.Valid = !point.Coordinate.IsNull()

	return (&p.Point).FromEWKB(point)
}

// Scan implements the SQL driver.Scanner interface.
func (p *Point) Scan(value interface{}) error {
	point := ewkb.Point{}
	if err := ewkb.Unmarshal(&point, value); err != nil {
		return err
	}

	return p.FromEWKB(point)
}

// Value implements the driver Valuer interface.
func (p Point) Value() (driver.Value, error) {
	return ewkb.Marshal(ewkb.Point(p))
}

// Value implements the driver Valuer interface.
func (p NullPoint) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return p.Point.Value()
}

// FromEWKB implements the ModelConverter interface.
func (p *Point) FromEWKB(from interface{}) error {
	pnt, ok := from.(ewkb.Point)
	if !ok {
		return ewkb.ErrWrongGeometryType
	}

	*p = Point(pnt)

	return nil
}
