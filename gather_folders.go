package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func gatherFolders(src string, depth int) (files []string, err error) {
	noRoot := viper.GetBool("noRoot")

	err = filepath.WalkDir(src, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if dir.IsDir() && (strings.Contains(path, ".git")) {
			if noRoot && filepath.Dir(path) == src {
				return fs.SkipDir
			}

			files = append(files, filepath.Dir(path))

			return fs.SkipDir
		} else if dir.IsDir() && depth >= 0 && len(strings.Split(path, string(os.PathSeparator))) > depth {
			return fs.SkipDir
		}

		return nil
	})

	return
}
