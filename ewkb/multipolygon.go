package ewkb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// MultiPolygon is a MULTILINESTRING in database.
//
// A MultiPolygon is a collection of non-overlapping, non-adjacent Polygons.
// Polygons in the collection may touch only at a finite number of points.
type MultiPolygon struct {
	SRID     *SystemReferenceID
	Polygons []Polygon
}

// Type implements the Geometry interface.
func (m MultiPolygon) Type() GeometryType {
	return GeometryTypeMultiPolygon
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (m *MultiPolygon) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != m.Type() {
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, m.Type())
	}

	m.SRID = record.SRID

	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	m.Polygons = make([]Polygon, size)

	for idx := range m.Polygons {
		polygon := &Polygon{}
		if err := (&Decoder{reader: record.DataStream}).Decode(polygon); err != nil {
			return err
		}

		m.Polygons[idx] = *polygon
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (m MultiPolygon) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	byteOrder.PutUint32(size, uint32(len(m.Polygons)))
	output = append(output, size...)

	buffer := bytes.NewBuffer(nil)

	for _, pnt := range m.Polygons {
		if err := (&Encoder{writer: buffer, byteOrder: byteOrder}).Encode(pnt); err != nil {
			return nil, err
		}
	}

	output = append(output, buffer.Bytes()...)

	return output, nil
}

// SystemReferenceID implements the Marshaler interface.
func (m MultiPolygon) SystemReferenceID() *SystemReferenceID {
	return m.SRID
}

// Layout implements the Marshaler interface.
func (m MultiPolygon) Layout() Layout {
	if len(m.Polygons) > 0 {
		return m.Polygons[0].Layout()
	}

	return layoutXY
}
