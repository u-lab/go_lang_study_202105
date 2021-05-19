package main

import (
	"fmt"
	"practice-golang-202105/pkg/service"
)

func main() {
	// CLang風な Hello World
	fmt.Printf("%s\n", "Hello World!")

	service.Play()

	// service.HelloWorld()

	// a := 1
	// b := 1

	// if a*b%5 == 0 {
	// 	fmt.Println("a*bは5の倍数")
	// } else {
	// 	fmt.Println("a*bは5の倍数ではない")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// t := 0
	// for {
	// 	fmt.Println(t)
	// 	t++

	// 	if t == 5 {
	// 		break
	// 	}
	// }
}
