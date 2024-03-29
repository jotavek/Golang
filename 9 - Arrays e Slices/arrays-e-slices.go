package main

import "fmt"

func main() {
	//Eu especifico o tamanho em colchetes e o tipo do array
	//Se eu não especificar o tamanho ele é considerado como um slice
	var array1 [5]string
	array1[0] = "posição 1"
	fmt.Println(array1)

	array2 := [5]string{"posição1", "posição2", "posição3", "posição4", "posição5"}
	fmt.Println(array2)

	//No slice eu não especifico o tamanho nos colchetes
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
	fmt.Println(slice)

	//Adicionando item em uma slice com append.
	slice = append(slice, 23)
	fmt.Println(slice)

	//Eu posso pegar uma fatia de um array e colocar em um slice, eu coloco os indices do que eu quero nos colchetes.
	slice2 := array2[0:2]
	fmt.Println(slice2)

	//Arrays Internos
	//Quando a cap maxima do array é atingida, ele dobra de tamanho automaticamente
	fmt.Println("____________")
	slice3 := make([]float32, 10, 11)
	slice3 = append(slice3, 11)
	slice3 = append(slice3, 12)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
