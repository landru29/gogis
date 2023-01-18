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

func TestCircularStringType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypeCircularString, ewkb.CircularString{}.Type())
}

func TestCircularStringUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		var linestring ewkb.CircularString

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeCircularString),
			),
		)
		require.NoError(t, err)
		assert.Len(t, linestring.CoordinateSet, 3)
		assert.Equal(t, linestring.CoordinateSet[0], ewkb.Coordinate{
			'x': -71.060316,
			'y': 48.432044,
			'z': 10,
			'm': 30,
		})
		assert.Equal(t, linestring.CoordinateSet[1], ewkb.Coordinate{
			'x': 5,
			'y': 6,
			'z': 7,
			'm': 8,
		})
		assert.Equal(t, linestring.CoordinateSet[2], ewkb.Coordinate{
			'x': 1,
			'y': 2,
			'z': 3,
			'm': 4,
		})
	})

	t.Run("XYZ", func(t *testing.T) {
		var linestring ewkb.CircularString

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"030000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40000000000000F03F00000000000000400000000000000840",
				withLayout(ewkb.Layout(2)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeCircularString),
			),
		)
		require.NoError(t, err)
		assert.Len(t, linestring.CoordinateSet, 3)
		assert.Equal(t, linestring.CoordinateSet[0], ewkb.Coordinate{
			'x': -71.060316,
			'y': 48.432044,
			'z': 10,
		})
		assert.Equal(t, linestring.CoordinateSet[1], ewkb.Coordinate{
			'x': 5,
			'y': 6,
			'z': 7,
		})
		assert.Equal(t, linestring.CoordinateSet[2], ewkb.Coordinate{
			'x': 1,
			'y': 2,
			'z': 3,
		})
	})

	t.Run("XY", func(t *testing.T) {
		var linestring ewkb.CircularString

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"030000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840000000000000F03F0000000000000040",
				withLayout(ewkb.Layout(0)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeCircularString),
			),
		)
		require.NoError(t, err)
		assert.Len(t, linestring.CoordinateSet, 3)
		assert.Equal(t, linestring.CoordinateSet[0], ewkb.Coordinate{
			'x': -71.060316,
			'y': 48.432044,
		})
		assert.Equal(t, linestring.CoordinateSet[1], ewkb.Coordinate{
			'x': 5,
			'y': 6,
		})
		assert.Equal(t, linestring.CoordinateSet[2], ewkb.Coordinate{
			'x': 1,
			'y': 2,
		})
	})
	t.Run("wrong type", func(t *testing.T) {
		var linestring ewkb.CircularString

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"030000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840000000000000F03F0000000000000040",
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

func TestCircularStringMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		linestring := ewkb.CircularString{
			CoordinateSet: []ewkb.Coordinate{
				{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10,
					'm': 30,
				},
				{
					'x': 5,
					'y': 6,
					'z': 7,
					'm': 8,
				},
				{
					'x': 1,
					'y': 2,
					'z': 3,
					'm': 4,
				},
			},
		}

		data, err := linestring.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XYZ", func(t *testing.T) {
		linestring := ewkb.CircularString{
			CoordinateSet: []ewkb.Coordinate{
				{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10,
				},
				{
					'x': 5,
					'y': 6,
					'z': 7,
				},
				{
					'x': 1,
					'y': 2,
					'z': 3,
				},
			},
		}

		data, err := linestring.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("030000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40000000000000F03F00000000000000400000000000000840"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XY", func(t *testing.T) {
		linestring := ewkb.CircularString{
			CoordinateSet: []ewkb.Coordinate{
				{
					'x': -71.060316,
					'y': 48.432044,
				},
				{
					'x': 5,
					'y': 6,
				},
				{
					'x': 1,
					'y': 2,
				},
			},
		}

		data, err := linestring.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("030000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840000000000000F03F0000000000000040"),
			hex.EncodeToString(data),
		)
	})
}

func TestCircularStringUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		binary      string
		expected    ewkb.Geometry
	}{
		{
			geometry:    &ewkb.CircularString{},
			strGeometry: "CIRCULARSTRING ZM(-71.060316 48.432044 10 30,5 6 7 8,1 2 3 4)",
			binary:      "01080000C0030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040",
			expected: &ewkb.CircularString{
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
						'm': 30,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
						'm': 8,
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
		{
			geometry:    &ewkb.CircularString{},
			strGeometry: "CIRCULARSTRING ZM(-71.060316 48.432044 10 30,5 6 7 8,1 2 3 4), 4326",
			binary:      "01080000E0E6100000030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040",
			expected: &ewkb.CircularString{
				SRID: &srid,
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
						'm': 30,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
						'm': 8,
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
		{
			geometry:    &ewkb.CircularString{},
			strGeometry: "CIRCULARSTRING Z(-71.060316 48.432044 10,5 6 7,1 2 3)",
			binary:      "0108000080030000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40000000000000F03F00000000000000400000000000000840",
			expected: &ewkb.CircularString{
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
					},
					{
						'x': 1,
						'y': 2,
						'z': 3,
					},
				},
			},
		},
		{
			geometry:    &ewkb.CircularString{},
			strGeometry: "CIRCULARSTRING Z(-71.060316 48.432044 10,5 6 7,1 2 3), 4326",
			binary:      "01080000A0E6100000030000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40000000000000F03F00000000000000400000000000000840",
			expected: &ewkb.CircularString{
				SRID: &srid,
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
					},
					{
						'x': 1,
						'y': 2,
						'z': 3,
					},
				},
			},
		},
		{
			geometry:    &ewkb.CircularString{},
			strGeometry: "CIRCULARSTRING(-71.060316 48.432044,5 6,1 2)",
			binary:      "0108000000030000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840000000000000F03F0000000000000040",
			expected: &ewkb.CircularString{
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
					},
					{
						'x': 5,
						'y': 6,
					},
					{
						'x': 1,
						'y': 2,
					},
				},
			},
		},
		{
			geometry:    &ewkb.CircularString{},
			strGeometry: "CIRCULARSTRING(-71.060316 48.432044,5 6,1 2), 4326",
			binary:      "0108000020E6100000030000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840000000000000F03F0000000000000040",
			expected: &ewkb.CircularString{
				SRID: &srid,
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
					},
					{
						'x': 5,
						'y': 6,
					},
					{
						'x': 1,
						'y': 2,
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

func TestCircularStringMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
	}{
		{
			strGeometry: "CIRCULARSTRING ZM(-71.060316 48.432044 10 30,5 6 7 8,1 2 3 4)",
			expected:    "01080000C0030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040",
			geometry: &ewkb.CircularString{
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
						'm': 30,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
						'm': 8,
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
		{
			strGeometry: "CIRCULARSTRING ZM(-71.060316 48.432044 10 30,5 6 7 8,1 2 3 4), 4326",
			expected:    "01080000E0E6100000030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040",
			geometry: &ewkb.CircularString{
				SRID: &srid,
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
						'm': 30,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
						'm': 8,
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
		{
			strGeometry: "CIRCULARSTRING Z(-71.060316 48.432044 10,5 6 7,1 2 3)",
			expected:    "0108000080030000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40000000000000F03F00000000000000400000000000000840",
			geometry: &ewkb.CircularString{
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
					},
					{
						'x': 1,
						'y': 2,
						'z': 3,
					},
				},
			},
		},
		{
			strGeometry: "CIRCULARSTRING Z(-71.060316 48.432044 10,5 6 7,1 2 3), 4326",
			expected:    "01080000A0E6100000030000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40000000000000F03F00000000000000400000000000000840",
			geometry: &ewkb.CircularString{
				SRID: &srid,
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
					},
					{
						'x': 1,
						'y': 2,
						'z': 3,
					},
				},
			},
		},
		{
			strGeometry: "CIRCULARSTRING(-71.060316 48.432044,5 6,1 2)",
			expected:    "0108000000030000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840000000000000F03F0000000000000040",
			geometry: &ewkb.CircularString{
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
					},
					{
						'x': 5,
						'y': 6,
					},
					{
						'x': 1,
						'y': 2,
					},
				},
			},
		},
		{
			strGeometry: "CIRCULARSTRING(-71.060316 48.432044,5 6,1 2), 4326",
			expected:    "0108000020E6100000030000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840000000000000F03F0000000000000040",
			geometry: &ewkb.CircularString{
				SRID: &srid,
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
					},
					{
						'x': 5,
						'y': 6,
					},
					{
						'x': 1,
						'y': 2,
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
