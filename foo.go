package foo

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("%s, 你好！ Version 2.0.0", name)
}
