# gogis

This package is an implementation of the PostGis database driver. Refer to documentation [https://pkg.go.dev/github.com/landru29/gogis](https://pkg.go.dev/github.com/landru29/gogis)

The following types are currently implemented:
* Point
* LineString
* Polygon
* MultiPoint
* MultiLineString
* MultiPolygon
* Triangle
* CircularString

## Example

```golang
package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/landru29/gogis"
	"github.com/landru29/gogis/ewkb"
)

func main() {
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

```

## Implements your own type

It's quite easy to implement your own type, based on `Extended Well Known Byte` format.

### Implementations on EWKB level

First implement the following:

```golang
type Custom struct {}

func (c *Custom) UnmarshalEWBK(ewkb.ExtendedWellKnownBytes) error {
    return nil
}

func (c Custom) MarshalEWBK(binary.ByteOrder) ([]byte, error) {
    return nil, nil
}

func (c Custom) SystemReferenceID() *SystemReferenceID {
    return nil
}

func (c Custom) Layout() Layout {
	return ewkb.Layout(0)
}

func (c Custom) Type() ewkb.GeometryType {
    return ewkb.GeometryType(42)
}
```

### Implementations on Database level

Now, you just have to implement the following:

```golang
type CustomSQL Custom

func (c *CustomSQL) Scan(value interface{}) error {
    custo := Custom{}

    if err := ewkb.Unmarshal(&custo, value); err != nil {
        return err
    }

    *c = CustomSQL(custo)

    return nil
}

func (c CustomSQL) Value() (driver.Value, error) {
    return ewkb.Marshal(Custom(c))
}
```