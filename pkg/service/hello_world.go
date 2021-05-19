package service

import "fmt"

// 関数をファイル外から使用できるかどうかを決めることができる！
// ファイル外から使用するためには、関数名を大文字から始める必要がある！
// 使用しないときは、関数名を小文字から始めれば良い。

func HelloWorld() {
	// Println でも出力できる
	fmt.Println("Hello World! My package is Service.")

	var demo string = "Hello World"
	fmt.Println(demo)

	s := "Hello World"
	fmt.Println(s)

	i := 1
	var bigInt uint64 = 1
	fmt.Println(i, bigInt)

	f := 1.0
	fmt.Println(f)

	b := true
	fmt.Println(b)
}

func VoidFunction1() {
	fmt.Println("aaa")
}

func voidFunction2() {
	fmt.Println("aaa")
}
