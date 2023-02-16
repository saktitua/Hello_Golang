package hello_golang

import "fmt"

type HelloGuys struct {
	Nama         string
	Age          int
	Address      string
	StatusMerrid string
}

func SayHello() {
	result := HelloGuys{"sakti", 20, "Medan Indonesia", "Merrid"}
	fmt.Println(result)
}
