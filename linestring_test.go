package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestLineStringScan(t *testing.T) {
	scanTest(t, testFixtureScan{
		title:   "linestring",
		rawData: []byte("01020000C0020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040"),
		scanner: &gogis.LineString{},
		expectedGeometry: &gogis.LineString{
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
		},
	})
}

func TestLineStringValue(t *testing.T) {
	valueTest(t, testFixtureValue{
		title:           "linestring",
		expectedRawData: []byte("01020000C0020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040"),
		valuer: &gogis.LineString{
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
		},
	})
}
