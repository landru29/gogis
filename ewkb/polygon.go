package ewkb

import (
	"encoding/binary"
	"fmt"
)

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
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, p.Type())
	}

	p.SRID = record.SRID

	return (&(p.CoordinateGroup)).UnmarshalEWBK(record)
}

// MarshalEWBK implements the Marshaler interface.
func (p Polygon) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	return p.CoordinateGroup.MarshalEWBK(byteOrder)
}

// SystemReferenceID implements the Marshaler interface.
func (p Polygon) SystemReferenceID() *SystemReferenceID {
	return p.SRID
}

// Layout implements the Marshaler interface.
func (p Polygon) Layout() Layout {
	indexes := []byte{}

	for idx0 := range p.CoordinateGroup {
		for idx1 := range p.CoordinateGroup[idx0] {
			return p.CoordinateGroup[idx0][idx1].Layout()
		}
	}

	return newLayoutFrom(indexes)
}
