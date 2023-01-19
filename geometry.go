package gogis

import (
	"bytes"
	"database/sql/driver"
	"encoding/hex"

	"github.com/landru29/gogis/ewkb"
)

// Geometry is any PostGIS geometry.
// This is used when user doesn't know which geometry to retrieve from database.
type Geometry struct {
	Type     ewkb.GeometryType
	Geometry ewkb.Geometry
	Valid    bool

	wellknown []ewkb.Geometry
}

// NewGeometry creates a new Geometry.
func NewGeometry(opts ...func(*Geometry)) *Geometry {
	output := &Geometry{
		wellknown: []ewkb.Geometry{
			&ewkb.Point{},
			&ewkb.LineString{},
			&ewkb.Polygon{},
			&ewkb.MultiPoint{},
			&ewkb.MultiLineString{},
			&ewkb.MultiPolygon{},
			&ewkb.Triangle{},
			&ewkb.CircularString{},
		},
	}

	for _, opt := range opts {
		opt(output)
	}

	return output
}

// WithWellKnownGeometry add custom Geometry to the wellknown.
func WithWellKnownGeometry(geometry ...ewkb.Geometry) func(*Geometry) {
	return func(shape *Geometry) {
		wellknown := []ewkb.Geometry{}
		wellknown = append(wellknown, geometry...)
		wellknown = append(wellknown, shape.wellknown...)

		shape.wellknown = wellknown
	}
}

// Scan implements the SQL driver.Scanner interface.
func (g *Geometry) Scan(value interface{}) error {
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

	record, err := ewkb.DecodeHeader(
		hex.NewDecoder(
			bytes.NewBuffer(dataByte),
		),
	)
	if err != nil {
		return err
	}

	g.Type = record.Type

	for _, geoShape := range g.wellknown {
		if geoShape.Type() == record.Type {
			g.Geometry = geoShape
			g.Valid = true

			return g.Geometry.UnmarshalEWBK(*record)
		}
	}

	return ewkb.ErrWrongGeometryType
}

// Value implements the driver Valuer interface.
func (g *Geometry) Value() (driver.Value, error) {
	if !g.Valid {
		return nil, nil
	}

	return ewkb.Marshal(g.Geometry)
}
