package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

func getStatus(folders []string) (statuses []Status, err error) {
	for _, path := range folders {
		status := Status{
			Path: path,
		}

		cmd := exec.Command("git", "status", "--short", "--branch")
		cmd.Dir = path
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}

		lines := bytes.Split(out, []byte("\n"))
		branch := lines[0]
		modified := len(lines) > 2

		re := regexp.MustCompile(`^## ([^\.]+)(?:\.\.\.)?`)
		matches := re.FindSubmatch(branch)
		status.Branch = strings.TrimSpace(string(matches[1]))

		re = regexp.MustCompile(`\.\.\.`)
		hasRemote := len(re.Find(branch)) > 0

		re = regexp.MustCompile(`\[.*ahead.*\]`)
		ahead := len(re.Find(branch)) > 0

		re = regexp.MustCompile(`\[.*behind.*\]`)
		behind := len(re.Find(branch)) > 0

		re = regexp.MustCompile(`\[gone\]`)
		gone := len(re.Find(branch)) > 0

		if modified {
			status.Modified = true
		}

		if hasRemote {
			if gone {
				status.Remote = Gone
			} else if ahead && behind {
				status.Remote = Diverged
			} else if ahead {
				status.Remote = LocalAhead
			} else if behind {
				status.Remote = RemoteAhead
			} else {
				status.Remote = InSync
			}
		} else {
			status.Remote = NoRemote
		}

		statuses = append(statuses, status)
	}

	return
}
