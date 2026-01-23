package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

func ParseDataFromFile(path string) (map[string]any, error) {
	var res map[string]any

	data, err := os.ReadFile(path)
	if err != nil {
		return res, err
	}

	ext := filepath.Ext(path)

	switch ext {
	case ".json":
		err = json.Unmarshal(data, &res)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &res)
	}
	if err != nil {
		return res, err
	}

	return res, nil
}
