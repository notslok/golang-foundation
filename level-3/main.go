package main

import "fmt"

func main() {
	/* ---------- pre-requisites

	// func init; condition; post{}
	// golang has no while/do-while loop

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// nested loop
	for i := 0; i < 5; i++ { // (init; condition; post)
		for j := i; j < 5; j++ {
			fmt.Println(i, j)
		}
	}
	// break and continue
	x := 1
	for x < 10 {
		if x >= 8 {
			break
		}
		if x%2 != 0 {
			x++
			continue
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Done.")

	// formatting and printing ascii and UTF-8 vals

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v \t %#U \n", i, i) // prints... Acii val - UTF-8 val pairs
	}

	// conditional statements and chaining (if - else if - else)
	if true {
		fmt.Println("print true\n")
	}
	if false {
		fmt.Println("print false\n")
	}
	if !false {
		fmt.Println("print not not false\n")
	}
	if !(2 != 2) {
		fmt.Println("2 and 2 are equal\n")
	}

	if x := 43; x == 2 { // falsy
		fmt.Println("inside in if\n")
	} else if x == 43 {
		fmt.Println("outside in else-if\n")
	} else {
		fmt.Println("in else\n")
	}

	// switch statements
	switch false {
	case false:
		fmt.Println("false case invoked\n") // no default fall-through in go
	case true:
		fmt.Println("true case invoked\n")
	}

	switch false {
	case false:
		fmt.Println("false case invoked\n")
		fallthrough // fallthrough explicitly invoked
	case true:
		fmt.Println("true case invoked\n")
	default: // can appear anywhere inside switch
		fmt.Println("default case fired\n")
	}
	// example from effective go
	c := byte('?')
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("true\n")
	default:
		fmt.Println("false\n")
	}
	// conditional operators
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false || false)
	fmt.Println(true || false)
	fmt.Println(!false)

	*/

	//=======================================================================

	// exercise - 1 --- print 1 to 10000

	/*
		for i := 1; i <= 10000; i++ {
			fmt.Printf("%v\t", i)
			if i%10 == 0 {
				fmt.Printf("\n")
			}
		}
	*/

	// exercise - 2 --- print every rune code of uppercase letter three times
	/*
		for i := 65; i <= 90; i++ {
			for j := 0; j < 3; j++ {
				fmt.Printf("%#U\n", i)
			}
			fmt.Println("")
		}

	*/

	// exercise - 3 --- use the syntax... for[condiition]{...} -> denoted using ebnf notations
	/*
		for i := 0; i != 10002; {
			fmt.Printf("%v\n", i)
			i++
		}
	*/

	// exercise - 4 --- use syntax ... for{}...used to impersonate while and do while loop in golang

	/*
		i := 0
		for {
			fmt.Printf("%d\n", i)
			i++
			if i == 20 {
				break
			}
		}
	*/

	// exercise - 5 --- print modulo of numbers between 10 and 1000 when div by 4
	/*
		for i := 10; i <= 10000; i++ {
			fmt.Printf("%v\n", i%4)
		}
	*/

	// exercise - 6 --- show if statement...done
	// exercise - 7 --- show if-else if-else ladder...done
	// exercise - 8 --- switch statements

	switch { // by default true
	case true:
		fmt.Printf("%s\n", "true case invoked")
	case false:
		fmt.Printf("%s\n", "false case invoked")
	}

	var c string = "hi"

	switch c {
	case "bye":
		fmt.Println("bye\n")
	case "hi":
		fmt.Println("hi\n")
	default:
		fmt.Println("default case\n")
	}
}
