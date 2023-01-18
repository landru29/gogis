package gogis_test

import (
	"context"
	"database/sql"
	"fmt"
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
		{
			title:                "multipoint",
			rawData:              []byte("01040000C00400000001010000C07B14AE47E1DA51C07B14AE47E15A45400000000000001040000000000000144001010000C0EC51B81E856B31C0F6285C8FC21545400000000000001040000000000000144001010000C0EC51B81E856B31C07B14AE47E1CA51400000000000001040000000000000144001010000C07B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440"),
			expectedGeometryType: ewkb.GeometryTypeMultiPoint,
			expectedGeometry: &ewkb.MultiPoint{
				Points: []ewkb.Point{
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

// This example shows how to read any geometry from database.
func Example_scanAny() {
	// Launch database:
	// $> docker run --name db -p 5432:5432 -e POSTGRES_PASSWORD=tester -e POSTGRES_USER=tester -e POSTGRES_DB=test -d postgis/postgis:15-master
	//
	// Create the table:
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "CREATE TABLE IF NOT EXISTS geometries (coordinate GEOMETRY);" -t
	//
	// Insert data:
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "INSERT INTO geometries(coordinate) VALUES (ST_GeomFromText('POINT ZM(10 20 30 50)', 4326))" -t
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "INSERT INTO geometries(coordinate) VALUES (ST_GeomFromText('MULTIPOINT((-71.42 42.71),(-17.42 42.17),(-17.42 71.17),(-71.42 42.71))', 4326))" -t
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "INSERT INTO geometries(coordinate) VALUES (ST_GeomFromText('POLYGON Z((-71.42 42.71 4,-17.42 42.17 4,-17.42 71.17 4,-71.42 42.71 4),(1 2 3,4 5 6,7 8 9,1 2 3))', 4326))" -t
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "INSERT INTO geometries(coordinate) VALUES (ST_GeomFromText('LINESTRING (-71.060316 48.432044, 5 6, 42 24)', 4326))" -t
	//
	// Do not forget the imports:
	// import (
	// 	"context"
	// 	"database/sql"

	// 	_ "github.com/lib/pq"

	// 	"github.com/landru29/gogis"
	// 	"github.com/landru29/gogis/ewkb"
	// )
	ctx := context.Background()

	// Connect to database.
	db, err := sql.Open("postgres", "postgresql://tester:tester@localhost/test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Prepare the query.
	rows, err := db.QueryContext(ctx, `
		SELECT
			coordinate
		FROM geometries
	`)
	if err != nil {
		panic(err)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	// Read data.
	for rows.Next() {
		// Here you can inject your own type:
		// geometry := gogis.NewFeometry(gogis.WithWellKnownGeometry(&myCustom1{}, &myCustom2{}))
		geometry := gogis.NewGeometry()

		err = rows.Scan(geometry)
		if err != nil {
			panic(err)
		}

		switch data := geometry.Geometry.(type) {
		case *ewkb.Point:
			// process point
			fmt.Printf("* point %+v\n", data)

		case *ewkb.Linestring:
			// process linestring
			fmt.Printf("* linestring %+v\n", data)

		case *ewkb.Polygon:
			// process polygon
			fmt.Printf("* polygon %+v\n", data)

		case *ewkb.MultiPoint:
			// process multipoint
			fmt.Printf("* multipoint %+v\n", data)

		default:
			// If you have your custom types, just add:
			// case *myCustom1:
			// process myCustom1
			// case *myCustom2:
			// process myCustom2
		}
	}
}
