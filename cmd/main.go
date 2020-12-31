/*
main.go
-John Taylor
Dec-31-2020

Insert an ellipses into the middle of a long string to shorten it

*/
package main

import (
	"fmt"
	"github.com/jftuga/ellipsis"
)

func main() {
	s := "abcdefghijklmnopqurstvwxyz"
	for i := 0; i <= 26; i++ {
		s := ellipsis.Shorten(s, i)
		fmt.Printf("%2d %s\n", i, s)
	}
}
