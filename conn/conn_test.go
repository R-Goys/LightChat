package connection

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	fmt.Println("Testing connection")
	fmt.Println("哈喽，我真的在测试哦！！！")
	pass := true
	if pass {
		panic("Testing connection")
	}
}
