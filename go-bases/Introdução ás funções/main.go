package main

import "fmt"

func main() {
	var salario float64
	fmt.Printf("Insira o valor do salário: ")
	fmt.Scanln(&salario)

	impostoSalario(salario)
}

func impostoSalario(salario float64) (deducaoPercentual float64, deducao float64, salarioDeduzido float64) {

	if salario >= 50000 && salario <= 149999 {
		deducaoPercentual = 17
	} else if salario >= 150000 {
		deducaoPercentual = 27
	} else {
		deducaoPercentual = 10
	}

	deducao = (salario * deducaoPercentual) / 100.00
	salarioDeduzido = salario - deducao

	fmt.Printf("Valor incial do salário: %.2f \n A dedução foi de %.0f%% \n O valor deduzido foi: %.2f \n O salário atual é de: %.2f", salario, deducaoPercentual, deducao, salarioDeduzido)

	return
}
