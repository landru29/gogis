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

func TestLinestringType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypeLineString, ewkb.Linestring{}.Type())
}

func TestLinestringUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		var linestring ewkb.Linestring

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeLineString),
			),
		)
		require.NoError(t, err)
		assert.Len(t, linestring.CoordinateSet, 2)
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
	})

	t.Run("XYZ", func(t *testing.T) {
		var linestring ewkb.Linestring

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"020000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40",
				withLayout(ewkb.Layout(2)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeLineString),
			),
		)
		require.NoError(t, err)
		assert.Len(t, linestring.CoordinateSet, 2)
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
	})

	t.Run("XY", func(t *testing.T) {
		var linestring ewkb.Linestring

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"020000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840",
				withLayout(ewkb.Layout(0)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypeLineString),
			),
		)
		require.NoError(t, err)
		assert.Len(t, linestring.CoordinateSet, 2)
		assert.Equal(t, linestring.CoordinateSet[0], ewkb.Coordinate{
			'x': -71.060316,
			'y': 48.432044,
		})
		assert.Equal(t, linestring.CoordinateSet[1], ewkb.Coordinate{
			'x': 5,
			'y': 6,
		})
	})
	t.Run("wrong type", func(t *testing.T) {
		var linestring ewkb.Linestring

		err := (&linestring).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040",
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

func TestLinestringMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		linestring := ewkb.Linestring{
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
			},
		}

		data, err := linestring.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XYZ", func(t *testing.T) {
		linestring := ewkb.Linestring{
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
			},
		}

		data, err := linestring.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("020000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XY", func(t *testing.T) {
		linestring := ewkb.Linestring{
			CoordinateSet: []ewkb.Coordinate{
				{
					'x': -71.060316,
					'y': 48.432044,
				},
				{
					'x': 5,
					'y': 6,
				},
			},
		}

		data, err := linestring.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("020000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840"),
			hex.EncodeToString(data),
		)
	})
}

func TestLinestringUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		binary      string
		expected    ewkb.Geometry
	}{
		{
			geometry:    &ewkb.Linestring{},
			strGeometry: "LINESTRING ZM(-71.060316 48.432044 10 30, 5 6 7 8)",
			binary:      "01020000C0020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040",
			expected: &ewkb.Linestring{
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
				},
			},
		},
		{
			geometry:    &ewkb.Linestring{},
			strGeometry: "LINESTRING ZM(-71.060316 48.432044 10 30, 5 6 7 8), 4326",
			binary:      "01020000E0E6100000020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040",
			expected: &ewkb.Linestring{
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
				},
			},
		},
		{
			geometry:    &ewkb.Linestring{},
			strGeometry: "LINESTRING Z(-71.060316 48.432044 10, 5 6 7)",
			binary:      "0102000080020000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40",
			expected: &ewkb.Linestring{
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
				},
			},
		},
		{
			geometry:    &ewkb.Linestring{},
			strGeometry: "LINESTRING Z(-71.060316 48.432044 10, 5 6 7), 4326",
			binary:      "01020000A0E6100000020000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40",
			expected: &ewkb.Linestring{
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
				},
			},
		},
		{
			geometry:    &ewkb.Linestring{},
			strGeometry: "LINESTRING (-71.060316 48.432044, 5 6)",
			binary:      "0102000000020000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840",
			expected: &ewkb.Linestring{
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
					},
					{
						'x': 5,
						'y': 6,
					},
				},
			},
		},
		{
			geometry:    &ewkb.Linestring{},
			strGeometry: "LINESTRING (-71.060316 48.432044, 5 6), 4326",
			binary:      "0102000020E6100000020000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840",
			expected: &ewkb.Linestring{
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

func TestLinestringMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
	}{
		{
			strGeometry: "LINESTRING ZM(-71.060316 48.432044 10 30, 5 6 7 8)",
			expected:    "01020000C0020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040",
			geometry: &ewkb.Linestring{
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
				},
			},
		},
		{
			strGeometry: "LINESTRING ZM(-71.060316 48.432044 10 30, 5 6 7 8), 4326",
			expected:    "01020000E0E6100000020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040",
			geometry: &ewkb.Linestring{
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
				},
			},
		},
		{
			strGeometry: "LINESTRING Z(-71.060316 48.432044 10, 5 6 7)",
			expected:    "0102000080020000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40",
			geometry: &ewkb.Linestring{
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
				},
			},
		},
		{
			strGeometry: "LINESTRING Z(-71.060316 48.432044 10, 5 6 7), 4326",
			expected:    "01020000A0E6100000020000003CDBA337DCC351C06D37C1374D3748400000000000002440000000000000144000000000000018400000000000001C40",
			geometry: &ewkb.Linestring{
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
				},
			},
		},
		{
			strGeometry: "LINESTRING (-71.060316 48.432044, 5 6)",
			expected:    "0102000000020000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840",
			geometry: &ewkb.Linestring{
				CoordinateSet: []ewkb.Coordinate{
					{
						'x': -71.060316,
						'y': 48.432044,
					},
					{
						'x': 5,
						'y': 6,
					},
				},
			},
		},
		{
			strGeometry: "LINESTRING (-71.060316 48.432044, 5 6), 4326",
			expected:    "0102000020E6100000020000003CDBA337DCC351C06D37C1374D37484000000000000014400000000000001840",
			geometry: &ewkb.Linestring{
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
