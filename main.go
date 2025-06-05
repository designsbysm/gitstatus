package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// go run . --src ./test --depth 1
func main() {
	if err := config(); err != nil {
		fmt.Println(err)
		return
	}

	depth := viper.GetInt("depth")
	fetch := viper.GetBool("fetch")
	src := viper.GetString("src")

	folders, err := gatherFolders(src, depth)
	if err != nil {
		fmt.Println(err)
		return
	}

	if fetch {
		err = fetchRemotes(folders)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	statuses, err := getStatus(folders)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputStatuses(statuses)
}
