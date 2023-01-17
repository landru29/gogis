# gogis

This package is an implementation of the PostGis database driver.

## Example

```golang
func main() {
	// CREATE TABLE IF NOT EXISTS points (
    //   coordinate GEOMETRY
	// );
	ctx := context.Background()

	// Connect to database
	db, err := sql.Open("postgres", "postgresql://tester:tester@localhost/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.PrepareContext(ctx, `
		INSERT INTO points(
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

	_, err = stmt.ExecContext(ctx,
		gogis.Point{
			Coordinates: map[byte]float64{
				'x': 42.42,
				'y': 24.24,
			},
			SRID: ewkb.WithSRID(ewkb.SystemReferenceWGS84),
		},
	)

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

func (c Custom) MarshalEWBK(ewkb.ExtendedWellKnownBytesHeader) ([]byte, error) {
    return nil, nil
}

func (c Custom) Header() ewkb.ExtendedWellKnownBytesHeader {
    return ewkb.ExtendedWellKnownBytesHeader{}
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