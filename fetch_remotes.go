package main

import (
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

func fetchRemotes(folders []string) error {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	// s.Prefix = "Fetching "
	s.UpdateCharSet(spinner.CharSets[11])
	s.Start()

	for _, path := range folders {
		cmd := exec.Command("git", "fetch", "--all", "--prune", "--tags", "--force")
		cmd.Dir = path
		err := cmd.Run()

		if err != nil {
			return err
		}
	}

	s.Stop()
	return nil
}
