package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	cases := []struct {
		name  string
		path1 string
		path2 string
		want  string
	}{
		{
			name:  "Basic JSON",
			path1: "../../testdata/fixture/file1.json",
			path2: "../../testdata/fixture/file2.json",
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
			name:  "Basic YAML",
			path1: "../../testdata/fixture/file1.yaml",
			path2: "../../testdata/fixture/file2.yaml",
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
			name:  "Empty file",
			path1: "../../testdata/fixture/file4.json",
			path2: "../../testdata/fixture/file5.json",
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
			name:  "Both files empty",
			path1: "../../testdata/fixture/file5.json",
			path2: "../../testdata/fixture/file5.json",
			want: `{
}
`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var err error
			var j1, j2 map[string]any
			j1, err = ParseDataFromFile(c.path1)
			if err != nil {
				t.Fatalf("не ожидали ошибку, получили %q", err.Error())
			}
			j2, err = ParseDataFromFile(c.path2)
			if err != nil {
				t.Fatalf("не ожидали ошибку, получили %q", err.Error())
			}
			res := Compare(j1, j2)
			got := FormatDiff(res)
			assert.Equal(t, c.want, got, "Строки должны быть одинаковыми")
		})
	}
}
