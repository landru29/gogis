package ewkb

import (
	"encoding/binary"
	"fmt"
)

const (
	triangleVerticesCount = 4

	triangleCoordinateGroupSize uint32 = 1

	// ErrTriangleWrongSize occurs when a triangle has more than 1 polygon,
	// or the polygon hasn't 3 vertices.
	ErrTriangleWrongSize = Error("triangle has a wrong size")
)

type Triangle struct {
	SRID *SystemReferenceID
	CoordinateSet
}

// Type implements the Geometry interface.
func (t Triangle) Type() GeometryType {
	return GeometryTypeTriangle
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (t *Triangle) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != t.Type() {
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, t.Type())
	}

	t.SRID = record.SRID

	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	if size != triangleCoordinateGroupSize {
		return ErrTriangleWrongSize
	}

	if err := (&(t.CoordinateSet)).UnmarshalEWBK(record); err != nil {
		return err
	}

	if len(t.CoordinateSet) != triangleVerticesCount {
		return ErrTriangleWrongSize
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (t Triangle) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	if len(t.CoordinateSet) != triangleVerticesCount {
		return nil, ErrTriangleWrongSize
	}

	size := make([]byte, size32bit)

	byteOrder.PutUint32(size, triangleCoordinateGroupSize)
	output = append(output, size...)

	data, err := t.CoordinateSet.MarshalEWBK(byteOrder)

	output = append(output, data...)

	return output, err
}

// SystemReferenceID implements the Marshaler interface.
func (t Triangle) SystemReferenceID() *SystemReferenceID {
	return t.SRID
}

// Layout implements the Marshaler interface.
func (t Triangle) Layout() Layout {
	indexes := []byte{}

	for idx := range t.CoordinateSet {
		return t.CoordinateSet[idx].Layout()
	}

	return newLayoutFrom(indexes)
}
