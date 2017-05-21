package main

import (
	"fmt"

	. "github.com/ZeroneVu/go.journey/hello/common"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("hello world")

	spew.Dump(Token{Value: "99754106633f94d350db34d548d6091a"})
}
