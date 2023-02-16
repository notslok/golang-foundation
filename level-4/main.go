package main

import "fmt"

func main() {

	/*---------------- Pre-requisites ------------------------------------

	// golang documentation recommends using slices over array
	var x [5]int  // in golang length of an array is part of its type
	var y [10]int // has different type from x
	x[3] = 9
	fmt.Println(x)      // o/p -> [0 0 0 9 0]
	fmt.Println(len(x)) // 5

	fmt.Printf("%T \t %T \n", x, y) // [5]int   [10]int

	// slices - composite literal------------------------------------

	//x := type{values} // composite literal
	x2 := []int{1, 2, 4, 5, 6} // a SLICE allows you to group together the values of same type
	fmt.Println(x2)
	fmt.Println(len(x2))
	fmt.Println(cap(x2)) // prints the capacity
	fmt.Println(x[3])
	// looping over a slice...using "range" keyword
	for i, v := range x2 { // i->index, v->value
		fmt.Println(i, v)
	}
	fmt.Println("Alternatively...\n")
	// similar to...
	for i := 0; i < len(x2); i++ {
		fmt.Println(i, x2[i])
	}

	// slicing a slice (:)
	fmt.Println(x2[1:])
	fmt.Println(x2[2:3]) // upto but not including 3rd value(1 base indexing)

	// appending to the slice
	x2 = append(x2, 23, 43, 67)
	fmt.Println(x2)

	y2 := []int{123, 345}
	x2 = append(x2, y2...) // appending slice to a slice...using the fact that append() takes in variadic parameters
	fmt.Println(x2)

	// deleting from slice

	x2 = append(x2[:2], x2[4:]...) // deleting 3rd and 4th element in slice x2
	fmt.Println(x2)

	x3 := make([]int, 10, 10) // use make when you know about the capacity limit of a slice...here we have created a slice of length 10 with capacity 100
	fmt.Println(x3)
	fmt.Println(len(x3))
	fmt.Println(cap(x3))
	//x3[10] = 33 // index out of range err
	x3[9] = 33
	fmt.Println(x3)
	x3 = append(x3, 10)
	fmt.Println(cap(x3)) //---> capacity doubles at this point as soon as we try append more values than the capacity defined in make expression above
	x3 = append(x3, 10)
	x3 = append(x3, 10)
	fmt.Println(cap(x3)) // ---> capacity remains same here i.e. 20
	fmt.Println(x3)

	// multi-dimensional slice
	names := []string{"shlok", "sneha"}
	iceCream := []string{"raspberry", "hazelnut"}
	fmt.Println(names)
	fmt.Println(iceCream)

	spreadSheet := [][]string{names, iceCream} // 2 dimensional slice
	fmt.Println(spreadSheet)

	// Maps
	m := map[string]int{
		"james":      32,
		"moneypenny": 31,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["shlok"]) // not stored so it will return zero

	v, ok := m["shlok"] // IMP: comma ok idiom
	fmt.Println(v)
	fmt.Println(ok) // o/p: false ... tells if the value key exists in the map or not

	if v, ok := m["moneypenny"]; ok {
		fmt.Printf("The key {%v} exists in the map, having value ... %d\n", "moneypenny", v)
	}

	// adding elements to map and getting the range of map
	m["sneha"] = 23       //  added an entry ... ("sneha", 23)
	for k, v := range m { // index, value := range ...
		fmt.Println(k, v)
	}

	// using range on slices...
	for i, v := range x3 {
		fmt.Println(i, v)
	}

	// deleting entry from a map
	fmt.Println(m)
	fmt.Println(`After deleting "james" entry...`)
	delete(m, "james")
	fmt.Println(m)
	fmt.Println(`After deleting a non-existing entry...`)
	delete(m, "askjdbak") // won't throw error even if that key doesn't exist in the map
	fmt.Println(m)
	*/
	//===========================================================================
	// exercise 1 --- array using composite literals

	arr := [5]int{}
	for i := range arr {
		arr[i] = i
	}
	for i := range arr {
		fmt.Println(arr[i])
	}
	fmt.Printf("%T \t %v \n", arr, arr)

	// exercise 2 --- slice using composite literal
	slice := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := range slice {
		slice[i] = i
	}
	for i := range slice {
		fmt.Println(slice[i])
	}
	fmt.Printf("%T \t %v\n", slice, slice)

	// exercise 3 --- simple slicing ...done
	// exercise 4 ---
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52) // appending 52 to slice x
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	aux := []int{56, 57, 58, 59, 60}
	x = append(x, aux...)
	fmt.Println(x)

	// exercise 5 --- DELETE from a slice
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	indianStates := make([]string, 36, 36) //defining length and capacity of slice using make() builtin func
	// appends after 36th index of the slice
	indianStates = append(indianStates, "Andhra Pradesh", "Arunachal Pradesh", "Assam", "Bihar", "Chhattisgarh", "Goa", "Gujarat", "Haryana", "Himachal Pradesh", "Jammu and Kashmir", "Jharkhand", "Karnataka", "Kerala", "Madhya Pradesh", "Maharashtra", "Manipur", "Meghalaya", "Mizoram", "Nagaland", "Odisha", "Punjab", "Rajasthan", "Sikkim", "Tamil Nadu", "Telangana", "Tripura", "Uttarakhand", "Uttar Pradesh", "West Bengal", "Andaman and Nicobar Islands", "Chandigarh", "Dadra and Nagar Haveli", "Daman and Diu", "Delhi", "Lakshadweep", "Puducherry")
	fmt.Println(indianStates)
	fmt.Println(len(indianStates))
	fmt.Println(cap(indianStates))
	for i, v := range indianStates {
		fmt.Println(i, v) // empty v till 35th index
	}

	fmt.Println("\n")
	fmt.Println("correct way below...\n")

	// correct way below...
	// creating a temporary slice for states
	states := []string{"Andhra Pradesh", "Arunachal Pradesh", "Assam", "Bihar", "Chhattisgarh", "Goa", "Gujarat", "Haryana", "Himachal Pradesh", "Jammu and Kashmir", "Jharkhand", "Karnataka", "Kerala", "Madhya Pradesh", "Maharashtra", "Manipur", "Meghalaya", "Mizoram", "Nagaland", "Odisha", "Punjab", "Rajasthan", "Sikkim", "Tamil Nadu", "Telangana", "Tripura", "Uttarakhand", "Uttar Pradesh", "West Bengal", "Andaman and Nicobar Islands", "Chandigarh", "Dadra and Nagar Haveli", "Daman and Diu", "Delhi", "Lakshadweep", "Puducherry"}
	for i, v := range states {
		indianStates[i] = v
	}
	for i := range indianStates {
		fmt.Println(i, indianStates[i])
	}

	// exercise 7 --- multi-dimensional slice
	james := []string{"James", "Bond", "Shaken, not stirred"}
	moneyPenny := []string{"Miss", "Moneypenny", "Helloooo, James"}
	fmt.Println(james)
	fmt.Println(moneyPenny)

	jamesAndMoneyPenny := [][]string{james, moneyPenny}
	fmt.Println(jamesAndMoneyPenny)

	for i := range jamesAndMoneyPenny {
		for j := range jamesAndMoneyPenny[i] {
			fmt.Println(jamesAndMoneyPenny[i][j])
		}
		fmt.Println("\n")
	}

	// exercise 8,8,10 --- Maps and its basic functionalities
	mp := map[string][]string{
		`bond_james`: {`Shaken not stirred`, `Martinis`, `Women`},
		`no_dr`:      {`being evil`, `icecream`, `sunset`},
	}

	for i, _ := range mp {
		fmt.Printf("values for the key %v are:\n", i)
		for j, _ := range mp[i] {
			fmt.Println(i, j, mp[i][j])
		}
	}

	// adding val for key "shlok"
	mp["shlok"] = []string{"nothing at all"}
	fmt.Println(mp)
	for i, _ := range mp {
		for _, val := range mp[i] {
			fmt.Println(i, val)
		}
	}
	// deleting entry corresponding to key "shlok"
	delete(mp, "shlok")
	fmt.Println(mp)
	for i, _ := range mp {
		for _, val := range mp[i] {
			fmt.Println(i, val)
		}
	}
}
