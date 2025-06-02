package main

type Status struct {
	Path     string
	Branch   string
	Modified bool
	Remote   RemoteStatus
}

type RemoteStatus int

const (
	NoRemote RemoteStatus = iota
	InSync
	LocalAhead
	RemoteAhead
	Diverged
	Gone
)
