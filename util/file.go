package util

import (
	"os"
	"os/exec"
	"path/filepath"
)

func ExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	re, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return re, nil
}
