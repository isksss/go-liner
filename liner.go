package liner

import (
	"strings"
)

// Make Line
func MakeLine(moji string, len int) string {
	hashes := strings.Repeat(moji, len)
	return hashes
}
