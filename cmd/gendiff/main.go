package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"code"

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
			format := cmd.String("format")
			diff, err := code.GenDiff(path1, path2, format)
			if err != nil {
				return fmt.Errorf("gendiff: %w", err)
			}
			fmt.Print(diff)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
