package foo

import (
	"fmt"
	"github.com/Jeffail/tunny"
)
func Greet(name string) string {
	return fmt.Sprintf("%s, 你好！ Version 3.0.0", name)
}
