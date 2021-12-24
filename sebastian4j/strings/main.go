package main

import (
	"bytes"
	"fmt"
	"math"
	"time"
	"unicode/utf8"
)

func main() {
	s := "Hello, 橙色"
	fmt.Println(s)
	// una rune no ocupa un byte
	fmt.Println("bytes: ", len(s))
	fmt.Println("runes: ", utf8.RuneCountInString(s))

	fmt.Println("s[1]: ", s[1])
	fmt.Println("s[1:]: ", s[1:])
	fmt.Println("s[2:4]: ", s[2:4])
	fmt.Println()
	var indx int
	for j := 0; j < utf8.RuneCountInString(s); j++ {
		indx += showString(s, indx)
	}

	fmt.Println()
	for i, r := range s {
		fmt.Printf("%d\t%c\t%d\n", i, r, r)
	}
	fmt.Println()
	fmt.Printf("% x\n", s)

	fmt.Println()
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
	fmt.Println(string(65))     // una rune, no el 65
	fmt.Println(string(0x4eac)) // una rune
	fmt.Println(string(123456)) // rune no valida

	fmt.Println()
	var buf bytes.Buffer // uso eficiente de byte slices
	buf.WriteByte('[')   // mejor para ascii
	buf.WriteRune('a')   // mejor para arbitrary rune
	buf.WriteString(" hola")
	fmt.Println(buf.String())
	fmt.Printf("%T %[1]v \n", time.Minute, time.Minute)

	const aaa = 3.14 // untype float, solo funciona con constantes
	var bbb = 3.14
	fmt.Printf("%T %v\n", aaa, aaa)
	fmt.Printf("%T %v\n", bbb, bbb)
	fmt.Printf("%T %v\n", math.Pi, math.Pi) // tambien es untyped

	var xx float32 = math.Pi
	var yy float64 = math.Pi
	fmt.Printf("%T %v\n", xx, yy)

	const untypedInt = 0xdeadbeef
	fmt.Printf("%T %v\n", untypedInt, untypedInt)
}

func showString(s string, indx int) int {
	r, size := utf8.DecodeLastRuneInString(s[0 : len(s)-indx])
	fmt.Printf("index: %d\t rune: %c\tbytes: %d\n", indx, r, size)
	return size
}
