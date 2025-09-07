package main

import "fmt"

func main() {

	studentNameRollMap := make(map[string]int)

	studentNameRollMap["Ronit"] = 1
	studentNameRollMap["George"] = 2
	studentNameRollMap["Luca"] = 3

	// Now find the Roll Number of Ronit

	fmt.Println(studentNameRollMap["Ronit"])
	studentNameRollMap["Luca"] = studentNameRollMap["Luca"] + 1
	fmt.Println(studentNameRollMap["Luca"])

	//The comma ok idiom in Golang

	commaOkMap := map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Println(commaOkMap)

	v, ok := commaOkMap["hello"]
	fmt.Println(v, ok)

	v1, ok1 := commaOkMap["world"]
	fmt.Println(v1, ok1)

	v2, ok2 := commaOkMap["goodBye"]
	fmt.Println(v2, ok2)
	/*
						Rather than assign the result of a map read to a single variable, with the comma ok
					idiom you assign the results of a map read to two variables. The first gets the value
					associated with the key. The second value returned is a bool. It is usually named ok. If
					ok is true, the key is present in the map. If ok is false, the key is not present. In this
					example, the code prints out 5 true, 0 true, and 0 false

		    5 true
			0 true
			0 false
	*/

	// Map as a Set in Golang

	Set := map[int]bool{}

	Slice := []int{1, 2, 3, 4, 5, 4, 5, 6, 2, 3, 7}

	for _, value := range Slice {
		Set[value] = true
	}
	fmt.Println(len(Slice), Slice)
	fmt.Println(len(Set), Set)

}
