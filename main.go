package main

import (
	"os"
)

func main() {
	basePath := "testbase"
	_ = os.MkdirAll(basePath, 0700)
	defer func() {
		_ = os.RemoveAll(basePath)
	}()
}
