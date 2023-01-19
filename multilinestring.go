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

	m.MultiLineString = MultiLineStringFromEWKB(multi)
	m.Valid = true

	return nil
}

// Scan implements the SQL driver.Scanner interface.
func (m *MultiLineString) Scan(value interface{}) error {
	multi := ewkb.MultiLineString{}

	if err := ewkb.Unmarshal(&multi, value); err != nil {
		return err
	}

	*m = MultiLineStringFromEWKB(multi)

	return nil
}

// Value implements the driver.Valuer interface.
func (m MultiLineString) Value() (driver.Value, error) {
	multi := ewkb.MultiLineString{
		LineStrings: make([]ewkb.LineString, len(m)),
	}

	if len(m) > 0 {
		multi.SRID = m.srid()
	}

	for idx, poly := range m {
		multi.LineStrings[idx] = linestringToEWKB(poly)
	}

	return ewkb.Marshal(multi)
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

// MultiLineStringFromEWKB converts EWKB to MultiLineString.
func MultiLineStringFromEWKB(multi ewkb.MultiLineString) MultiLineString {
	polySet := make([]LineString, len(multi.LineStrings))

	for idx0, poly := range multi.LineStrings {
		polySet[idx0] = LinestringFromEWKB(poly)
	}

	return MultiLineString(polySet)
}
