package ewkb

// Linestring is a set of lines.
type Linestring struct {
	SRID   *SystemReferenceID
	Points []Point
}

// Type implements the Geometry interface.
func (l Linestring) Type() GeometryType {
	return GeometryTypeLineString
}

// UnmarshalEWBK implements the Unmarshaler interface.
func (l *Linestring) UnmarshalEWBK(record ExtendedWellKnownBytes) error {
	if record.Type != l.Type() {
		return ErrWrongGeometryType
	}

	size, err := record.ReadUint32()
	if err != nil {
		return err
	}

	l.SRID = record.SRID
	record.Type = GeometryTypePoint

	l.Points = make([]Point, size)
	for idx := range l.Points {
		err := (&(l.Points[idx])).UnmarshalEWBK(record)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalEWBK implements the Marshaler interface.
func (l Linestring) MarshalEWBK(header ExtendedWellKnownBytesHeader) ([]byte, error) {
	output := []byte{}

	size := make([]byte, size32bit)

	header.ByteOrder.PutUint32(size, uint32(len(l.Points)))
	output = append(output, size...)

	for _, point := range l.Points {
		dataByte, err := point.MarshalEWBK(header)
		if err != nil {
			return nil, err
		}

		output = append(output, dataByte...)
	}

	return output, nil
}

// Header implements the Marshaler interface.
func (l Linestring) Header() ExtendedWellKnownBytesHeader {
	var header ExtendedWellKnownBytesHeader

	if len(l.Points) > 0 {
		header = l.Points[0].Header()
	}

	header.SRID = l.SRID
	header.Type = GeometryTypeLineString

	return header
}
