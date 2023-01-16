// Package gogis is the SQL driver for PostGIS (https://postgis.net/).
// It implements some of the PostGIS geometry, and let users to implements
// their own types:
//
//	import (
//	    "io"
//
//	    "github.com/landru29/gogis/ewkb"
//	)
//
//	type Custom struct {
//	    databytes []byte
//	}
//
//	func (c *Custom) UnmarshalEWBK(ewkb.ExtendedWellKnownBytes) error {
//	   data, err := io.ReadAll()
//	   c.dataBytes = data
//	   return err
//	}
//
//	func (c Custom) MarshalEWBK(ewkb.ExtendedWellKnownBytesHeader) ([]byte, error) {
//	    return c.dataBytes, nil
//	}
//
//	func (c Custom) Header() ewkb.ExtendedWellKnownBytesHeader {
//	    return ewkb.ExtendedWellKnownBytesHeader{}
//	}
//
//	func (c Custom) Type() ewkb.GeometryType {
//	    return ewkb.GeometryType(42)
//	}
//
//	// CustomSQL is to used with "sql" package.
//	type CustomSQL Custom
//
//	func (c *CustomSQL) Scan(value interface{}) error {
//		custo := Custom{}
//
//		if err := ewkb.Unmarshal(&custo, value); err != nil {
//			return err
//		}
//
//		*c = CustomSQL(custo)
//
//		return nil
//	}
//
//	func (c CustomSQL) Value() (driver.Value, error) {
//	    return ewkb.Marshal(Custom(c))
//	}
package gogis
