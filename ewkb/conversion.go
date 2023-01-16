package ewkb

import (
	"encoding/binary"
	"math"
)

// Float64FromBytes convert a 8-bytes array to float64.
func float64FromBytes(bytes []byte, byteOrder binary.ByteOrder) float64 {
	bits := byteOrder.Uint64(bytes)

	float := math.Float64frombits(bits)

	return float
}

// Float64Bytes converts float64 to 8-bytes array.
func float64Bytes(float float64, byteOrder binary.ByteOrder) []byte {
	bits := math.Float64bits(float)

	bytes := make([]byte, 8) //nolint: gomnd

	byteOrder.PutUint64(bytes, bits)

	return bytes
}
