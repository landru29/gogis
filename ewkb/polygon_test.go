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

func TestPolygonType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypePolygon, ewkb.Polygon{}.Type())
}

func TestPolygonUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		var polygon ewkb.Polygon

		err := (&polygon).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"01000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePolygon),
			),
		)
		require.NoError(t, err)
		assert.Len(t, polygon.CoordinateGroup, 1)
		assert.Len(t, polygon.CoordinateGroup[0], 4)
		assert.Equal(t, polygon.CoordinateGroup[0][0], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
			'z': 4,
			'm': 5,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][1], ewkb.Coordinate{
			'x': -17.42,
			'y': 42.17,
			'z': 4,
			'm': 5,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][2], ewkb.Coordinate{
			'x': -17.42,
			'y': 71.17,
			'z': 4,
			'm': 5,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][3], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
			'z': 4,
			'm': 5,
		})
	})

	t.Run("XYZ", func(t *testing.T) {
		var polygon ewkb.Polygon

		err := (&polygon).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"01000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A45400000000000001040",
				withLayout(ewkb.Layout(2)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePolygon),
			),
		)
		require.NoError(t, err)
		assert.Len(t, polygon.CoordinateGroup, 1)
		assert.Len(t, polygon.CoordinateGroup[0], 4)
		assert.Equal(t, polygon.CoordinateGroup[0][0], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
			'z': 4,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][1], ewkb.Coordinate{
			'x': -17.42,
			'y': 42.17,
			'z': 4,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][2], ewkb.Coordinate{
			'x': -17.42,
			'y': 71.17,
			'z': 4,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][3], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
			'z': 4,
		})
	})

	t.Run("XY", func(t *testing.T) {
		var polygon ewkb.Polygon

		err := (&polygon).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"01000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A4540",
				withLayout(ewkb.Layout(0)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePolygon),
			),
		)
		require.NoError(t, err)
		assert.Len(t, polygon.CoordinateGroup, 1)
		assert.Len(t, polygon.CoordinateGroup[0], 4)
		assert.Equal(t, polygon.CoordinateGroup[0][0], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][1], ewkb.Coordinate{
			'x': -17.42,
			'y': 42.17,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][2], ewkb.Coordinate{
			'x': -17.42,
			'y': 71.17,
		})
		assert.Equal(t, polygon.CoordinateGroup[0][3], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
		})
	})

	t.Run("wrong type", func(t *testing.T) {
		var polygon ewkb.Polygon

		err := (&polygon).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"01000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A4540",
				withLayout(ewkb.Layout(0)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePoint),
			),
		)
		require.Error(t, err)
		assert.ErrorIs(t, ewkb.ErrWrongGeometryType, err)
	})
}

func TestPolygonMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		polygon := ewkb.Polygon{
			CoordinateGroup: ewkb.CoordinateGroup{
				{
					{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
						'm': 5,
					},
					{
						'x': -17.42,
						'y': 42.17,
						'z': 4,
						'm': 5,
					},
					{
						'x': -17.42,
						'y': 71.17,
						'z': 4,
						'm': 5,
					},
					{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
						'm': 5,
					},
				},
			},
		}

		data, err := polygon.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("01000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XYZ", func(t *testing.T) {
		polygon := ewkb.Polygon{
			CoordinateGroup: ewkb.CoordinateGroup{
				{
					{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
					},
					{
						'x': -17.42,
						'y': 42.17,
						'z': 4,
					},
					{
						'x': -17.42,
						'y': 71.17,
						'z': 4,
					},
					{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
					},
				},
			},
		}

		data, err := polygon.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("01000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A45400000000000001040"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XY", func(t *testing.T) {
		polygon := ewkb.Polygon{
			CoordinateGroup: ewkb.CoordinateGroup{
				{
					{
						'x': -71.42,
						'y': 42.71,
					},
					{
						'x': -17.42,
						'y': 42.17,
					},
					{
						'x': -17.42,
						'y': 71.17,
					},
					{
						'x': -71.42,
						'y': 42.71,
					},
				},
			},
		}

		data, err := polygon.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("01000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A4540"),
			hex.EncodeToString(data),
		)
	})
}

func TestPolygonUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		binary      string
		expected    ewkb.Geometry
	}{
		{
			geometry:    &ewkb.Polygon{},
			strGeometry: "POLYGON ZM((-71.42 42.71 4 5,-17.42 42.17 4 5,-17.42 71.17 4 5,-71.42 42.71 4 5),(1 2 3 4,4 5 6 7,7 8 9 0,1 2 3 4))",
			binary:      "01030000C002000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144004000000000000000000F03F0000000000000040000000000000084000000000000010400000000000001040000000000000144000000000000018400000000000001C400000000000001C40000000000000204000000000000022400000000000000000000000000000F03F000000000000004000000000000008400000000000001040",
			expected: &ewkb.Polygon{
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
							'm': 7,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
							'm': 0,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.Polygon{},
			strGeometry: "POLYGON ZM((-71.42 42.71 4 5,-17.42 42.17 4 5,-17.42 71.17 4 5,-71.42 42.71 4 5),(1 2 3 4,4 5 6 7,7 8 9 0,1 2 3 4)), 4326",
			binary:      "01030000E0E610000002000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144004000000000000000000F03F0000000000000040000000000000084000000000000010400000000000001040000000000000144000000000000018400000000000001C400000000000001C40000000000000204000000000000022400000000000000000000000000000F03F000000000000004000000000000008400000000000001040",
			expected: &ewkb.Polygon{
				SRID: &srid,
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
							'm': 7,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
							'm': 0,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.Polygon{},
			strGeometry: "POLYGON Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3))",
			binary:      "010300008002000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A4540000000000000104004000000000000000000F03F000000000000004000000000000008400000000000001040000000000000144000000000000018400000000000001C4000000000000020400000000000002240000000000000F03F00000000000000400000000000000840",
			expected: &ewkb.Polygon{
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.Polygon{},
			strGeometry: "POLYGON Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3)),4326",
			binary:      "01030000A0E610000002000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A4540000000000000104004000000000000000000F03F000000000000004000000000000008400000000000001040000000000000144000000000000018400000000000001C4000000000000020400000000000002240000000000000F03F00000000000000400000000000000840",
			expected: &ewkb.Polygon{
				SRID: &srid,
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.Polygon{},
			strGeometry: "POLYGON ((-71.42 42.71,-17.42 42.17,-17.42 71.17,-71.42 42.71),(1 2,4 5,7 8,1 2))",
			binary:      "010300000002000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A454004000000000000000000F03F0000000000000040000000000000104000000000000014400000000000001C400000000000002040000000000000F03F0000000000000040",
			expected: &ewkb.Polygon{
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
						},
						{
							'x': -17.42,
							'y': 42.17,
						},
						{
							'x': -17.42,
							'y': 71.17,
						},
						{
							'x': -71.42,
							'y': 42.71,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
						},
						{
							'x': 4,
							'y': 5,
						},
						{
							'x': 7,
							'y': 8,
						},
						{
							'x': 1,
							'y': 2,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.Polygon{},
			strGeometry: "POLYGON ((-71.42 42.71,-17.42 42.17,-17.42 71.17,-71.42 42.71),(1 2,4 5,7 8,1 2)),4326",
			binary:      "0103000020E610000002000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A454004000000000000000000F03F0000000000000040000000000000104000000000000014400000000000001C400000000000002040000000000000F03F0000000000000040",
			expected: &ewkb.Polygon{
				SRID: &srid,
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
						},
						{
							'x': -17.42,
							'y': 42.17,
						},
						{
							'x': -17.42,
							'y': 71.17,
						},
						{
							'x': -71.42,
							'y': 42.71,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
						},
						{
							'x': 4,
							'y': 5,
						},
						{
							'x': 7,
							'y': 8,
						},
						{
							'x': 1,
							'y': 2,
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

			assert.Equal(t, fixture.expected, fixture.geometry)
		})
	}
}

func TestPolygonMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
	}{
		{
			strGeometry: "POLYGON ZM((-71.42 42.71 4 5,-17.42 42.17 4 5,-17.42 71.17 4 5,-71.42 42.71 4 5),(1 2 3 4,4 5 6 7,7 8 9 0,1 2 3 4))",
			expected:    "01030000C002000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144004000000000000000000F03F0000000000000040000000000000084000000000000010400000000000001040000000000000144000000000000018400000000000001C400000000000001C40000000000000204000000000000022400000000000000000000000000000F03F000000000000004000000000000008400000000000001040",
			geometry: &ewkb.Polygon{
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
							'm': 7,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
							'm': 0,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
					},
				},
			},
		},
		{
			strGeometry: "POLYGON ZM((-71.42 42.71 4 5,-17.42 42.17 4 5,-17.42 71.17 4 5,-71.42 42.71 4 5),(1 2 3 4,4 5 6 7,7 8 9 0,1 2 3 4)), 4326",
			expected:    "01030000E0E610000002000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144004000000000000000000F03F0000000000000040000000000000084000000000000010400000000000001040000000000000144000000000000018400000000000001C400000000000001C40000000000000204000000000000022400000000000000000000000000000F03F000000000000004000000000000008400000000000001040",
			geometry: &ewkb.Polygon{
				SRID: &srid,
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
							'm': 7,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
							'm': 0,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
					},
				},
			},
		},
		{
			strGeometry: "POLYGON Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3))",
			expected:    "010300008002000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A4540000000000000104004000000000000000000F03F000000000000004000000000000008400000000000001040000000000000144000000000000018400000000000001C4000000000000020400000000000002240000000000000F03F00000000000000400000000000000840",
			geometry: &ewkb.Polygon{
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
						},
					},
				},
			},
		},
		{
			strGeometry: "POLYGON Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3)),4326",
			expected:    "01030000A0E610000002000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A4540000000000000104004000000000000000000F03F000000000000004000000000000008400000000000001040000000000000144000000000000018400000000000001C4000000000000020400000000000002240000000000000F03F00000000000000400000000000000840",
			geometry: &ewkb.Polygon{
				SRID: &srid,
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
						},
					},
				},
			},
		},
		{
			strGeometry: "POLYGON ((-71.42 42.71,-17.42 42.17,-17.42 71.17,-71.42 42.71),(1 2,4 5,7 8,1 2))",
			expected:    "010300000002000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A454004000000000000000000F03F0000000000000040000000000000104000000000000014400000000000001C400000000000002040000000000000F03F0000000000000040",
			geometry: &ewkb.Polygon{
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
						},
						{
							'x': -17.42,
							'y': 42.17,
						},
						{
							'x': -17.42,
							'y': 71.17,
						},
						{
							'x': -71.42,
							'y': 42.71,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
						},
						{
							'x': 4,
							'y': 5,
						},
						{
							'x': 7,
							'y': 8,
						},
						{
							'x': 1,
							'y': 2,
						},
					},
				},
			},
		},
		{
			strGeometry: "POLYGON ((-71.42 42.71,-17.42 42.17,-17.42 71.17,-71.42 42.71),(1 2,4 5,7 8,1 2)),4326",
			expected:    "0103000020E610000002000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A454004000000000000000000F03F0000000000000040000000000000104000000000000014400000000000001C400000000000002040000000000000F03F0000000000000040",
			geometry: &ewkb.Polygon{
				SRID: &srid,
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
						},
						{
							'x': -17.42,
							'y': 42.17,
						},
						{
							'x': -17.42,
							'y': 71.17,
						},
						{
							'x': -71.42,
							'y': 42.71,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
						},
						{
							'x': 4,
							'y': 5,
						},
						{
							'x': 7,
							'y': 8,
						},
						{
							'x': 1,
							'y': 2,
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
