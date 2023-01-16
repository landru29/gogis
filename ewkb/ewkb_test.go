package ewkb_test

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"testing"

	"github.com/landru29/gogis/ewkb"
	"github.com/stretchr/testify/require"
)

type configurator func(*ewkb.ExtendedWellKnownBytesHeader)

func newExtendedWellKnownBytesHeader(opt ...configurator) ewkb.ExtendedWellKnownBytesHeader {
	output := ewkb.ExtendedWellKnownBytesHeader{
		SRID:      nil,
		ByteOrder: binary.LittleEndian,
		Type:      ewkb.GeometryType(0),
		Layout:    ewkb.Layout(0),
	}

	for _, conf := range opt {
		conf(&output)
	}

	return output
}

func newExtendedWellKnownBytes(t *testing.T, ewkbStr string, opt ...configurator) ewkb.ExtendedWellKnownBytes {
	t.Helper()

	output := ewkb.ExtendedWellKnownBytes{
		ExtendedWellKnownBytesHeader: newExtendedWellKnownBytesHeader(opt...),
	}

	if ewkbStr == "" {
		output.IsNil = true
	} else {
		data, err := hex.DecodeString(ewkbStr)
		require.NoError(t, err)

		output.DataStream = bytes.NewBuffer(data)
	}

	return output
}

func withSRID(srid ewkb.SystemReferenceID) configurator {
	return func(header *ewkb.ExtendedWellKnownBytesHeader) {
		header.SRID = &srid
	}
}

func withByteOrder(byteOrder binary.ByteOrder) configurator {
	return func(header *ewkb.ExtendedWellKnownBytesHeader) {
		header.ByteOrder = byteOrder
	}
}

func withType(geometryType ewkb.GeometryType) configurator {
	return func(header *ewkb.ExtendedWellKnownBytesHeader) {
		header.Type = geometryType
	}
}

func withLayout(layout ewkb.Layout) configurator {
	return func(header *ewkb.ExtendedWellKnownBytesHeader) {
		header.Layout = layout
	}
}
