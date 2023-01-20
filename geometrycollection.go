package gogis

import (
	"database/sql/driver"

	"github.com/landru29/gogis/ewkb"
)

// GeometryCollection is a lat lng position in database.
type GeometryCollection struct {
	Collection []ModelConverter
	Valid      bool
	SRID       *ewkb.SystemReferenceID

	wellknown BindSet
}

// NewGeometryCollection creates a new empty collection.
func NewGeometryCollection(opts ...func(interface{})) *GeometryCollection {
	output := &GeometryCollection{}

	for _, opt := range opts {
		opt(output)
	}

	if len(opts) == 0 {
		output.wellknown = DefaultWellKnownBinding()
	}

	return output
}

// WithGeometry adds geometry to the collection.
func WithGeometry(geometry ...ModelConverter) func(interface{}) {
	return func(geo interface{}) {
		if collection, ok := geo.(*GeometryCollection); ok {
			collection.Collection = append(collection.Collection, geometry...)

			if len(collection.Collection) > 0 {
				collection.Valid = true
			}
		}
	}
}

// WithSystemReferenceID specifies the system reference ID.
func WithSystemReferenceID(srid ewkb.SystemReferenceID) func(interface{}) {
	return func(geo interface{}) {
		if collection, ok := geo.(*GeometryCollection); ok {
			collection.SRID = &srid
		}
	}
}

// Scan implements the SQL driver.Scanner interface.
func (g *GeometryCollection) Scan(value interface{}) error {
	if dataBytes, ok := value.([]byte); ok && dataBytes == nil {
		return nil
	}

	wellKnown := make([]ewkb.Geometry, len(g.wellknown))
	for idx, binding := range g.wellknown {
		wellKnown[idx] = binding.ewkbType
	}

	collection := ewkb.NewGeometryCollection(wellKnown...)

	if err := ewkb.Unmarshal(collection, value); err != nil {
		return err
	}

	return g.FromEWKB(collection)
}

// Value implements the driver Valuer interface.
func (g GeometryCollection) Value() (driver.Value, error) {
	if len(g.Collection) == 0 || !g.Valid {
		return nil, nil
	}

	return ewkb.Marshal(g.ToEWKB())
}

// FromEWKB implements the ModelConverter interface.
func (g *GeometryCollection) FromEWKB(from interface{}) error {
	collection, ok := fromPtr(from).(ewkb.GeometryCollection)
	if !ok {
		return ewkb.ErrWrongGeometryType
	}

	g.SRID = collection.SRID

	g.Collection = make([]ModelConverter, len(collection.Collection))
	for idx, geo := range collection.Collection {
		converter, err := g.wellknown.pick(geo.Type())
		if err != nil {
			return err
		}

		if err := converter.FromEWKB(geo); err != nil {
			return err
		}

		g.Collection[idx] = converter
	}

	if len(g.Collection) > 0 {
		g.Valid = true
	}

	return nil
}

// ToEWKB implements the ModelConverter interface.
func (g GeometryCollection) ToEWKB() ewkb.Geometry { //nolint: ireturn
	collection := ewkb.NewGeometryCollection()

	collection.Collection = make([]ewkb.Geometry, len(g.Collection))

	collection.SRID = g.SRID

	for idx, geo := range g.Collection {
		collection.Collection[idx] = geo.ToEWKB()
	}

	return collection
}
