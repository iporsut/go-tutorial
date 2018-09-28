package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
	THB
)

var symbol = [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥", THB: "฿"}

func main() {
	fmt.Println(THB, symbol[THB])
}
