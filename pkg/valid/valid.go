package valid

import (
	"fmt"
	"net"
)

// InclusiveRangeError returns a string explanation of a numeric "must be
// between" validation failure.
func InclusiveRangeError(lo, hi int) string {
	return fmt.Sprintf(`must be between %d and %d, inclusive`, lo, hi)
}

// IsValidPortNum tests that the argument is a valid, non-zero port number.
func IsValidPortNum(port int) []string {
	if 1 <= port && port <= 65535 {
		return nil
	}
	return []string{InclusiveRangeError(1, 65535)}
}

func IsValidIP(value string) []string {
	if net.ParseIP(value) == nil {
		return []string{"must be a valid IP address, (e.g. 10.9.8.7 or 2001:db8::ffff)"}
	}
	return nil
}
