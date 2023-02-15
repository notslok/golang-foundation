package main

// untyped const...type determined during compile time
const (
	a = 42
	b = "string is here"
)

/*
// typed const...
const(
	a int = 42
	b string = "string is here"
)
*/
func main() {
	// level - 2 prereq ---------------------
	/*
		var x rune = 9  // rune = int32
		var y byte = 60 // byte = uint8
		fmt.Println(x)
		fmt.Println(y)

		fmt.Println(runtime.GOOS)   // gets the os of the machine
		fmt.Println(runtime.GOARCH) // gets the architecture of the machine

		str := "Hello playground"
		fmt.Println(str)
		fmt.Printf("%T\n", str)

		byteStr := []byte(str) // converting str into slice of byte/int8..will give us an array of acii vals corresponding to each char in the string
		fmt.Println(byteStr)
		fmt.Printf("%T\n", byteStr)

		for i := 0; i < len(str); i++ {
			fmt.Printf("%#U\n", str[i]) // printing utf-8 values of each char in str
			fmt.Printf("%#x\n", str[i]) // printing hex values of each char in str
		}

		// const
		fmt.Println(a)
		fmt.Println(b)
		fmt.Printf("%T\n", a) // int
		fmt.Printf("%T\n", b) // string

		// iota ... a predefined identifier
		const (
			a2 = iota
			b2 = iota
			c2
			d2
		)
		fmt.Println(a2)              // 0
		fmt.Println(b2)              // 1
		fmt.Println(c2)              // 2
		fmt.Println(d2)              // 3...self increments
		fmt.Printf("%T\t%T", a2, b2) // 0 1 ... self increments

		// bit shifts
		// naive way...
		var num uint32 = 1
		for itr := 0; itr < 10; itr++ {
			fmt.Printf("%b\n", num)
			num = num << 1 // num = num*2...left bit shift
		}
		//using iota...
		const (
			_  = iota // dropping first iota value i.e. 0
			kb = 1 << (iota * 10)
			mb = 1 << (iota * 10)
			gb = 1 << (iota * 10)
		)

		fmt.Printf("%d \t\t %b\n", kb, kb)
		fmt.Printf("%d \t\t %b\n", mb, mb)
		fmt.Printf("%d \t\t %b\n", gb, gb)
	*/
	//==========================================================

	// exercise 1 --- hex, binary and decimal format printing
	/*
		num := 32
		fmt.Printf("%d \t %d \t %#x \n", num, num, num)
	*/

	// exercise 2 --- conditional operators
	/*
		a := 2 == 2
		b := 2 <= 4
		c := 2 >= 4
		d := 1 != 0
		e := 1 < 1
		f := 1 > 1

		fmt.Printf("%t \t %t \t %t \t %t \t %t \t %t \n", a, b, c, d, e, f)

	*/

	// exercise 3 --- typed and untyped const...already done above
	// exercise 4 --- bit shifts and fmt
	/*
		var num uint32 = 246
		fmt.Printf("%d \t %b \t %#x \n", num, num, num)
		num <<= 1 // shorthand for num = num<<1
		fmt.Printf("%d \t %b \t %#x \n", num, num, num)
	*/
	// exercise 5 --- raw string literal
	/*
			a := `here is
				  string literal
		           in action!!!`
			fmt.Printf("%s\n", a)

	*/

	// exercise 6 --- iota.. done above
}

