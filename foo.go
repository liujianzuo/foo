package foo

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"runtime"
)
func Greet(name string) string {
	numCPUs := runtime.NumCPU()
	fmt.Println("numCPUs: ", numCPUs)
	return fmt.Sprintf("%s, 你好！ Version 3.0.0", name)
}
