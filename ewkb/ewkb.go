// Package ewkb decodes Extended Well-Known Byte format.
//
// EWKB is encoded in hexadecimal. There are 2 parts:
//
//   - The header
//
//   - The data
//
// # HEADER
//
// Byte 0: 0 means big endian, 1 means little endian.
//
// Bytes 1-4:
//
//   - bit 64 means with Z.
//
//   - bit 63 means with M.
//
//   - bit 62 means with SRID (System Reference ID).
//
//   - bits 0-61 are for the type of the geometry.
//
// If SRID bit is 1, then the 4 following bytes are the SRID (32bit unsigned integer).
//
// After that, come the data part.
//
// # DATA
//
//   - Point: depending on the the format (XY, XYM or XYZM), there are 2, 3 or 4 float64
//     (8 bytes).
//   - Linestring: It's an array of Point. First of all, there is a uint32 (4 bytes) for
//     the dimension of the array. The points are read as below.
package ewkb

// SystemReferenceID is the identifier of the system reference for projection.
type SystemReferenceID uint32

// GeometryType is the type of the geometry (bits 0-61 of the bytes 1-4 of the header).
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

// Unmarshaler is the byte array to Geometry converter.
type Unmarshaler interface {
	// UnmarshalEWBK is to extract Geometry information from the EWKB record.
	UnmarshalEWBK(ExtendedWellKnownBytes) error
}

// Marshaler is the Geometry to byte array converter.
type Marshaler interface {
	// MarshalEWBK must only generate the data part of the EWKB (not the header part).
	MarshalEWBK(ExtendedWellKnownBytesHeader) ([]byte, error)

	// Header builds a header record (used to generate the first bytes of the EWKB).
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
