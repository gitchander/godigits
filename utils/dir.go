package utils

import (
	"fmt"
	"os"
)

func MkdirIfNotExist(name string) error {
	fi, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(name, os.ModePerm)
		}
		return err
	}
	if !(fi.IsDir()) {
		return fmt.Errorf("name %q is not a directory", name)
	}
	return nil
}

func MustMkdirIfNotExist(name string) {
	err := MkdirIfNotExist(name)
	if err != nil {
		panic(err)
	}
}
