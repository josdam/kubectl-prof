package fake

import (
	"fmt"
)

type RuntimeFake struct {
	RootFileSystemLocationResultError bool
	PIDResultError                    bool
}

func NewRuntimeFake() *RuntimeFake {
	return &RuntimeFake{}
}

func (r *RuntimeFake) WithRootFileSystemLocationResultError() *RuntimeFake {
	r.RootFileSystemLocationResultError = true
	return r
}

func (r *RuntimeFake) WithPIDResultError() *RuntimeFake {
	r.PIDResultError = true
	return r
}

func (r *RuntimeFake) RootFileSystemLocation(containerID string) (string, error) {
	if r.RootFileSystemLocationResultError {
		return "", fmt.Errorf("fake RootFileSystemLocation with error")
	}
	return fmt.Sprintf("/root/fs/%s", containerID), nil
}

func (r *RuntimeFake) PID(containerID string) (string, error) {
	if r.PIDResultError {
		return "", fmt.Errorf("fake PID with error")
	}
	return fmt.Sprintf("PID_%s", containerID), nil
}
