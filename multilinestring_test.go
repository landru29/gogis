package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestMultiLineStringScan(t *testing.T) {
	scanTest(t, testFixtureScan{
		title:   "multilinestring",
		rawData: []byte("01050000C00200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240"),
		scanner: &gogis.MultiLineString{},
		expectedGeometry: &gogis.MultiLineString{
			{
				{
					Coordinate: ewkb.Coordinate{
						'x': 42.42,
						'y': -24.24,
						'z': 42.24,
						'm': -24.42,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': 5,
						'y': 6,
						'z': 7,
						'm': 8,
					},
				},
			},
			{
				{
					Coordinate: ewkb.Coordinate{
						'x': 142.42,
						'y': -424.24,
						'z': 142.24,
						'm': -124.42,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': 15,
						'y': 16,
						'z': 17,
						'm': 18,
					},
				},
			},
		},
	})
}

func TestMultiLineStringValue(t *testing.T) {
	valueTest(t, testFixtureValue{
		title:           "multilinestring",
		expectedRawData: []byte("01050000C00200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240"),
		valuer: &gogis.MultiLineString{
			{
				{
					Coordinate: ewkb.Coordinate{
						'x': 42.42,
						'y': -24.24,
						'z': 42.24,
						'm': -24.42,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': 5,
						'y': 6,
						'z': 7,
						'm': 8,
					},
				},
			},
			{
				{
					Coordinate: ewkb.Coordinate{
						'x': 142.42,
						'y': -424.24,
						'z': 142.24,
						'm': -124.42,
					},
				},
				{
					Coordinate: ewkb.Coordinate{
						'x': 15,
						'y': 16,
						'z': 17,
						'm': 18,
					},
				},
			},
		},
	})
}
