package character_encoding

import "strings"

func convertChar(c uint8) byte {
	letterConvert := c - 63
	isLetter := (letterConvert >= 65 && letterConvert <= 90) || (letterConvert >= 97 && letterConvert <= 122)

	if isLetter {
		return letterConvert
	}

	if c == 0x50 {
		return 0
	}

	return 32
}

func ConvertString(s []uint8) string {
	var result [11]byte

	for i := 0; i < 11; i++ {
		result[i] = convertChar(s[i])
	}

	return strings.TrimRight(string(result[:]), "\x00")
}
