package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestMultiPointScan(t *testing.T) {
	scanTest(t, testFixtureScan{
		title:   "multipoint",
		rawData: []byte("01040000C00400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440"),
		scanner: &gogis.MultiPoint{},
		expectedGeometry: &gogis.MultiPoint{
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
		},
	})
}

func TestMultiPointValue(t *testing.T) {
	valueTest(t, testFixtureValue{
		title:           "multipoint",
		expectedRawData: []byte("01040000C00400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440"),
		valuer: &gogis.MultiPoint{
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
		},
	})
}
