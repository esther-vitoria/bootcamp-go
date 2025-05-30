package main

import "fmt"

func main() {
	numeroMes := 7
	meses := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	numeroMes--

	if numeroMes >= 0 && numeroMes < 12 {
		fmt.Println(meses[numeroMes])
	} else {
		fmt.Print("Número invalido!")
	}

}
