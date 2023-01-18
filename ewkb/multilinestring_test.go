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

func TestMultiLineStringType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypeMultiLineString, ewkb.MultiLineString{}.Type())
}

func TestMultiLineStringUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		var multipoint ewkb.MultiLineString

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"0200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiLineString),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiLineString{
			LineStrings: []ewkb.LineString{
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 42.42,
							'y': -24.24,
							'z': 42.24,
							'm': -24.42,
						},
						{
							'x': 5,
							'y': 6,
							'z': 7,
							'm': 8,
						},
					},
				},
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 142.42,
							'y': -424.24,
							'z': 142.24,
							'm': -124.42,
						},
						{
							'x': 15,
							'y': 16,
							'z': 17,
							'm': 18,
						},
					},
				},
			},
		})
	})

	t.Run("XYZ", func(t *testing.T) {
		var multipoint ewkb.MultiLineString

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"02000000010200008002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540000000000000144000000000000018400000000000001C400102000080020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761400000000000002E4000000000000030400000000000003140",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiLineString),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiLineString{
			LineStrings: []ewkb.LineString{
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 42.42,
							'y': -24.24,
							'z': 42.24,
						},
						{
							'x': 5,
							'y': 6,
							'z': 7,
						},
					},
				},
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 142.42,
							'y': -424.24,
							'z': 142.24,
						},
						{
							'x': 15,
							'y': 16,
							'z': 17,
						},
					},
				},
			},
		})
	})

	t.Run("XY", func(t *testing.T) {
		var multipoint ewkb.MultiLineString

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"02000000010200000002000000F6285C8FC23545403D0AD7A3703D38C0000000000000144000000000000018400102000000020000003D0AD7A370CD6140A4703D0AD7837AC00000000000002E400000000000003040",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiLineString),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiLineString{
			LineStrings: []ewkb.LineString{
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 42.42,
							'y': -24.24,
						},
						{
							'x': 5,
							'y': 6,
						},
					},
				},
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 142.42,
							'y': -424.24,
						},
						{
							'x': 15,
							'y': 16,
						},
					},
				},
			},
		})
	})

	t.Run("wrong type", func(t *testing.T) {
		var multipoint ewkb.MultiLineString

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"02000000010200000002000000F6285C8FC23545403D0AD7A3703D38C0000000000000144000000000000018400102000000020000003D0AD7A370CD6140A4703D0AD7837AC00000000000002E400000000000003040",
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

func TestMultiLineStringMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		point := ewkb.MultiLineString{
			LineStrings: []ewkb.LineString{
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 42.42,
							'y': -24.24,
							'z': 42.24,
							'm': -24.42,
						},
						{
							'x': 5,
							'y': 6,
							'z': 7,
							'm': 8,
						},
					},
				},
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 142.42,
							'y': -424.24,
							'z': 142.24,
							'm': -124.42,
						},
						{
							'x': 15,
							'y': 16,
							'z': 17,
							'm': 18,
						},
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("0200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240"),
			hex.EncodeToString(data),
		)
	})
	t.Run("XYZ", func(t *testing.T) {
		point := ewkb.MultiLineString{
			LineStrings: []ewkb.LineString{
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 42.42,
							'y': -24.24,
							'z': 42.24,
						},
						{
							'x': 5,
							'y': 6,
							'z': 7,
						},
					},
				},
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 142.42,
							'y': -424.24,
							'z': 142.24,
						},
						{
							'x': 15,
							'y': 16,
							'z': 17,
						},
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("02000000010200008002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540000000000000144000000000000018400000000000001C400102000080020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761400000000000002E4000000000000030400000000000003140"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XY", func(t *testing.T) {
		point := ewkb.MultiLineString{
			LineStrings: []ewkb.LineString{
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 42.42,
							'y': -24.24,
						},
						{
							'x': 5,
							'y': 6,
						},
					},
				},
				{
					CoordinateSet: []ewkb.Coordinate{
						{
							'x': 142.42,
							'y': -424.24,
						},
						{
							'x': 15,
							'y': 16,
						},
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("02000000010200000002000000F6285C8FC23545403D0AD7A3703D38C0000000000000144000000000000018400102000000020000003D0AD7A370CD6140A4703D0AD7837AC00000000000002E400000000000003040"),
			hex.EncodeToString(data),
		)
	})
}

func TestMultiLineStringUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		binary      string
		expected    ewkb.Geometry
	}{
		{
			geometry:    &ewkb.MultiLineString{},
			strGeometry: "MULTILINESTRING((42.42 -24.24 42.24 -24.42,5 6 7 8),(142.42 -424.24 142.24 -124.42,15 16 17 18))",
			binary:      "01050000C00200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240",
			expected: &ewkb.MultiLineString{
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
								'm': -24.42,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
								'm': 8,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
								'm': -124.42,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
								'm': 18,
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiLineString{},
			strGeometry: "MULTILINESTRING((42.42 -24.24 42.24 -24.42,5 6 7 8),(142.42 -424.24 142.24 -124.42,15 16 17 18)),4326",
			binary:      "01050000E0E61000000200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240",
			expected: &ewkb.MultiLineString{
				SRID: &srid,
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
								'm': -24.42,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
								'm': 8,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
								'm': -124.42,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
								'm': 18,
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiLineString{},
			strGeometry: "MULTILINESTRING((42.42 -24.24 42.24,5 6 7),(142.42 -424.24 142.24,15 16 17))",
			binary:      "010500008002000000010200008002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540000000000000144000000000000018400000000000001C400102000080020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761400000000000002E4000000000000030400000000000003140",
			expected: &ewkb.MultiLineString{
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiLineString{},
			strGeometry: "MULTILINESTRING((42.42 -24.24 42.24,5 6 7),(142.42 -424.24 142.24,15 16 17)), 4326",
			binary:      "01050000A0E610000002000000010200008002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540000000000000144000000000000018400000000000001C400102000080020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761400000000000002E4000000000000030400000000000003140",
			expected: &ewkb.MultiLineString{
				SRID: &srid,
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiLineString{},
			strGeometry: "MULTILINESTRING((42.42 -24.24,5 6),(142.42 -424.24,15 16))",
			binary:      "010500000002000000010200000002000000F6285C8FC23545403D0AD7A3703D38C0000000000000144000000000000018400102000000020000003D0AD7A370CD6140A4703D0AD7837AC00000000000002E400000000000003040",
			expected: &ewkb.MultiLineString{
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
							},
							{
								'x': 5,
								'y': 6,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
							},
							{
								'x': 15,
								'y': 16,
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiLineString{},
			strGeometry: "MULTILINESTRING((42.42 -24.24,5 6),(142.42 -424.24,15 16)), 4326",
			binary:      "0105000020E610000002000000010200000002000000F6285C8FC23545403D0AD7A3703D38C0000000000000144000000000000018400102000000020000003D0AD7A370CD6140A4703D0AD7837AC00000000000002E400000000000003040",
			expected: &ewkb.MultiLineString{
				SRID: &srid,
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
							},
							{
								'x': 5,
								'y': 6,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
							},
							{
								'x': 15,
								'y': 16,
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

			assert.Equal(t, fixture.expected, fixture.geometry)
		})
	}
}

func TestMultiLineStringMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
	}{
		{
			strGeometry: "MULTILINESTRING((42.42 -24.24 42.24 -24.42,5 6 7 8),(142.42 -424.24 142.24 -124.42,15 16 17 18))",
			expected:    "01050000C00200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240",
			geometry: &ewkb.MultiLineString{
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
								'm': -24.42,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
								'm': 8,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
								'm': -124.42,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
								'm': 18,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTILINESTRING((42.42 -24.24 42.24 -24.42,5 6 7 8),(142.42 -424.24 142.24 -124.42,15 16 17 18)),4326",
			expected:    "01050000E0E61000000200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240",
			geometry: &ewkb.MultiLineString{
				SRID: &srid,
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
								'm': -24.42,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
								'm': 8,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
								'm': -124.42,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
								'm': 18,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTILINESTRING((42.42 -24.24 42.24,5 6 7),(142.42 -424.24 142.24,15 16 17))",
			expected:    "010500008002000000010200008002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540000000000000144000000000000018400000000000001C400102000080020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761400000000000002E4000000000000030400000000000003140",
			geometry: &ewkb.MultiLineString{
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTILINESTRING((42.42 -24.24 42.24,5 6 7),(142.42 -424.24 142.24,15 16 17)), 4326",
			expected:    "01050000A0E610000002000000010200008002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540000000000000144000000000000018400000000000001C400102000080020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761400000000000002E4000000000000030400000000000003140",
			geometry: &ewkb.MultiLineString{
				SRID: &srid,
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTILINESTRING((42.42 -24.24,5 6),(142.42 -424.24,15 16))",
			expected:    "010500000002000000010200000002000000F6285C8FC23545403D0AD7A3703D38C0000000000000144000000000000018400102000000020000003D0AD7A370CD6140A4703D0AD7837AC00000000000002E400000000000003040",
			geometry: &ewkb.MultiLineString{
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
							},
							{
								'x': 5,
								'y': 6,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
							},
							{
								'x': 15,
								'y': 16,
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTILINESTRING((42.42 -24.24,5 6),(142.42 -424.24,15 16)), 4326",
			expected:    "0105000020E610000002000000010200000002000000F6285C8FC23545403D0AD7A3703D38C0000000000000144000000000000018400102000000020000003D0AD7A370CD6140A4703D0AD7837AC00000000000002E400000000000003040",
			geometry: &ewkb.MultiLineString{
				SRID: &srid,
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
							},
							{
								'x': 5,
								'y': 6,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
							},
							{
								'x': 15,
								'y': 16,
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
