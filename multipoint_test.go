package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestMultiPoint(t *testing.T) {
	fixture := gogis.MultiPoint{
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
	}

	dataByte := []byte("01040000C00400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440")

	t.Run("scan with data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          &gogis.MultiPoint{},
			expectedGeometry: &fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.NullMultiPoint{},
			expectedGeometry: &gogis.NullMultiPoint{},
		})
	})

	t.Run("scan valid data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData: dataByte,
			scanner: &gogis.NullMultiPoint{},
			expectedGeometry: &gogis.NullMultiPoint{
				Valid:      true,
				MultiPoint: fixture,
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
			valuer:          gogis.NullMultiPoint{},
		})
	})

	t.Run("value valid data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer: gogis.NullMultiPoint{
				MultiPoint: fixture,
				Valid:      true,
			},
		})
	})
}
