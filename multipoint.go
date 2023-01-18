package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// MultiPoint is MULTIPOINT in database.
type MultiPoint []Point

// NullMultiPoint represents a MultiPoint that may be null.
// NullMultiPoint implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var multi gogis.NullMultiPoint
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&multi)
//	...
//	if multi.Valid {
//	   // use multi.MultiPoint
//	} else {
//	   // NULL value
//	}
type NullMultiPoint struct {
	MultiPoint MultiPoint
	Valid      bool
}

// Scan implements the SQL driver.Scanner interface.
func (m *NullMultiPoint) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	multi := ewkb.MultiPoint{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	pointSet := make([]Point, len(multi.Points))

	for idx, pnt := range multi.Points {
		pointSet[idx] = Point(pnt)
	}

	m.MultiPoint = MultiPoint(pointSet)
	m.Valid = true

	return nil
}

// Scan implements the SQL driver.Scanner interface.
func (m *MultiPoint) Scan(value interface{}) error {
	multi := ewkb.MultiPoint{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	pointSet := make([]Point, len(multi.Points))

	for idx, pnt := range multi.Points {
		pointSet[idx] = Point(pnt)
	}

	*m = MultiPoint(pointSet)

	return nil
}

// Value implements the driver.Valuer interface.
func (m MultiPoint) Value() (driver.Value, error) {
	multi := ewkb.MultiPoint{
		Points: make([]ewkb.Point, len(m)),
	}

	if len(m) > 0 {
		multi.SRID = m[0].SRID
	}

	for idx, pnt := range m {
		multi.Points[idx] = ewkb.Point(pnt)
	}

	return ewkb.Marshal(multi)
}

// Value implements the driver.Valuer interface.
func (m NullMultiPoint) Value() (driver.Value, error) {
	if !m.Valid {
		return nil, nil
	}

	return m.MultiPoint.Value()
}
