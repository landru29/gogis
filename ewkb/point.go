package ewkb

import (
	"encoding/binary"
	"fmt"
)

// Point is a POINT in database.
//
// A Point is a 0-dimensional geometry that represents a single location in coordinate space.
type Point struct {
	SRID *SystemReferenceID
	Coordinate
}

// Type implements the Geometry interface.
func (p Point) Type() GeometryType {
	return GeometryTypePoint
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (p *Point) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != p.Type() {
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, p.Type())
	}

	p.SRID = record.SRID

	if err := (&(p.Coordinate)).UnmarshalEWBK(record); err != nil {
		return err
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (p Point) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	return p.Coordinate.MarshalEWBK(byteOrder)
}

// SystemReferenceID implements the Marshaler interface.
func (p Point) SystemReferenceID() *SystemReferenceID {
	return p.SRID
}

// Layout implements the Marshaler interface.
func (p Point) Layout() Layout {
	return p.Coordinate.Layout()
}
