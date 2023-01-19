package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestTriangle(t *testing.T) {
	fixture := gogis.Triangle{
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
	}

	dataByte := []byte("01110000C001000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440")

	t.Run("scan with data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          &gogis.Triangle{},
			expectedGeometry: &fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.NullTriangle{},
			expectedGeometry: &gogis.NullTriangle{},
		})
	})

	t.Run("scan valid data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData: dataByte,
			scanner: &gogis.NullTriangle{},
			expectedGeometry: &gogis.NullTriangle{
				Valid:    true,
				Triangle: fixture,
			},
		})
	})

	t.Run("value with data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer:          fixture,
		})
	})

	t.Run("value null data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: nil,
			valuer:          gogis.NullTriangle{},
		})
	})

	t.Run("value valid data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer: gogis.NullTriangle{
				Triangle: fixture,
				Valid:    true,
			},
		})
	})
}
