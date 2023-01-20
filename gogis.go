// Package gogis is the SQL driver for PostGIS (https://postgis.net/).
// It implements some of the PostGIS geometry:
//
//	import (
//		"context"
//		"database/sql"
//		_ "github.com/lib/pq"
//		"github.com/landru29/gogis"
//		"github.com/landru29/gogis/ewkb"
//	)
//
//	func main() {
//		ctx := context.Background()
//
//		db, err := sql.Open("postgres", "postgresql://tester:tester@localhost/test?sslmode=disable")
//		if err != nil {
//			panic(err)
//		}
//
//		rows, err := db.QueryContext(ctx, `
//			SELECT
//				coordinate
//			FROM geometries
//		`)
//		if err != nil {
//			panic(err)
//		}
//
//		if err := rows.Err(); err != nil {
//			panic(err)
//		}
//
//		defer func() {
//			_ = rows.Close()
//		}()
//
//		output := []gogis.Point{}
//
//		for rows.Next() {
//			pnt := gogis.NullPoint{}
//
//			err = rows.Scan(&pnt)
//			if err != nil {
//				panic(err)
//			}
//
//			if pnt.Valid {
//				output = append(output, pnt.Point)
//			}
//		}
//
//		fmt.Println(output)
//	}
//
// It also let users to implements
// their own types:
//
//		import (
//		    "io"
//
//		    "github.com/landru29/gogis/ewkb"
//		)
//
//		type Custom struct {
//		    databytes []byte
//		}
//
//		func (c *Custom) UnmarshalEWBK(ewkb.ExtendedWellKnownBytes) error {
//		   data, err := io.ReadAll()
//		   c.dataBytes = data
//		   return err
//		}
//
//		func (c Custom) MarshalEWBK(binary.ByteOrder) ([]byte, error) {
//		    return c.dataBytes, nil
//		}
//
//	    func (c Custom) Layout() ewkb.Layout {
//		    return ewkb.Layout(0)
//	    }
//
//		func (c Custom) Type() ewkb.GeometryType {
//		    return ewkb.GeometryType(42)
//		}
//
//		// CustomSQL is to used with "sql" package.
//		type CustomSQL Custom
//
//		func (c *CustomSQL) Scan(value interface{}) error {
//			custo := Custom{}
//
//			if err := ewkb.Unmarshal(&custo, value); err != nil {
//				return err
//			}
//
//			*c = CustomSQL(custo)
//
//			return nil
//		}
//
//		func (c CustomSQL) Value() (driver.Value, error) {
//		    return ewkb.Marshal(Custom(c))
//		}
package gogis

// ModelConverter is the converter from EWKB to Model.
type ModelConverter interface {
	FromEWKB(geometry interface{}) error
}
