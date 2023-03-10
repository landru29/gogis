package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// MultiLineString is MULTILINESTRING in database.
type MultiLineString []LineString

// NullMultiLineString represents a MultiLineString that may be null.
// NullMultiLineString implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var multi gogis.NullMultiLineString
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&multi)
//	...
//	if multi.Valid {
//	   // use multi.MultiLineString
//	} else {
//	   // NULL value
//	}
type NullMultiLineString struct {
	MultiLineString MultiLineString
	Valid           bool
}

// Scan implements the SQL driver.Scanner interface.
func (m *NullMultiLineString) Scan(value interface{}) error {
	if dataBytes, ok := value.([]byte); ok && dataBytes == nil {
		return nil
	}

	multi := ewkb.MultiLineString{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	m.Valid = true

	return (&m.MultiLineString).FromEWKB(multi)
}

// Scan implements the SQL driver.Scanner interface.
func (m *MultiLineString) Scan(value interface{}) error {
	multi := ewkb.MultiLineString{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	return m.FromEWKB(multi)
}

// Value implements the driver.Valuer interface.
func (m MultiLineString) Value() (driver.Value, error) {
	return ewkb.Marshal(m.ToEWKB())
}

// Value implements the driver.Valuer interface.
func (m NullMultiLineString) Value() (driver.Value, error) {
	if !m.Valid {
		return nil, nil
	}

	return m.MultiLineString.Value()
}

func (m MultiLineString) srid() *ewkb.SystemReferenceID {
	for _, poly := range m {
		for _, pnt := range poly {
			return pnt.SRID
		}
	}

	return nil
}

// ToEWKB implements the ModelConverter interface.
func (m MultiLineString) ToEWKB() ewkb.Geometry { //nolint: ireturn
	multi := ewkb.MultiLineString{
		LineStrings: make([]ewkb.LineString, len(m)),
	}

	multi.SRID = m.srid()

	for idx, poly := range m {
		multiline, _ := poly.ToEWKB().(*ewkb.LineString)
		multi.LineStrings[idx] = *multiline
	}

	return &multi
}

// FromEWKB implements the ModelConverter interface.
func (m *MultiLineString) FromEWKB(from interface{}) error {
	multi, ok := fromPtr(from).(ewkb.MultiLineString)
	if !ok {
		return ewkb.ErrWrongGeometryType
	}

	polySet := make([]LineString, len(multi.LineStrings))

	for idx0, poly := range multi.LineStrings {
		if err := (&polySet[idx0]).FromEWKB(poly); err != nil {
			return err
		}
	}

	*m = MultiLineString(polySet)

	return nil
}

// Geometry converts to a generic geometry.
func (m MultiLineString) Geometry(opts ...func(interface{})) Geometry {
	output := Geometry{
		Type:     ewkb.GeometryTypeMultiLineString,
		Geometry: &m,
		Valid:    true,
	}

	for _, opt := range opts {
		opt(&output)
	}

	return output
}
