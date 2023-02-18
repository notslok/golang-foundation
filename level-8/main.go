// application
package main

import (
	"encoding/json"
	"fmt"
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

// ---------------------------- exercise 1
type user struct {
	First string
	Age   int
}

// ---------------------------- exercise 2
type peeps struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"age"`
	Sayings []string `json:"Sayings"`
}

// ------------------------------ exercise 3
type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// ---------------------------------- exercise 5
type user3 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

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

	// exercise - 1 ----------------------------------------------
	fmt.Printf("\n\n\n exercise 1 \n\n\n")
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// Marshal
	bs, err = json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", bs)         // raw json output in  []byte/byte slice type
	fmt.Printf("%+v\n", string(bs)) // json output in string type

	// exercise - 2 ---------------- Unmarshal------------------
	fmt.Printf("\n\n\n exercise 2 \n\n\n")
	j := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(j)
	dump := []peeps{}

	err2 := json.Unmarshal([]byte(j), &dump)
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Printf("\n here is the unmarshalled dump...\n%+v\n", dump)

	// exercise 3 --------------------------------------------------------------------------
	fmt.Printf("\n\n\n exercise 3 \n\n\n")
	u12 := user2{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u22 := user2{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u32 := user2{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users2 := []user2{u12, u22, u32}

	fmt.Println(users2)

	// TASK: encode to JSON the []user sending the results to Stdout
	json.NewEncoder(os.Stdout).Encode(users2) // encodes to json and dumps to stdout

	// exercise 4 -------------------------------------------------------------------------
	fmt.Printf("\n\n\n exercise 4 \n\n\n")
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)

	// exercise 5 ---------------------------------------------------------------------------
	fmt.Printf("\n\n\n exercise 5 \n\n\n")

	u13 := user3{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u23 := user3{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u33 := user3{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users3 := []user3{u13, u23, u33}

	fmt.Println(users3)
	/*
		 TASK:
			sort the []user by
			● age
			● last
			● print everything in a way that is pleasant
	*/

	// ------------- refer notes....
}
