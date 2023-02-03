package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestGeometryArray(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84
	fixture := &gogis.GeometryArray{
		gogis.Point{
			Coordinate: ewkb.Coordinate{
				'x': 2,
				'y': 3,
				'z': 4,
				'm': 5,
			},
			SRID: &srid,
		}.Geometry(),
		gogis.LineString{
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': 2,
					'y': 3,
					'z': 4,
					'm': 5,
				},
				SRID: &srid,
			},
			gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': 3,
					'y': 4,
					'z': 5,
					'm': 6,
				},
				SRID: &srid,
			},
		}.Geometry(),
	}

	dataByte := []byte("{01010000E0E61000000000000000000040000000000000084000000000000010400000000000001440:01020000E0E61000000200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840}")

	t.Run("scan with data", func(t *testing.T) {
		scanner := &gogis.GeometryArray{}

		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          scanner,
			expectedGeometry: fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.GeometryArray{},
			expectedGeometry: &gogis.GeometryArray{},
		})
	})

	t.Run("value with data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: dataByte,
			valuer:          fixture,
		})
	})

	t.Run("value null data", func(t *testing.T) {
		valueTest(t, testFixtureValue{
			expectedRawData: nil,
			valuer:          gogis.GeometryArray{},
		})
	})
}
