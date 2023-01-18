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
		if _, err := io.ReadFull(dataStream, floatBytes); err != nil {
			return err
		}

		out[idx] = float64FromBytes(floatBytes, byteOrder)
	}

	*p = out

	return nil
}

func (p point) isNull() bool {
	for _, value := range p {
		if value != value {
			return true
		}
	}

	return false
}
