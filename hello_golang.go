package hello_golang

import "fmt"

type HelloGuys struct {
	Nama    string
	Age     int
	Address string
}

func SayHello() {
	result := HelloGuys{"sakti", 20, "Medan Indonesia"}
	fmt.Println(result)
}
