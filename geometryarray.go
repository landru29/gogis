package gogis

import (
	"bytes"
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

const (
	geometryArrayMinSize = 2
)

// GeometryArray is an array of geometries.
type GeometryArray []Geometry

// Scan implements the SQL driver.Scanner interface.
func (g *GeometryArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if strData, ok := value.(string); ok {
		return g.Scan([]byte(strData))
	}

	dataByte, ok := value.([]byte)
	if !ok {
		return ewkb.ErrIncompatibleFormat
	}

	if dataByte == nil {
		return nil
	}

	if len(dataByte) < geometryArrayMinSize {
		return ewkb.ErrIncompatibleFormat
	}

	splitted := bytes.Split(dataByte[1:len(dataByte)-1], []byte(":"))

	*g = make(GeometryArray, len(splitted))

	for idx, datab := range splitted {
		geom := Geometry{}
		if err := (&geom).Scan(datab); err != nil {
			return err
		}

		(*g)[idx] = geom
	}

	return nil
}

// Value implements the driver Valuer interface.
func (g GeometryArray) Value() (driver.Value, error) {
	if len(g) == 0 {
		return nil, nil
	}

	ewkbData := make([][]byte, len(g))

	for idx, geo := range g {
		drvValue, err := geo.Value()
		if err != nil {
			return nil, err
		}

		val, _ := drvValue.([]byte)

		if idx == 0 {
			ewkbData[idx] = append([]byte("{"), val...)
		}

		if idx == len(g)-1 {
			ewkbData[idx] = append(val, '}')
		}
	}

	return bytes.Join(ewkbData, []byte(":")), nil
}
