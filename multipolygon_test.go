package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestMultiPolygon(t *testing.T) {
	fixture := gogis.MultiPolygon{
		{
			{
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': -7.03,
						'y': 2.08,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 3.99,
						'y': -7.38,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 0.59,
						'y': 4.27,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': -7.03,
						'y': 2.08,
					},
				},
			},
			{
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 9.37,
						'y': 54.44,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 86.33,
						'y': -31.49,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 38.34,
						'y': -92.61,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 9.37,
						'y': 54.44,
					},
				},
			},
		},

		{
			{
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': -17.03,
						'y': 12.08,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 13.99,
						'y': -17.38,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 10.59,
						'y': 14.27,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': -17.03,
						'y': 12.08,
					},
				},
			},
			{
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 19.37,
						'y': 154.44,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 186.33,
						'y': -131.49,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 138.34,
						'y': -192.61,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 19.37,
						'y': 154.44,
					},
				},
			},
		},
	}

	dataByte := []byte("010600000002000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340")

	t.Run("scan with data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          &gogis.MultiPolygon{},
			expectedGeometry: &fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.NullMultiPolygon{},
			expectedGeometry: &gogis.NullMultiPolygon{},
		})
	})

	t.Run("scan valid data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData: dataByte,
			scanner: &gogis.NullMultiPolygon{},
			expectedGeometry: &gogis.NullMultiPolygon{
				Valid:        true,
				MultiPolygon: fixture,
			},
		})
	})

	t.Run("value with data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer:          &fixture,
		})
	})

	t.Run("value null data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: nil,
			valuer:          gogis.NullMultiPolygon{},
		})
	})

	t.Run("value valid data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer: gogis.NullMultiPolygon{
				MultiPolygon: fixture,
				Valid:        true,
			},
		})
	})
}
