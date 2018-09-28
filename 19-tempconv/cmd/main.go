package main

import (
	"fmt"
	"go-tutorial/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("%g\n", boilingF-tempconv.FreezingC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	fmt.Println(c == f)
	fmt.Println(c == tempconv.Celsius(f))

	c = tempconv.FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
