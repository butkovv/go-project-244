package formatters

import (
	"fmt"
	"sort"
	"strings"

	"code/modules/compare"
)

func formatValuePlain(value any) any {
	if stringValue, ok := value.(string); ok {
		value = "'" + stringValue + "'"
	}
	if compare.IsMap(value) {
		value = "[complex value]"
	}
	if value == nil {
		value = "null"
	}
	return value
}

func formatDiffPlain(diff map[string]compare.DiffEntry, parentDir string) string {
	diffKeys := make([]string, 0, len(diff))
	for key := range diff {
		diffKeys = append(diffKeys, key)
	}
	sort.Strings(diffKeys)

	res := []string{}
	for _, key := range diffKeys {
		v := diff[key]
		property := key
		if len(parentDir) > 0 {
			property = parentDir + "." + property
		}

		v.FirstValue = formatValuePlain(v.FirstValue)
		v.SecondValue = formatValuePlain(v.SecondValue)
		if v.Status == compare.StatusAdded {
			added := fmt.Sprintf("Property '%s' was added with value: %v", property, v.SecondValue)
			res = append(res, added)
		}
		if v.Status == compare.StatusRemoved {
			removed := fmt.Sprintf("Property '%s' was removed", property)
			res = append(res, removed)
		}
		if v.Status == compare.StatusChanged && !v.IsNested {
			changed := fmt.Sprintf("Property '%s' was updated. From %v to %v", property, v.FirstValue, v.SecondValue)
			res = append(res, changed)
		}
		if v.Status == compare.StatusChanged && v.IsNested {
			changed := formatDiffPlain(v.Diff, property)
			res = append(res, changed)
		}
	}
	return strings.Join(res, "\n")
}

func FormatPlain(diff map[string]compare.DiffEntry) string {
	parentDir := ""
	return formatDiffPlain(diff, parentDir)
}
