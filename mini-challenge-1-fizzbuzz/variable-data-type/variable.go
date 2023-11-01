package main

import "fmt"

func main() {
	// with data type
	// var name string
	// name = "tiara"
	// fmt.Printf("halo nama saya %s", name)

	// var number int
	// number = 1
	// fmt.Println(number)

	// with out data type

	// newName := "tiara"
	// fmt.Println(newName)

	// age := 1
	// fmt.Print(age)

	// multiple declaration variable

	// firstName, LastName, age := "tiara", "juli", 1
	// var NewfirstName, NewLastName, _ = "tiara", "juli", 1
	// var FName, LName, AgePerson string
	// FName = "yaya"
	// fmt.Println(FName)

	// const (
	// 	salery = 10000
	// 	numb   = iota
	// 	numb1
	// 	numb2
	// )
	// fmt.Println(numb, numb1, numb2)

	// // fmt.Println(salery + numb)
	// // fmt.Println(salery + numb)

	// status:= true
	// fmt.Println(!status)

	// func main() {
	// outerLoop:
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println("Perulangan ke - ", i+1)
	// 		for j := 0; j < 3; j++ {
	// 			if i == 2 {
	// 				break outerLoop
	// 			}
	// 			fmt.Print(j, " ")
	// 		}
	// 		fmt.Print("\n")
	// 	}
	// }

	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("j")
	// 	}
	// 	fmt.Println("i")
	// }

	// for i := 1; i <= 5; i++ {
	/  	}
	// 	fmt.Println(i)
	// }
	var number = [5]int{1, 4, 19, 16, 25}
	for _, v := range number {
		fmt.Println("angka :", v)
	}

}
