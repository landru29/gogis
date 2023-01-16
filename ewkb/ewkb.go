// Package ewkb decodes Extended Well-Known Byte format.
package ewkb

// SystemReferenceID is the identifier of the system reference for projection.
type SystemReferenceID uint32

// GeometryType is the type of the geometry.
type GeometryType uint8

const (
	// SystemReferenceWGS84 stands for GCS WGS 84.
	SystemReferenceWGS84 SystemReferenceID = 4326

	// SystemReferenceUTMZone stands for UTM Zone 17N NAD 27.
	SystemReferenceUTMZone SystemReferenceID = 26717

	// SystemReferenceTennesseeZone stands for SPCS Tennessee Zone NAD 83.
	SystemReferenceTennesseeZone SystemReferenceID = 6576

	// GeometryTypePoint stands for point.
	GeometryTypePoint GeometryType = 1

	// GeometryTypeLineString stands for lineString.
	GeometryTypeLineString GeometryType = 2

	// GeometryTypePolygon stands for polygon.
	GeometryTypePolygon GeometryType = 3

	// GeometryTypeMultiPoint stands for multiPoint.
	GeometryTypeMultiPoint GeometryType = 4

	// GeometryTypeMultiLineString stands for multiLineString.
	GeometryTypeMultiLineString GeometryType = 5

	// GeometryTypeMultiPolygon stands for multiPolygon.
	GeometryTypeMultiPolygon GeometryType = 6

	// GeometryTypeGeometryCollection stands for geometryCollection.
	GeometryTypeGeometryCollection GeometryType = 7

	// GeometryTypePolyhedralSurface stands for polyhedralSurface.
	GeometryTypePolyhedralSurface GeometryType = 15

	// GeometryTypeTin stands for tin.
	GeometryTypeTin GeometryType = 16

	// GeometryTypeTriangle stands for triangle.
	GeometryTypeTriangle GeometryType = 17
)

// Unmarshaler is the byte array to EWKB converter.
type Unmarshaler interface {
	UnmarshalEWBK(ExtendedWellKnownBytes) error
}

// Marshaler is the EWKB to byte array converter.
type Marshaler interface {
	MarshalEWBK(ExtendedWellKnownBytesHeader) ([]byte, error)
	Header() ExtendedWellKnownBytesHeader
}

// Geometry is a geometrical shape.
type Geometry interface {
	Type() GeometryType
	Unmarshaler
	Marshaler
}

// WithSRID converts SystemReferenceID to pointer.
func WithSRID(srid SystemReferenceID) *SystemReferenceID {
	return &srid
}
