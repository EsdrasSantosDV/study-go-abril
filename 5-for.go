package main

import "fmt"

func main() {

	//PRA PERCORRER
	//UMA LISTA OU ARRAY, USAMOS O FOR

	//pra fazer estilo um while
	//var i = 1
	//for i <= 100000000 {
	//	println(i)
	//	i = i + 1
	//}

	//pra fazer de forma tipo c
	//for j := 0; j <= 100; j++ {
	//	println(j)
	//}
	//
	//SE NÃO QUISER USAR O I++
	//ESTILO PYTHON
	//for i := range 100000 {
	//	fmt.Println("range", i)
	//}

	//tem o break
	for {
		fmt.Println("loop")
		break
	}
	//TEM CONTINUE
	for n := range 60 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
