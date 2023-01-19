package ewkb

import (
	"encoding/binary"
	"fmt"
)

const (
	// ErrCircularStringWrongSize occurs when the number of verticies is not odd or is less than 3.
	ErrCircularStringWrongSize = Error("circularstring has a wrong size (odd number of vertices >1)")
)

// CircularString is a CIRCULARSTRING in database.
//
// CircularString is the basic curve type, similar to a LineString in the linear world.
// A single arc segment is specified by three points: the start and end points (first
// and third) and some other point on the arc. To specify a closed circle the start and
// end points are the same and the middle point is the opposite point on the circle
// diameter (which is the center of the arc). In a sequence of arcs the end point of
// the previous arc is the start point of the next arc, just like the segments of a
// LineString. This means that a CircularString must have an odd number of points
// greater than 1.
type CircularString struct {
	SRID *SystemReferenceID
	CoordinateSet
}

// Type implements the Geometry interface.
func (c CircularString) Type() GeometryType {
	return GeometryTypeCircularString
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (c *CircularString) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != c.Type() {
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, c.Type())
	}

	c.SRID = record.SRID

	if err := c.CoordinateSet.UnmarshalEWBK(record); err != nil {
		return err
	}

	size := len(c.CoordinateSet)

	if size%2 == 0 || size < 3 {
		return fmt.Errorf("%w: found %d vertices", ErrCircularStringWrongSize, size)
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (c CircularString) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	size := len(c.CoordinateSet)

	if size%2 == 0 || size < 3 {
		return nil, fmt.Errorf("%w: found %d vertices", ErrCircularStringWrongSize, size)
	}

	return c.CoordinateSet.MarshalEWBK(byteOrder)
}

// SystemReferenceID implements the Marshaler interface.
func (c CircularString) SystemReferenceID() *SystemReferenceID {
	return c.SRID
}

// Layout implements the Marshaler interface.
func (c CircularString) Layout() Layout {
	indexes := []byte{}

	if len(c.CoordinateSet) > 0 {
		return c.CoordinateSet[0].Layout()
	}

	return newLayoutFrom(indexes)
}
