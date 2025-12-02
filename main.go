package main

import "fmt"

func main() {
	//pointer and ref
	var a int = 1080
	var ptr *int = &a

	fmt.Println("Value variabel a : ", a)
	fmt.Println("Address variabel a : ", &a)
	fmt.Println("Value pointer : ", *ptr)
	fmt.Println("Address pointer : ", ptr)

	//ref
	fmt.Println("=====================")
	var b string = "Antigravity"
	var ref = &b
	fmt.Println("Value variabel b : ", b)
	fmt.Println("Address variabel b : ", &b)
	fmt.Println("Value ref : ", *ref)
	fmt.Println("Address ref : ", ref)

	//pointer aritmatika
	fmt.Println("=====================")
	var c int = 10
	var d *int = &c
	fmt.Println("Value variabel c : ", c)
	fmt.Println("Address variabel c : ", &c)
	fmt.Println("Value pointer : ", *ptr)
	fmt.Println("Address pointer : ", ptr)
}