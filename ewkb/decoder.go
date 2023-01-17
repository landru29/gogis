package ewkb

import (
	"bytes"
	"encoding/binary"
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
	DataStream *bytes.Buffer
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

// Marshal converts ExtendedWellKnownBytes to EWKB array of bytes.
func (e ExtendedWellKnownBytes) Marshal() []byte {
	output := make([]byte, 1+size32bit)

	output[0] = map[binary.ByteOrder]byte{
		binary.BigEndian:    0,
		binary.LittleEndian: 1,
	}[e.ByteOrder]

	withSRID := uint32(0)

	if e.SRID != nil {
		sridBytes := make([]byte, size32bit)

		e.ByteOrder.PutUint32(sridBytes, uint32(*e.SRID))

		output = append(output, sridBytes...) //nolint: makezero

		withSRID = ewkbSRID
	}

	e.ByteOrder.PutUint32(output[1:], e.Layout.Uint32()+uint32(e.Type)+withSRID)

	output = append(output, e.DataStream.Bytes()...) //nolint: makezero

	return toHex(output)
}

// DecodeHeader decodes EWKB header.
func DecodeHeader( /*data interface{}*/ reader io.Reader) (*ExtendedWellKnownBytes, error) {
	if data == nil {
		return &ExtendedWellKnownBytes{IsNil: true}, nil
	}

	if strData, ok := data.(string); ok {
		return DecodeHeader([]byte(strData))
	}

	dataByte, ok := data.([]byte)
	if !ok {
		return nil, ErrIncompatibleFormat
	}

	decodedData, err := fromHex(dataByte)
	if err != nil {
		return nil, err
	}

	var byteOrder binary.ByteOrder

	switch decodedData[0] {
	case bigEndian:
		byteOrder = binary.BigEndian
	case littleEndian:
		byteOrder = binary.LittleEndian
	default:
		return nil, ErrWrongByteOrder
	}

	header := byteOrder.Uint32(decodedData[1:5])

	output := ExtendedWellKnownBytes{
		ExtendedWellKnownBytesHeader: ExtendedWellKnownBytesHeader{
			Layout:    Layout((header & (ewkbZ | ewkbM)) >> 30), //nolint: gomnd
			Type:      GeometryType(header &^ (ewkbZ | ewkbM | ewkbSRID)),
			ByteOrder: byteOrder,
		},
		DataStream: bytes.NewBuffer(decodedData[5:]),
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

// Marshal converts Geometry to EWKB array of bytes.
func Marshal(geoShape Marshaler) ([]byte, error) {
	header := geoShape.Header()

	if header.ByteOrder == nil {
		header.ByteOrder = binary.LittleEndian
	}

	data, err := geoShape.MarshalEWBK(header)
	if err != nil {
		return nil, err
	}

	return ExtendedWellKnownBytes{
		ExtendedWellKnownBytesHeader: header,
		DataStream:                   bytes.NewBuffer(data),
	}.Marshal(), nil
}

// Unmarshal converts EWKB array of bytes to Geometry.
func Unmarshal(geoShape Unmarshaler, value interface{}) error {
	if value == nil {
		return nil
	}

	record, err := DecodeHeader(value)
	if err != nil {
		return err
	}

	return geoShape.UnmarshalEWBK(*record)
}
