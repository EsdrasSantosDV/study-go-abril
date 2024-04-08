package main

import "fmt"

func main() {
	var esdras = "Esdras"
	var idade = 23

	//O GO JA INFERE O TIPO DAS VARIAVEIS
	fmt.Println(esdras)
	fmt.Println(idade)

	//multiplas variaveis
	var altura, peso, sobrenome = 1.80, 90, "Santos"

	fmt.Println(altura)
	fmt.Println(peso)
	fmt.Println(sobrenome)

	//Variáveis ​​declaradas sem uma
	//inicialização correspondente têm valor zero .
	//Por exemplo, o valor zero para an inté 0.
	var e int
	fmt.Println(e)

	/*
		A :=sintaxe é uma abreviação para declarar e
		inicializar uma variável, por exemplo,
		para var f string = "esdras "neste caso.
		Esta sintaxe está disponível apenas dentro de funções.
		sem precisar colocar o var
	*/
	dindo := "esdras"

	fmt.Println(dindo)
}
