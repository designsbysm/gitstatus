package main

import (
	"flag"
	"os"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func config() error {
	var changes bool
	var depth int
	var fetch bool
	var noRoot bool
	var src string

	flag.BoolVar(&changes, "changes", false, "fetch remote branches")
	flag.IntVar(&depth, "depth", 1, "max recursive folder depth")
	flag.BoolVar(&fetch, "fetch", false, "fetch remote branches")
	flag.BoolVar(&noRoot, "no-root", false, "skip root folder, only search subfolders")
	flag.StringVar(&src, "src", ".", "source folder, defaults to current directory")
	flag.Parse()

	viper.Set("changes", changes)
	viper.Set("depth", depth)
	viper.Set("fetch", fetch)
	viper.Set("noRoot", noRoot)
	viper.Set("src", src)

	// loggers
	if err := timber.New(
		os.Stdout,
		timber.LevelAll,
		"",
		timber.FlagColorful,
	); err != nil {
		return err
	}

	return nil
}
