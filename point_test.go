package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestPointScan(t *testing.T) {
	scanTest(t, testFixtureScan{
		title:   "point",
		rawData: []byte("01010000C03CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40"),
		scanner: &gogis.Point{},
		expectedGeometry: &gogis.Point{
			Coordinate: ewkb.Coordinate{
				'x': -71.060316,
				'y': 48.432044,
				'z': 10.0,
				'm': 30.0,
			},
		},
	})
}

func TestPointValue(t *testing.T) {
	valueTest(t, testFixtureValue{
		title:           "point",
		expectedRawData: []byte("01010000C03CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40"),
		valuer: &gogis.Point{
			Coordinate: ewkb.Coordinate{
				'x': -71.060316,
				'y': 48.432044,
				'z': 10.0,
				'm': 30.0,
			},
		},
	})
}
