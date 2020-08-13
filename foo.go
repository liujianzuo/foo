package foo

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("%s, 你好！ Version 5.0.0", name)
}
