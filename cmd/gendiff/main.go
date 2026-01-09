/*
hexlet-path-size calculates size of a given path.

Given a file, it calculates the size of the file; given a directory, it calculates the size of all files in
that directory. By default it does not scan the directory recursively and does not include hidden files.

Usage:

	hexlet-path-size [path] [flags]

The flags are:

	   --recursive, -r
		 									recursive size of directories (default: false)
	   --human, -H
		 									human-readable sizes (auto-select unit) (default: false)
	   --all, -a
		 									include hidden files and directories (default: false)
	   --help, -h
		 									show help
*/
package main

import (
	"context"
	"log/slog"
	"os"

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
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error(err.Error())
	}
}
