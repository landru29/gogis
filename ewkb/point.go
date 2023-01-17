package ewkb

import "encoding/binary"

// Point is a lat lng position in database.
type Point struct {
	SRID        *SystemReferenceID
	Coordinates map[byte]float64
}

// Type implements the Geometry interface.
func (p Point) Type() GeometryType {
	return GeometryTypePoint
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (p *Point) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != p.Type() {
		return ErrWrongGeometryType
	}

	var pnt point

	if err := (&pnt).read(record.DataStream, record.Layout.Size(), record.ByteOrder); err != nil {
		return err
	}

	*p = Point{
		SRID:        record.SRID,
		Coordinates: map[byte]float64{},
	}

	for idx, char := range record.Layout.Format() {
		p.Coordinates[byte(char)] = pnt[idx]
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (p Point) MarshalEWBK(header ExtendedWellKnownBytesHeader) ([]byte, error) {
	output := []byte{}

	for _, name := range header.Layout.Format() {
		bytes := float64Bytes(p.Coordinates[byte(name)], header.ByteOrder)
		output = append(output, bytes...)
	}

	return output, nil
}

// Header implements the Marshaler interface.
func (p Point) Header() ExtendedWellKnownBytesHeader {
	indexes := []byte{}
	for idx := range p.Coordinates {
		indexes = append(indexes, idx)
	}

	return ExtendedWellKnownBytesHeader{
		Type:      p.Type(),
		Layout:    newLayoutFrom(indexes),
		ByteOrder: binary.LittleEndian,
		SRID:      p.SRID,
	}
}
