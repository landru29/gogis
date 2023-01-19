package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// Triangle is TRIANGLE in database.
type Triangle []Point

// NullTriangle represents a Triangle that may be null.
// NullTriangle implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var triangle gogis.NullTriangle
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&triangle)
//	...
//	if triangle.Valid {
//	   // use triangle.Triangle
//	} else {
//	   // NULL value
//	}
type NullTriangle struct {
	Triangle Triangle
	Valid    bool
}

// Scan implements the SQL driver.Scanner interface.
func (t *NullTriangle) Scan(value interface{}) error {
	if dataBytes, ok := value.([]byte); ok && dataBytes == nil {
		return nil
	}

	triangle := ewkb.Triangle{}

	if err := ewkb.Unmarshal(&triangle, value); err != nil {
		return err
	}

	t.Triangle = TriangleFromEWKB(triangle)
	t.Valid = true

	return nil
}

// Scan implements the SQL driver.Scanner interface.
func (t *Triangle) Scan(value interface{}) error {
	triangle := ewkb.Triangle{}

	if err := ewkb.Unmarshal(&triangle, value); err != nil {
		return err
	}

	*t = TriangleFromEWKB(triangle)

	return nil
}

// Value implements the driver.Valuer interface.
func (t Triangle) Value() (driver.Value, error) {
	var srid *ewkb.SystemReferenceID

	triangle := ewkb.Triangle{
		CoordinateSet: make(ewkb.CoordinateSet, len(t)),
	}

	for idx, pnt := range t {
		triangle.CoordinateSet[idx] = pnt.Coordinate
	}

	triangle.SRID = srid

	return ewkb.Marshal(triangle)
}

// Value implements the driver.Valuer interface.
func (t NullTriangle) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}

	return t.Triangle.Value()
}

// TriangleFromEWKB converts EWKB to Triangle.
func TriangleFromEWKB(triangle ewkb.Triangle) Triangle {
	poly := make([]Point, len(triangle.CoordinateSet))

	for idx, pnt := range triangle.CoordinateSet {
		poly[idx].Coordinate = pnt
	}

	return Triangle(poly)
}
