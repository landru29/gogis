package ewkb

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"
)

// Marshal converts Geometry to EWKB array of bytes.
func Marshal(geoShape Marshaler) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	err := NewEncoder(buffer).Encode(geoShape)

	return buffer.Bytes(), err
}

// Encoder is a Extended Well Known Byte encoder.
type Encoder struct {
	writer    io.Writer
	byteOrder binary.ByteOrder
}

// NewDecoder creates a EWKB decoder.
func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{
		writer:    hex.NewEncoder(writer),
		byteOrder: binary.LittleEndian,
	}
}

func (e *Encoder) Encode(geoShape Marshaler) error {
	output := make([]byte, 1+size32bit)

	output[0] = map[binary.ByteOrder]byte{
		binary.BigEndian:    0,
		binary.LittleEndian: 1,
	}[e.byteOrder]

	withSRID := uint32(0)

	srid := geoShape.SystemReferenceID()
	if srid != nil {
		sridBytes := make([]byte, size32bit)

		e.byteOrder.PutUint32(sridBytes, uint32(*srid))

		output = append(output, sridBytes...) //nolint: makezero

		withSRID = ewkbSRID
	}

	e.byteOrder.PutUint32(output[1:], geoShape.Layout().Uint32()+uint32(geoShape.Type())+withSRID)

	if _, err := e.writer.Write(output); err != nil {
		return err
	}

	data, err := geoShape.MarshalEWBK(e.byteOrder)
	if err != nil {
		return err
	}

	_, err = io.Copy(e.writer, bytes.NewBuffer(data))

	return err
}
