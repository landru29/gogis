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

func TestGeometryScan(t *testing.T) {
	fixtures := []struct {
		title                string
		rawData              []byte
		expectedGeometryType ewkb.GeometryType
		expectedGeometry     ewkb.Geometry
		//scanner              sql.Scanner
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
			expectedGeometry: &ewkb.LineString{
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
			expectedGeometry: &ewkb.LineString{
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
		{
			title:                "multilinestring",
			rawData:              []byte("01050000C00200000001020000C002000000F6285C8FC23545403D0AD7A3703D38C01F85EB51B81E4540EC51B81E856B38C0000000000000144000000000000018400000000000001C40000000000000204001020000C0020000003D0AD7A370CD6140A4703D0AD7837AC048E17A14AEC761407B14AE47E11A5FC00000000000002E40000000000000304000000000000031400000000000003240"),
			expectedGeometryType: ewkb.GeometryTypeMultiLineString,
			expectedGeometry: &ewkb.MultiLineString{
				LineStrings: []ewkb.LineString{
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 42.42,
								'y': -24.24,
								'z': 42.24,
								'm': -24.42,
							},
							{
								'x': 5,
								'y': 6,
								'z': 7,
								'm': 8,
							},
						},
					},
					{
						CoordinateSet: []ewkb.Coordinate{
							{
								'x': 142.42,
								'y': -424.24,
								'z': 142.24,
								'm': -124.42,
							},
							{
								'x': 15,
								'y': 16,
								'z': 17,
								'm': 18,
							},
						},
					},
				},
			},
		},
		{
			title:                "multipolygon",
			rawData:              []byte("010600000002000000010300000002000000040000001F85EB51B81E1CC0A4703D0AD7A30040EC51B81E85EB0F4085EB51B81E851DC0E17A14AE47E1E23F14AE47E17A1411401F85EB51B81E1CC0A4703D0AD7A30040040000003D0AD7A370BD2240B81E85EB51384B4085EB51B81E9555403D0AD7A3707D3FC0EC51B81E852B4340D7A3703D0A2757C03D0AD7A370BD2240B81E85EB51384B400103000000020000000400000048E17A14AE0731C0295C8FC2F52828407B14AE47E1FA2B40E17A14AE476131C0AE47E17A142E25400AD7A3703D8A2C4048E17A14AE0731C0295C8FC2F5282840040000001F85EB51B85E3340AE47E17A144E6340C3F5285C8F4A674048E17A14AE6F60C07B14AE47E14A6140EC51B81E851368C01F85EB51B85E3340AE47E17A144E6340"),
			expectedGeometryType: ewkb.GeometryTypeMultiPolygon,
			expectedGeometry: &ewkb.MultiPolygon{
				Polygons: []ewkb.Polygon{
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -7.03,
									'y': 2.08,
								},
								{
									'x': 3.99,
									'y': -7.38,
								},
								{
									'x': 0.59,
									'y': 4.27,
								},
								{
									'x': -7.03,
									'y': 2.08,
								},
							},
							{
								{
									'x': 9.37,
									'y': 54.44,
								},
								{
									'x': 86.33,
									'y': -31.49,
								},
								{
									'x': 38.34,
									'y': -92.61,
								},
								{
									'x': 9.37,
									'y': 54.44,
								},
							},
						},
					},
					{
						CoordinateGroup: ewkb.CoordinateGroup{
							{
								{
									'x': -17.03,
									'y': 12.08,
								},
								{
									'x': 13.99,
									'y': -17.38,
								},
								{
									'x': 10.59,
									'y': 14.27,
								},
								{
									'x': -17.03,
									'y': 12.08,
								},
							},
							{
								{
									'x': 19.37,
									'y': 154.44,
								},
								{
									'x': 186.33,
									'y': -131.49,
								},
								{
									'x': 138.34,
									'y': -192.61,
								},
								{
									'x': 19.37,
									'y': 154.44,
								},
							},
						},
					},
				},
			},
		},
		{
			title:                "triangle",
			rawData:              []byte("01110000C001000000040000007B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440EC51B81E856B31C0F6285C8FC215454000000000000010400000000000001440EC51B81E856B31C07B14AE47E1CA5140000000000000104000000000000014407B14AE47E1DA51C07B14AE47E15A454000000000000010400000000000001440"),
			expectedGeometryType: ewkb.GeometryTypeTriangle,
			expectedGeometry: &ewkb.Triangle{
				CoordinateSet: ewkb.CoordinateSet{
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
			},
		},
		{
			title:                "circularstring",
			rawData:              []byte("01080000C0030000003CDBA337DCC351C06D37C1374D37484000000000000024400000000000003E40000000000000144000000000000018400000000000001C400000000000002040000000000000F03F000000000000004000000000000008400000000000001040"),
			expectedGeometryType: ewkb.GeometryTypeCircularString,
			expectedGeometry: &ewkb.CircularString{
				CoordinateSet: []ewkb.Coordinate{
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
					{
						'x': 1,
						'y': 2,
						'z': 3,
						'm': 4,
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

		case *ewkb.LineString:
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
