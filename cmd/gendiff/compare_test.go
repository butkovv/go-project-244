package main

import "testing"

func TestCompare(t *testing.T) {
	cases := []struct {
		name  string
		path1 string
		path2 string
		want  string
	}{
		{
			name:  "Basic",
			path1: "../../testdata/fixture/file1.json",
			path2: "../../testdata/fixture/file2.json",
			want: `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
`,
		},
		{
			name:  "Basic 2",
			path1: "../../testdata/fixture/file3.json",
			path2: "../../testdata/fixture/file4.json",
			want: `{
    active: false
  - category: A
  + category: B
  - id: 101
  + id: 202
  - name: alpha
  + name: beta
  + priority: 2
  - score: 88.5
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
			j1, err = ParseJsonFromFile(c.path1)
			if err != nil {
				t.Fatalf("не ожидали ошибку, получили %q", err.Error())
			}
			j2, err = ParseJsonFromFile(c.path2)
			if err != nil {
				t.Fatalf("не ожидали ошибку, получили %q", err.Error())
			}
			got := Compare(j1, j2)
			if got != c.want {
				t.Errorf(`Compare(%v, %v) = %v, ожидали %v`, j1, j2, got, c.want)
			}
		})
	}
}
