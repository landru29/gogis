package ewkb

import "encoding/binary"

type Polygon struct {
	SRID *SystemReferenceID
	CoordinateGroup
}

// Type implements the Geometry interface.
func (p Polygon) Type() GeometryType {
	return GeometryTypePolygon
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (p *Polygon) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != p.Type() {
		return ErrWrongGeometryType
	}

	return (&(p.CoordinateGroup)).UnmarshalEWBK(record)
}

// MarshalEWBK implements the Marshaler interface.
func (p Polygon) MarshalEWBK(header ExtendedWellKnownBytesHeader) ([]byte, error) {
	return p.CoordinateGroup.MarshalEWBK(header)
}

// Header implements the Marshaler interface.
func (p Polygon) Header() ExtendedWellKnownBytesHeader {
	indexes := []byte{}

	if len(p.CoordinateGroup) > 0 && len(p.CoordinateGroup[0]) > 0 {
		for idx1 := range p.CoordinateGroup[0][0] {
			indexes = append(indexes, idx1)
		}
	}

	return ExtendedWellKnownBytesHeader{
		Type:      p.Type(),
		Layout:    newLayoutFrom(indexes),
		ByteOrder: binary.LittleEndian,
		SRID:      p.SRID,
	}
}
