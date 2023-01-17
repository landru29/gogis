package gogis_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const dbQueryString = "SELECT coordinate FROM geometry"

func TestGeometryScan(t *testing.T) {
	fixtures := []struct {
		title                string
		rawData              []byte
		expectedGeometryType ewkb.GeometryType
		expectedGeometry     ewkb.Geometry
	}{
		{
			title:                "point",
			rawData:              []byte("01010000C03CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40"),
			expectedGeometryType: ewkb.GeometryTypePoint,
			expectedGeometry: &ewkb.Point{
				Coordinates: map[byte]float64{
					'x': -71.060316,
					'y': 48.432044,
					'z': 10.0,
					'm': 30.0,
				},
			},
		},
		{
			title:                "linestring",
			rawData:              []byte("01020000C0020000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040"),
			expectedGeometryType: ewkb.GeometryTypeLineString,
			expectedGeometry: &ewkb.Linestring{
				Points: []ewkb.Point{
					{
						Coordinates: map[byte]float64{
							'x': -71.060316,
							'y': 48.432044,
							'z': 10,
							'm': 30,
						},
					},
					{
						Coordinates: map[byte]float64{
							'x': 5,
							'y': 6,
							'z': 7,
							'm': 8,
						},
					},
				},
			},
		},
	}

	for idx := range fixtures {
		fixture := fixtures[idx]

		t.Run(fixture.title, func(t *testing.T) {

			dbSQL, mock, err := sqlmock.New(
				sqlmock.QueryMatcherOption(sqlmock.QueryMatcherFunc(matcher)),
			)
			require.NoError(t, err)

			mock.ExpectQuery(dbQueryString).WillReturnRows(
				sqlmock.NewRows([]string{"coordinate"}).
					AddRow(fixture.rawData))

			rows, err := dbSQL.Query(dbQueryString)
			require.NoError(t, err)
			require.NoError(t, rows.Err())

			defer func() {
				_ = rows.Close()
			}()

			geometry := gogis.NewGeometry()

			if rows.Next() {
				require.NoError(t, rows.Scan(geometry))
			}

			assert.Equal(t, fixture.expectedGeometryType, geometry.Type)
			assert.Equal(t, fixture.expectedGeometry, geometry.Geometry)
		})
	}
}
