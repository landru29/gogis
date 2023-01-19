package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestPolygon(t *testing.T) {
	fixture := gogis.Polygon{
		{
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': -71.42,
					'y': 42.71,
					'z': 4,
					'm': 5,
				},
			},
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': -17.42,
					'y': 42.17,
					'z': 4,
					'm': 5,
				},
			},
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': -17.42,
					'y': 71.17,
					'z': 4,
					'm': 5,
				},
			},
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': -71.42,
					'y': 42.71,
					'z': 4,
					'm': 5,
				},
			},
		},
		{
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': 1,
					'y': 2,
					'z': 3,
					'm': 4,
				},
			},
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': 4,
					'y': 5,
					'z': 6,
					'm': 7,
				},
			},
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': 7,
					'y': 8,
					'z': 9,
					'm': 0,
				},
			},
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': 1,
					'y': 2,
					'z': 3,
					'm': 4,
				},
			},
		},
	}

	dataByte := []byte("01030000C002000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144004000000000000000000F03F0000000000000040000000000000084000000000000010400000000000001040000000000000144000000000000018400000000000001C400000000000001C40000000000000204000000000000022400000000000000000000000000000F03F000000000000004000000000000008400000000000001040")

	t.Run("scan with data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          &gogis.Polygon{},
			expectedGeometry: &fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.NullPolygon{},
			expectedGeometry: &gogis.NullPolygon{},
		})
	})

	t.Run("scan valid data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData: dataByte,
			scanner: &gogis.NullPolygon{},
			expectedGeometry: &gogis.NullPolygon{
				Valid:   true,
				Polygon: fixture,
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
			valuer:          gogis.NullPolygon{},
		})
	})

	t.Run("value valid data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer: gogis.NullPolygon{
				Polygon: fixture,
				Valid:   true,
			},
		})
	})
}
