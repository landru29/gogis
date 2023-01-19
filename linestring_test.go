package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestLineStringScan(t *testing.T) {
	fixture := gogis.LineString{
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
	}

	dataByte := []byte("01020000C0020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040")

	t.Run("scan with data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          &gogis.LineString{},
			expectedGeometry: &fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.NullLineString{},
			expectedGeometry: &gogis.NullLineString{},
		})
	})

	t.Run("scan valid data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData: dataByte,
			scanner: &gogis.NullLineString{},
			expectedGeometry: &gogis.NullLineString{
				Valid:      true,
				LineString: fixture,
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
			valuer:          gogis.NullLineString{},
		})
	})

	t.Run("value valid data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer: gogis.NullLineString{
				LineString: fixture,
				Valid:      true,
			},
		})
	})
}
