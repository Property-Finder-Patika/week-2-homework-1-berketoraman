package main

import "fmt"

func main() {
	i := 42
	f := 3.14
	b := true
	s := "Hello"
	r := 'A'

	fmt.Printf("i is %T\n", i)
	fmt.Printf("f is %T\n", f)
	fmt.Printf("b is %T\n", b)
	fmt.Printf("s is %T\n", s)

	// int32 = rune (type alias)
	fmt.Printf("r is %T\n", r)
	
	/* OUTPUT
i is int
f is float64
b is bool
s is string
r is int32
*/
}
