package ewkb_test

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/landru29/gogis/ewkb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeometricalCollectionType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypeGeometryCollection, ewkb.GeometryCollection{}.Type())
}

func TestGeometricalCollectionUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		collection := ewkb.NewGeometryCollection()

		err := (collection).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"0200000001010000C0000000000000004000000000000008400000000000001040000000000000144001020000C00200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeGeometryCollection),
			),
		)
		require.NoError(t, err)
		assert.Len(t, collection.Collection, 2)
		assert.Equal(t, collection.Collection[0], &ewkb.Point{
			Coordinate: ewkb.Coordinate{
				'x': 2,
				'y': 3,
				'z': 4,
				'm': 5,
			},
		})
		assert.Equal(t, collection.Collection[1], &ewkb.LineString{
			CoordinateSet: ewkb.CoordinateSet{
				ewkb.Coordinate{
					'x': 2,
					'y': 3,
					'z': 4,
					'm': 5,
				},
				ewkb.Coordinate{
					'x': 3,
					'y': 4,
					'z': 5,
					'm': 6,
				},
			},
		})
	})

	t.Run("XYZ", func(t *testing.T) {
		collection := ewkb.NewGeometryCollection()

		err := (collection).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"020000000101000080000000000000004000000000000008400000000000001040010200008002000000000000000000004000000000000008400000000000001040000000000000084000000000000010400000000000001440",
				withLayout(ewkb.Layout(2)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeGeometryCollection),
			),
		)
		require.NoError(t, err)
		assert.Len(t, collection.Collection, 2)
		assert.Equal(t, collection.Collection[0], &ewkb.Point{
			Coordinate: ewkb.Coordinate{
				'x': 2,
				'y': 3,
				'z': 4,
			},
		})
		assert.Equal(t, collection.Collection[1], &ewkb.LineString{
			CoordinateSet: ewkb.CoordinateSet{
				ewkb.Coordinate{
					'x': 2,
					'y': 3,
					'z': 4,
				},
				ewkb.Coordinate{
					'x': 3,
					'y': 4,
					'z': 5,
				},
			},
		})
	})

	t.Run("XY", func(t *testing.T) {
		collection := ewkb.NewGeometryCollection()

		err := (collection).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"020000000101000000000000000000004000000000000008400102000000020000000000000000000040000000000000084000000000000008400000000000001040",
				withLayout(ewkb.Layout(0)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeGeometryCollection),
			),
		)
		require.NoError(t, err)
		assert.Len(t, collection.Collection, 2)
		assert.Equal(t, collection.Collection[0], &ewkb.Point{
			Coordinate: ewkb.Coordinate{
				'x': 2,
				'y': 3,
			},
		})
		assert.Equal(t, collection.Collection[1], &ewkb.LineString{
			CoordinateSet: ewkb.CoordinateSet{
				ewkb.Coordinate{
					'x': 2,
					'y': 3,
				},
				ewkb.Coordinate{
					'x': 3,
					'y': 4,
				},
			},
		})
	})

	t.Run("wrong type", func(t *testing.T) {
		var linestring ewkb.LineString

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"0200000001010000C0000000000000004000000000000008400000000000001040000000000000144001020000C00200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeTin),
			),
		)
		require.Error(t, err)
		assert.ErrorIs(t, err, ewkb.ErrWrongGeometryType)
	})
}

func TestGeometryCollectionMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		collection := ewkb.GeometryCollection{
			Collection: []ewkb.Geometry{
				&ewkb.Point{
					Coordinate: ewkb.Coordinate{
						'x': 2,
						'y': 3,
						'z': 4,
						'm': 5,
					},
				},
				&ewkb.LineString{
					CoordinateSet: ewkb.CoordinateSet{
						ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
							'm': 5,
						},
						ewkb.Coordinate{
							'x': 3,
							'y': 4,
							'z': 5,
							'm': 6,
						},
					},
				},
			},
		}

		data, err := collection.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("0200000001010000C0000000000000004000000000000008400000000000001040000000000000144001020000C00200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XYZ", func(t *testing.T) {
		collection := ewkb.GeometryCollection{
			Collection: []ewkb.Geometry{
				&ewkb.Point{
					Coordinate: ewkb.Coordinate{
						'x': 2,
						'y': 3,
						'z': 4,
					},
				},
				&ewkb.LineString{
					CoordinateSet: ewkb.CoordinateSet{
						ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
						},
						ewkb.Coordinate{
							'x': 3,
							'y': 4,
							'z': 5,
						},
					},
				},
			},
		}

		data, err := collection.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("020000000101000080000000000000004000000000000008400000000000001040010200008002000000000000000000004000000000000008400000000000001040000000000000084000000000000010400000000000001440"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XY", func(t *testing.T) {
		collection := ewkb.GeometryCollection{
			Collection: []ewkb.Geometry{
				&ewkb.Point{
					Coordinate: ewkb.Coordinate{
						'x': 2,
						'y': 3,
					},
				},
				&ewkb.LineString{
					CoordinateSet: ewkb.CoordinateSet{
						ewkb.Coordinate{
							'x': 2,
							'y': 3,
						},
						ewkb.Coordinate{
							'x': 3,
							'y': 4,
						},
					},
				},
			},
		}

		data, err := collection.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("020000000101000000000000000000004000000000000008400102000000020000000000000000000040000000000000084000000000000008400000000000001040"),
			hex.EncodeToString(data),
		)
	})
}

func TestGeometryCollectionUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    *ewkb.GeometryCollection
		strGeometry string
		binary      string
		expected    *ewkb.GeometryCollection
	}{
		{
			geometry:    ewkb.NewGeometryCollection(),
			strGeometry: "GEOMETRYCOLLECTION ZM( POINT(2 3 4 5), LINESTRING(2 3 4 5, 3 4 5 6))",
			binary:      "01070000C00200000001010000C0000000000000004000000000000008400000000000001040000000000000144001020000C00200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840",
			expected: &ewkb.GeometryCollection{
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
							'm': 5,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
								'z': 4,
								'm': 5,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
								'z': 5,
								'm': 6,
							},
						},
					},
				},
			},
		},
		{
			geometry:    ewkb.NewGeometryCollection(),
			strGeometry: "GEOMETRYCOLLECTION ZM( POINT(2 3 4 5), LINESTRING(2 3 4 5, 3 4 5 6)), 4326",
			binary:      "01070000E0E61000000200000001010000C0000000000000004000000000000008400000000000001040000000000000144001020000C00200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840",
			expected: &ewkb.GeometryCollection{
				SRID: &srid,
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
							'm': 5,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
								'z': 4,
								'm': 5,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
								'z': 5,
								'm': 6,
							},
						},
					},
				},
			},
		},
		{
			geometry:    ewkb.NewGeometryCollection(),
			strGeometry: "GEOMETRYCOLLECTION Z( POINT(2 3 4), LINESTRING(2 3 4, 3 4 5))",
			binary:      "0107000080020000000101000080000000000000004000000000000008400000000000001040010200008002000000000000000000004000000000000008400000000000001040000000000000084000000000000010400000000000001440",
			expected: &ewkb.GeometryCollection{
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
								'z': 4,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
								'z': 5,
							},
						},
					},
				},
			},
		},
		{
			geometry:    ewkb.NewGeometryCollection(),
			strGeometry: "GEOMETRYCOLLECTION Z( POINT(2 3 4), LINESTRING(2 3 4, 3 4 5)), 4326",
			binary:      "01070000A0E6100000020000000101000080000000000000004000000000000008400000000000001040010200008002000000000000000000004000000000000008400000000000001040000000000000084000000000000010400000000000001440",
			expected: &ewkb.GeometryCollection{
				SRID: &srid,
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
								'z': 4,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
								'z': 5,
							},
						},
					},
				},
			},
		},
		{
			geometry:    ewkb.NewGeometryCollection(),
			strGeometry: "GEOMETRYCOLLECTION( POINT(2 3), LINESTRING(2 3, 3 4))",
			binary:      "0107000000020000000101000000000000000000004000000000000008400102000000020000000000000000000040000000000000084000000000000008400000000000001040",
			expected: &ewkb.GeometryCollection{
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
							},
						},
					},
				},
			},
		},
		{
			geometry:    ewkb.NewGeometryCollection(),
			strGeometry: "GEOMETRYCOLLECTION( POINT(2 3), LINESTRING(2 3, 3 4)), 4326",
			binary:      "0107000020E6100000020000000101000000000000000000004000000000000008400102000000020000000000000000000040000000000000084000000000000008400000000000001040",
			expected: &ewkb.GeometryCollection{
				SRID: &srid,
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
							},
						},
					},
				},
			},
		},
	}

	for idx, elt := range fixtures {
		fixture := elt

		t.Run(fmt.Sprintf("%d - %s", idx, fixture.strGeometry), func(t *testing.T) {
			assert.NoError(t, ewkb.Unmarshal(fixture.geometry, fixture.binary))

			assert.Equal(t, fixture.expected.SRID, fixture.geometry.SRID)
			assert.Equal(t, fixture.expected.Collection, fixture.geometry.Collection)
		})
	}
}

func TestGeometryCollectionMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
	}{
		{
			strGeometry: "GEOMETRYCOLLECTION ZM( POINT(2 3 4 5), LINESTRING(2 3 4 5, 3 4 5 6))",
			expected:    "01070000C00200000001010000C0000000000000004000000000000008400000000000001040000000000000144001020000C00200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840",
			geometry: &ewkb.GeometryCollection{
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
							'm': 5,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
								'z': 4,
								'm': 5,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
								'z': 5,
								'm': 6,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "GEOMETRYCOLLECTION ZM( POINT(2 3 4 5), LINESTRING(2 3 4 5, 3 4 5 6)), 4326",
			expected:    "01070000E0E61000000200000001010000C0000000000000004000000000000008400000000000001040000000000000144001020000C00200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840",
			geometry: &ewkb.GeometryCollection{
				SRID: &srid,
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
							'm': 5,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
								'z': 4,
								'm': 5,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
								'z': 5,
								'm': 6,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "GEOMETRYCOLLECTION Z( POINT(2 3 4), LINESTRING(2 3 4, 3 4 5))",
			expected:    "0107000080020000000101000080000000000000004000000000000008400000000000001040010200008002000000000000000000004000000000000008400000000000001040000000000000084000000000000010400000000000001440",
			geometry: &ewkb.GeometryCollection{
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
								'z': 4,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
								'z': 5,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "GEOMETRYCOLLECTION Z( POINT(2 3 4), LINESTRING(2 3 4, 3 4 5)), 4326",
			expected:    "01070000A0E6100000020000000101000080000000000000004000000000000008400000000000001040010200008002000000000000000000004000000000000008400000000000001040000000000000084000000000000010400000000000001440",
			geometry: &ewkb.GeometryCollection{
				SRID: &srid,
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
							'z': 4,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
								'z': 4,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
								'z': 5,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "GEOMETRYCOLLECTION( POINT(2 3), LINESTRING(2 3, 3 4))",
			expected:    "0107000000020000000101000000000000000000004000000000000008400102000000020000000000000000000040000000000000084000000000000008400000000000001040",
			geometry: &ewkb.GeometryCollection{
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "GEOMETRYCOLLECTION( POINT(2 3), LINESTRING(2 3, 3 4)), 4326",
			expected:    "0107000020E6100000020000000101000000000000000000004000000000000008400102000000020000000000000000000040000000000000084000000000000008400000000000001040",
			geometry: &ewkb.GeometryCollection{
				SRID: &srid,
				Collection: []ewkb.Geometry{
					&ewkb.Point{
						Coordinate: ewkb.Coordinate{
							'x': 2,
							'y': 3,
						},
					},
					&ewkb.LineString{
						CoordinateSet: ewkb.CoordinateSet{
							ewkb.Coordinate{
								'x': 2,
								'y': 3,
							},
							ewkb.Coordinate{
								'x': 3,
								'y': 4,
							},
						},
					},
				},
			},
		},
	}

	for idx, elt := range fixtures {
		fixture := elt

		t.Run(fmt.Sprintf("%d - %s", idx, fixture.strGeometry), func(t *testing.T) {
			output, err := ewkb.Marshal(fixture.geometry)
			assert.NoError(t, err)

			assert.Equal(t, strings.ToLower(fixture.expected), string(output))
		})
	}
}
