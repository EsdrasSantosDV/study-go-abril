package main

import "fmt"

func main() {
	var vetor [5]int
	fmt.Println("vetor", vetor)

	vetor[3] = 100
	fmt.Println("vetor", vetor)
	for i := range vetor {
		println("vetor[", i, "] = ", vetor[i])
	}

	vetorporposicao := [5]int{1, 2, 3, 4}
	fmt.Println("dcl:", vetorporposicao)
	//matriz
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("matriz: ", twoD)
}
