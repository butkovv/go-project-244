package compare

import (
	"maps"
	"reflect"
	"slices"
)

type DiffStatus string

type DiffEntry struct {
	Status      string
	FirstValue  any
	SecondValue any
	IsNested    bool
	Diff        map[string]DiffEntry
}

const (
	StatusAdded     DiffStatus = "+"
	StatusRemoved   DiffStatus = "-"
	StatusUnchanged DiffStatus = " "
	StatusChanged   DiffStatus = " "
)

func IsMap(v any) bool {
	if v == nil {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func Compare(first, second map[string]any) map[string]DiffEntry {
	res := make(map[string]DiffEntry)

	keys := append(slices.Collect(maps.Keys(first)), slices.Collect(maps.Keys(second))...)
	slices.Sort(keys)
	keys = slices.Compact(keys)
	slices.Sort(keys)
	for _, k := range keys {
		v1, ok1 := first[k]
		v2, ok2 := second[k]
		if !ok1 && ok2 {
			res[k] = DiffEntry{Status: "added", FirstValue: nil, SecondValue: v2, IsNested: false, Diff: nil}
		}
		if !ok2 && ok1 {
			res[k] = DiffEntry{Status: "removed", FirstValue: v1, SecondValue: nil, IsNested: false, Diff: nil}
		}
		if ok1 && ok2 && IsMap(v1) && IsMap(v2) {
			m1, ok1 := v1.(map[string]any)
			m2, ok2 := v2.(map[string]any)
			if reflect.DeepEqual(m1, m2) {
				res[k] = DiffEntry{Status: "unchanged", FirstValue: v1, SecondValue: v2, IsNested: false, Diff: nil}
			} else if ok1 && ok2 {
				res[k] = DiffEntry{Status: "changed", FirstValue: v1, SecondValue: v2, IsNested: true, Diff: Compare(m1, m2)}
			}
		}
		if ok1 && ok2 && ((IsMap(v1) && !IsMap(v2)) || (!IsMap(v1) && IsMap(v2)) || (!IsMap(v1) && !IsMap(v2) && v1 != v2)) {
			res[k] = DiffEntry{Status: "changed", FirstValue: v1, SecondValue: v2, IsNested: false, Diff: nil}
		}
		if ok1 && ok2 && !IsMap(v1) && !IsMap(v2) && v1 == v2 {
			res[k] = DiffEntry{Status: "unchanged", FirstValue: v1, SecondValue: v2, IsNested: false, Diff: nil}
		}
	}
	return res
}
