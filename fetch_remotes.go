package main

import (
	"os/exec"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/designsbysm/timber/v2"
)

func fetch(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command("git", "fetch", "--all", "--prune", "--tags", "--force")
	cmd.Dir = path

	if err := cmd.Run(); err != nil {
		timber.Error(err)
	}
}

func fetchRemotes(folders []string) error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()

	var wg sync.WaitGroup

	for _, path := range folders {
		wg.Add(1)
		go fetch(path, &wg)
	}

	wg.Wait()
	s.Stop()

	return nil
}
