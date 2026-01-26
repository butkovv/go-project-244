package code

import (
	"code/modules/compare"
	"code/modules/formatters"
	"code/modules/parsers"
)

func GenDiff(path1, path2 string) (string, error) {
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
	fr := formatters.FormatDiff(res)
	return fr, nil
}
