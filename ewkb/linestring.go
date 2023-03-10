package ewkb

import (
	"encoding/binary"
	"fmt"
)

// LineString is a set of lines.
//
// A LineString is a 1-dimensional line formed by a contiguous sequence
// of line segments. Each line segment is defined by two points, with
// the end point of one segment forming the start point of the next
// segment. An OGC-valid LineString has either zero or two or more points,
// but PostGIS also allows single-point LineStrings. LineStrings may cross
// themselves (self-intersect). A LineString is closed if the start and
// end points are the same. A LineString is simple if it does not
// self-intersect.
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
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, l.Type())
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
