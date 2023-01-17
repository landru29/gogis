package ewkb

import "encoding/binary"

// Point is a POINT in database.
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
		return ErrWrongGeometryType
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
