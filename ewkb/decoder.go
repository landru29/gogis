package ewkb

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io"
	"math"
)

const (
	bigEndian    = 0
	littleEndian = 1

	ewkbZ    uint32 = 0x80000000
	ewkbM    uint32 = 0x40000000
	ewkbSRID uint32 = 0x20000000

	size32bit = 4
	size64bit = 8
)

// ExtendedWellKnownBytesHeader is the header from EWKB data.
type ExtendedWellKnownBytesHeader struct {
	SRID      *SystemReferenceID
	ByteOrder binary.ByteOrder
	Type      GeometryType
	Layout    Layout
}

// ExtendedWellKnownBytes is the EWKB.
type ExtendedWellKnownBytes struct {
	ExtendedWellKnownBytesHeader
	DataStream io.Reader
	IsNil      bool
}

// ReadUint32 reads 32-bit unsigned integer from the stream.
func (e ExtendedWellKnownBytes) ReadUint32() (uint32, error) {
	data := make([]byte, size32bit)

	_, err := e.DataStream.Read(data)

	return e.ByteOrder.Uint32(data), err
}

// ReadFloat64 reads 64-bit float from the stream.
func (e ExtendedWellKnownBytes) ReadFloat64() (float64, error) {
	data := make([]byte, size64bit)

	_, err := e.DataStream.Read(data)

	bits := e.ByteOrder.Uint64(data)

	return math.Float64frombits(bits), err
}

// DecodeHeader decodes EWKB header.
func DecodeHeader(reader io.Reader) (*ExtendedWellKnownBytes, error) {
	var firstByte = make([]byte, 1)
	_, err := reader.Read(firstByte)
	if err == io.EOF {
		return &ExtendedWellKnownBytes{IsNil: true}, nil
	}
	if err != nil {
		return nil, err
	}

	var byteOrder binary.ByteOrder

	switch firstByte[0] {
	case bigEndian:
		byteOrder = binary.BigEndian
	case littleEndian:
		byteOrder = binary.LittleEndian
	default:
		return nil, ErrWrongByteOrder
	}

	var controlByte = make([]byte, 4)
	if _, err := reader.Read(controlByte); err != nil {
		return nil, err
	}

	header := byteOrder.Uint32(controlByte)

	output := ExtendedWellKnownBytes{
		ExtendedWellKnownBytesHeader: ExtendedWellKnownBytesHeader{
			Layout:    Layout((header & (ewkbZ | ewkbM)) >> 30), //nolint: gomnd
			Type:      GeometryType(header &^ (ewkbZ | ewkbM | ewkbSRID)),
			ByteOrder: byteOrder,
		},
		DataStream: reader,
	}

	if header&ewkbSRID != 0 {
		srid, err := output.ReadUint32()
		if err != nil {
			return nil, err
		}

		output.SRID = (*SystemReferenceID)(&srid)
	}

	return &output, nil
}

// Unmarshal converts EWKB array of bytes to Geometry.
func Unmarshal(geoShape Unmarshaler, value interface{}) error {
	if value == nil {
		return nil
	}

	if strData, ok := value.(string); ok {
		return Unmarshal(geoShape, []byte(strData))
	}

	dataByte, ok := value.([]byte)
	if !ok {
		return ErrIncompatibleFormat
	}

	return NewDecoder(bytes.NewBuffer(dataByte)).Decode(geoShape)
}

// Decoder is a Extended Well Known Byte decoder.
type Decoder struct {
	reader io.Reader
}

// NewDecoder creates a EWKB decoder.
func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{
		reader: hex.NewDecoder(reader),
	}
}

// Decode decodes to a Geometry.
func (d *Decoder) Decode(geoShape Unmarshaler) error {
	record, err := DecodeHeader(d.reader)
	if err != nil {
		return err
	}

	return geoShape.UnmarshalEWBK(*record)
}
