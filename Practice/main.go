package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func intro() {
	// var a string = "world!"
	// var nameOne = "Alvan"
	// var nameTwo string
	//b := "new"
	/*
		multiline comment

	*/
	// fmt.Println(a, nameOne, nameTwo)

	age := 30
	name := "Alvan"

	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name)

	fmt.Printf("my age is of type of %T, and my name is of type %T \n", age, name)

	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)

	fmt.Println("extracted strings are:", str)
}

func arrayAndSlices() {
	var ages = [3]int{1, 2, 3}
	ages[1] = 6
	ages[2] = 8
	ages[0] = 7

	fmt.Println(ages, len(ages))

	// slices

	names := []string{"Alvan", "John", "Mike"}
	names = append(names, "luigi")
	var cut1 = names[:3]
	var cut2 = names[:]
	fmt.Println(names)
	fmt.Println(cut1)
	fmt.Println(cut2)
}

func goStl() {
	// name := "hello from here"
	// fmt.Println(name)
	// var name_1 = strings.Contains(name, "hello!")
	// var name_2 = strings.ReplaceAll(name, "hello", "hi")
	// fmt.Println(name_2)
	// fmt.Println(strings.ToUpper(name))
	// fmt.Println(strings.ToLower(name))

	ages := []int{12, 45, 32, 76, 23, 63}
	fmt.Println(ages)
	sort.Ints(ages)
	fmt.Println(ages)
	fmt.Println(sort.SearchInts(ages, 122))

}

func loops() {
	// x := 0

	// for x < 5 {
	// 	fmt.Println("The value of x is ", x)
	// 	x++
	// }

	names := []string{"alvan", "mike", "john", "phil"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Printf("The name at position %v is %v \n", i, strings.ToUpper(names[i]))
	// }

	for _, value := range names[:len(names)-2] {
		fmt.Printf("The name is %v \n", strings.ToUpper(value))
	}

}

func conditionals() {
	age := 21

	fmt.Println(age < 20)
	fmt.Println(age > 30)
	fmt.Println(age <= 20)
	fmt.Println(age >= 20)
	fmt.Println(age == 20)
	fmt.Println(age != 20)

	if age < 20 || age > 20 || age == 20 {
		fmt.Println("here")
	} else {
		fmt.Println("not here")
	}
}

func sayHello(n string) {
	fmt.Printf("hello from %v \n", n)
}

func combine(n []string, f func(n string)) {
	for _, v := range n {
		f(v)
	}
}

func calCircleArea(n float64) float64 {
	return math.Pi * n * n
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	value := strings.Split(s, " ")

	var initials []string
	for _, v := range value {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	} else {
		return initials[0], "_"
	}
}

func maps() {
	menu := map[string]float64{
		"egg":   2.222,
		"sugar": 23.00,
		"milk":  5000,
	}
	fmt.Println(menu)
	fmt.Println(menu["egg"])

	phoneNumber := map[int]string{
		233:  "alvan",
		2322: "mike",
	}
	fmt.Println(phoneNumber)

	for key, value := range menu {
		fmt.Println(key, "_", value)
	}
	for key, value := range phoneNumber {
		fmt.Println(key, "_", value)
	}
}
func main() {
	// intro()
	// arrayAndSlices()
	// conditionals()
	sayHello("alvan")
	combine([]string{"Alvan", "Josh", "will"}, sayHello)
	n := calCircleArea(10.22)
	fmt.Printf("the area of the circle is %0.3f \n", n)

	f1, f2 := getInitials("Alvan Young")
	s1, s2 := getInitials("Alvan")

	fmt.Printf("The first value is %v, and the second value is %v \n", f1, f2)
	fmt.Printf("The first value is %v, and the second value is %v \n", s1, s2)

	fmt.Println("====================================================")

	maps()
}
