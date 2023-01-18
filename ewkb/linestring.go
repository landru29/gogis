package ewkb

import "encoding/binary"

// LineString is a set of lines.
type LineString struct {
	SRID *SystemReferenceID
	CoordinateSet
}

// Type implements the Geometry interface.
func (l LineString) Type() GeometryType {
	return GeometryTypeLineString
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (l *LineString) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != l.Type() {
		return ErrWrongGeometryType
	}

	l.SRID = record.SRID

	return l.CoordinateSet.UnmarshalEWBK(record)
}

// MarshalEWBK implements the Marshaler interface.
func (l LineString) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	return l.CoordinateSet.MarshalEWBK(byteOrder)
}

// SystemReferenceID implements the Marshaler interface.
func (l LineString) SystemReferenceID() *SystemReferenceID {
	return l.SRID
}

// Layout implements the Marshaler interface.
func (l LineString) Layout() Layout {
	indexes := []byte{}

	if len(l.CoordinateSet) > 0 {
		return l.CoordinateSet[0].Layout()
	}

	return newLayoutFrom(indexes)
}
