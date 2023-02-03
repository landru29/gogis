package gogis

import (
	"bytes"
	"database/sql/driver"
	"encoding/hex"

	"github.com/landru29/gogis/ewkb"
)

var globalWellknownBindings = DefaultWellKnownBinding() //nolint: gochecknoglobals

// AppendWellKnownBinding add new binding to the globazl list.
func AppendWellKnownBinding(binding Binding) {
	out := append([]Binding{}, binding)
	globalWellknownBindings = append(out, globalWellknownBindings...)
}

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

	wellknown := g.wellknown

	if len(wellknown) == 0 {
		wellknown = globalWellknownBindings
	}

	for _, bind := range wellknown {
		if bind.ewkbType.Type() == record.Type {
			if err := bind.ewkbType.UnmarshalEWBK(*record); err != nil {
				return err
			}

			g.Geometry = bind.modelType

			err := bind.modelType.FromEWKB(bind.ewkbType)
			if err == nil {
				g.Valid = true
			}

			return err
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
