package formatters

import (
	"testing"

	"code/internal/compare"

	"github.com/stretchr/testify/assert"
)

func TestFormatStylish(t *testing.T) {
	cases := []struct {
		name string
		diff map[string]compare.DiffEntry
		want string
	}{
		{
			name: "Basic",
			diff: map[string]compare.DiffEntry{"common": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"setting1": "Value 1", "setting2": 200, "setting3": true, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}}, SecondValue: map[string]interface{}{"follow": false, "setting1": "Value 1", "setting3": interface{}(nil), "setting4": "blah blah", "setting5": map[string]interface{}{"key5": "value5"}, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}}, IsNested: true, Diff: map[string]compare.DiffEntry{"follow": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: false, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting1": {Status: compare.StatusUnchanged, FirstValue: "Value 1", SecondValue: "Value 1", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting2": {Status: compare.StatusRemoved, FirstValue: 200, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting3": {Status: compare.StatusChanged, FirstValue: true, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting4": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: "blah blah", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting5": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"key5": "value5"}, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting6": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}, SecondValue: map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}, IsNested: true, Diff: map[string]compare.DiffEntry{"doge": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"wow": ""}, SecondValue: map[string]interface{}{"wow": "so much"}, IsNested: true, Diff: map[string]compare.DiffEntry{"wow": {Status: compare.StatusChanged, FirstValue: "", SecondValue: "so much", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}, "key": {Status: compare.StatusUnchanged, FirstValue: "value", SecondValue: "value", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "ops": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: "vops", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}}}, "group1": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"baz": "bas", "foo": "bar", "nest": map[string]interface{}{"key": "value"}}, SecondValue: map[string]interface{}{"baz": "bars", "foo": "bar", "nest": "str"}, IsNested: true, Diff: map[string]compare.DiffEntry{"baz": {Status: compare.StatusChanged, FirstValue: "bas", SecondValue: "bars", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "foo": {Status: compare.StatusUnchanged, FirstValue: "bar", SecondValue: "bar", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "nest": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"key": "value"}, SecondValue: "str", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}, "group2": {Status: compare.StatusRemoved, FirstValue: map[string]interface{}{"abc": 12345, "deep": map[string]interface{}{"id": 45}}, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "group3": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"deep": map[string]interface{}{"id": map[string]interface{}{"number": 45}}, "fee": 100500}, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}},
			want: "{\n    common: {\n      + follow: false\n        setting1: Value 1\n      - setting2: 200\n      - setting3: true\n      + setting3: null\n      + setting4: blah blah\n      + setting5: {\n            key5: value5\n        }\n        setting6: {\n            doge: {\n              - wow: \n              + wow: so much\n            }\n            key: value\n          + ops: vops\n        }\n    }\n    group1: {\n      - baz: bas\n      + baz: bars\n        foo: bar\n      - nest: {\n            key: value\n        }\n      + nest: str\n    }\n  - group2: {\n        abc: 12345\n        deep: {\n            id: 45\n        }\n    }\n  + group3: {\n        deep: {\n            id: {\n                number: 45\n            }\n        }\n        fee: 100500\n    }\n}",
		},
		{
			name: "Empty dataset",
			diff: map[string]compare.DiffEntry{"active": {Status: compare.StatusRemoved, FirstValue: false, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "category": {Status: compare.StatusRemoved, FirstValue: "B", SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "id": {Status: compare.StatusRemoved, FirstValue: 202, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "name": {Status: compare.StatusRemoved, FirstValue: "beta", SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "priority": {Status: compare.StatusRemoved, FirstValue: 2, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}},
			want: "{\n  - active: false\n  - category: B\n  - id: 202\n  - name: beta\n  - priority: 2\n}",
		},
		{
			name: "Empty diff",
			diff: map[string]compare.DiffEntry{},
			want: `{
}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FormatStylish(c.diff)
			assert.Equal(t, c.want, got, "Strings must match")
		})
	}
}

func TestFormatPlain(t *testing.T) {
	cases := []struct {
		name string
		diff map[string]compare.DiffEntry
		want string
	}{
		{
			name: "Basic",
			diff: map[string]compare.DiffEntry{"common": {Status: "changed", FirstValue: map[string]interface{}{"setting1": "Value 1", "setting2": 200, "setting3": true, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}}, SecondValue: map[string]interface{}{"follow": false, "setting1": "Value 1", "setting3": interface{}(nil), "setting4": "blah blah", "setting5": map[string]interface{}{"key5": "value5"}, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}}, IsNested: true, Diff: map[string]compare.DiffEntry{"follow": {Status: "added", FirstValue: interface{}(nil), SecondValue: false, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting1": {Status: "unchanged", FirstValue: "Value 1", SecondValue: "Value 1", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting2": {Status: "removed", FirstValue: 200, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting3": {Status: "changed", FirstValue: true, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting4": {Status: "added", FirstValue: interface{}(nil), SecondValue: "blah blah", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting5": {Status: "added", FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"key5": "value5"}, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting6": {Status: "changed", FirstValue: map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}, SecondValue: map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}, IsNested: true, Diff: map[string]compare.DiffEntry{"doge": {Status: "changed", FirstValue: map[string]interface{}{"wow": ""}, SecondValue: map[string]interface{}{"wow": "so much"}, IsNested: true, Diff: map[string]compare.DiffEntry{"wow": {Status: "changed", FirstValue: "", SecondValue: "so much", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}, "key": {Status: "unchanged", FirstValue: "value", SecondValue: "value", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "ops": {Status: "added", FirstValue: interface{}(nil), SecondValue: "vops", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}}}, "group1": {Status: "changed", FirstValue: map[string]interface{}{"baz": "bas", "foo": "bar", "nest": map[string]interface{}{"key": "value"}}, SecondValue: map[string]interface{}{"baz": "bars", "foo": "bar", "nest": "str"}, IsNested: true, Diff: map[string]compare.DiffEntry{"baz": {Status: "changed", FirstValue: "bas", SecondValue: "bars", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "foo": {Status: "unchanged", FirstValue: "bar", SecondValue: "bar", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "nest": {Status: "changed", FirstValue: map[string]interface{}{"key": "value"}, SecondValue: "str", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}, "group2": {Status: "removed", FirstValue: map[string]interface{}{"abc": 12345, "deep": map[string]interface{}{"id": 45}}, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "group3": {Status: "added", FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"deep": map[string]interface{}{"id": map[string]interface{}{"number": 45}}, "fee": 100500}, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}},
			want: `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FormatPlain(c.diff)
			assert.Equal(t, c.want, got, "Strings must match")
		})
	}
}

func TestFormatJSON(t *testing.T) {
	cases := []struct {
		name string
		diff map[string]compare.DiffEntry
		want string
	}{
		{
			name: "Basic",
			diff: map[string]compare.DiffEntry{"common": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"setting1": "Value 1", "setting2": 200, "setting3": true, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}}, SecondValue: map[string]interface{}{"follow": false, "setting1": "Value 1", "setting3": interface{}(nil), "setting4": "blah blah", "setting5": map[string]interface{}{"key5": "value5"}, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}}, IsNested: true, Diff: map[string]compare.DiffEntry{"follow": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: false, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting1": {Status: compare.StatusUnchanged, FirstValue: "Value 1", SecondValue: "Value 1", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting2": {Status: compare.StatusRemoved, FirstValue: 200, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting3": {Status: compare.StatusChanged, FirstValue: true, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting4": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: "blah blah", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting5": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"key5": "value5"}, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "setting6": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}, SecondValue: map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}, IsNested: true, Diff: map[string]compare.DiffEntry{"doge": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"wow": ""}, SecondValue: map[string]interface{}{"wow": "so much"}, IsNested: true, Diff: map[string]compare.DiffEntry{"wow": {Status: compare.StatusChanged, FirstValue: "", SecondValue: "so much", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}, "key": {Status: compare.StatusUnchanged, FirstValue: "value", SecondValue: "value", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "ops": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: "vops", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}}}, "group1": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"baz": "bas", "foo": "bar", "nest": map[string]interface{}{"key": "value"}}, SecondValue: map[string]interface{}{"baz": "bars", "foo": "bar", "nest": "str"}, IsNested: true, Diff: map[string]compare.DiffEntry{"baz": {Status: compare.StatusChanged, FirstValue: "bas", SecondValue: "bars", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "foo": {Status: compare.StatusUnchanged, FirstValue: "bar", SecondValue: "bar", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "nest": {Status: compare.StatusChanged, FirstValue: map[string]interface{}{"key": "value"}, SecondValue: "str", IsNested: false, Diff: map[string]compare.DiffEntry(nil)}}}, "group2": {Status: compare.StatusRemoved, FirstValue: map[string]interface{}{"abc": 12345, "deep": map[string]interface{}{"id": 45}}, SecondValue: interface{}(nil), IsNested: false, Diff: map[string]compare.DiffEntry(nil)}, "group3": {Status: compare.StatusAdded, FirstValue: interface{}(nil), SecondValue: map[string]interface{}{"deep": map[string]interface{}{"id": map[string]interface{}{"number": 45}}, "fee": 100500}, IsNested: false, Diff: map[string]compare.DiffEntry(nil)}},
			want: `{
 "root": [
  {
   "key": "common",
   "status": "changed",
   "old_value": {
    "setting1": "Value 1",
    "setting2": 200,
    "setting3": true,
    "setting6": {
     "doge": {
      "wow": ""
     },
     "key": "value"
    }
   },
   "new_value": {
    "follow": false,
    "setting1": "Value 1",
    "setting3": null,
    "setting4": "blah blah",
    "setting5": {
     "key5": "value5"
    },
    "setting6": {
     "doge": {
      "wow": "so much"
     },
     "key": "value",
     "ops": "vops"
    }
   },
   "children": [
    {
     "key": "follow",
     "status": "added",
     "new_value": false
    },
    {
     "key": "setting1",
     "status": "unchanged",
     "old_value": "Value 1",
     "new_value": "Value 1"
    },
    {
     "key": "setting2",
     "status": "removed",
     "old_value": 200
    },
    {
     "key": "setting3",
     "status": "changed",
     "old_value": true
    },
    {
     "key": "setting4",
     "status": "added",
     "new_value": "blah blah"
    },
    {
     "key": "setting5",
     "status": "added",
     "new_value": {
      "key5": "value5"
     }
    },
    {
     "key": "setting6",
     "status": "changed",
     "old_value": {
      "doge": {
       "wow": ""
      },
      "key": "value"
     },
     "new_value": {
      "doge": {
       "wow": "so much"
      },
      "key": "value",
      "ops": "vops"
     },
     "children": [
      {
       "key": "doge",
       "status": "changed",
       "old_value": {
        "wow": ""
       },
       "new_value": {
        "wow": "so much"
       },
       "children": [
        {
         "key": "wow",
         "status": "changed",
         "old_value": "",
         "new_value": "so much"
        }
       ]
      },
      {
       "key": "key",
       "status": "unchanged",
       "old_value": "value",
       "new_value": "value"
      },
      {
       "key": "ops",
       "status": "added",
       "new_value": "vops"
      }
     ]
    }
   ]
  },
  {
   "key": "group1",
   "status": "changed",
   "old_value": {
    "baz": "bas",
    "foo": "bar",
    "nest": {
     "key": "value"
    }
   },
   "new_value": {
    "baz": "bars",
    "foo": "bar",
    "nest": "str"
   },
   "children": [
    {
     "key": "baz",
     "status": "changed",
     "old_value": "bas",
     "new_value": "bars"
    },
    {
     "key": "foo",
     "status": "unchanged",
     "old_value": "bar",
     "new_value": "bar"
    },
    {
     "key": "nest",
     "status": "changed",
     "old_value": {
      "key": "value"
     },
     "new_value": "str"
    }
   ]
  },
  {
   "key": "group2",
   "status": "removed",
   "old_value": {
    "abc": 12345,
    "deep": {
     "id": 45
    }
   }
  },
  {
   "key": "group3",
   "status": "added",
   "new_value": {
    "deep": {
     "id": {
      "number": 45
     }
    },
    "fee": 100500
   }
  }
 ]
}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FormatJson(c.diff)
			assert.Equal(t, c.want, got, "Strings must match")
		})
	}
}
