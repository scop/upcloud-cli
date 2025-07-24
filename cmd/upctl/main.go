package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/UpCloudLtd/upcloud-cli/v3/internal/core"
)

func main() {
	fmt.Println(parseAllocateBad1("10"))
	exitCode := core.Execute()
	os.Exit(exitCode)
}

func parseAllocateBad1(wanted string) int32 {
	parsed, err := strconv.Atoi(wanted)
	if err != nil {
		panic(err)
	}
	return int32(parsed)
}
