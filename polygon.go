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
func (p *NullPolygon) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	polygon := ewkb.Polygon{}

	if err := ewkb.Unmarshal(&polygon, value); err != nil {
		return err
	}

	p.Polygon = polygonFromEWKB(polygon)
	p.Valid = true

	return nil
}

// Scan implements the SQL driver.Scanner interface.
func (p *Polygon) Scan(value interface{}) error {
	polygon := ewkb.Polygon{}

	if err := ewkb.Unmarshal(&polygon, value); err != nil {
		return err
	}

	*p = polygonFromEWKB(polygon)

	return nil
}

// Value implements the driver.Valuer interface.
func (p Polygon) Value() (driver.Value, error) {
	var srid *ewkb.SystemReferenceID

	polygon := ewkb.Polygon{
		CoordinateGroup: make(ewkb.CoordinateGroup, len(p)),
	}

	for idx0, ring := range p {
		polygon.CoordinateGroup[idx0] = make(ewkb.CoordinateSet, len(ring))
		for idx1, pnt := range ring {
			srid = pnt.SRID
			polygon.CoordinateGroup[idx0][idx1] = pnt.Coordinate
		}
	}

	polygon.SRID = srid

	return ewkb.Marshal(polygonToEWKB(p))
}

// Value implements the driver.Valuer interface.
func (p NullPolygon) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return p.Polygon.Value()
}

func polygonFromEWKB(polygon ewkb.Polygon) Polygon {
	ringSet := make([]LineString, len(polygon.CoordinateGroup))

	for idx0, ring := range polygon.CoordinateGroup {
		pointSet := make([]Point, len(ring))
		for idx1, pnt := range ring {
			pointSet[idx1].Coordinate = pnt
		}
		ringSet[idx0] = LineString(pointSet)
	}

	return Polygon(ringSet)
}

func polygonToEWKB(p Polygon) ewkb.Polygon {
	var srid *ewkb.SystemReferenceID

	polygon := ewkb.Polygon{
		CoordinateGroup: make(ewkb.CoordinateGroup, len(p)),
	}

	for idx0, ring := range p {
		polygon.CoordinateGroup[idx0] = make(ewkb.CoordinateSet, len(ring))
		for idx1, pnt := range ring {
			srid = pnt.SRID
			polygon.CoordinateGroup[idx0][idx1] = pnt.Coordinate
		}
	}

	polygon.SRID = srid

	return polygon
}
