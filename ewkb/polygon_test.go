package ewkb_test

import (
	"encoding/binary"
	"encoding/hex"
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
