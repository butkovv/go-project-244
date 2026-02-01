package code

import (
	"code/internal/compare"
	"code/internal/formatters"
	"code/internal/parsers"
)

// GenDiff builds a diff between two files and returns its string representation.
//
// The function reads data from files at path1 and path2, parses it into a data
// structure, compares the resulting values, and formats the result according to
// the specified format. The comparison is performed as a deep comparison of the
// data structure, while the rendering choice is delegated to the formatter.
//
// The format parameter specifies the output format ("plain",
// "stylish", "json". The exact set of supported values depends on the formatters
// package).
//
// Returned errors relate to reading/parsing the input files; behavior for an
// unsupported format is determined by formatters.Render.
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
