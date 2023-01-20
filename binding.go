package gogis

import (
	"reflect"

	"github.com/landru29/gogis/ewkb"
)

// Binding is a type binding.
type Binding struct {
	ewkbType  ewkb.Geometry
	modelType ModelConverter
}

// BindSet is a set of bindings.
type BindSet []Binding

// Bind creates a binding.
func Bind(ewkbType ewkb.Geometry, modelType ModelConverter) Binding {
	return Binding{
		ewkbType:  ewkbType,
		modelType: modelType,
	}
}

// DefaultWellKnownBinding is the default binding.
func DefaultWellKnownBinding() BindSet {
	return []Binding{
		Bind(&ewkb.Point{}, &Point{}),
		Bind(&ewkb.LineString{}, &LineString{}),
		Bind(&ewkb.Polygon{}, &Polygon{}),
		Bind(&ewkb.MultiPoint{}, &MultiPoint{}),
		Bind(&ewkb.MultiLineString{}, &MultiLineString{}),
		Bind(&ewkb.MultiPolygon{}, &MultiPolygon{}),
		Bind(&ewkb.Triangle{}, &Triangle{}),
		Bind(&ewkb.CircularString{}, &CircularString{}),
		Bind(&ewkb.GeometryCollection{}, &GeometryCollection{}),
	}
}

// WithWellKnownGeometry add custom Geometry to the wellknown.
func WithWellKnownGeometry(binding ...Binding) func(interface{}) {
	return func(shape interface{}) {
		wellknown := []Binding{}
		wellknown = append(wellknown, binding...)

		switch out := shape.(type) {
		case *Geometry:
			wellknown = append(wellknown, out.wellknown...)

			out.wellknown = wellknown
		case *GeometryCollection:
			wellknown = append(wellknown, out.wellknown...)

			out.wellknown = wellknown
		}
	}
}

func (b BindSet) pick(geoType ewkb.GeometryType) (ModelConverter, error) { //nolint: ireturn
	for _, bind := range b {
		if bind.ewkbType.Type() == geoType {
			newGeo := reflect.New(reflect.TypeOf(bind.modelType).Elem())

			out, _ := newGeo.Interface().(ModelConverter)

			return out, nil
		}
	}

	return nil, ewkb.ErrWrongGeometryType
}
