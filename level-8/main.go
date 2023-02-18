// application
package main

import (
	"encoding/json"
	"fmt"
	//"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"sort"
)

// for marshal
type person struct {
	// making the first letter of each field name capital, to expose them to outside functions like Marshal()
	// "json: ..." tags in-front of each field name tells the "which field to consider" while unmarshalling and "with what key name"
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

// ---
type Person struct {
	// exposed field name so that it can be accessed by sort function
	First string
	Age   int
}

type ByAge []Person

// any type which has following 3 methods(Len(), Swap(), Less()) can be passed to sort() function
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age } // for sorting w.r.t name we would use ... { return a[i].First < a[j].First }

// =====================================================================
func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	// Marshaling ------------------------------------------------------------
	bs, err := json.Marshal(people) // returns  <json output,err>
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))         // converts byte slices to their string equivalents and then prints it
	fmt.Println(bs)                 // prints byte slices as it is
	fmt.Printf("%+v\n", string(bs)) // %+v --> formatting for json data type

	// Unmarshal -------------------------------------------------------------
	agentList := []person{} // data structure to store json fields
	err = json.Unmarshal(bs, &agentList)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("All the data got from unmarshalling: %v\n", agentList)

	for i, v := range agentList {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	// Writer interface ...!!! refer notes and docs(must) -------------------
	// refer lec-140(writer interface again)
	fmt.Println("Hello, Playground")
	fmt.Fprintln(os.Stdout, "Hello, Playground") // equivalent to fmt.Println

	io.WriteString(os.Stdout, "Hello, Playground\n") // takes in a writer and empty interface{}

	// Sort-------------------------------------------------------------------
	s := []int{5, 1, 5, 7, 3, 0}
	sort.Ints(s) // doesnt returns anything...deals with the address of the string directly
	fmt.Println(s)

	str := []string{"B for boy,", "A for apple,", "C for car,"}
	sort.Strings(str)
	fmt.Println(str)

	// custom sort  ------------------------------------------------------------
	pp1 := Person{"James", 32}
	pp2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	peoples := []Person{pp1, pp2, p3, p4}

	fmt.Println(peoples)
	// here we do sorting by age by default... to sort by name modify the methods Less()... which defines the comparison logic while sorting
	sort.Sort(ByAge(peoples)) // type conversion to ByAge() so, methods: Len(), Swap(), Less() belongs to it
	fmt.Println(peoples)

	// bcrypt (preview) ---------------------------------------------------------
	s2 := `password123`
	//bs2, err := bcrypt.GenerateFromPassword([]byte(s2), bcrypt.MinCost) // encrypting
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)
	//fmt.Println(bs2) // encrypted output

	//loginPword1 := `password1234`

	//err = bcrypt.CompareHashAndPassword(bs2, []byte(loginPword1)) // decrypting
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}
