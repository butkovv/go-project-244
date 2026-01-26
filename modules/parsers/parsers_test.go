package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFromFile(t *testing.T) {
	cases := []struct {
		name string
		path string
		want map[string]interface{}
	}{
		{
			name: "Basic JSON",
			path: "../../testdata/fixture/file1.json",
			want: map[string]interface{}{"common": map[string]interface{}{"setting1": "Value 1", "setting2": float64(200), "setting3": true, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": ""}, "key": "value"}}, "group1": map[string]interface{}{"baz": "bas", "foo": "bar", "nest": map[string]interface{}{"key": "value"}}, "group2": map[string]interface{}{"abc": float64(12345), "deep": map[string]interface{}{"id": float64(45)}}},
		},
		{
			name: "Basic YAML",
			path: "../../testdata/fixture/file2.yaml",
			want: map[string]interface{}{"common": map[string]interface{}{"follow": false, "setting1": "Value 1", "setting3": interface{}(nil), "setting4": "blah blah", "setting5": map[string]interface{}{"key5": "value5"}, "setting6": map[string]interface{}{"doge": map[string]interface{}{"wow": "so much"}, "key": "value", "ops": "vops"}}, "group1": map[string]interface{}{"baz": "bars", "foo": "bar", "nest": "str"}, "group3": map[string]interface{}{"deep": map[string]interface{}{"id": map[string]interface{}{"number": uint64(45)}}, "fee": uint64(100500)}},
		},
		{
			name: "Empty JSON",
			path: "../../testdata/fixture/file5.json",
			want: map[string]interface{}{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ParseDataFromFile(c.path)
			if err != nil {
				t.Fatalf("Error parsing file, %v", err.Error())
			}
			assert.Equal(t, c.want, got, "Maps must match")
		})
	}
}
