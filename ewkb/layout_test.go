package ewkb_test

import (
	"testing"

	"github.com/landru29/gogis/ewkb"
	"github.com/stretchr/testify/assert"
)

func TestLayoutFormat(t *testing.T) {
	assert.Equal(t, "xy", ewkb.Layout(0).Format())
	assert.Equal(t, "xym", ewkb.Layout(1).Format())
	assert.Equal(t, "xyz", ewkb.Layout(2).Format())
	assert.Equal(t, "xyzm", ewkb.Layout(3).Format())
}

func TestLayoutSize(t *testing.T) {
	assert.Equal(t, uint32(2), ewkb.Layout(0).Size())
	assert.Equal(t, uint32(3), ewkb.Layout(1).Size())
	assert.Equal(t, uint32(3), ewkb.Layout(2).Size())
	assert.Equal(t, uint32(4), ewkb.Layout(3).Size())
}

func TestLayoutWith(t *testing.T) {
	assert.Equal(t, ewkb.Layout(0), ewkb.LayoutWith(false, false))
	assert.Equal(t, ewkb.Layout(1), ewkb.LayoutWith(true, false))
	assert.Equal(t, ewkb.Layout(2), ewkb.LayoutWith(false, true))
	assert.Equal(t, ewkb.Layout(3), ewkb.LayoutWith(true, true))
}
