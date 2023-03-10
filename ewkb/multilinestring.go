package ewkb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// MultiLineString is a MULTILINESTRING in database.
//
// A MultiLineString is a collection of LineStrings. A MultiLineString is
// closed if each of its elements is closed.
type MultiLineString struct {
	SRID        *SystemReferenceID
	LineStrings []LineString
}

// Type implements the Geometry interface.
func (m MultiLineString) Type() GeometryType {
	return GeometryTypeMultiLineString
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (m *MultiLineString) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != m.Type() {
		return fmt.Errorf("%w: found %d, expected %d", ErrWrongGeometryType, record.Type, m.Type())
	}

	m.SRID = record.SRID

	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	m.LineStrings = make([]LineString, size)

	for idx := range m.LineStrings {
		lineStr := &LineString{}
		if err := (&Decoder{reader: record.DataStream}).Decode(lineStr); err != nil {
			return err
		}

		m.LineStrings[idx] = *lineStr
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (m MultiLineString) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	byteOrder.PutUint32(size, uint32(len(m.LineStrings)))
	output = append(output, size...)

	buffer := bytes.NewBuffer(nil)

	for _, pnt := range m.LineStrings {
		if err := (&Encoder{writer: buffer, byteOrder: byteOrder}).Encode(pnt); err != nil {
			return nil, err
		}
	}

	output = append(output, buffer.Bytes()...)

	return output, nil
}

// SystemReferenceID implements the Marshaler interface.
func (m MultiLineString) SystemReferenceID() *SystemReferenceID {
	return m.SRID
}

// Layout implements the Marshaler interface.
func (m MultiLineString) Layout() Layout {
	if len(m.LineStrings) > 0 {
		return m.LineStrings[0].Layout()
	}

	return layoutXY
}
