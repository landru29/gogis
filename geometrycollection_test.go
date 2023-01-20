package gogis_test

import (
	"testing"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func TestGeometryCollection(t *testing.T) {
	srid := ewkb.SystemReferenceWGS84
	fixture := gogis.NewGeometryCollection(
		gogis.WithWellKnownGeometry(gogis.DefaultWellKnownBinding()...),
		gogis.WithSystemReferenceID(srid),
		gogis.WithGeometry(
			&gogis.Point{
				Coordinate: ewkb.Coordinate{
					'x': 2,
					'y': 3,
					'z': 4,
					'm': 5,
				},
			},
			&gogis.LineString{
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 2,
						'y': 3,
						'z': 4,
						'm': 5,
					},
				},
				gogis.Point{
					Coordinate: ewkb.Coordinate{
						'x': 3,
						'y': 4,
						'z': 5,
						'm': 6,
					},
				},
			},
		),
	)

	dataByte := []byte("01070000E0E61000000200000001010000C0000000000000004000000000000008400000000000001040000000000000144001020000C00200000000000000000000400000000000000840000000000000104000000000000014400000000000000840000000000000104000000000000014400000000000001840")

	t.Run("scan with data", func(t *testing.T) {
		scanner := gogis.NewGeometryCollection(gogis.WithWellKnownGeometry(gogis.DefaultWellKnownBinding()...))

		scanTest(t, testFixtureScan{
			rawData:          dataByte,
			scanner:          scanner,
			expectedGeometry: fixture,
		})
	})

	t.Run("scan null data", func(t *testing.T) {
		scanTest(t, testFixtureScan{
			rawData:          nil,
			scanner:          &gogis.GeometryCollection{},
			expectedGeometry: &gogis.GeometryCollection{},
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
			valuer:          gogis.GeometryCollection{},
		})
	})
}
