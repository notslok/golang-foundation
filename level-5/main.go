// structs

package main

import "fmt"

// struct helps in composing together different data types
// structs are aka composite/aggregate/complex data type
type person struct {
	// identifier list:
	firstName string
	lastName  string
	age       int
}
type class struct { // embedded struct
	person    // anonymous/embedded field -> acc. to docs
	className string
	isActive  bool
}

func main() {
	/* prerequisites ----------------------------------------------------------
	// in go we dont use words like object and class... we call it creating a "value of type person"
	p1 := person{
		firstName: "Shlok",
		lastName:  "Jha",
		age:       23,
	}
	p2 := person{
		firstName: "sne",
		lastName:  "trola",
		age:       23,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)

	// 	Embedded structs------------------------------------------
	c1 := class{
		person: person{
			firstName: p1.firstName,
			lastName:  p1.lastName,
			age:       p1.age,
		},
		className: "A div",
		isActive:  true,
	}
	fmt.Println(c1)
	fmt.Println(c1.person.firstName, c1.className)

	c2 := class{
		person:    p2,
		className: "I div",
		isActive:  true,
	}
	fmt.Println(c2)
	fmt.Println(c2.person.firstName, c2.person)

	{ // start of new scope
		var x, y int
		x, y = 22, 33
		fmt.Println(x, y)
		x, y = y, x // swaps val
		fmt.Println(x, y)
	} // end of this scope

	// anonymous structs ... has on spot definition and declaration--------------------------

	p3 := struct { // struct definition
		first  string
		last   string
		shipNo int
	}{ // struct identifier value assignment
		first:  "capitane",
		last:   "sppok",
		shipNo: 377,
	}
	fmt.Println(p3)
	*/

	//============================================================================

	// exercise 1 ---

	type person struct {
		firstName   string
		lastName    string
		favIceCream []string
	}

	p1 := person{
		firstName:   "shlok",
		lastName:    "jha",
		favIceCream: []string{"hazlenut", "caramel", "chocolate"},
	}
	p2 := person{
		firstName:   "sne",
		lastName:    "trola",
		favIceCream: []string{"raspberry", "Trixy cup", "chocolate"},
	}

	fmt.Println("Shlok's fav:\n")
	for _, v := range p1.favIceCream {
		fmt.Println(v)
	}
	fmt.Printf("\n")
	fmt.Println("Sne's fav:\n")
	for _, v := range p2.favIceCream {
		fmt.Println(v)
	}

	// exercise 2 ---
	personMap := map[string]person{ // <string, type(person)> mapping
		p1.lastName: p1,
		p2.lastName: p2,
	}
	// initialising map as follows will give error ... "panic: assignment to entry in nil map"
	//personMap[p1.lastName] = p1
	//personMap[p2.lastName] = p2

	for i, v := range personMap {
		fmt.Println(i, v.firstName)
		fmt.Println(i, v.lastName)
		fmt.Println(i, v.favIceCream)
		fmt.Printf("\n")
	}

	// exercise 3 ---
	type vehicle struct {
		doors int
		color string
	}
	type truck struct {
		vehicle     // anonymous/embedded field
		fourWheeler bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{ // defining identifier "t1" of type truck using composite literal
		vehicle{
			doors: 2,
			color: "red",
		},
		true,
	}
	s1 := sedan{
		vehicle{
			doors: 4,
			color: "blue",
		},
		false,
	}
	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println(t1.doors)
	fmt.Println(s1.luxury)

	// Anonymous structs ------------------------------------------------------
	anonymousStruct := struct {
		carName string
		regNum  int
		mileage float32
	}{
		carName: "skoda fabia",
		regNum:  9797,
		mileage: 32.4,
	}
	fmt.Println(anonymousStruct)
	fmt.Println(anonymousStruct.carName, anonymousStruct.mileage)
}
