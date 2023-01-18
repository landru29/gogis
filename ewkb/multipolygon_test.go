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

func TestMultiPolygonType(t *testing.T) {
	assert.Equal(t, ewkb.GeometryTypeMultiPolygon, ewkb.MultiPolygon{}.Type())
}

func TestMultiPolygonUnmarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		var multipoint ewkb.MultiPolygon

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"0200000001030000C002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C000000000005087C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F51540B81E85EB51B80CC01F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484085EB51B81E9555403D0AD7A3707D3FC00000000000204340F6285C8FC2B55140EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E15840D7A3703D0A7748C03D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484001030000C0020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC07B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC00000000000489BC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E40AE47E17A142E2BC048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC0040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E15A65407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F06840F6285C8FC29D62C01F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiPolygon),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiPolygon{
			Polygons: []ewkb.Polygon{
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -7.03,
								'y': 2.08,
								'z': 4.58,
								'm': -5.9,
							},
							{
								'x': 3.99,
								'y': -7.38,
								'z': -4.53,
								'm': -746,
							},
							{
								'x': 0.59,
								'y': 4.27,
								'z': 5.49,
								'm': -3.59,
							},
							{
								'x': -7.03,
								'y': 2.08,
								'z': 4.58,
								'm': -5.9,
							},
						},
						{
							{
								'x': 9.37,
								'y': 54.44,
								'z': -75.29,
								'm': 49.19,
							},
							{
								'x': 86.33,
								'y': -31.49,
								'z': 38.25,
								'm': 70.84,
							},
							{
								'x': 38.34,
								'y': -92.61,
								'z': 99.52,
								'm': -48.93,
							},
							{
								'x': 9.37,
								'y': 54.44,
								'z': -75.29,
								'm': 49.19,
							},
						},
					},
				},
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -17.03,
								'y': 12.08,
								'z': 14.58,
								'm': -15.9,
							},
							{
								'x': 13.99,
								'y': -17.38,
								'z': -14.53,
								'm': -1746,
							},
							{
								'x': 10.59,
								'y': 14.27,
								'z': 15.49,
								'm': -13.59,
							},
							{
								'x': -17.03,
								'y': 12.08,
								'z': 14.58,
								'm': -15.9,
							},
						},
						{
							{
								'x': 19.37,
								'y': 154.44,
								'z': -175.29,
								'm': 149.19,
							},
							{
								'x': 186.33,
								'y': -131.49,
								'z': 138.25,
								'm': 170.84,
							},
							{
								'x': 138.34,
								'y': -192.61,
								'z': 199.52,
								'm': -148.93,
							},
							{
								'x': 19.37,
								'y': 154.44,
								'z': -175.29,
								'm': 149.19,
							},
						},
					},
				},
			},
		})
	})

	t.Run("XYZ", func(t *testing.T) {
		var multipoint ewkb.MultiPolygon

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"02000000010300008002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F515401F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C085EB51B81E9555403D0AD7A3707D3FC00000000000204340EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E158403D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C00103000080020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D407B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E4048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F068401F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiPolygon),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiPolygon{
			Polygons: []ewkb.Polygon{
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -7.03,
								'y': 2.08,
								'z': 4.58,
							},
							{
								'x': 3.99,
								'y': -7.38,
								'z': -4.53,
							},
							{
								'x': 0.59,
								'y': 4.27,
								'z': 5.49,
							},
							{
								'x': -7.03,
								'y': 2.08,
								'z': 4.58,
							},
						},
						{
							{
								'x': 9.37,
								'y': 54.44,
								'z': -75.29,
							},
							{
								'x': 86.33,
								'y': -31.49,
								'z': 38.25,
							},
							{
								'x': 38.34,
								'y': -92.61,
								'z': 99.52,
							},
							{
								'x': 9.37,
								'y': 54.44,
								'z': -75.29,
							},
						},
					},
				},
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -17.03,
								'y': 12.08,
								'z': 14.58,
							},
							{
								'x': 13.99,
								'y': -17.38,
								'z': -14.53,
							},
							{
								'x': 10.59,
								'y': 14.27,
								'z': 15.49,
							},
							{
								'x': -17.03,
								'y': 12.08,
								'z': 14.58,
							},
						},
						{
							{
								'x': 19.37,
								'y': 154.44,
								'z': -175.29,
							},
							{
								'x': 186.33,
								'y': -131.49,
								'z': 138.25,
							},
							{
								'x': 138.34,
								'y': -192.61,
								'z': 199.52,
							},
							{
								'x': 19.37,
								'y': 154.44,
								'z': -175.29,
							},
						},
					},
				},
			},
		})
	})

	t.Run("XY", func(t *testing.T) {
		var multipoint ewkb.MultiPolygon

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"02000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340",
				withLayout(ewkb.Layout(3)),
				withByteOrder(binary.LittleEndian),
				withType(ewkb.GeometryTypeMultiPolygon),
			),
		)
		require.NoError(t, err)
		assert.Equal(t, multipoint, ewkb.MultiPolygon{
			Polygons: []ewkb.Polygon{
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -7.03,
								'y': 2.08,
							},
							{
								'x': 3.99,
								'y': -7.38,
							},
							{
								'x': 0.59,
								'y': 4.27,
							},
							{
								'x': -7.03,
								'y': 2.08,
							},
						},
						{
							{
								'x': 9.37,
								'y': 54.44,
							},
							{
								'x': 86.33,
								'y': -31.49,
							},
							{
								'x': 38.34,
								'y': -92.61,
							},
							{
								'x': 9.37,
								'y': 54.44,
							},
						},
					},
				},
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -17.03,
								'y': 12.08,
							},
							{
								'x': 13.99,
								'y': -17.38,
							},
							{
								'x': 10.59,
								'y': 14.27,
							},
							{
								'x': -17.03,
								'y': 12.08,
							},
						},
						{
							{
								'x': 19.37,
								'y': 154.44,
							},
							{
								'x': 186.33,
								'y': -131.49,
							},
							{
								'x': 138.34,
								'y': -192.61,
							},
							{
								'x': 19.37,
								'y': 154.44,
							},
						},
					},
				},
			},
		})
	})

	t.Run("wrong type", func(t *testing.T) {
		var multipoint ewkb.MultiPolygon

		err := (&multipoint).UnmarshalEWBK(
			newExtendedWellKnownBytes(
				t,
				"02000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340",
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

func TestMultiPolygonMarshalEWBK(t *testing.T) {
	t.Run("XYZM", func(t *testing.T) {
		point := ewkb.MultiPolygon{
			Polygons: []ewkb.Polygon{
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -7.03,
								'y': 2.08,
								'z': 4.58,
								'm': -5.9,
							},
							{
								'x': 3.99,
								'y': -7.38,
								'z': -4.53,
								'm': -746,
							},
							{
								'x': 0.59,
								'y': 4.27,
								'z': 5.49,
								'm': -3.59,
							},
							{
								'x': -7.03,
								'y': 2.08,
								'z': 4.58,
								'm': -5.9,
							},
						},
						{
							{
								'x': 9.37,
								'y': 54.44,
								'z': -75.29,
								'm': 49.19,
							},
							{
								'x': 86.33,
								'y': -31.49,
								'z': 38.25,
								'm': 70.84,
							},
							{
								'x': 38.34,
								'y': -92.61,
								'z': 99.52,
								'm': -48.93,
							},
							{
								'x': 9.37,
								'y': 54.44,
								'z': -75.29,
								'm': 49.19,
							},
						},
					},
				},
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -17.03,
								'y': 12.08,
								'z': 14.58,
								'm': -15.9,
							},
							{
								'x': 13.99,
								'y': -17.38,
								'z': -14.53,
								'm': -1746,
							},
							{
								'x': 10.59,
								'y': 14.27,
								'z': 15.49,
								'm': -13.59,
							},
							{
								'x': -17.03,
								'y': 12.08,
								'z': 14.58,
								'm': -15.9,
							},
						},
						{
							{
								'x': 19.37,
								'y': 154.44,
								'z': -175.29,
								'm': 149.19,
							},
							{
								'x': 186.33,
								'y': -131.49,
								'z': 138.25,
								'm': 170.84,
							},
							{
								'x': 138.34,
								'y': -192.61,
								'z': 199.52,
								'm': -148.93,
							},
							{
								'x': 19.37,
								'y': 154.44,
								'z': -175.29,
								'm': 149.19,
							},
						},
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("0200000001030000C002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C000000000005087C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F51540B81E85EB51B80CC01F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484085EB51B81E9555403D0AD7A3707D3FC00000000000204340F6285C8FC2B55140EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E15840D7A3703D0A7748C03D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484001030000C0020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC07B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC00000000000489BC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E40AE47E17A142E2BC048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC0040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E15A65407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F06840F6285C8FC29D62C01F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240"),
			hex.EncodeToString(data),
		)
	})
	t.Run("XYZ", func(t *testing.T) {
		point := ewkb.MultiPolygon{
			Polygons: []ewkb.Polygon{
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -7.03,
								'y': 2.08,
								'z': 4.58,
							},
							{
								'x': 3.99,
								'y': -7.38,
								'z': -4.53,
							},
							{
								'x': 0.59,
								'y': 4.27,
								'z': 5.49,
							},
							{
								'x': -7.03,
								'y': 2.08,
								'z': 4.58,
							},
						},
						{
							{
								'x': 9.37,
								'y': 54.44,
								'z': -75.29,
							},
							{
								'x': 86.33,
								'y': -31.49,
								'z': 38.25,
							},
							{
								'x': 38.34,
								'y': -92.61,
								'z': 99.52,
							},
							{
								'x': 9.37,
								'y': 54.44,
								'z': -75.29,
							},
						},
					},
				},
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -17.03,
								'y': 12.08,
								'z': 14.58,
							},
							{
								'x': 13.99,
								'y': -17.38,
								'z': -14.53,
							},
							{
								'x': 10.59,
								'y': 14.27,
								'z': 15.49,
							},
							{
								'x': -17.03,
								'y': 12.08,
								'z': 14.58,
							},
						},
						{
							{
								'x': 19.37,
								'y': 154.44,
								'z': -175.29,
							},
							{
								'x': 186.33,
								'y': -131.49,
								'z': 138.25,
							},
							{
								'x': 138.34,
								'y': -192.61,
								'z': 199.52,
							},
							{
								'x': 19.37,
								'y': 154.44,
								'z': -175.29,
							},
						},
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("02000000010300008002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F515401F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C085EB51B81E9555403D0AD7A3707D3FC00000000000204340EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E158403D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C00103000080020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D407B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E4048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F068401F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0"),
			hex.EncodeToString(data),
		)
	})

	t.Run("XY", func(t *testing.T) {
		point := ewkb.MultiPolygon{
			Polygons: []ewkb.Polygon{
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -7.03,
								'y': 2.08,
							},
							{
								'x': 3.99,
								'y': -7.38,
							},
							{
								'x': 0.59,
								'y': 4.27,
							},
							{
								'x': -7.03,
								'y': 2.08,
							},
						},
						{
							{
								'x': 9.37,
								'y': 54.44,
							},
							{
								'x': 86.33,
								'y': -31.49,
							},
							{
								'x': 38.34,
								'y': -92.61,
							},
							{
								'x': 9.37,
								'y': 54.44,
							},
						},
					},
				},
				{
					CoordinateGroup: ewkb.CoordinateGroup{
						{
							{
								'x': -17.03,
								'y': 12.08,
							},
							{
								'x': 13.99,
								'y': -17.38,
							},
							{
								'x': 10.59,
								'y': 14.27,
							},
							{
								'x': -17.03,
								'y': 12.08,
							},
						},
						{
							{
								'x': 19.37,
								'y': 154.44,
							},
							{
								'x': 186.33,
								'y': -131.49,
							},
							{
								'x': 138.34,
								'y': -192.61,
							},
							{
								'x': 19.37,
								'y': 154.44,
							},
						},
					},
				},
			},
		}

		data, err := point.MarshalEWBK(binary.LittleEndian)
		assert.NoError(t, err)
		assert.Equal(
			t,
			strings.ToLower("02000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340"),
			hex.EncodeToString(data),
		)
	})
}

func TestMultiPolygonUnmarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		binary      string
		expected    ewkb.Geometry
	}{
		{
			geometry:    &ewkb.MultiPolygon{},
			strGeometry: "MULTIPOLYGON ZM(((-7.03 2.08 4.58 -5.9,3.99 -7.38 -4.53 -746,0.59 4.27 5.49 -3.59,-7.03 2.08 4.58 -5.9),(9.37 54.44 -75.29 49.19,86.33 -31.49 38.25 70.84,38.34 -92.61 99.52 -48.93,9.37 54.44 -75.29 49.19)),((-17.03 12.08 14.58 -15.9,13.99 -17.38 -14.53 -1746,10.59 14.27 15.49 -13.59,-17.03 12.08 14.58 -15.9),(19.37 154.44 -175.29 149.19,186.33 -131.49 138.25 170.84,138.34 -192.61 199.52 -148.93,19.37 154.44 -175.29 149.19)))",
			binary:      "01060000C00200000001030000C002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C000000000005087C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F51540B81E85EB51B80CC01F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484085EB51B81E9555403D0AD7A3707D3FC00000000000204340F6285C8FC2B55140EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E15840D7A3703D0A7748C03D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484001030000C0020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC07B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC00000000000489BC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E40AE47E17A142E2BC048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC0040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E15A65407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F06840F6285C8FC29D62C01F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240",
			expected: &ewkb.MultiPolygon{
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
									'm': -5.9,
								},
								{
									'x': 3.99,
									'y': -7.38,
									'z': -4.53,
									'm': -746,
								},
								{
									'x': 0.59,
									'y': 4.27,
									'z': 5.49,
									'm': -3.59,
								},
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
									'm': -5.9,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
									'm': 49.19,
								},
								{
									'x': 86.33,
									'y': -31.49,
									'z': 38.25,
									'm': 70.84,
								},
								{
									'x': 38.34,
									'y': -92.61,
									'z': 99.52,
									'm': -48.93,
								},
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
									'm': 49.19,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
									'm': -15.9,
								},
								{
									'x': 13.99,
									'y': -17.38,
									'z': -14.53,
									'm': -1746,
								},
								{
									'x': 10.59,
									'y': 14.27,
									'z': 15.49,
									'm': -13.59,
								},
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
									'm': -15.9,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
									'm': 149.19,
								},
								{
									'x': 186.33,
									'y': -131.49,
									'z': 138.25,
									'm': 170.84,
								},
								{
									'x': 138.34,
									'y': -192.61,
									'z': 199.52,
									'm': -148.93,
								},
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
									'm': 149.19,
								},
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPolygon{},
			strGeometry: "MULTIPOLYGON ZM(((-7.03 2.08 4.58 -5.9,3.99 -7.38 -4.53 -746,0.59 4.27 5.49 -3.59,-7.03 2.08 4.58 -5.9),(9.37 54.44 -75.29 49.19,86.33 -31.49 38.25 70.84,38.34 -92.61 99.52 -48.93,9.37 54.44 -75.29 49.19)),((-17.03 12.08 14.58 -15.9,13.99 -17.38 -14.53 -1746,10.59 14.27 15.49 -13.59,-17.03 12.08 14.58 -15.9),(19.37 154.44 -175.29 149.19,186.33 -131.49 138.25 170.84,138.34 -192.61 199.52 -148.93,19.37 154.44 -175.29 149.19))),4326",
			binary:      "01060000E0E61000000200000001030000C002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C000000000005087C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F51540B81E85EB51B80CC01F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484085EB51B81E9555403D0AD7A3707D3FC00000000000204340F6285C8FC2B55140EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E15840D7A3703D0A7748C03D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484001030000C0020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC07B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC00000000000489BC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E40AE47E17A142E2BC048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC0040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E15A65407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F06840F6285C8FC29D62C01F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240",
			expected: &ewkb.MultiPolygon{
				SRID: &srid,
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
									'm': -5.9,
								},
								{
									'x': 3.99,
									'y': -7.38,
									'z': -4.53,
									'm': -746,
								},
								{
									'x': 0.59,
									'y': 4.27,
									'z': 5.49,
									'm': -3.59,
								},
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
									'm': -5.9,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
									'm': 49.19,
								},
								{
									'x': 86.33,
									'y': -31.49,
									'z': 38.25,
									'm': 70.84,
								},
								{
									'x': 38.34,
									'y': -92.61,
									'z': 99.52,
									'm': -48.93,
								},
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
									'm': 49.19,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
									'm': -15.9,
								},
								{
									'x': 13.99,
									'y': -17.38,
									'z': -14.53,
									'm': -1746,
								},
								{
									'x': 10.59,
									'y': 14.27,
									'z': 15.49,
									'm': -13.59,
								},
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
									'm': -15.9,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
									'm': 149.19,
								},
								{
									'x': 186.33,
									'y': -131.49,
									'z': 138.25,
									'm': 170.84,
								},
								{
									'x': 138.34,
									'y': -192.61,
									'z': 199.52,
									'm': -148.93,
								},
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
									'm': 149.19,
								},
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPolygon{},
			strGeometry: "MULTIPOLYGON Z(((-7.03 2.08 4.58,3.99 -7.38 -4.53,0.59 4.27 5.49,-7.03 2.08 4.58),(9.37 54.44 -75.29,86.33 -31.49 38.25,38.34 -92.61 99.52,9.37 54.44 -75.29)),((-17.03 12.08 14.58,13.99 -17.38 -14.53,10.59 14.27 15.49,-17.03 12.08 14.58),(19.37 154.44 -175.29,186.33 -131.49 138.25,138.34 -192.61 199.52,19.37 154.44 -175.29)))",
			binary:      "010600008002000000010300008002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F515401F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C085EB51B81E9555403D0AD7A3707D3FC00000000000204340EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E158403D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C00103000080020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D407B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E4048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F068401F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0",
			expected: &ewkb.MultiPolygon{
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
								},
								{
									'x': 3.99,
									'y': -7.38,
									'z': -4.53,
								},
								{
									'x': 0.59,
									'y': 4.27,
									'z': 5.49,
								},
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
								},
								{
									'x': 86.33,
									'y': -31.49,
									'z': 38.25,
								},
								{
									'x': 38.34,
									'y': -92.61,
									'z': 99.52,
								},
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
								},
								{
									'x': 13.99,
									'y': -17.38,
									'z': -14.53,
								},
								{
									'x': 10.59,
									'y': 14.27,
									'z': 15.49,
								},
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
								},
								{
									'x': 186.33,
									'y': -131.49,
									'z': 138.25,
								},
								{
									'x': 138.34,
									'y': -192.61,
									'z': 199.52,
								},
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
								},
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPolygon{},
			strGeometry: "MULTIPOLYGON Z(((-7.03 2.08 4.58,3.99 -7.38 -4.53,0.59 4.27 5.49,-7.03 2.08 4.58),(9.37 54.44 -75.29,86.33 -31.49 38.25,38.34 -92.61 99.52,9.37 54.44 -75.29)),((-17.03 12.08 14.58,13.99 -17.38 -14.53,10.59 14.27 15.49,-17.03 12.08 14.58),(19.37 154.44 -175.29,186.33 -131.49 138.25,138.34 -192.61 199.52,19.37 154.44 -175.29))),4326",
			binary:      "01060000A0E610000002000000010300008002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F515401F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C085EB51B81E9555403D0AD7A3707D3FC00000000000204340EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E158403D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C00103000080020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D407B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E4048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F068401F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0",
			expected: &ewkb.MultiPolygon{
				SRID: &srid,
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
								},
								{
									'x': 3.99,
									'y': -7.38,
									'z': -4.53,
								},
								{
									'x': 0.59,
									'y': 4.27,
									'z': 5.49,
								},
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
								},
								{
									'x': 86.33,
									'y': -31.49,
									'z': 38.25,
								},
								{
									'x': 38.34,
									'y': -92.61,
									'z': 99.52,
								},
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
								},
								{
									'x': 13.99,
									'y': -17.38,
									'z': -14.53,
								},
								{
									'x': 10.59,
									'y': 14.27,
									'z': 15.49,
								},
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
								},
								{
									'x': 186.33,
									'y': -131.49,
									'z': 138.25,
								},
								{
									'x': 138.34,
									'y': -192.61,
									'z': 199.52,
								},
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
								},
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPolygon{},
			strGeometry: "MULTIPOLYGON(((-7.03 2.08,3.99 -7.38,0.59 4.27,-7.03 2.08),(9.37 54.44,86.33 -31.49,38.34 -92.61,9.37 54.44)),((-17.03 12.08,13.99 -17.38,10.59 14.27,-17.03 12.08),(19.37 154.44,186.33 -131.49,138.34 -192.61,19.37 154.44)))",
			binary:      "010600000002000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340",
			expected: &ewkb.MultiPolygon{
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
								},
								{
									'x': 3.99,
									'y': -7.38,
								},
								{
									'x': 0.59,
									'y': 4.27,
								},
								{
									'x': -7.03,
									'y': 2.08,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
								},
								{
									'x': 86.33,
									'y': -31.49,
								},
								{
									'x': 38.34,
									'y': -92.61,
								},
								{
									'x': 9.37,
									'y': 54.44,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
								},
								{
									'x': 13.99,
									'y': -17.38,
								},
								{
									'x': 10.59,
									'y': 14.27,
								},
								{
									'x': -17.03,
									'y': 12.08,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
								},
								{
									'x': 186.33,
									'y': -131.49,
								},
								{
									'x': 138.34,
									'y': -192.61,
								},
								{
									'x': 19.37,
									'y': 154.44,
								},
							},
						},
					},
				},
			},
		},
		{
			geometry:    &ewkb.MultiPolygon{},
			strGeometry: "MULTIPOLYGON(((-7.03 2.08,3.99 -7.38,0.59 4.27,-7.03 2.08),(9.37 54.44,86.33 -31.49,38.34 -92.61,9.37 54.44)),((-17.03 12.08,13.99 -17.38,10.59 14.27,-17.03 12.08),(19.37 154.44,186.33 -131.49,138.34 -192.61,19.37 154.44))),4326",
			binary:      "0106000020E610000002000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340",
			expected: &ewkb.MultiPolygon{
				SRID: &srid,
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
								},
								{
									'x': 3.99,
									'y': -7.38,
								},
								{
									'x': 0.59,
									'y': 4.27,
								},
								{
									'x': -7.03,
									'y': 2.08,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
								},
								{
									'x': 86.33,
									'y': -31.49,
								},
								{
									'x': 38.34,
									'y': -92.61,
								},
								{
									'x': 9.37,
									'y': 54.44,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
								},
								{
									'x': 13.99,
									'y': -17.38,
								},
								{
									'x': 10.59,
									'y': 14.27,
								},
								{
									'x': -17.03,
									'y': 12.08,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
								},
								{
									'x': 186.33,
									'y': -131.49,
								},
								{
									'x': 138.34,
									'y': -192.61,
								},
								{
									'x': 19.37,
									'y': 154.44,
								},
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

func TestMultiPolygonMarshal(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84

	fixtures := []struct {
		geometry    ewkb.Geometry
		strGeometry string
		expected    string
	}{
		{
			strGeometry: "MULTIPOLYGON ZM(((-7.03 2.08 4.58 -5.9,3.99 -7.38 -4.53 -746,0.59 4.27 5.49 -3.59,-7.03 2.08 4.58 -5.9),(9.37 54.44 -75.29 49.19,86.33 -31.49 38.25 70.84,38.34 -92.61 99.52 -48.93,9.37 54.44 -75.29 49.19)),((-17.03 12.08 14.58 -15.9,13.99 -17.38 -14.53 -1746,10.59 14.27 15.49 -13.59,-17.03 12.08 14.58 -15.9),(19.37 154.44 -175.29 149.19,186.33 -131.49 138.25 170.84,138.34 -192.61 199.52 -148.93,19.37 154.44 -175.29 149.19)))",
			expected:    "01060000C00200000001030000C002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C000000000005087C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F51540B81E85EB51B80CC01F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484085EB51B81E9555403D0AD7A3707D3FC00000000000204340F6285C8FC2B55140EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E15840D7A3703D0A7748C03D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484001030000C0020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC07B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC00000000000489BC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E40AE47E17A142E2BC048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC0040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E15A65407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F06840F6285C8FC29D62C01F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240",
			geometry: &ewkb.MultiPolygon{
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
									'm': -5.9,
								},
								{
									'x': 3.99,
									'y': -7.38,
									'z': -4.53,
									'm': -746,
								},
								{
									'x': 0.59,
									'y': 4.27,
									'z': 5.49,
									'm': -3.59,
								},
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
									'm': -5.9,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
									'm': 49.19,
								},
								{
									'x': 86.33,
									'y': -31.49,
									'z': 38.25,
									'm': 70.84,
								},
								{
									'x': 38.34,
									'y': -92.61,
									'z': 99.52,
									'm': -48.93,
								},
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
									'm': 49.19,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
									'm': -15.9,
								},
								{
									'x': 13.99,
									'y': -17.38,
									'z': -14.53,
									'm': -1746,
								},
								{
									'x': 10.59,
									'y': 14.27,
									'z': 15.49,
									'm': -13.59,
								},
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
									'm': -15.9,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
									'm': 149.19,
								},
								{
									'x': 186.33,
									'y': -131.49,
									'z': 138.25,
									'm': 170.84,
								},
								{
									'x': 138.34,
									'y': -192.61,
									'z': 199.52,
									'm': -148.93,
								},
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
									'm': 149.19,
								},
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOLYGON ZM(((-7.03 2.08 4.58 -5.9,3.99 -7.38 -4.53 -746,0.59 4.27 5.49 -3.59,-7.03 2.08 4.58 -5.9),(9.37 54.44 -75.29 49.19,86.33 -31.49 38.25 70.84,38.34 -92.61 99.52 -48.93,9.37 54.44 -75.29 49.19)),((-17.03 12.08 14.58 -15.9,13.99 -17.38 -14.53 -1746,10.59 14.27 15.49 -13.59,-17.03 12.08 14.58 -15.9),(19.37 154.44 -175.29 149.19,186.33 -131.49 138.25 170.84,138.34 -192.61 199.52 -148.93,19.37 154.44 -175.29 149.19))),4326",
			expected:    "01060000E0E61000000200000001030000C002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C000000000005087C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F51540B81E85EB51B80CC01F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB5112409A999999999917C0040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484085EB51B81E9555403D0AD7A3707D3FC00000000000204340F6285C8FC2B55140EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E15840D7A3703D0A7748C03D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C0B81E85EB5198484001030000C0020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC07B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC00000000000489BC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E40AE47E17A142E2BC048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40CDCCCCCCCCCC2FC0040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E15A65407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F06840F6285C8FC29D62C01F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0AE47E17A14A66240",
			geometry: &ewkb.MultiPolygon{
				SRID: &srid,
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
									'm': -5.9,
								},
								{
									'x': 3.99,
									'y': -7.38,
									'z': -4.53,
									'm': -746,
								},
								{
									'x': 0.59,
									'y': 4.27,
									'z': 5.49,
									'm': -3.59,
								},
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
									'm': -5.9,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
									'm': 49.19,
								},
								{
									'x': 86.33,
									'y': -31.49,
									'z': 38.25,
									'm': 70.84,
								},
								{
									'x': 38.34,
									'y': -92.61,
									'z': 99.52,
									'm': -48.93,
								},
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
									'm': 49.19,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
									'm': -15.9,
								},
								{
									'x': 13.99,
									'y': -17.38,
									'z': -14.53,
									'm': -1746,
								},
								{
									'x': 10.59,
									'y': 14.27,
									'z': 15.49,
									'm': -13.59,
								},
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
									'm': -15.9,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
									'm': 149.19,
								},
								{
									'x': 186.33,
									'y': -131.49,
									'z': 138.25,
									'm': 170.84,
								},
								{
									'x': 138.34,
									'y': -192.61,
									'z': 199.52,
									'm': -148.93,
								},
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
									'm': 149.19,
								},
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOLYGON Z(((-7.03 2.08 4.58,3.99 -7.38 -4.53,0.59 4.27 5.49,-7.03 2.08 4.58),(9.37 54.44 -75.29,86.33 -31.49 38.25,38.34 -92.61 99.52,9.37 54.44 -75.29)),((-17.03 12.08 14.58,13.99 -17.38 -14.53,10.59 14.27 15.49,-17.03 12.08 14.58),(19.37 154.44 -175.29,186.33 -131.49 138.25,138.34 -192.61 199.52,19.37 154.44 -175.29)))",
			expected:    "010600008002000000010300008002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F515401F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C085EB51B81E9555403D0AD7A3707D3FC00000000000204340EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E158403D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C00103000080020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D407B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E4048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F068401F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0",
			geometry: &ewkb.MultiPolygon{
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
								},
								{
									'x': 3.99,
									'y': -7.38,
									'z': -4.53,
								},
								{
									'x': 0.59,
									'y': 4.27,
									'z': 5.49,
								},
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
								},
								{
									'x': 86.33,
									'y': -31.49,
									'z': 38.25,
								},
								{
									'x': 38.34,
									'y': -92.61,
									'z': 99.52,
								},
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
								},
								{
									'x': 13.99,
									'y': -17.38,
									'z': -14.53,
								},
								{
									'x': 10.59,
									'y': 14.27,
									'z': 15.49,
								},
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
								},
								{
									'x': 186.33,
									'y': -131.49,
									'z': 138.25,
								},
								{
									'x': 138.34,
									'y': -192.61,
									'z': 199.52,
								},
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
								},
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOLYGON Z(((-7.03 2.08 4.58,3.99 -7.38 -4.53,0.59 4.27 5.49,-7.03 2.08 4.58),(9.37 54.44 -75.29,86.33 -31.49 38.25,38.34 -92.61 99.52,9.37 54.44 -75.29)),((-17.03 12.08 14.58,13.99 -17.38 -14.53,10.59 14.27 15.49,-17.03 12.08 14.58),(19.37 154.44 -175.29,186.33 -131.49 138.25,138.34 -192.61 199.52,19.37 154.44 -175.29))),4326",
			expected:    "01060000A0E610000002000000010300008002000000040000001F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240EC51B81E85EB0F4085EB51B81E851DC01F85EB51B81E12C0E17A14AE47E1E23F14AE47E17A141140F6285C8FC2F515401F85EB51B81E1CC0A4703D0AD7A3004052B81E85EB511240040000003D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C085EB51B81E9555403D0AD7A3707D3FC00000000000204340EC51B81E852B4340D7A3703D0A2757C0E17A14AE47E158403D0AD7A370BD2240B81E85EB51384B40C3F5285C8FD252C00103000080020000000400000048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D407B14AE47E1FA2B40E17A14AE476131C08FC2F5285C0F2DC0AE47E17A142E25400AD7A3703D8A2C407B14AE47E1FA2E4048E17A14AE0731C0295C8FC2F5282840295C8FC2F5282D40040000001F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0C3F5285C8F4A674048E17A14AE6F60C000000000004861407B14AE47E14A6140EC51B81E851368C0713D0AD7A3F068401F85EB51B85E3340AE47E17A144E6340E17A14AE47E965C0",
			geometry: &ewkb.MultiPolygon{
				SRID: &srid,
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
								},
								{
									'x': 3.99,
									'y': -7.38,
									'z': -4.53,
								},
								{
									'x': 0.59,
									'y': 4.27,
									'z': 5.49,
								},
								{
									'x': -7.03,
									'y': 2.08,
									'z': 4.58,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
								},
								{
									'x': 86.33,
									'y': -31.49,
									'z': 38.25,
								},
								{
									'x': 38.34,
									'y': -92.61,
									'z': 99.52,
								},
								{
									'x': 9.37,
									'y': 54.44,
									'z': -75.29,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
								},
								{
									'x': 13.99,
									'y': -17.38,
									'z': -14.53,
								},
								{
									'x': 10.59,
									'y': 14.27,
									'z': 15.49,
								},
								{
									'x': -17.03,
									'y': 12.08,
									'z': 14.58,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
								},
								{
									'x': 186.33,
									'y': -131.49,
									'z': 138.25,
								},
								{
									'x': 138.34,
									'y': -192.61,
									'z': 199.52,
								},
								{
									'x': 19.37,
									'y': 154.44,
									'z': -175.29,
								},
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOLYGON(((-7.03 2.08,3.99 -7.38,0.59 4.27,-7.03 2.08),(9.37 54.44,86.33 -31.49,38.34 -92.61,9.37 54.44)),((-17.03 12.08,13.99 -17.38,10.59 14.27,-17.03 12.08),(19.37 154.44,186.33 -131.49,138.34 -192.61,19.37 154.44)))",
			expected:    "010600000002000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340",
			geometry: &ewkb.MultiPolygon{
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
								},
								{
									'x': 3.99,
									'y': -7.38,
								},
								{
									'x': 0.59,
									'y': 4.27,
								},
								{
									'x': -7.03,
									'y': 2.08,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
								},
								{
									'x': 86.33,
									'y': -31.49,
								},
								{
									'x': 38.34,
									'y': -92.61,
								},
								{
									'x': 9.37,
									'y': 54.44,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
								},
								{
									'x': 13.99,
									'y': -17.38,
								},
								{
									'x': 10.59,
									'y': 14.27,
								},
								{
									'x': -17.03,
									'y': 12.08,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
								},
								{
									'x': 186.33,
									'y': -131.49,
								},
								{
									'x': 138.34,
									'y': -192.61,
								},
								{
									'x': 19.37,
									'y': 154.44,
								},
							},
						},
					},
				},
			},
		},
		{
			strGeometry: "MULTIPOLYGON(((-7.03 2.08,3.99 -7.38,0.59 4.27,-7.03 2.08),(9.37 54.44,86.33 -31.49,38.34 -92.61,9.37 54.44)),((-17.03 12.08,13.99 -17.38,10.59 14.27,-17.03 12.08),(19.37 154.44,186.33 -131.49,138.34 -192.61,19.37 154.44))),4326",
			expected:    "0106000020E610000002000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340",
			geometry: &ewkb.MultiPolygon{
				SRID: &srid,
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
								},
								{
									'x': 3.99,
									'y': -7.38,
								},
								{
									'x': 0.59,
									'y': 4.27,
								},
								{
									'x': -7.03,
									'y': 2.08,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
								},
								{
									'x': 86.33,
									'y': -31.49,
								},
								{
									'x': 38.34,
									'y': -92.61,
								},
								{
									'x': 9.37,
									'y': 54.44,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
								},
								{
									'x': 13.99,
									'y': -17.38,
								},
								{
									'x': 10.59,
									'y': 14.27,
								},
								{
									'x': -17.03,
									'y': 12.08,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
								},
								{
									'x': 186.33,
									'y': -131.49,
								},
								{
									'x': 138.34,
									'y': -192.61,
								},
								{
									'x': 19.37,
									'y': 154.44,
								},
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
