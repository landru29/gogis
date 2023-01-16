package ewkb

import "encoding/hex"

func fromHex(bytes []byte) ([]byte, error) {
	byteCount, err := hex.Decode(bytes, bytes)
	if err != nil {
		return nil, err
	}

	if byteCount == 0 {
		// field is null
		return nil, nil
	}

	return bytes[:byteCount], nil
}

func toHex(bytes []byte) []byte {
	output := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(output, bytes)

	return output
}
