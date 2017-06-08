package main

import (
	"fmt"
	"reflect"

	. "github.com/ZeroneVu/go.journey/hello/common"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("hello world")

	spew.Dump(Token{Value: "99754106633f94d350db34d548d6091a"})

	onCallBack(1, Executable(func(args ...interface{}) {
		spew.Dump(args)
		fmt.Println("Foo =", fmt.Sprint(args...))
	}))

	onCallBack(2, Executable(func(args ...interface{}) {
		spew.Dump(args)
		fmt.Println("Foo =", fmt.Sprint(args...))
	}))

	onCallBack(3, Executable(func(args ...interface{}) {
		spew.Dump(args)
		fmt.Println("Foo =", fmt.Sprint(args...))

		// checking https://stackoverflow.com/a/44027953/2540986 for more information
		// basically []interface{} and []string has different layout under memory presentation
		// so it's no otherway to use interface.
		switch reflect.TypeOf(args[0]).Kind() {
		case reflect.Slice:

			t := reflect.ValueOf(args[0])
			s := make([]string, t.Len())

			for i := 0; i < t.Len(); i++ {
				fmt.Println(t.Index(i))
				spew.Dump(t.Index(i))
				s[i] = fmt.Sprint(t.Index(i))
			}

			spew.Dump(s)
		}

	}))
}

func onCallBack(noOfArgs int, d Doer) {
	fmt.Println("I'm printting some nonsense BEFORE it happen")

	switch noOfArgs {
	case 1:
		d.Execute(1)
	case 2:
		d.Execute(1, 2)
	case 3:
		d.Execute("1", 2, "hello world")
	}

	fmt.Println("I'm printting some nonsense AFTER it happen")
}
