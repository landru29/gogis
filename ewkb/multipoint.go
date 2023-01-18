package ewkb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Multipoint is a MULTIPOINT in database.
type MultiPoint struct {
	SRID   *SystemReferenceID
	Points []Point
}

// Type implements the Geometry interface.
func (m MultiPoint) Type() GeometryType {
	return GeometryTypeMultiPoint
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (m *MultiPoint) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != m.Type() {
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, m.Type())
	}

	m.SRID = record.SRID

	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	m.Points = make([]Point, size)

	for idx := range m.Points {
		pnt := &Point{}
		if err := (&Decoder{reader: record.DataStream}).Decode(pnt); err != nil {
			return err
		}

		m.Points[idx] = *pnt
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (m MultiPoint) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	byteOrder.PutUint32(size, uint32(len(m.Points)))
	output = append(output, size...)

	buffer := bytes.NewBuffer(nil)
	for _, pnt := range m.Points {
		if err := (&Encoder{writer: buffer, byteOrder: byteOrder}).Encode(pnt); err != nil {
			return nil, err
		}
	}

	output = append(output, buffer.Bytes()...)

	return output, nil
}

// SystemReferenceID implements the Marshaler interface.
func (m MultiPoint) SystemReferenceID() *SystemReferenceID {
	return m.SRID
}

// Layout implements the Marshaler interface.
func (m MultiPoint) Layout() Layout {
	if len(m.Points) > 0 {
		return m.Points[0].Layout()
	}
	return layoutXY
}
