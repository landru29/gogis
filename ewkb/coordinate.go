package ewkb

import (
	"encoding/binary"
	"math"
)

// Coordinate is coordinate system.
type Coordinate map[byte]float64

// IsNull checks if coordinate is null.
func (c Coordinate) IsNull() bool {
	if len(c) == 0 {
		return true
	}

	for _, coord := range c {
		if coord != coord {
			return true
		}
	}

	return false
}

// NewNullCoordinate creates a null coordinate system.
func NewNullCoordinate(layout Layout) Coordinate {
	output := Coordinate{}
	for _, name := range layout.Format() {
		output[byte(name)] = math.NaN()
	}

	return output
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (c *Coordinate) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	var pnt point

	if err := (&pnt).read(record.DataStream, record.Layout.Size(), record.ByteOrder); err != nil {
		return err
	}

	*c = Coordinate{}

	if pnt.isNull() {
		*c = NewNullCoordinate(record.Layout)

		return nil
	}

	for idx, char := range record.Layout.Format() {
		(*c)[byte(char)] = pnt[idx]
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (c Coordinate) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	for _, name := range c.Layout().Format() {
		bytes := float64Bytes(c[byte(name)], byteOrder)
		output = append(output, bytes...)
	}

	return output, nil
}

// Layout implements the Marshaler interface.
func (c Coordinate) Layout() Layout {
	indexes := []byte{}
	for idx := range c {
		indexes = append(indexes, idx)
	}

	return newLayoutFrom(indexes)
}

// CoordinateSet is a set of coordinates.
type CoordinateSet []Coordinate

// UnmarshalEWBK implements the Unmarshaler interface.
func (c *CoordinateSet) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	*c = make(CoordinateSet, size)
	for idx := range *c {
		err := (&((*c)[idx])).UnmarshalEWBK(record)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (c CoordinateSet) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	byteOrder.PutUint32(size, uint32(len(c)))
	output = append(output, size...)

	for _, point := range c {
		dataByte, err := point.MarshalEWBK(byteOrder)
		if err != nil {
			return nil, err
		}

		output = append(output, dataByte...)
	}

	return output, nil
}

// CoordinateGroup is a group of set of coordinates.
type CoordinateGroup []CoordinateSet

// UnmarshalEWBK implements the Unmarshaler interface.
func (c *CoordinateGroup) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	*c = make(CoordinateGroup, size)
	for idx := range *c {
		err := (&((*c)[idx])).UnmarshalEWBK(record)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (c CoordinateGroup) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	byteOrder.PutUint32(size, uint32(len(c)))
	output = append(output, size...)

	for _, point := range c {
		dataByte, err := point.MarshalEWBK(byteOrder)
		if err != nil {
			return nil, err
		}

		output = append(output, dataByte...)
	}

	return output, nil
}
