package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestCircularStringScan(t *testing.T) {
	scanTest(t, testFixtureScan{
		title:   "circularstring",
		rawData: []byte("01080000C0030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040"),
		scanner: &gogis.CircularString{},
		expectedGeometry: &gogis.CircularString{
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
		},
	})
}

func TestCircularStringValue(t *testing.T) {
	valueTest(t, testFixtureValue{
		title:           "circularstring",
		expectedRawData: []byte("01080000C0030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040"),
		valuer: &gogis.CircularString{
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
		},
	})
}
