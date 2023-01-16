package ewkb

import (
	"encoding/binary"
	"io"
)

type point []float64

func (p *point) read(dataStream io.Reader, size uint32, byteOrder binary.ByteOrder) error {
	out := make(point, size)

	for idx := 0; idx < int(size); idx++ {
		floatBytes := make([]byte, size64bit)
		if _, err := dataStream.Read(floatBytes); err != nil {
			return err
		}

		out[idx] = Float64FromBytes(floatBytes, byteOrder)
	}

	*p = out

	return nil
}
