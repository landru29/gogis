package ewkb

// Layout is the EWKB layout.
type Layout uint8

const (
	layoutXY   Layout = 0
	layoutXYM  Layout = 1
	layoutXYZ  Layout = 2
	layoutXYZM Layout = 3
)

// LayoutWith builds a layout.
func LayoutWith(hasM bool, hasZ bool) Layout {
	out := layoutXY
	if hasM {
		out += layoutXYM
	}

	if hasZ {
		out += layoutXYZ
	}

	return out
}

// Format is the coordinate format.
func (l Layout) Format() string {
	switch l {
	case layoutXY:
		return "xy"
	case layoutXYM:
		return "xym"
	case layoutXYZ:
		return "xyz"
	case layoutXYZM:
		return "xyzm"
	}

	return ""
}

// Size is the number of coordinates in the layout.
func (l Layout) Size() uint32 {
	return uint32(len(l.Format()))
}

// Uint32 convert layout in uint32.
func (l Layout) Uint32() uint32 {
	return uint32(l) << 30 //nolint: gomnd
}

func newLayoutFrom(indexes []byte) Layout {
	var (
		hasM bool
		hasZ bool
	)

	for _, idx := range indexes {
		if idx == 'z' {
			hasZ = true
		}

		if idx == 'm' {
			hasM = true
		}
	}

	if hasM && hasZ {
		return layoutXYZM
	}

	if hasZ {
		return layoutXYZ
	}

	if hasM {
		return layoutXYM
	}

	return layoutXY
}
