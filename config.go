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
	var install bool
	var noRoot bool
	var pull bool
	var src string

	flag.BoolVar(&changes, "changes", false, "fetch remote branches")
	flag.IntVar(&depth, "depth", 0, "max recursive folder depth")
	flag.BoolVar(&fetch, "fetch", false, "fetch remote branches")
	flag.BoolVar(&install, "install", false, "install develkpment dependencies")
	flag.BoolVar(&noRoot, "no-root", false, "skip root folder, only search subfolders")
	flag.BoolVar(&pull, "pull", false, "pull changes from remote branches")
	flag.StringVar(&src, "src", ".", "source folder")
	flag.Parse()

	viper.Set("changes", changes)
	viper.Set("depth", depth)
	viper.Set("fetch", fetch)
	viper.Set("install", install)
	viper.Set("noRoot", noRoot)
	viper.Set("pull", pull)
	viper.Set("src", src)

	if src == "." && len(flag.Args()) > 0 {
		viper.Set("src", flag.Args()[0])
	}

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
