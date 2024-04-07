package main

import (
	"fmt"
	"strconv"
)

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

	// 型変換
	var i5 int = 1
	fl64 := float64(i5)
	fmt.Println(fl64)
	fmt.Printf("i5 = %T\n", i5)
	fmt.Printf("fl64 = %T\n", fl64)

	i6 := int(fl64)
	fmt.Printf("i6 = %T\n", i6)
	fl := 10.5

	i7 := int(fl)
	fmt.Printf("i7 = %T\n", i7)
	fmt.Println(i7)


	// 文字列から数値
	var s1 string = "100"
	fmt.Printf("s1 = %T\n", s1)
	i8, _ := strconv.Atoi(s1)

	fmt.Println(i8)
	fmt.Printf("i8 = %T\n", i8)

	// 数値から文字列
	var i9 int = 200
	s5 := strconv.Itoa(i9)
	fmt.Println(s5)
	fmt.Printf("s5 = %T\n", s5)

	// 文字列からバイト配列
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	// バイト配列から文字列
	h2 := string(b)
	fmt.Println(h2)
}