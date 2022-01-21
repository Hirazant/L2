package main

import (
	"fmt"
)

func Foo() *error {
	var err *error = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)

}
