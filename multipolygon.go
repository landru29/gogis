package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// MultiPolygon is MULTIPOLYGON in database.
type MultiPolygon []Polygon

// NullMultiPolygon represents a MultiPolygon that may be null.
// NullMultiPolygon implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var multi gogis.NullMultiPolygon
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&multi)
//	...
//	if multi.Valid {
//	   // use multi.MultiPolygon
//	} else {
//	   // NULL value
//	}
type NullMultiPolygon struct {
	MultiPolygon MultiPolygon
	Valid        bool
}

// Scan implements the SQL driver.Scanner interface.
func (p *NullMultiPolygon) Scan(value interface{}) error {
	if dataBytes, ok := value.([]byte); ok && dataBytes == nil {
		return nil
	}

	multi := ewkb.MultiPolygon{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	p.MultiPolygon = MultiPolygonFromEWKB(multi)
	p.Valid = true

	return nil
}

// Scan implements the SQL driver.Scanner interface.
func (p *MultiPolygon) Scan(value interface{}) error {
	multi := ewkb.MultiPolygon{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	*p = MultiPolygonFromEWKB(multi)

	return nil
}

// Value implements the driver.Valuer interface.
func (p MultiPolygon) Value() (driver.Value, error) {
	multi := ewkb.MultiPolygon{
		Polygons: make([]ewkb.Polygon, len(p)),
	}

	if len(p) > 0 {
		multi.SRID = p.srid()
	}

	for idx, poly := range p {
		multi.Polygons[idx] = polygonToEWKB(poly)
	}

	return ewkb.Marshal(multi)
}

// Value implements the driver.Valuer interface.
func (p NullMultiPolygon) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return p.MultiPolygon.Value()
}

func (p MultiPolygon) srid() *ewkb.SystemReferenceID {
	for _, poly := range p {
		for _, line := range poly {
			for _, pnt := range line {
				return pnt.SRID
			}
		}
	}

	return nil
}

// MultiPolygonFromEWKB converts EWKB to MultiPolygon.
func MultiPolygonFromEWKB(multi ewkb.MultiPolygon) MultiPolygon {
	polySet := make([]Polygon, len(multi.Polygons))

	for idx0, poly := range multi.Polygons {
		polySet[idx0] = PolygonFromEWKB(poly)
	}

	return MultiPolygon(polySet)
}
