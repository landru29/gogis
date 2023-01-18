package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// Polygon is POLYGON in database.
type Polygon []LineString

// NullPolygon represents a Polygon that may be null.
// NullPolygon implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var poly gogis.NullPolygon
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&poly)
//	...
//	if poly.Valid {
//	   // use poly.Polygon
//	} else {
//	   // NULL value
//	}
type NullPolygon struct {
	Polygon Polygon
	Valid   bool
}

// Scan implements the SQL driver.Scanner interface.
func (l *NullPolygon) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	polygon := ewkb.Polygon{}

	if err := ewkb.Unmarshal(&polygon, value); err != nil {
		return err
	}

	ringSet := make([]LineString, len(polygon.CoordinateGroup))

	for idx0, ring := range polygon.CoordinateGroup {
		pointSet := make([]Point, len(ring))
		for idx1, pnt := range ring {
			pointSet[idx1].Coordinate = pnt
		}
		ringSet[idx0] = LineString(pointSet)
	}

	l.Polygon = Polygon(ringSet)
	l.Valid = true

	return nil
}

// Scan implements the SQL driver.Scanner interface.
func (l *Polygon) Scan(value interface{}) error {
	polygon := ewkb.Polygon{}

	if err := ewkb.Unmarshal(&polygon, value); err != nil {
		return err
	}

	ringSet := make([]LineString, len(polygon.CoordinateGroup))

	for idx0, ring := range polygon.CoordinateGroup {
		pointSet := make([]Point, len(ring))
		for idx1, pnt := range ring {
			pointSet[idx1].Coordinate = pnt
		}
		ringSet[idx0] = LineString(pointSet)
	}

	*l = Polygon(ringSet)

	return nil
}

// Value implements the driver.Valuer interface.
func (l Polygon) Value() (driver.Value, error) {
	var srid *ewkb.SystemReferenceID

	polygon := ewkb.Polygon{
		CoordinateGroup: make(ewkb.CoordinateGroup, len(l)),
	}

	for idx0, ring := range l {
		polygon.CoordinateGroup[idx0] = make(ewkb.CoordinateSet, len(ring))
		for idx1, pnt := range ring {
			srid = pnt.SRID
			polygon.CoordinateGroup[idx0][idx1] = pnt.Coordinate
		}
	}

	polygon.SRID = srid

	return ewkb.Marshal(polygon)
}

// Value implements the driver.Valuer interface.
func (l NullPolygon) Value() (driver.Value, error) {
	if !l.Valid {
		return nil, nil
	}

	return l.Polygon.Value()
}
