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
	Geometry interface{}
	Valid    bool

	wellknown []Binding
}

// Binding is a type binding.
type Binding struct {
	ewkbType  ewkb.Geometry
	modelType ModelConverter
}

// Bind creates a binding.
func Bind(ewkbType ewkb.Geometry, modelType ModelConverter) Binding {
	return Binding{
		ewkbType:  ewkbType,
		modelType: modelType,
	}
}

// NewGeometry creates a new Geometry.
func NewGeometry(opts ...func(*Geometry)) *Geometry {
	output := &Geometry{
		wellknown: []Binding{
			Bind(&ewkb.Point{}, &Point{}),
			Bind(&ewkb.LineString{}, &LineString{}),
			Bind(&ewkb.Polygon{}, &Polygon{}),
			Bind(&ewkb.MultiPoint{}, &MultiPoint{}),
			Bind(&ewkb.MultiLineString{}, &MultiLineString{}),
			Bind(&ewkb.MultiPolygon{}, &MultiPolygon{}),
			Bind(&ewkb.Triangle{}, &Triangle{}),
			Bind(&ewkb.CircularString{}, &CircularString{}),
		},
	}

	for _, opt := range opts {
		opt(output)
	}

	return output
}

// WithWellKnownGeometry add custom Geometry to the wellknown.
func WithWellKnownGeometry(binding ...Binding) func(*Geometry) {
	return func(shape *Geometry) {
		wellknown := []Binding{}
		wellknown = append(wellknown, binding...)
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

	for _, bind := range g.wellknown {
		if bind.ewkbType.Type() == record.Type {
			if err := bind.ewkbType.UnmarshalEWBK(*record); err != nil {
				return err
			}

			g.Geometry = bind.modelType

			return bind.modelType.FromEWKB(bind.ewkbType)
		}
	}

	return ewkb.ErrWrongGeometryType
}

// Value implements the driver Valuer interface.
func (g *Geometry) Value() (driver.Value, error) {
	if !g.Valid {
		return nil, nil
	}

	converter, ok := g.Geometry.(ModelConverter)
	if !ok {
		return nil, ewkb.ErrIncompatibleFormat
	}

	return ewkb.Marshal(converter.ToEWKB())
}
