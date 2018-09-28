package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	{
		s := "Hello, ประเทศไทย"
		fmt.Println(len(s))
		fmt.Println(utf8.RuneCountInString(s))

		for i := 0; i < len(s); {
			r, size := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("%d\t%c\n", i, r)
			i += size
		}

		for i, r := range "Hello, ประเทศไทย" {
			fmt.Printf("%d\t%q\t%d\n", i, r, r)
		}
	}

	{
		s := "สวัสดี วันเสาร์"
		fmt.Printf("% x\n", s)
		r := []rune(s)
		fmt.Printf("%x\n", r)
		fmt.Println(string(r))

		fmt.Println(string(65))
		fmt.Println(string(0x4eac))

		fmt.Println(string(1234567))
	}

	{
		x := 123
		y := fmt.Sprintf("%d", x)
		fmt.Println(y, strconv.Itoa(x))

		fmt.Println(strconv.FormatInt(int64(x), 2))
		fmt.Printf("x=%b", x)
	}

	{
		var err error
		x, _ := strconv.Atoi("123")
		fmt.Println(x)
		y, _ := strconv.ParseInt("123", 10, 64)
		fmt.Println(y)

		x, err = strconv.Atoi("10.0")
		if err != nil {
			log.Printf("convert: %s\n", err)
		} else {
			fmt.Println(x)
		}

		y, err = strconv.ParseInt("10.0", 10, 64)
		if err != nil {
			log.Printf("convert: %s\n", err)
		} else {
			fmt.Println(y)
		}
	}
}
