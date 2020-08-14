package foo

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"runtime"
)
func Greet(name string) string {
	numCPUs := runtime.NumCPU()
	fmt.Println("numCPUs: ", numCPUs)
	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var result []byte

		// TODO: Something CPU heavy with payload

		return result
	})
	defer pool.Close()
	fmt.Println("sdfsf: ", pool)
	return fmt.Sprintf("%s, 你好！ Version 3.0.0", name)
}
