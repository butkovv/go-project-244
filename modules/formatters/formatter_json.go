package formatters

import (
	"encoding/json"
	"fmt"
	"sort"

	"code/modules/compare"
)

type JSONDiffElement struct {
	Key      string             `json:"key"`
	Status   compare.DiffStatus `json:"status"`
	OldValue any                `json:"old_value,omitempty"`
	NewValue any                `json:"new_value,omitempty"`
	Children []JSONDiffElement  `json:"children,omitempty"`
}

func formatDiffJSON(diff map[string]compare.DiffEntry) []JSONDiffElement {
	d := []JSONDiffElement{}
	diffKeys := make([]string, 0, len(diff))
	for key := range diff {
		diffKeys = append(diffKeys, key)
	}
	sort.Strings(diffKeys)
	for _, key := range diffKeys {
		el := JSONDiffElement{
			Key:      key,
			Status:   diff[key].Status,
			OldValue: diff[key].FirstValue,
			NewValue: diff[key].SecondValue,
			Children: formatDiffJSON(diff[key].Diff),
		}
		d = append(d, el)
	}
	return d
}

func FormatJson(diff map[string]compare.DiffEntry) string {
	res := make(map[string]any)
	res["root"] = formatDiffJSON(diff)
	jsonData, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		fmt.Printf("error marshaling json: %v", err.Error())
	}
	return string(jsonData)
}
