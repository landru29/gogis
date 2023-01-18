package gogis_test

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func Example_scanPoint() {
	// Launch database:
	// $> docker run --name db -p 5432:5432 -e POSTGRES_PASSWORD=tester -e POSTGRES_USER=tester -e POSTGRES_DB=test -d postgis/postgis:15-master
	//
	// Create the table:
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "CREATE TABLE IF NOT EXISTS geometries (coordinate GEOMETRY);" -t
	//
	// Insert data:
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "INSERT INTO geometries(coordinate) VALUES (ST_GeomFromText('POINT ZM(10 20 30 50)', 4326))" -t
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

	output := []gogis.Point{}

	// Read data.
	for rows.Next() {
		pnt := gogis.NullPoint{}

		err = rows.Scan(&pnt)
		if err != nil {
			panic(err)
		}

		if pnt.Valid {
			output = append(output, gogis.Point(pnt.Point))
		}
	}

	// Display the result.
	fmt.Println(output)
}

func Example_scanLineString() {
	// Launch database:
	// $> docker run --name db -p 5432:5432 -e POSTGRES_PASSWORD=tester -e POSTGRES_USER=tester -e POSTGRES_DB=test -d postgis/postgis:15-master
	//
	// Create the table:
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "CREATE TABLE IF NOT EXISTS geometries (coordinate GEOMETRY);" -t
	//
	// Insert data:
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

	output := []gogis.LineString{}

	// Read data.
	for rows.Next() {
		pnt := gogis.NullLineString{}

		err = rows.Scan(&pnt)
		if err != nil {
			panic(err)
		}

		if pnt.Valid {
			output = append(output, gogis.LineString(pnt.LineString))
		}
	}

	// Display the result.
	fmt.Println(output)
}

func Example_insertPoint() {
	// Launch database:
	// $> docker run --name db -p 5432:5432 -e POSTGRES_PASSWORD=tester -e POSTGRES_USER=tester -e POSTGRES_DB=test -d postgis/postgis:15-master
	//
	// Create the table:
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "CREATE TABLE IF NOT EXISTS geometries (coordinate GEOMETRY);" -t
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
	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO geometries(
			coordinate
		) VALUES(
			$1
		)
	`)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = stmt.Close()
	}()

	// Execute the query.
	if _, err = stmt.ExecContext(ctx,
		gogis.Point{
			Coordinate: ewkb.Coordinate{
				'x': 42.42,
				'y': 24.24,
			},
			SRID: ewkb.WithSRID(ewkb.SystemReferenceWGS84),
		},
	); err != nil {
		panic(err)
	}
}

func Example_insertLineString() {
	// Launch database:
	// $> docker run --name db -p 5432:5432 -e POSTGRES_PASSWORD=tester -e POSTGRES_USER=tester -e POSTGRES_DB=test -d postgis/postgis:15-master
	//
	// Create the table:
	// $> docker exec -i db psql -h 0.0.0.0 -p 5432 -U tester -d test -c "CREATE TABLE IF NOT EXISTS geometries (coordinate GEOMETRY);" -t
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
	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO geometries(
			coordinate
		) VALUES(
			$1
		)
	`)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = stmt.Close()
	}()

	// Execute the query.
	if _, err = stmt.ExecContext(ctx,
		gogis.LineString{
			{
				Coordinate: ewkb.Coordinate{
					'x': 42.42,
					'y': 24.24,
				},
				SRID: ewkb.WithSRID(ewkb.SystemReferenceWGS84),
			},
			{
				Coordinate: ewkb.Coordinate{
					'x': 10,
					'y': 30,
				},
				SRID: ewkb.WithSRID(ewkb.SystemReferenceWGS84),
			},
		},
	); err != nil {
		panic(err)
	}
}
