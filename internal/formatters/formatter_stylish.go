package formatters

import (
	"fmt"
	"sort"
	"strings"

	"code/internal/compare"
)

const INDENT_LENGTH = 4

func formatMapStylish(m map[string]any, level int) string {
	mapKeys := make([]string, 0, len(m))
	for key := range m {
		mapKeys = append(mapKeys, key)
	}
	sort.Strings(mapKeys)
	res := "{\n"
	for _, key := range mapKeys {
		v := m[key]
		var value, offset, marker string
		if level == 0 {
			offset = strings.Repeat(" ", INDENT_LENGTH-2)
		} else {
			offset = strings.Repeat(" ", level*INDENT_LENGTH-2)
		}
		if mapValue, ok := v.(map[string]any); ok {
			value = formatMapStylish(mapValue, level+1)
		} else {
			value = fmt.Sprintf("%v\n", v)
		}
		marker = "  "
		res += offset + marker + key + ": " + value
	}
	offset := strings.Repeat(" ", (level-1)*INDENT_LENGTH)
	res += offset + "}\n"
	return res
}

func formatLineStylish(key string, v any, marker string, level int, offset string) string {
	var value string
	if mapValue, ok := v.(map[string]any); ok {
		value = formatMapStylish(mapValue, level+1)
	} else if v == nil {
		value = "null\n"
	} else {
		value = fmt.Sprintf("%v\n", v)
	}
	return offset + marker + key + ": " + value
}

func formatDiffStylish(diff map[string]compare.DiffEntry, level int) string {
	diffKeys := make([]string, 0, len(diff))
	for key := range diff {
		diffKeys = append(diffKeys, key)
	}
	sort.Strings(diffKeys)

	res := "{\n"
	for _, key := range diffKeys {
		v := diff[key]
		var offset string
		if level == 0 {
			offset = strings.Repeat(" ", INDENT_LENGTH-2)
		} else {
			offset = strings.Repeat(" ", level*INDENT_LENGTH-2)
		}
		var marker string
		if v.Status == compare.StatusAdded {
			marker = "+ "
			added := formatLineStylish(key, v.SecondValue, marker, level, offset)
			res += added
		}
		if v.Status == compare.StatusRemoved {
			marker = "- "
			removed := formatLineStylish(key, v.FirstValue, marker, level, offset)
			res += removed
		}
		if v.Status == compare.StatusUnchanged {
			marker = "  "
			unchanged := formatLineStylish(key, v.FirstValue, marker, level, offset)
			res += unchanged
		}
		if v.Status == compare.StatusChanged && !v.IsNested {
			marker = "- "
			removed := formatLineStylish(key, v.FirstValue, marker, level, offset)
			res += removed
			marker = "+ "
			added := formatLineStylish(key, v.SecondValue, marker, level, offset)
			res += added
		}
		if v.Status == compare.StatusChanged && v.IsNested {
			changed := formatDiffStylish(v.Diff, level+1)
			marker = "  "
			res += offset + marker + key + ": " + changed + "\n"
		}
	}
	var offset string
	if level == 0 {
		offset = strings.Repeat(" ", INDENT_LENGTH-2)
	} else {
		offset = strings.Repeat(" ", (level-1)*INDENT_LENGTH)
	}
	res += offset + "}"
	return res
}

// FormatStylish formats the diff in stylish format.
func FormatStylish(diff map[string]compare.DiffEntry) string {
	level := 1
	return formatDiffStylish(diff, level)
}
