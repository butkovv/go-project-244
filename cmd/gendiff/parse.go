package main

import (
	// "path/filepath"
	"encoding/json"
	"os"
)

func ParseJsonFromFile(path string) (map[string]any, error) {
	// if filepath.IsLocal(path) {
	// 	absPath, err := filepath.Abs(path)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	path = absPath
	// }
	var res map[string]any
	data, err := os.ReadFile(path)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
