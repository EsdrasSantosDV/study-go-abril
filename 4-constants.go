package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	//PRA DECLARAR UMA FCONSTANTE NO GO, COLOCAMOS CONST
	const esdras = "Esdras"
	fmt.Println(s)

	const n = 500000000

	const d = 3e19 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
