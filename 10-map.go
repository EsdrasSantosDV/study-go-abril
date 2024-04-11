package main

import "fmt"

func main() {

	//CIRAR MEU PODEROSIMO HASHMAP
	//primeiro e o tipo da chave, segundo o tipo do valor

	dindo := make(map[string]string)
	m := make(map[string]int)
	m["k1"] = 10
	m["k2"] = 120
	dindo["k1"] = "esdras"
	dindo["k2"] = "esdras2"
	fmt.Println(m["k1"])
	fmt.Println(m["k2"])
	fmt.Println(dindo["k1"])
	fmt.Println(m)

	//pra deletar um valor do hashmap
	delete(m, "k2")
	fmt.Println(m)
	for s2 := range m {
		fmt.Println("key:", s2, "value:", m[s2])
	}

}
