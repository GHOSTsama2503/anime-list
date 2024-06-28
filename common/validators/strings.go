package validators

import "strings"

func StringIsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
