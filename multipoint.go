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
	if dataBytes, ok := value.([]byte); ok && dataBytes == nil {
		return nil
	}

	multi := ewkb.MultiPoint{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	m.Valid = true

	return (&m.MultiPoint).FromEWKB(multi)
}

// Scan implements the SQL driver.Scanner interface.
func (m *MultiPoint) Scan(value interface{}) error {
	multi := ewkb.MultiPoint{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	return m.FromEWKB(multi)
}

// Value implements the driver.Valuer interface.
func (m MultiPoint) Value() (driver.Value, error) {
	return ewkb.Marshal(m.ToEWKB())
}

// Value implements the driver.Valuer interface.
func (m NullMultiPoint) Value() (driver.Value, error) {
	if !m.Valid {
		return nil, nil
	}

	return m.MultiPoint.Value()
}

// ToEWKB implements the ModelConverter interface.
func (m MultiPoint) ToEWKB() ewkb.Marshaler { //nolint: ireturn
	multi := ewkb.MultiPoint{
		Points: make([]ewkb.Point, len(m)),
	}

	if len(m) > 0 {
		multi.SRID = m[0].SRID
	}

	for idx, pnt := range m {
		multi.Points[idx] = ewkb.Point(pnt)
	}

	return multi
}

// FromEWKB implements the ModelConverter interface.
func (m *MultiPoint) FromEWKB(from interface{}) error {
	multi, ok := fromPtr(from).(ewkb.MultiPoint)
	if !ok {
		return ewkb.ErrWrongGeometryType
	}

	pointSet := make([]Point, len(multi.Points))

	for idx, pnt := range multi.Points {
		pointSet[idx] = Point(pnt)
	}

	*m = MultiPoint(pointSet)

	return nil
}
