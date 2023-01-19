package ewkb_test

import (
	"math"
	"testing"

	"github.com/landru29/gogis/ewkb"
	"github.com/stretchr/testify/assert"
)

func TestCoordinateIsNull(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		assert.True(t, ewkb.Coordinate{}.IsNull())
	})

	t.Run("all NaN", func(t *testing.T) {
		assert.True(t, ewkb.Coordinate{
			'x': math.NaN(),
			'y': math.NaN(),
		}.IsNull())
	})

	t.Run("one NaN", func(t *testing.T) {
		assert.True(t, ewkb.Coordinate{
			'x': math.NaN(),
			'y': 42.0,
		}.IsNull())
	})

	t.Run("no NaN", func(t *testing.T) {
		assert.False(t, ewkb.Coordinate{
			'x': 35.3,
			'y': 42.0,
		}.IsNull())
	})
}

func TestNewNullCoordinate(t *testing.T) {
	coord := ewkb.NewNullCoordinate(ewkb.LayoutWith(true, false))

	assert.True(t, coord['x'] != coord['x'])
	assert.True(t, coord['y'] != coord['y'])
	assert.True(t, coord['z'] == coord['z'])
	assert.True(t, coord['m'] != coord['m'])
}
