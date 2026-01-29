package code

import (
	"code/internal/compare"
	"code/internal/formatters"
	"code/internal/parsers"
)

// GenDiff returns a formatted diff between files at path1 and path2.
func GenDiff(path1, path2, format string) (string, error) {
	var err error
	var original, changed map[string]any
	original, err = parsers.ParseDataFromFile(path1)
	if err != nil {
		return "", err
	}
	changed, err = parsers.ParseDataFromFile(path2)
	if err != nil {
		return "", err
	}
	res := compare.Compare(original, changed)
	fr := formatters.Render(res, format)
	return fr, nil
}
