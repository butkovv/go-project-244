package main

import (
	"fmt"
	"sort"
	"strings"
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
			offset = ""
		} else {
			offset = strings.Repeat(" ", level*INDENT_LENGTH)
		}
		if mapValue, ok := v.(map[string]any); ok {
			value = formatMapStylish(mapValue, level+1)
		} else {
			value = fmt.Sprintf("%v\n", v)
		}
		marker = "  "
		res += offset + marker + key + ": " + value
	}
	offset := strings.Repeat(" ", (level)*INDENT_LENGTH-2)
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

func FormatStylish(diff map[string]DiffEntry, level int) string {
	diffKeys := make([]string, 0, len(diff))
	for key := range diff {
		diffKeys = append(diffKeys, key)
	}
	sort.Strings(diffKeys)

	res := "{\n"
	for _, key := range diffKeys {
		v := diff[key]
		var offset string
		if v.Level == 0 {
			offset = ""
		} else {
			offset = strings.Repeat(" ", v.Level*INDENT_LENGTH)
		}
		var marker string
		if v.Status == "added" {
			marker = "+ "
			added := formatLineStylish(key, v.SecondValue, marker, v.Level, offset)
			res += added
		}
		if v.Status == "removed" {
			marker = "- "
			removed := formatLineStylish(key, v.FirstValue, marker, v.Level, offset)
			res += removed
		}
		if v.Status == "unchanged" {
			marker = "  "
			unchanged := formatLineStylish(key, v.FirstValue, marker, v.Level, offset)
			res += unchanged
		}
		if v.Status == "changed" && !v.IsNested {
			marker = "- "
			removed := formatLineStylish(key, v.FirstValue, marker, v.Level, offset)
			res += removed
			marker = "+ "
			added := formatLineStylish(key, v.SecondValue, marker, v.Level, offset)
			res += added
		}
		if v.Status == "changed" && v.IsNested {
			changed := FormatStylish(v.Diff, v.Level+1)
			marker = "  "
			res += offset + marker + key + ": " + changed
		}
	}
	var offset string
	if level == 0 {
		offset = ""
	} else {
		offset = strings.Repeat(" ", level*INDENT_LENGTH-2)
	}
	res += offset + "}\n"
	return res
}

func FormatDiff(diff map[string]DiffEntry) string {
	level := 0
	return FormatStylish(diff, level)
}
