package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1:=0
	f2:=1
	fmt.Println(f1)
	return func() int {
		f3 := f1+f2
		f1=f2
		f2=f3
		return f3
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


























=========================================================================

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	f1:=0
	f2:=1
	fmt.Println(f1)
	return func(f3 int)int {
		f3 = f1+f2
		f1=f2
		f2=f3
		return f3
	
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}