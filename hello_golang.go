package hello_golang

import "fmt"

type HelloGuys struct {
	Nama string
	Age  int
}

func SayHello() {
	result := HelloGuys{"sakti", 20}
	fmt.Println(result)
}
