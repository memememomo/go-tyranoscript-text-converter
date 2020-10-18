package main

import (
	"fmt"
	"os"
)

const ConfigFileName = "studio_config.json"

func main() {
	fmt.Println(readAndProcess(ConfigFileName, os.Args[1]))
}
