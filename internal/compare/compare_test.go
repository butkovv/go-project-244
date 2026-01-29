package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	cases := []struct {
		name     string
		dataset1 map[string]interface{}
		dataset2 map[string]interface{}
		want     map[string]DiffEntry
	}{
		{
			name:     "Basic",
			dataset1: map[string]interface{}{"common": map[string]interface{}{"setting1": "Value 1", "setting2": 200, "setting3": true, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}}, "group1": map[string]interface{}{"baz": "bas", "foo": "bar", "nest": map[string]interface{}{"key": "value"}}, "group2": map[string]interface{}{"abc": 12345, "deep": map[string]interface{}{"id": 45}}},
			dataset2: map[string]interface{}{"common": map[string]interface{}{"follow": false, "setting1": "Value 1", "setting3": interface{}(nil), "setting4": "blah blah", "setting5": map[string]interface{}{"key5": "value5"}, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}}, "group1": map[string]interface{}{"baz": "bars", "foo": "bar", "nest": "str"}, "group3": map[string]interface{}{"deep": map[string]interface{}{"id": map[string]interface{}{"number": 45}}, "fee": 100500}},
			want:     map[string]DiffEntry{"common": {Status: StatusChanged, FirstValue: map[string]interface{}{"setting1": "Value 1", "setting2": 200, "setting3": true, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}}, SecondValue: map[string]interface{}{"follow": false, "setting1": "Value 1", "setting3": interface{}(nil), "setting4": "blah blah", "setting5": map[string]interface{}{"key5": "value5"}, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}}, IsNested: true, Diff: map[string]DiffEntry{"follow": {Status: StatusAdded, FirstValue: interface{}(nil), SecondValue: false, IsNested: false, Diff: map[string]DiffEntry(nil)}, "setting1": {Status: StatusUnchanged, FirstValue: "Value 1", SecondValue: "Value 1", IsNested: false, Diff: map[string]DiffEntry(nil)}, "setting2": {Status: StatusRemoved, FirstValue: 200, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]DiffEntry(nil)}, "setting3": {Status: StatusChanged, FirstValue: true, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]DiffEntry(nil)}, "setting4": {Status: StatusAdded, FirstValue: interface{}(nil), SecondValue: "blah blah", IsNested: false, Diff: map[string]DiffEntry(nil)}, "setting5": {Status: StatusAdded, FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"key5": "value5"}, IsNested: false, Diff: map[string]DiffEntry(nil)}, "setting6": {Status: StatusChanged, FirstValue: map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}, SecondValue: map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}, IsNested: true, Diff: map[string]DiffEntry{"doge": {Status: StatusChanged, FirstValue: map[string]interface{}{"wow": ""}, SecondValue: map[string]interface{}{"wow": "so much"}, IsNested: true, Diff: map[string]DiffEntry{"wow": {Status: StatusChanged, FirstValue: "", SecondValue: "so much", IsNested: false, Diff: map[string]DiffEntry(nil)}}}, "key": {Status: StatusUnchanged, FirstValue: "value", SecondValue: "value", IsNested: false, Diff: map[string]DiffEntry(nil)}, "ops": {Status: StatusAdded, FirstValue: interface{}(nil), SecondValue: "vops", IsNested: false, Diff: map[string]DiffEntry(nil)}}}}}, "group1": {Status: StatusChanged, FirstValue: map[string]interface{}{"baz": "bas", "foo": "bar", "nest": map[string]interface{}{"key": "value"}}, SecondValue: map[string]interface{}{"baz": "bars", "foo": "bar", "nest": "str"}, IsNested: true, Diff: map[string]DiffEntry{"baz": {Status: StatusChanged, FirstValue: "bas", SecondValue: "bars", IsNested: false, Diff: map[string]DiffEntry(nil)}, "foo": {Status: StatusUnchanged, FirstValue: "bar", SecondValue: "bar", IsNested: false, Diff: map[string]DiffEntry(nil)}, "nest": {Status: StatusChanged, FirstValue: map[string]interface{}{"key": "value"}, SecondValue: "str", IsNested: false, Diff: map[string]DiffEntry(nil)}}}, "group2": {Status: StatusRemoved, FirstValue: map[string]interface{}{"abc": 12345, "deep": map[string]interface{}{"id": 45}}, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]DiffEntry(nil)}, "group3": {Status: StatusAdded, FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"deep": map[string]interface{}{"id": map[string]interface{}{"number": 45}}, "fee": 100500}, IsNested: false, Diff: map[string]DiffEntry(nil)}},
		},
		{
			name:     "Empty dataset",
			dataset1: map[string]interface{}{"active": false, "category": "B", "id": 202, "name": "beta", "priority": 2},
			dataset2: map[string]interface{}{},
			want:     map[string]DiffEntry{"active": {Status: StatusRemoved, FirstValue: false, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]DiffEntry(nil)}, "category": {Status: StatusRemoved, FirstValue: "B", SecondValue: interface{}(nil), IsNested: false, Diff: map[string]DiffEntry(nil)}, "id": {Status: StatusRemoved, FirstValue: 202, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]DiffEntry(nil)}, "name": {Status: StatusRemoved, FirstValue: "beta", SecondValue: interface{}(nil), IsNested: false, Diff: map[string]DiffEntry(nil)}, "priority": {Status: StatusRemoved, FirstValue: 2, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]DiffEntry(nil)}},
		},
		{
			name:     "Both datasets empty",
			dataset1: map[string]interface{}{},
			dataset2: map[string]interface{}{},
			want:     map[string]DiffEntry{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Compare(c.dataset1, c.dataset2)
			assert.Equal(t, c.want, got, "Diffs must match")
		})
	}
}
