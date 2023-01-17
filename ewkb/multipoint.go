package ewkb

import (
	"bytes"
	"encoding/binary"
)

// Multipoint is a MULTIPOINT in database.
type MultiPoint struct {
	SRID   *SystemReferenceID
	Points []Point
}

// Type implements the Geometry interface.
func (p MultiPoint) Type() GeometryType {
	return GeometryTypeMultiPoint
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (p *MultiPoint) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != p.Type() {
		return ErrWrongGeometryType
	}

	p.SRID = record.SRID

	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	p.Points = make([]Point, size)

	for idx := range p.Points {
		record.Type = GeometryTypePoint

		pnt := &Point{}
		if err := (&Decoder{reader: record.DataStream}).Decode(pnt); err != nil {
			return err
		}

		p.Points[idx] = *pnt
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (p MultiPoint) MarshalEWBK(byteOrder binary.ByteOrder) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	byteOrder.PutUint32(size, uint32(len(p.Points)))
	output = append(output, size...)

	buffer := bytes.NewBuffer(nil)
	for _, pnt := range p.Points {
		if err := (&Encoder{writer: buffer, byteOrder: byteOrder}).Encode(pnt); err != nil {
			return nil, err
		}
	}

	output = append(output, buffer.Bytes()...)

	return output, nil
}

// SystemReferenceID implements the Marshaler interface.
func (p MultiPoint) SystemReferenceID() *SystemReferenceID {
	return p.SRID
}

// Layout implements the Marshaler interface.
func (p MultiPoint) Layout() Layout {
	if len(p.Points) > 0 {
		return p.Points[0].Layout()
	}
	return layoutXY
}
