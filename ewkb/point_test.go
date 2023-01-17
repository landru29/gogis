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

func TestPointType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypePoint, ewkb.Point{}.Type())
}

func TestPointUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		var point ewkb.Point

		err := (&point).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"3CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePoint),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, point.Coordinates, map[byte]float64{
			'x': -71.060316,
			'y': 48.432044,
			'z': 10,
			'm': 30,
		})
	})

	t.Run("XYZ", func(t *testing.T) {
		var point ewkb.Point

		err := (&point).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"3CDBA337DCC351C06D37C1374D3748400000000000002440",
				withLayout(ewkb.Layout(2)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePoint),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, point.Coordinates, map[byte]float64{
			'x': -71.060316,
			'y': 48.432044,
			'z': 10,
		})
	})

	t.Run("XY", func(t *testing.T) {
		var point ewkb.Point

		err := (&point).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"3CDBA337DCC351C06D37C1374D374840",
				withLayout(ewkb.Layout(0)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePoint),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, point.Coordinates, map[byte]float64{
			'x': -71.060316,
			'y': 48.432044,
		})
	})

	t.Run("wrong type", func(t *testing.T) {
		var point ewkb.Point

		err := (&point).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"3CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40",
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

func TestPointMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		point := ewkb.Point{
			Coordinates: map[byte]float64{
				'x': -71.060316,
				'y': 48.432044,
				'z': 10,
				'm': 30,
			},
		}

		data, err := point.MarshalEWBK(
			newExtendedWellKnownBytesHeader(
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePoint),
			),
		)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("3CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XYZ", func(t *testing.T) {
		point := ewkb.Point{
			Coordinates: map[byte]float64{
				'x': -71.060316,
				'y': 48.432044,
				'z': 10,
			},
		}

		data, err := point.MarshalEWBK(
			newExtendedWellKnownBytesHeader(
				withLayout(ewkb.Layout(2)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePoint),
			),
		)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("3CDBA337DCC351C06D37C1374D3748400000000000002440"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XY", func(t *testing.T) {
		point := ewkb.Point{
			Coordinates: map[byte]float64{
				'x': -71.060316,
				'y': 48.432044,
			},
		}

		data, err := point.MarshalEWBK(
			newExtendedWellKnownBytesHeader(
				withLayout(ewkb.Layout(0)),
				withByteOrder(binary.LittleEndian),
				withSRID(ewkb.SystemReferenceWGS84),
				withType(ewkb.GeometryTypePoint),
			),
		)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("3CDBA337DCC351C06D37C1374D374840"),
			hex.EncodeToString(data),
		)
	})
}

func TestPointUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		binary      string
		expected    ewkb.Geometry
	}{
		{
			geometry:    &ewkb.Point{},
			strGeometry: "POINT ZM(-71.060316 48.432044 10 30)",
			binary:      "01010000C03CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40",
			expected: &ewkb.Point{
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
					'm': 30.0,
				},
			},
		},
		{
			geometry:    &ewkb.Point{},
			strGeometry: "POINT ZM(-71.060316 48.432044 10 30), 4326",
			binary:      "01010000E0E61000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40",
			expected: &ewkb.Point{
				SRID: &srid,
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
					'm': 30.0,
				},
			},
		},
		{
			geometry:    &ewkb.Point{},
			strGeometry: "POINT Z(-71.060316 48.432044 10)",
			binary:      "01010000803CDBA337DCC351C06D37C1374D3748400000000000002440",
			expected: &ewkb.Point{
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
				},
			},
		},
		{
			geometry:    &ewkb.Point{},
			strGeometry: "POINT Z(-71.060316 48.432044 10), 4326",
			binary:      "01010000A0E61000003CDBA337DCC351C06D37C1374D3748400000000000002440",
			expected: &ewkb.Point{
				SRID: &srid,
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
				},
			},
		},
		{
			geometry:    &ewkb.Point{},
			strGeometry: "POINT (-71.060316 48.432044)",
			binary:      "01010000003CDBA337DCC351C06D37C1374D374840",
			expected: &ewkb.Point{
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
				},
			},
		},
		{
			geometry:    &ewkb.Point{},
			strGeometry: "POINT (-71.060316 48.432044), 4326",
			binary:      "0101000020E61000003CDBA337DCC351C06D37C1374D374840",
			expected: &ewkb.Point{
				SRID: &srid,
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
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

func TestPointMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
		srid        *ewkb.SystemReferenceID
	}{
		{
			strGeometry: "POINT ZM(-71.060316 48.432044 10 30)",
			expected:    "01010000C03CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40",
			geometry: &ewkb.Point{
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
					'm': 30.0,
				},
			},
		},
		{
			strGeometry: "POINT ZM(-71.060316 48.432044 10 30), 4326",
			expected:    "01010000E0E61000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40",
			geometry: &ewkb.Point{
				SRID: &srid,
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
					'm': 30.0,
				},
			},
		},
		{
			strGeometry: "POINT Z(-71.060316 48.432044 10)",
			expected:    "01010000803CDBA337DCC351C06D37C1374D3748400000000000002440",
			geometry: &ewkb.Point{
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
				},
			},
		},
		{
			strGeometry: "POINT Z(-71.060316 48.432044 10), 4326",
			expected:    "01010000A0E61000003CDBA337DCC351C06D37C1374D3748400000000000002440",
			geometry: &ewkb.Point{
				SRID: &srid,
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
				},
			},
		},
		{
			strGeometry: "POINT (-71.060316 48.432044)",
			expected:    "01010000003CDBA337DCC351C06D37C1374D374840",
			geometry: &ewkb.Point{
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
				},
			},
		},
		{
			strGeometry: "POINT (-71.060316 48.432044), 4326",
			expected:    "0101000020E61000003CDBA337DCC351C06D37C1374D374840",
			geometry: &ewkb.Point{
				SRID: &srid,
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
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
