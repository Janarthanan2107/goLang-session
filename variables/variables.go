// package name should be declare at first as main
package main

// packages need to import at second
import (
	"fmt"
)

// main function also a mandatory thing to use
func main() {

	// NOTE: if any variable is declared we have to use that variable once in our code otherwise we wont declare any needs ,if you forget to use a declared variable in a code it throw's error.

	// * using var we can re-assign
	var a int
	a = 10 //output:10

	// * we can directly assign values
	var b int = 20 //output:10,20

	// * we can declare values with out a data type also
	var c = 30 //output:10,20,30
	// c = 40 //if we re-declare c output is 40

	// shorthand operator
	// using ( := ) we a declare and assign a value
	// NOTE: if any variable is already declared with same variable using ( := ), error is ==> no new variables on left side of :=

	d := 40 //output:10,20,30,40

	// we can assign multiple variable at the same time
	e, f := 50, 60 //output:10,20,30,40,50,60

	fmt.Println(a, b, c, d, e, f)
}
