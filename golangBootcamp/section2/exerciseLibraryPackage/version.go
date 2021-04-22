package version

import (
	"fmt"
	"runtime"
)

func CurrentGoV() string {
	return runtime.Version()
}

func PrintGoV() {
	fmt.Println(runtime.Version())
}
