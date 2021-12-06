package util

import "strings"

// SplitArgs splt args into slice, and remove if item is empty
func SplitArgs(args string) []string {
	items := strings.Split(args, " ")

	list := make([]string, 0, len(items))

	for _, it := range items {
		trim := strings.TrimSpace(it)
		if len(trim) > 0 {
			list = append(list, trim)
		}
	}

	return list
}
