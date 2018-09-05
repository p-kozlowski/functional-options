package main

import (
	"fmt"
)

func vararg(args ...string) {
	fmt.Printf("%v\n", args)

	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	vararg()
	vararg("hello", "world")
}
