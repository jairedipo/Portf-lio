package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print((" Linha."))

	fmt.Println("\nNova Linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x) // Inválido, pois não é possível somar dois tipos diferentes no print
	xs := fmt.Sprint(x) // Retorna uma string do que seria printado
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.3f\n", x) // Formatação e arredondamento

	// Outros parâmetros para a função printf
	a := 1
	b := 1.99999
	c := false
	d := "opa"
	fmt.Printf("%d %f %.1f %t %s\n", a, b, b, c, d)
	fmt.Printf("%v %v %v %v\n", a, b, c, d) // %v aceita alguns tipos no print

	// Outra função
	e := fmt.Sprintf("O número %f arredondado é %.0f", b, b) // Retorna uma string do que seria printado com formatação
	fmt.Println(e)
}
