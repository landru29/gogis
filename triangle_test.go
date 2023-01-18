package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestTriangleScan(t *testing.T) {
	scanTest(t, testFixtureScan{
		title:   "triangle",
		rawData: []byte("01110000C001000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440"),
		scanner: &gogis.Triangle{},
		expectedGeometry: &gogis.Triangle{
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
	})
}

func TestTriangleValue(t *testing.T) {
	valueTest(t, testFixtureValue{
		title:           "triangle",
		expectedRawData: []byte("01110000C001000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440"),
		valuer: &gogis.Triangle{
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
	})
}
