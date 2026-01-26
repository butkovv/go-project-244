package formatters

import (
	"testing"

	"code/modules/compare"

	"github.com/stretchr/testify/assert"
)

func TestFormatStylish(t *testing.T) {
	cases := []struct {
		name string
		diff map[string]compare.DiffEntry
		want string
	}{
		{
			name: "Basic JSON",
			diff: map[string]compare.DiffEntry{"common": {Status: "changed", FirstValue: map[string]interface{}{"setting1": "Value 1", "setting2": 200, "setting3": true, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}}, SecondValue: map[string]interface{}{"follow": false, "setting1": "Value 1", "setting3": interface{}(nil), "setting4": "blah blah", "setting5": map[string]interface{}{"key5": "value5"}, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}}, Level: 0, IsNested: true, Diff: map[string]compare.DiffEntry{"follow": {Status: "added", FirstValue: interface{}(nil), SecondValue: false, Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting1": {Status: "unchanged", FirstValue: "Value 1", SecondValue: "Value 1", Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting2": {Status: "removed", FirstValue: 200, SecondValue: interface{}(nil), Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting3": {Status: "changed", FirstValue: true, SecondValue: interface{}(nil), Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting4": {Status: "added", FirstValue: interface{}(nil), SecondValue: "blah blah", Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting5": {Status: "added", FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"key5": "value5"}, Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting6": {Status: "changed", FirstValue: map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}, SecondValue: map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}, Level: 1, IsNested: true, Diff: map[string]compare.DiffEntry{"doge": {Status: "changed", FirstValue: map[string]interface{}{"wow": ""}, SecondValue: map[string]interface{}{"wow": "so much"}, Level: 2, IsNested: true, Diff: map[string]compare.DiffEntry{"wow": {Status: "changed", FirstValue: "", SecondValue: "so much", Level: 3, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}, "key": {Status: "unchanged", FirstValue: "value", SecondValue: "value", Level: 2, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "ops": {Status: "added", FirstValue: interface{}(nil), SecondValue: "vops", Level: 2, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}}}, "group1": {Status: "changed", FirstValue: map[string]interface{}{"baz": "bas", "foo": "bar", "nest": map[string]interface{}{"key": "value"}}, SecondValue: map[string]interface{}{"baz": "bars", "foo": "bar", "nest": "str"}, Level: 0, IsNested: true, Diff: map[string]compare.DiffEntry{"baz": {Status: "changed", FirstValue: "bas", SecondValue: "bars", Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "foo": {Status: "unchanged", FirstValue: "bar", SecondValue: "bar", Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "nest": {Status: "changed", FirstValue: map[string]interface{}{"key": "value"}, SecondValue: "str", Level: 1, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}, "group2": {Status: "removed", FirstValue: map[string]interface{}{"abc": 12345, "deep": map[string]interface{}{"id": 45}}, SecondValue: interface{}(nil), Level: 0, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "group3": {Status: "added", FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"deep": map[string]interface{}{"id": map[string]interface{}{"number": 45}}, "fee": 100500}, Level: 0, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}},
			want: `{
  common: {
    + follow: false
      setting1: Value 1
    - setting2: 200
    - setting3: true
    + setting3: null
    + setting4: blah blah
    + setting5: {
          key5: value5
      }
      setting6: {
          doge: {
            - wow: 
            + wow: so much
          }
          key: value
        + ops: vops
      }
  }
  group1: {
    - baz: bas
    + baz: bars
      foo: bar
    - nest: {
          key: value
      }
    + nest: str
  }
- group2: {
      abc: 12345
      deep: {
          id: 45
      }
  }
+ group3: {
      deep: {
          id: {
              number: 45
          }
      }
      fee: 100500
  }
}
`,
		},
		{
			name: "Empty file",
			diff: map[string]compare.DiffEntry{"active": {Status: "removed", FirstValue: false, SecondValue: interface{}(nil), Level: 0, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "category": {Status: "removed", FirstValue: "B", SecondValue: interface{}(nil), Level: 0, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "id": {Status: "removed", FirstValue: 202, SecondValue: interface{}(nil), Level: 0, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "name": {Status: "removed", FirstValue: "beta", SecondValue: interface{}(nil), Level: 0, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "priority": {Status: "removed", FirstValue: 2, SecondValue: interface{}(nil), Level: 0, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}},
			want: `{
- active: false
- category: B
- id: 202
- name: beta
- priority: 2
}
`,
		},
		{
			name: "Both files empty",
			diff: map[string]compare.DiffEntry{},
			want: `{
}
`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FormatDiff(c.diff)
			assert.Equal(t, c.want, got, "Строки должны быть одинаковыми")
		})
	}
}
