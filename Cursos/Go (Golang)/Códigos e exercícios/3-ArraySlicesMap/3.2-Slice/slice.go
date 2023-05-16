package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // Array
	s1 := []int{1, 2, 3}  // Slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:4] // Slice aponta para o array
	fmt.Println(a2, s2)

	// Alterar um afeta o outro
	a2[2] = 7
	s2[0] = 8
	fmt.Println(a2, s2)

	s3 := s2[:1] // Slice de slice
	fmt.Println(s2, s3)

	// Criando slice usando a função make
	s := make([]int, 10) // cria um slice de int de tamanho 10
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 15) // O terceiro elemento define a capacidade do vetor associado
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5) // Adiciona elementos ao slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 6) // Caso a capacidade do vetor seja ultrapassada, a capacidade é aumentada automáticamente
	fmt.Println(s, len(s), cap(s))

	// Dois slices referenciando o mesmo array interno
	s1 = append(s, 1)
	s = append(s, 4)
	fmt.Println(s, s1)

	var slice1 []int
	slice1 = append(slice1, 3, 4, 5)
	slice2 := make([]int, 2)
	// Copia os elementos de um slice para outro
	copy(slice2, slice1) // Como slice2 é menor, vai copiar apenas o que couber
	fmt.Println(slice1, slice2)

	slice2[1] = 13
	fmt.Println(slice1, slice2) // Nesse caso não apontam para o mesmo vetor

}
