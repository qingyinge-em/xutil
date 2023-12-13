// go test -v
package xutil

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("rand string:", GetRandomString(4))
	fmt.Println("rand int:", GetRandomInt(1, 10))
	fmt.Println("current datetime:", GetCurrentDateTime())
}
