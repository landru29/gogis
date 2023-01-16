package gogis_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func Example_scanPoint() {
	ctx := context.Background()
	// CREATE TABLE IF NOT EXISTS geometries (
	// 	coordinate GEOMETRY
	// );

	// Connect to database.
	db, err := sql.Open("postgres", "postgresql://tester:tester@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the query.
	rows, err := db.QueryContext(ctx, `
		SELECT
			coordinate
		FROM geometries
	`)
	if err != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}

		if pnt.Valid {
			output = append(output, gogis.Point(pnt.Point))
		}
	}

	// Display the result.
	fmt.Println(output)
}

func Example_scanLinestring() {
	ctx := context.Background()
	// CREATE TABLE IF NOT EXISTS geometries (
	// 	coordinate GEOMETRY
	// );

	// Connect to database.
	db, err := sql.Open("postgres", "postgresql://tester:tester@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the query.
	rows, err := db.QueryContext(ctx, `
		SELECT
			coordinate
		FROM geometries
	`)
	if err != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	output := []gogis.Linestring{}

	// Read data.
	for rows.Next() {
		pnt := gogis.NullLinestring{}

		err = rows.Scan(&pnt)
		if err != nil {
			log.Fatal(err)
		}

		if pnt.Valid {
			output = append(output, gogis.Linestring(pnt.Linestring))
		}
	}

	// Display the result.
	fmt.Println(output)
}

func Example_insertPoint() {
	ctx := context.Background()
	// CREATE TABLE IF NOT EXISTS geometries (
	// 	coordinate GEOMETRY
	// );

	// Connect to database.
	db, err := sql.Open("postgres", "postgresql://tester:tester@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	defer func() {
		_ = stmt.Close()
	}()

	// Execute the query.
	if _, err = stmt.ExecContext(ctx,
		gogis.Point{
			Coordinates: map[byte]float64{
				'x': 42.42,
				'y': 24.24,
			},
			SRID: ewkb.WithSRID(ewkb.SystemReferenceWGS84),
		},
	); err != nil {
		log.Fatal(err)
	}
}

func Example_insertLinestring() {
	ctx := context.Background()
	// CREATE TABLE IF NOT EXISTS geometries (
	// 	coordinate GEOMETRY
	// );

	// Connect to database.
	db, err := sql.Open("postgres", "postgresql://tester:tester@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	defer func() {
		_ = stmt.Close()
	}()

	// Execute the query.
	if _, err = stmt.ExecContext(ctx,
		gogis.Linestring{
			{
				Coordinates: map[byte]float64{
					'x': 42.42,
					'y': 24.24,
				},
				SRID: ewkb.WithSRID(ewkb.SystemReferenceWGS84),
			},
			{
				Coordinates: map[byte]float64{
					'x': 10,
					'y': 30,
				},
				SRID: ewkb.WithSRID(ewkb.SystemReferenceWGS84),
			},
		},
	); err != nil {
		log.Fatal(err)
	}
}
