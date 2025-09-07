package main

import "fmt"

func main() {
	fmt.Println("hello World")

	/*
		Slices : Slices in golang is  a data structure which is built on top of array . Array are not flexible enough means when we predefined
		the size of an array and later if we need to insert more data to it , Array will not allow us to do it . Hence it will create a mess
		for development

		var a [4]int => Here it means that it as an array having integer value as sequence and size of , where accessing an elemnt is a[n] nth elemet of array

		Array Literals =>
		var a = [4]int{1,2,3,4} This is array literal
		 **** => An array variable denotes the entire array; it is not a pointer to the first array element (as would be the case in C)
		var a = [...]int{1,2,3,4} in This case both the time the compiler will automatically assume the size of the array.

		< Here Comes the Slices below >

		A slice literal is declared just like an array literal, except you leave out the element count:

		Representing a slice => names = []string{ "ronit", "roy", "rama", "roy", "misti", "roy"}

		Here we have not specified the size of the slices , Slices internally handles it , Slices has two internal properties , one is size and another is capacity

		size means the number of element it has currently and the capacaity means the number maximum it can hold .

		Like here in the `names` slice the size of the slice is 6 but if we look into its capacity its is 8 , capacity increases in 2^n where n is the nearest to reaching 2's power

		Accessing the size and capacity of Slice => len(names) => 6 and cap(names) => 8

		Another way of Representing Slices => name :=make([]string, x, y) [where x is size and y is capacity]

		Slicing a slice :> names[a:b] => it returns the ath to b-1th element , if we omit b then we will get full slice

		<****> The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.

		<Here Comes The Slices internals >
	*/

	names := []string{"ronit", "roy", "rama", "roy", "misti", "roy"}

	fmt.Println(names) // prints :[ronit roy rama roy misti roy]

	newNames := make([]string, 6)

	copy(newNames, names) // copy(destination slice, source slice) , it creates the copy of the former slice

	fmt.Println(newNames) // [ronit roy rama roy misti roy]

	newNames[5] = "roy"

	fmt.Println(names, newNames) // copy cretaes the new reference of the slice to a new array but belo we will see how slicing does the opposite

	sliceNames := names[2:]

	sliceNames[3] = "nath"

	fmt.Println(names, newNames, sliceNames)

	/*
		Now if we go under the hood of the slices we can now make a visula represemtation of how a slice is working


		[ *ptr element] ----------
								 | points to the array first element
								\ /
		[ length ]             [ronit][roy][rama][roy][misti][roy]

		[ capacity ]
	*/
	/* append() => it helps slices to grow */
	appendSlices := make([]int, 7)
	appendSlices = append(appendSlices, 1, 2, 35, 69, 0)
	fmt.Println(appendSlices) // [0 0 0 0 0 0 0 1 2 35 69 0]

	x := make([]int, 0, 10)
	x = append(x, 10)
	fmt.Println(x) //[10]
}
