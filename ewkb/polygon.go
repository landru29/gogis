package ewkb

type Polygon struct {
	SRID  *SystemReferenceID
	Rings []Linestring
}

// Type implements the Geometry interface.
func (p Polygon) Type() GeometryType {
	return GeometryTypePolygon
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (p *Polygon) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != p.Type() {
		return ErrWrongGeometryType
	}

	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	p.SRID = record.SRID

	p.Rings = make([]Linestring, size)
	for idx := range p.Rings {
		record.Type = GeometryTypeLineString
		err := (&(p.Rings[idx])).UnmarshalEWBK(record)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (p Polygon) MarshalEWBK(header ExtendedWellKnownBytesHeader) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	header.ByteOrder.PutUint32(size, uint32(len(p.Rings)))
	output = append(output, size...)

	for _, rings := range p.Rings {
		dataByte, err := rings.MarshalEWBK(header)
		if err != nil {
			return nil, err
		}

		output = append(output, dataByte...)
	}

	return output, nil
}

// Header implements the Marshaler interface.
func (p Polygon) Header() ExtendedWellKnownBytesHeader {
	var header ExtendedWellKnownBytesHeader

	if len(p.Rings) > 0 {
		header = p.Rings[0].Header()
	}

	header.SRID = p.SRID
	header.Type = GeometryTypeLineString

	return header
}
