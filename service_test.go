package main

import (
	"fmt"
	"testing"
)

func TestReadAndProcess(t *testing.T) {
	fmt.Println(readAndProcess("studio_config.json", "target_file1.ks"))
}
