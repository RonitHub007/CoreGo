package main

import "fmt"

func main() {
	fmt.Println("Hello Strings")

	s := "Hello 😊"
	s2 := s[4:7]
	s3 := s[:5]
	s4 := s[6:]

	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	/*
			Hello Strings
		o �
		Hello
		😊
	*/

	fmt.Println(len(s)) // 10

	s1 := "Hello 0"
	fmt.Println(len(s1)) //7

	var sp string = "Hello, 😊"
	var bs []byte = []byte(sp)
	var rs []rune = []rune(sp)

	fmt.Println(bs)
	fmt.Println(rs)

	/*
			7
		bs => [72 101 108 108 111 44 32 240 159 152 138]
		rs => [72 101 108 108 111 44 32 128522]
	*/
}
