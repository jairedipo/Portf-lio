package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(20000))

	// sem sinal uint8, uint16, uint32, uint64 ('u' -> unsigned)
	var b1 byte = 255 // byte = uint8
	fmt.Println("O byte é", reflect.TypeOf(b1))

	// com sinal int8, int16, int32, int64
	i1 := math.MaxInt64
	i2 := 234
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i2)) // int reflete a arquitetura da máquina, x32 ou x64

	var i3 rune = 'a'                           // representa o mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i3)) // int32
	fmt.Println(i3)                             // 97

	// numeros reais float32 float64
	var x = float32(49.99)                                          // força a variável a ser daquele tipo
	fmt.Println("O tipo de x é", reflect.TypeOf(x))                 // float32
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64 -> padrão

	// boolean bool
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) // false

	// string
	s1 := "Olá mundo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	mundo
	!`
	fmt.Println("O tamanho da string é", len(s2)) // considera os tabs e quebra de linha

	// char não existe
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char)) // int32
	fmt.Println("O valor de char é", char)                // 97

	// zeros
	var a int
	var b float64
	var c bool
	var d string
	var e *int // ponteiro de inteiro

	fmt.Printf("%v %v %v %v %v\n", a, b, c, d, e) // <nil> == null
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)   // %q mostra as aspas da string
}
