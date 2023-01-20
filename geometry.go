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

	wellknown BindSet
}

// NewGeometry creates a new Geometry.
func NewGeometry(opts ...func(interface{})) *Geometry {
	output := &Geometry{}

	for _, opt := range opts {
		opt(output)
	}

	if len(opts) == 0 {
		output.wellknown = DefaultWellKnownBinding()
	}

	return output
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
	if !g.Valid || g.Geometry == nil {
		return nil, nil
	}

	converter, ok := g.Geometry.(ModelConverter)
	if !ok {
		return nil, ewkb.ErrIncompatibleFormat
	}

	return ewkb.Marshal(converter.ToEWKB())
}
