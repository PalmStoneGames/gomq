package zmtp

import "fmt"

func toNullPaddedString(str string, dst []byte) {
	if len(str) > len(dst) {
		panic(fmt.Sprintf("Null padded string of length %v cannot accomodate content of length %v", len(str), len(str)))
	}

	copy(dst, str)
}

func fromNullPaddedString(slice []byte) string {
	str := ""
	for _, b := range slice {
		if b == 0 {
			break
		}

		str += string(b)
	}

	return str
}

func toByteBool(b bool) byte {
	if b {
		return byte(0x01)
	} else {
		return byte(0x00)
	}
}

func fromByteBool(b byte) (bool, error) {
	if b == byte(0x00) {
		return false, nil
	}

	if b == byte(0x01) {
		return true, nil
	}

	return false, fmt.Errorf("Invalid boolean byte %b", b)
}
