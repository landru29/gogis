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

func TestMultiPointType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypeMultiPoint, ewkb.MultiPoint{}.Type())
}

func TestMultiPointUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		var multipoint ewkb.MultiPoint

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"0400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiPoint),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiPoint{
			Points: []ewkb.Point{
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
						'm': 5,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 42.17,
						'z': 4,
						'm': 5,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 71.17,
						'z': 4,
						'm': 5,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
						'm': 5,
					},
				},
			},
		})
	})

	t.Run("XYZ", func(t *testing.T) {
		var multipoint ewkb.MultiPoint

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"0400000001010000807B14AE47E1DA51C07B14AE47E15A454000000000000010400101000080EC51B81E856B31C0F6285C8FC215454000000000000010400101000080EC51B81E856B31C07B14AE47E1CA5140000000000000104001010000807B14AE47E1DA51C07B14AE47E15A45400000000000001040",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiPoint),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiPoint{
			Points: []ewkb.Point{
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 42.17,
						'z': 4,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 71.17,
						'z': 4,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
					},
				},
			},
		})
	})

	t.Run("XY", func(t *testing.T) {
		var multipoint ewkb.MultiPoint

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"0400000001010000007B14AE47E1DA51C07B14AE47E15A45400101000000EC51B81E856B31C0F6285C8FC21545400101000000EC51B81E856B31C07B14AE47E1CA514001010000007B14AE47E1DA51C07B14AE47E15A4540",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiPoint),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiPoint{
			Points: []ewkb.Point{
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 42.17,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 71.17,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
					},
				},
			},
		})
	})

	t.Run("wrong type", func(t *testing.T) {
		var multipoint ewkb.MultiPoint

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"0400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeTin),
			),
		)
		require.Error(t, err)
		assert.ErrorIs(t, ewkb.ErrWrongGeometryType, err)
	})
}

func TestMultiPointMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		point := ewkb.MultiPoint{
			Points: []ewkb.Point{
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
						'm': 5,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 42.17,
						'z': 4,
						'm': 5,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 71.17,
						'z': 4,
						'm': 5,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
						'm': 5,
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("0400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440"),
			hex.EncodeToString(data),
		)
	})
	t.Run("XYZ", func(t *testing.T) {
		point := ewkb.MultiPoint{
			Points: []ewkb.Point{
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 42.17,
						'z': 4,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 71.17,
						'z': 4,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
						'z': 4,
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("0400000001010000807B14AE47E1DA51C07B14AE47E15A454000000000000010400101000080EC51B81E856B31C0F6285C8FC215454000000000000010400101000080EC51B81E856B31C07B14AE47E1CA5140000000000000104001010000807B14AE47E1DA51C07B14AE47E15A45400000000000001040"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XY", func(t *testing.T) {
		point := ewkb.MultiPoint{
			Points: []ewkb.Point{
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 42.17,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -17.42,
						'y': 71.17,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': -71.42,
						'y': 42.71,
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("0400000001010000007B14AE47E1DA51C07B14AE47E15A45400101000000EC51B81E856B31C0F6285C8FC21545400101000000EC51B81E856B31C07B14AE47E1CA514001010000007B14AE47E1DA51C07B14AE47E15A4540"),
			hex.EncodeToString(data),
		)
	})
}

func TestMultiPointUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		binary      string
		expected    ewkb.Geometry
	}{
		{
			geometry:    &ewkb.MultiPoint{},
			strGeometry: "MULTIPOINT((-71.42 42.71 4 5),(-17.42 42.17 4 5),(-17.42 71.17 4 5),(-71.42 42.71 4 5))",
			binary:      "01040000C00400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
			expected: &ewkb.MultiPoint{
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPoint{},
			strGeometry: "MULTIPOINT((-71.42 42.71 4 5),(-17.42 42.17 4 5),(-17.42 71.17 4 5),(-71.42 42.71 4 5)),4326",
			binary:      "01040000E0E61000000400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
			expected: &ewkb.MultiPoint{
				SRID: &srid,
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPoint{},
			strGeometry: "MULTIPOINT((-71.42 42.71 4),(-17.42 42.17 4),(-17.42 71.17 4),(-71.42 42.71 4))",
			binary:      "01040000800400000001010000807B14AE47E1DA51C07B14AE47E15A454000000000000010400101000080EC51B81E856B31C0F6285C8FC215454000000000000010400101000080EC51B81E856B31C07B14AE47E1CA5140000000000000104001010000807B14AE47E1DA51C07B14AE47E15A45400000000000001040",
			expected: &ewkb.MultiPoint{
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPoint{},
			strGeometry: "MULTIPOINT((-71.42 42.71 4),(-17.42 42.17 4),(-17.42 71.17 4),(-71.42 42.71 4)), 4326",
			binary:      "01040000A0E61000000400000001010000807B14AE47E1DA51C07B14AE47E15A454000000000000010400101000080EC51B81E856B31C0F6285C8FC215454000000000000010400101000080EC51B81E856B31C07B14AE47E1CA5140000000000000104001010000807B14AE47E1DA51C07B14AE47E15A45400000000000001040",
			expected: &ewkb.MultiPoint{
				SRID: &srid,
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPoint{},
			strGeometry: "MULTIPOINT((-71.42 42.71),(-17.42 42.17),(-17.42 71.17),(-71.42 42.71))",
			binary:      "01040000000400000001010000007B14AE47E1DA51C07B14AE47E15A45400101000000EC51B81E856B31C0F6285C8FC21545400101000000EC51B81E856B31C07B14AE47E1CA514001010000007B14AE47E1DA51C07B14AE47E15A4540",
			expected: &ewkb.MultiPoint{
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPoint{},
			strGeometry: "MULTIPOINT((-71.42 42.71),(-17.42 42.17),(-17.42 71.17),(-71.42 42.71)), 4326",
			binary:      "0104000020E61000000400000001010000007B14AE47E1DA51C07B14AE47E15A45400101000000EC51B81E856B31C0F6285C8FC21545400101000000EC51B81E856B31C07B14AE47E1CA514001010000007B14AE47E1DA51C07B14AE47E15A4540",
			expected: &ewkb.MultiPoint{
				SRID: &srid,
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
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

func TestMultiPointMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
	}{
		{
			strGeometry: "MULTIPOINT((-71.42 42.71 4 5),(-17.42 42.17 4 5),(-17.42 71.17 4 5),(-71.42 42.71 4 5))",
			expected:    "01040000C00400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
			geometry: &ewkb.MultiPoint{
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOINT((-71.42 42.71 4 5),(-17.42 42.17 4 5),(-17.42 71.17 4 5),(-71.42 42.71 4 5)),4326",
			expected:    "01040000E0E61000000400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440",
			geometry: &ewkb.MultiPoint{
				SRID: &srid,
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOINT((-71.42 42.71 4),(-17.42 42.17 4),(-17.42 71.17 4),(-71.42 42.71 4))",
			expected:    "01040000800400000001010000807B14AE47E1DA51C07B14AE47E15A454000000000000010400101000080EC51B81E856B31C0F6285C8FC215454000000000000010400101000080EC51B81E856B31C07B14AE47E1CA5140000000000000104001010000807B14AE47E1DA51C07B14AE47E15A45400000000000001040",
			geometry: &ewkb.MultiPoint{
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOINT((-71.42 42.71 4),(-17.42 42.17 4),(-17.42 71.17 4),(-71.42 42.71 4)), 4326",
			expected:    "01040000A0E61000000400000001010000807B14AE47E1DA51C07B14AE47E15A454000000000000010400101000080EC51B81E856B31C0F6285C8FC215454000000000000010400101000080EC51B81E856B31C07B14AE47E1CA5140000000000000104001010000807B14AE47E1DA51C07B14AE47E15A45400000000000001040",
			geometry: &ewkb.MultiPoint{
				SRID: &srid,
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOINT((-71.42 42.71),(-17.42 42.17),(-17.42 71.17),(-71.42 42.71))",
			expected:    "01040000000400000001010000007B14AE47E1DA51C07B14AE47E15A45400101000000EC51B81E856B31C0F6285C8FC21545400101000000EC51B81E856B31C07B14AE47E1CA514001010000007B14AE47E1DA51C07B14AE47E15A4540",
			geometry: &ewkb.MultiPoint{
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOINT((-71.42 42.71),(-17.42 42.17),(-17.42 71.17),(-71.42 42.71)), 4326",
			expected:    "0104000020E61000000400000001010000007B14AE47E1DA51C07B14AE47E15A45400101000000EC51B81E856B31C0F6285C8FC21545400101000000EC51B81E856B31C07B14AE47E1CA514001010000007B14AE47E1DA51C07B14AE47E15A4540",
			geometry: &ewkb.MultiPoint{
				SRID: &srid,
				Points: []ewkb.Point{
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 42.17,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -17.42,
							'y': 71.17,
						},
					},
					{
						Coordinate: ewkb.Coordinate{
							'x': -71.42,
							'y': 42.71,
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
