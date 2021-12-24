package main

import "fmt"

// https://www.electrosoftcloud.com/operaciones-bit-a-bit-bitwise/

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) //  {1, 5}
	fmt.Printf("%08b\n", y) //  {1, 2}

	fmt.Printf("%08b\n", x&y)  // la intersección {1}
	fmt.Printf("%08b\n", x|y)  // la unión {1,2,5}
	fmt.Printf("%08b\n", x^y)  // la diferencia simetrica {2,5}
	fmt.Printf("%08b\n", x&^y) // la diferencia {5}
	fmt.Printf("%08b\n", x<<1) // set {2, 6}
	fmt.Printf("%08b\n", x>>1) // set {0, 4}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // 1, 5
		}
	}

	fmt.Println("--------------------")
	/*
		0 & 0 == 0
		0 & 1 == 0
		1 & 0 == 0
		1 & 1 == 1

		0 | 0 == 0
		0 | 1 == 1
		1 | 0 == 1
		1 | 1 == 1

		0 ^ 0 == 0
		0 ^ 1 == 1
		1 ^ 0 == 1
		1 ^ 1 == 0

		not = invierte los bits, bit a bit (unario)
	*/
	fmt.Println("&")
	fmt.Println("0000000000101111")
	fmt.Println("0000000001001011")
	fmt.Println("0000000000001011")

	fmt.Println("|")
	fmt.Println("0000000000101111")
	fmt.Println("0000000001001011")
	fmt.Println("0000000001101111")

	fmt.Println("^")
	fmt.Println("0000000000101111")
	fmt.Println("0000000001001011")
	fmt.Println("0000000001101111")

	fmt.Println("not")
	fmt.Println("0000000001100111")
	fmt.Println("1111111110011000")
	/*
		el primer bit es el signo (no unsigned):
		1 = negativo
		0 = positivo

		bit shift

		int a = 5;        // binario: 0000000000000101
		int b = a << 3;   // binario: 0000000000101000, o 40 en decimal
		int c = b >> 3;   // binario: 0000000000000101, vuelve a ser igual que a (5 en decimal)

		int x = -16;     // binario: 1111111111110000
		int y = x >> 3;  // binario: 1111111111111110

		int x = -16;               // binario: 1111111111110000
		int y = unsigned(x) >> 3;  // binario: 0001111111111110

	*/
}
