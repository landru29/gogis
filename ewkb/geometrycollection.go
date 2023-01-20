package ewkb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

const (
	// ErrMissingWellKnownGeometry is when no well known geometries were specified.
	ErrMissingWellKnownGeometry = Error("missing well known geometries")
)

// GeometryCollection is a heterogeneous (mixed) collection of geometries.
type GeometryCollection struct {
	SRID       *SystemReferenceID
	Collection []Geometry

	wellKnownGeometry []Geometry
}

func (g GeometryCollection) pick(geoType GeometryType) (Geometry, error) { //nolint: ireturn
	if len(g.wellKnownGeometry) == 0 {
		return nil, ErrMissingWellKnownGeometry
	}

	for _, geo := range g.wellKnownGeometry {
		if geo.Type() == geoType {
			newGeo := reflect.New(reflect.TypeOf(geo).Elem())

			out, _ := newGeo.Interface().(Geometry)

			return out, nil
		}
	}

	return nil, ErrWrongGeometryType
}

// DefaultWellKnownGeometry is the default well known geometry set.
func DefaultWellKnownGeometry() []Geometry {
	return []Geometry{
		&Point{},
		&LineString{},
		&Polygon{},
		&MultiPoint{},
		&MultiLineString{},
		&MultiPolygon{},
		&Triangle{},
		&CircularString{},
		&GeometryCollection{},
	}
}

// NewGeometryCollection creates a new empty collection of geometries.
func NewGeometryCollection(wellknownGeometry ...Geometry) *GeometryCollection {
	if len(wellknownGeometry) == 0 {
		return &GeometryCollection{wellKnownGeometry: DefaultWellKnownGeometry()}
	}

	return &GeometryCollection{wellKnownGeometry: wellknownGeometry}
}

// Type implements the Geometry interface.
func (g GeometryCollection) Type() GeometryType {
	return GeometryTypeGeometryCollection
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (g *GeometryCollection) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != g.Type() {
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, g.Type())
	}

	g.SRID = record.SRID

	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	g.Collection = make([]Geometry, size)

	for idx := range g.Collection {
		dataSet, err := DecodeHeader(record.DataStream)
		if err != nil {
			return err
		}

		geometry, err := g.pick(dataSet.Type)
		if err != nil {
			return err
		}

		record.ExtendedWellKnownBytesHeader.Type = dataSet.Type

		if err := geometry.UnmarshalEWBK(*dataSet); err != nil {
			return err
		}

		g.Collection[idx] = geometry
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (g GeometryCollection) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	byteOrder.PutUint32(size, uint32(len(g.Collection)))
	output = append(output, size...)

	buffer := bytes.NewBuffer(nil)

	for _, geo := range g.Collection {
		if err := (&Encoder{writer: buffer, byteOrder: byteOrder, ignoreSRID: true}).Encode(geo); err != nil {
			return nil, err
		}
	}

	output = append(output, buffer.Bytes()...)

	return output, nil
}

// SystemReferenceID implements the Marshaler interface.
func (g GeometryCollection) SystemReferenceID() *SystemReferenceID {
	return g.SRID
}

// Layout implements the Marshaler interface.
func (g GeometryCollection) Layout() Layout {
	for _, geo := range g.Collection {
		return geo.Layout()
	}

	return 0
}
