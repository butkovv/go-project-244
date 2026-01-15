package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path1 := cmd.Args().Get(0)
			path2 := cmd.Args().Get(1)
			ext1 := filepath.Ext(path1)
			ext2 := filepath.Ext(path2)
			if ext1 != ext2 {
				return errors.New("файлы имеют разные расширения")
			}
			if ext1 == ".json" {
				var err error
				var json1, json2 map[string]any
				json1, err = ParseJsonFromFile(path1)
				if err != nil {
					return err
				}
				json2, err = ParseJsonFromFile(path2)
				if err != nil {
					return err
				}
				res := Compare(json1, json2)
				fmt.Print(res)
			}
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error(err.Error())
	}
}
