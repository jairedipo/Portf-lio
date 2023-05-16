package main

import "fmt"

func main() {
	aprovados := make(map[int]string) // map[chave]valor

	aprovados[12345678978] = "Maria"
	aprovados[98765432121] = "Pedro"
	aprovados[54540654124] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // For usando a chave e o valor
		fmt.Printf("%s (CPF %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978) // Remove um registro do map
	fmt.Println(aprovados)
	fmt.Println(aprovados[12345678978])

	// Declarando map já com valores iniciais
	funcESalarios := map[string]float64{
		"José João":      12345.67,
		"Gabriela Silva": 23456.78,
		"Pedro Junior":   1200.00,
	}

	funcESalarios["Rafael Filho"] = 1350.00 // É possível adicionar elementos a qualquer momento
	delete(funcESalarios, "Inexistente")    // Deletar algo que não está presente no map não gera erro

	for nome, salario := range funcESalarios {
		fmt.Printf("%s R$%.2f\n", nome, salario)
	}

	// Maps aninhados
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 12345.67,
			"Guga":     6789.00,
		},
		"J": {
			"José": 4567.89,
		},
		"P": {
			"Pedro": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	// Exemplo de for aninhado para acessar o map aninhado
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(letra, nome, salario)
		}
	}
}
