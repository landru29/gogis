package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// LineString is LINESTRING in database.
type LineString []Point

// NullLineString represents a LineString that may be null.
// NullLineString implements the SQL driver.Scanner interface so it
// can be used as a scan destination:
//
//	var line gogis.NullLineString
//	err := db.QueryRow("SELECT coordinate FROM foo WHERE id=?", id).Scan(&line)
//	...
//	if line.Valid {
//	   // use line.LineString
//	} else {
//	   // NULL value
//	}
type NullLineString struct {
	LineString LineString
	Valid      bool
}

// Scan implements the SQL driver.Scanner interface.
func (l *NullLineString) Scan(value interface{}) error {
	if dataBytes, ok := value.([]byte); ok && dataBytes == nil {
		return nil
	}

	linestring := ewkb.LineString{}

	if err := ewkb.Unmarshal(&linestring, value); err != nil {
		return err
	}

	l.Valid = true

	return (&l.LineString).FromEWKB(linestring)
}

// Scan implements the SQL driver.Scanner interface.
func (l *LineString) Scan(value interface{}) error {
	linestring := ewkb.LineString{}

	if err := ewkb.Unmarshal(&linestring, value); err != nil {
		return err
	}

	return l.FromEWKB(linestring)
}

// Value implements the driver.Valuer interface.
func (l LineString) Value() (driver.Value, error) {
	return ewkb.Marshal(linestringToEWKB(l))
}

// Value implements the driver.Valuer interface.
func (l NullLineString) Value() (driver.Value, error) {
	if !l.Valid {
		return nil, nil
	}

	return l.LineString.Value()
}

// FromEWKB implements the ModelConverter interface.
func (l *LineString) FromEWKB(from interface{}) error {
	linestring, ok := from.(ewkb.LineString)
	if !ok {
		return ewkb.ErrWrongGeometryType
	}

	pointSet := make([]Point, len(linestring.CoordinateSet))

	for idx, pnt := range linestring.CoordinateSet {
		pointSet[idx].Coordinate = pnt
	}

	*l = LineString(pointSet)

	return nil
}

func linestringToEWKB(line LineString) ewkb.LineString {
	linestring := ewkb.LineString{
		CoordinateSet: make(ewkb.CoordinateSet, len(line)),
	}

	if len(line) > 0 {
		linestring.SRID = line[0].SRID
	}

	for idx, pnt := range line {
		linestring.CoordinateSet[idx] = pnt.Coordinate
	}

	return linestring
}
