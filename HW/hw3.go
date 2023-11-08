package main

import "fmt"

func main() {
	name := "Jinn"
	age := 26
	str1 := "Happy"
	str2 := " Birthday"
	res := fmt.Sprintf("%s %d %s %s", str1, age, str2, name)

	fmt.Println(res)
}
