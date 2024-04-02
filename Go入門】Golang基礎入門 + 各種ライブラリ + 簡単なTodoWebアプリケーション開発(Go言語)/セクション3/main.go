package main

import "fmt"

// 基本的には型を指定する。 バグの混入を防ぐ。

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {

	// 明示的な定義
	//var 変数名型=値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 暗黙的な定義
	//变数名 := 值
	i4 := 400
	fmt.Println(i4)

	// i4 = 450
	// fmt.Println(i4)
	outer()
}