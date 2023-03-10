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

	t.Valid = true

	return (&t.Triangle).FromEWKB(triangle)
}

// Scan implements the SQL driver.Scanner interface.
func (t *Triangle) Scan(value interface{}) error {
	triangle := ewkb.Triangle{}

	if err := ewkb.Unmarshal(&triangle, value); err != nil {
		return err
	}

	return t.FromEWKB(triangle)
}

// Value implements the driver.Valuer interface.
func (t Triangle) Value() (driver.Value, error) {
	return ewkb.Marshal(t.ToEWKB())
}

// Value implements the driver.Valuer interface.
func (t NullTriangle) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}

	return t.Triangle.Value()
}

// FromEWKB implements the ModelConverter interface.
func (t *Triangle) FromEWKB(from interface{}) error {
	triangle, ok := fromPtr(from).(ewkb.Triangle)
	if !ok {
		return ewkb.ErrWrongGeometryType
	}

	poly := make([]Point, len(triangle.CoordinateSet))

	for idx, pnt := range triangle.CoordinateSet {
		poly[idx].Coordinate = pnt
		poly[idx].SRID = triangle.SRID
	}

	*t = Triangle(poly)

	return nil
}

// ToEWKB implements the ModelConverter interface.
func (t Triangle) ToEWKB() ewkb.Geometry { //nolint: ireturn
	var srid *ewkb.SystemReferenceID

	triangle := ewkb.Triangle{
		CoordinateSet: make(ewkb.CoordinateSet, len(t)),
	}

	for idx, pnt := range t {
		triangle.CoordinateSet[idx] = pnt.Coordinate
		srid = pnt.SRID
	}

	triangle.SRID = srid

	return &triangle
}

// Geometry converts to a generic geometry.
func (t Triangle) Geometry(opts ...func(interface{})) Geometry {
	output := Geometry{
		Type:     ewkb.GeometryTypeTriangle,
		Geometry: &t,
		Valid:    true,
	}

	for _, opt := range opts {
		opt(&output)
	}

	return output
}
