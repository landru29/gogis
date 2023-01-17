package ewkb

import "encoding/binary"

// Linestring is a set of lines.
type Linestring struct {
	SRID *SystemReferenceID
	CoordinateSet
}

// Type implements the Geometry interface.
func (l Linestring) Type() GeometryType {
	return GeometryTypeLineString
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (l *Linestring) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != l.Type() {
		return ErrWrongGeometryType
	}

	l.SRID = record.SRID

	return l.CoordinateSet.UnmarshalEWBK(record)
}

// MarshalEWBK implements the Marshaler interface.
func (l Linestring) MarshalEWBK(header ExtendedWellKnownBytesHeader) ([]byte, error) {
	return l.CoordinateSet.MarshalEWBK(header)
}

// Header implements the Marshaler interface.
func (l Linestring) Header() ExtendedWellKnownBytesHeader {
	indexes := []byte{}

	if len(l.CoordinateSet) > 0 {
		for idx := range l.CoordinateSet[0] {
			indexes = append(indexes, idx)
		}
	}

	return ExtendedWellKnownBytesHeader{
		Type:      l.Type(),
		Layout:    newLayoutFrom(indexes),
		ByteOrder: binary.LittleEndian,
		SRID:      l.SRID,
	}
}
