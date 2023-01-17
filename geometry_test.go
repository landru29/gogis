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
				Coordinate: ewkb.Coordinate{
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
				CoordinateSet: ewkb.CoordinateSet{
					{
						'x': -71.060316,
						'y': 48.432044,
						'z': 10,
						'm': 30,
					},
					{
						'x': 5,
						'y': 6,
						'z': 7,
						'm': 8,
					},
				},
			},
		},
		{
			title:                "linestring empty",
			rawData:              []byte("010200000000000000"),
			expectedGeometryType: ewkb.GeometryTypeLineString,
			expectedGeometry: &ewkb.Linestring{
				CoordinateSet: ewkb.CoordinateSet{},
			},
		},
		{
			title:                "polygon",
			rawData:              []byte("01030000C002000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144004000000000000000000F03F0000000000000040000000000000084000000000000010400000000000001040000000000000144000000000000018400000000000001C400000000000001C40000000000000204000000000000022400000000000000000000000000000F03F000000000000004000000000000008400000000000001040"),
			expectedGeometryType: ewkb.GeometryTypePolygon,
			expectedGeometry: &ewkb.Polygon{
				CoordinateGroup: ewkb.CoordinateGroup{
					{
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 42.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -17.42,
							'y': 71.17,
							'z': 4,
							'm': 5,
						},
						{
							'x': -71.42,
							'y': 42.71,
							'z': 4,
							'm': 5,
						},
					},
					{
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
						},
						{
							'x': 4,
							'y': 5,
							'z': 6,
							'm': 7,
						},
						{
							'x': 7,
							'y': 8,
							'z': 9,
							'm': 0,
						},
						{
							'x': 1,
							'y': 2,
							'z': 3,
							'm': 4,
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
