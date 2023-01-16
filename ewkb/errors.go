package ewkb

// Error is a database error.
type Error string

const (
	// ErrWrongByteOrder occurs when byte order is not recognized.
	ErrWrongByteOrder = Error("wrong byte order")

	// ErrIncompatibleFormat occurs when geometry formats are incompatible.
	ErrIncompatibleFormat = Error("incompatible format")

	// ErrWrongGeometryType occurs when geometry type is not the expected one.
	ErrWrongGeometryType = Error("wrong geometry type")
)

func (e Error) Error() string {
	return string(e)
}
