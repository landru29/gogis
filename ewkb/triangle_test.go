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

func TestTriangleType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypeTriangle, ewkb.Triangle{}.Type())
}

func TestTriangleUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		var polygon ewkb.Triangle

		err := (&polygon).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"01000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeTriangle),
			),
		)
		require.NoError(t, err)
		assert.Len(t, polygon.CoordinateSet, 4)
		assert.Equal(t, polygon.CoordinateSet[0], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
			'z': 4,
			'm': 5,
		})
		assert.Equal(t, polygon.CoordinateSet[1], ewkb.Coordinate{
			'x': -17.42,
			'y': 42.17,
			'z': 4,
			'm': 5,
		})
		assert.Equal(t, polygon.CoordinateSet[2], ewkb.Coordinate{
			'x': -17.42,
			'y': 71.17,
			'z': 4,
			'm': 5,
		})
		assert.Equal(t, polygon.CoordinateSet[3], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
			'z': 4,
			'm': 5,
		})
	})

	t.Run("XYZ", func(t *testing.T) {
		var polygon ewkb.Triangle

		err := (&polygon).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"01000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A45400000000000001040",
				withLayout(ewkb.Layout(2)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeTriangle),
			),
		)
		require.NoError(t, err)
		assert.Len(t, polygon.CoordinateSet, 4)
		assert.Equal(t, polygon.CoordinateSet[0], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
			'z': 4,
		})
		assert.Equal(t, polygon.CoordinateSet[1], ewkb.Coordinate{
			'x': -17.42,
			'y': 42.17,
			'z': 4,
		})
		assert.Equal(t, polygon.CoordinateSet[2], ewkb.Coordinate{
			'x': -17.42,
			'y': 71.17,
			'z': 4,
		})
		assert.Equal(t, polygon.CoordinateSet[3], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
			'z': 4,
		})
	})

	t.Run("XY", func(t *testing.T) {
		var polygon ewkb.Triangle

		err := (&polygon).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"01000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A4540",
				withLayout(ewkb.Layout(0)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeTriangle),
			),
		)
		require.NoError(t, err)
		assert.Len(t, polygon.CoordinateSet, 4)
		assert.Equal(t, polygon.CoordinateSet[0], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
		})
		assert.Equal(t, polygon.CoordinateSet[1], ewkb.Coordinate{
			'x': -17.42,
			'y': 42.17,
		})
		assert.Equal(t, polygon.CoordinateSet[2], ewkb.Coordinate{
			'x': -17.42,
			'y': 71.17,
		})
		assert.Equal(t, polygon.CoordinateSet[3], ewkb.Coordinate{
			'x': -71.42,
			'y': 42.71,
		})
	})

	t.Run("wrong type", func(t *testing.T) {
		var polygon ewkb.Triangle

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
		assert.ErrorIs(t, err, ewkb.ErrWrongGeometryType)
	})
}

func TestTriangleMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		polygon := ewkb.Triangle{
			CoordinateSet: ewkb.CoordinateSet{
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
		polygon := ewkb.Triangle{
			CoordinateSet: ewkb.CoordinateSet{
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
		polygon := ewkb.Triangle{
			CoordinateSet: ewkb.CoordinateSet{
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

func TestTriangleUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		binary      string
		expected    ewkb.Geometry
	}{
		{
			geometry:    &ewkb.Triangle{},
			strGeometry: "TRIANGLE ZM((-71.42 42.71 4 5,-17.42 42.17 4 5,-17.42 71.17 4 5,-71.42 42.71 4 5),(1 2 3 4,4 5 6 7,7 8 9 0,1 2 3 4))",
			binary:      "01110000C001000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
			expected: &ewkb.Triangle{
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			geometry:    &ewkb.Triangle{},
			strGeometry: "TRIANGLE ZM((-71.42 42.71 4 5,-17.42 42.17 4 5,-17.42 71.17 4 5,-71.42 42.71 4 5),(1 2 3 4,4 5 6 7,7 8 9 0,1 2 3 4)), 4326",
			binary:      "01110000E0E610000001000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
			expected: &ewkb.Triangle{
				SRID: &srid,
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			geometry:    &ewkb.Triangle{},
			strGeometry: "TRIANGLE Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3))",
			binary:      "011100008001000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A45400000000000001040",
			expected: &ewkb.Triangle{
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			geometry:    &ewkb.Triangle{},
			strGeometry: "TRIANGLE Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3)),4326",
			binary:      "01110000A0E610000001000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A45400000000000001040",
			expected: &ewkb.Triangle{
				SRID: &srid,
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			geometry:    &ewkb.Triangle{},
			strGeometry: "TRIANGLE ((-71.42 42.71,-17.42 42.17,-17.42 71.17,-71.42 42.71))",
			binary:      "011100000001000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A4540",
			expected: &ewkb.Triangle{
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			geometry:    &ewkb.Triangle{},
			strGeometry: "TRIANGLE ((-71.42 42.71,-17.42 42.17,-17.42 71.17,-71.42 42.71)),4326",
			binary:      "0111000020E610000001000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A4540",
			expected: &ewkb.Triangle{
				SRID: &srid,
				CoordinateSet: ewkb.CoordinateSet{
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

func TestTriangleMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
	}{
		{
			strGeometry: "TRIANGLE ZM((-71.42 42.71 4 5,-17.42 42.17 4 5,-17.42 71.17 4 5,-71.42 42.71 4 5),(1 2 3 4,4 5 6 7,7 8 9 0,1 2 3 4))",
			expected:    "01110000C001000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
			geometry: &ewkb.Triangle{
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			strGeometry: "TRIANGLE ZM((-71.42 42.71 4 5,-17.42 42.17 4 5,-17.42 71.17 4 5,-71.42 42.71 4 5),(1 2 3 4,4 5 6 7,7 8 9 0,1 2 3 4)), 4326",
			expected:    "01110000E0E610000001000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
			geometry: &ewkb.Triangle{
				SRID: &srid,
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			strGeometry: "TRIANGLE Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3))",
			expected:    "011100008001000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A45400000000000001040",
			geometry: &ewkb.Triangle{
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			strGeometry: "TRIANGLE Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3)),4326",
			expected:    "01110000A0E610000001000000040000007B14AE47E1DA51C07B14AE47E15A45400000000000001040EC51B81E856B31C0F6285C8FC21545400000000000001040EC51B81E856B31C07B14AE47E1CA514000000000000010407B14AE47E1DA51C07B14AE47E15A45400000000000001040",
			geometry: &ewkb.Triangle{
				SRID: &srid,
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			strGeometry: "TRIANGLE ((-71.42 42.71,-17.42 42.17,-17.42 71.17,-71.42 42.71))",
			expected:    "011100000001000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A4540",
			geometry: &ewkb.Triangle{
				CoordinateSet: ewkb.CoordinateSet{
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
		},
		{
			strGeometry: "TRIANGLE ((-71.42 42.71,-17.42 42.17,-17.42 71.17,-71.42 42.71)),4326",
			expected:    "0111000020E610000001000000040000007B14AE47E1DA51C07B14AE47E15A4540EC51B81E856B31C0F6285C8FC2154540EC51B81E856B31C07B14AE47E1CA51407B14AE47E1DA51C07B14AE47E15A4540",
			geometry: &ewkb.Triangle{
				SRID: &srid,
				CoordinateSet: ewkb.CoordinateSet{
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
