package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestPoint(t *testing.T) {
	fixture := gogis.Point{
		Coordinate: ewkb.Coordinate{
			'x': -71.060316,
			'y': 48.432044,
			'z': 10.0,
			'm': 30.0,
		},
	}

	dataByte := []byte("01010000C03CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40")

	t.Run("scan with data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          &gogis.Point{},
			expectedGeometry: &fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.NullPoint{},
			expectedGeometry: &gogis.NullPoint{},
		})
	})

	t.Run("scan valid data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData: dataByte,
			scanner: &gogis.NullPoint{},
			expectedGeometry: &gogis.NullPoint{
				Valid: true,
				Point: fixture,
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
			valuer:          gogis.NullPoint{},
		})
	})

	t.Run("value valid data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer: gogis.NullPoint{
				Point: fixture,
				Valid: true,
			},
		})
	})
}
