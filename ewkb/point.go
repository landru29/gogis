package ewkb

import "encoding/binary"

// Point is a lat lng position in database.
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
func (p Point) MarshalEWBK(header ExtendedWellKnownBytesHeader) ([]byte, error) {
	return p.Coordinate.MarshalEWBK(header)
}

// Header implements the Marshaler interface.
func (p Point) Header() ExtendedWellKnownBytesHeader {
	indexes := []byte{}
	for idx := range p.Coordinate {
		indexes = append(indexes, idx)
	}

	return ExtendedWellKnownBytesHeader{
		Type:      p.Type(),
		Layout:    newLayoutFrom(indexes),
		ByteOrder: binary.LittleEndian,
		SRID:      p.SRID,
	}
}
