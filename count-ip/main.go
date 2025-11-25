package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Dir(b)

	logFilePath := filepath.Join(path, "service.log")	
}


func 