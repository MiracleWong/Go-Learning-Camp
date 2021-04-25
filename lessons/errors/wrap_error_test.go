package main

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"testing"
)

func TestReadConfig(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		t.Logf("original err: %T --> %v\nstack trace: \n%+v\n", errors.Cause(err), errors.Cause(err), err)
	}
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()
	return nil, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	if err != nil {
		return config, errors.WithMessage(err, "could not read config")
	}
	return config, nil
}

