package main

import "fmt"

func main() {

	bs := make([]byte, 10)
	bl := 0
	for n := 0; n < 10; n++ {
		bl += copy(bs[bl:], "x")
		// fmt.Println(bl)
		// fmt.Println(bs)
	}

	str := []byte("foobar banana keks")
	fmt.Printf("%s\n", str[:])
	fmt.Printf("%s\n", str[0:])
	fmt.Printf("%s\n", str[:0])

}
