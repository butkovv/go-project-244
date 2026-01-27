package formatters

import (
	"code/modules/compare"
)

func Render(diff map[string]compare.DiffEntry, format string) string {
	if format == "plain" {
		return FormatPlain(diff)
	}
	if format == "stylish" {
		return FormatStylish(diff)
	}
	if format == "json" {
		return FormatJson(diff)
	}
	return FormatStylish(diff)
}
