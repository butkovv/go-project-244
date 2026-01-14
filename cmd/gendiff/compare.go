package main

import (
	"fmt"
	"maps"
	"slices"
)

type DiffStatus string

const (
	StatusAdded     DiffStatus = "+"
	StatusRemoved   DiffStatus = "-"
	StatusUnchanged DiffStatus = " "
)

func createLine(key string, value any, status DiffStatus) string {
	return fmt.Sprintf("  %s %s: %v\n", status, key, value)
}

func Compare(first, second map[string]any) string {
	res := "{\n"

	keys := append(slices.Collect(maps.Keys(first)), slices.Collect(maps.Keys(second))...)
	slices.Sort(keys)
	keys = slices.Compact(keys)
	slices.Sort(keys)

	for _, k := range keys {
		v1, ok1 := first[k]
		v2, ok2 := second[k]
		if !ok1 && ok2 {
			res += createLine(k, v2, StatusAdded)
		}
		if !ok2 && ok1 {
			res += createLine(k, v1, StatusRemoved)
		}
		if ok1 && ok2 && v1 != v2 {
			res += createLine(k, v1, StatusRemoved)
			res += createLine(k, v2, StatusAdded)
		}
		if ok1 && ok2 && v1 == v2 {
			res += createLine(k, v1, StatusUnchanged)
		}
	}
	res += "}\n"
	return res
}
