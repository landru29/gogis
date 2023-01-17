package ewkb

// GeometryType is the type of the geometry (bits 0-61 of the bytes 1-4 of the header).
type GeometryType uint8

const (
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

	// GeometryTypeCircle stands for circle.
	GeometryTypeCircle GeometryType = 8

	// GeometryTypeCompound stands for compound.
	GeometryTypeCompound GeometryType = 9

	// GeometryTypeCurvePoly stands for curvepoly.
	GeometryTypeCurvePoly GeometryType = 10

	// GeometryTypeMultiCurve stands for multicurve.
	GeometryTypeMultiCurve GeometryType = 11

	// GeometryTypeMultiSurface stands for multisurface.
	GeometryTypeMultiSurface GeometryType = 12

	// GeometryTypePolyhedralSurface stands for polyhedralSurface.
	GeometryTypePolyhedralSurface GeometryType = 13

	// GeometryTypeTriangle stands for triangle.
	GeometryTypeTriangle GeometryType = 14

	// GeometryTypeTin stands for tin.
	GeometryTypeTin GeometryType = 15
)

// Geometry is a geometrical shape.
type Geometry interface {
	Unmarshaler
	Marshaler
}
