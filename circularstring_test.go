package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestCircularString(t *testing.T) {
	fixture := gogis.CircularString{
		gogis.Point{
			Coordinate: ewkb.Coordinate{
				'x': -71.060316,
				'y': 48.432044,
				'z': 10,
				'm': 30,
			},
		},
		gogis.Point{
			Coordinate: ewkb.Coordinate{
				'x': 5,
				'y': 6,
				'z': 7,
				'm': 8,
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
	}

	dataByte := []byte("01080000C0030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040")

	t.Run("scan with data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          &gogis.CircularString{},
			expectedGeometry: &fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.NullCircularString{},
			expectedGeometry: &gogis.NullCircularString{},
		})
	})

	t.Run("scan valid data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData: dataByte,
			scanner: &gogis.NullCircularString{},
			expectedGeometry: &gogis.NullCircularString{
				Valid:          true,
				CircularString: fixture,
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
			valuer:          gogis.NullCircularString{},
		})
	})

	t.Run("value valid data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer: gogis.NullCircularString{
				CircularString: fixture,
				Valid:          true,
			},
		})
	})
}
