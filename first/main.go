package main

import "fmt"

func main()  {
	isappgood := true
	isfalse := false
	a := 54
	b := 3.14
	c := "uma string qualquer"
	var name string = "Leonardo"
	fmt.Printf("%T\n", isappgood)
	fmt.Printf("%T\n", isfalse)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println("Meu nome é", name)
	var nota, total int = 90, 100
	println("voce tirou", nota, "de", total)
	println("ola", name, "voce tirou", nota, "de", total)
	var (
		sobrenome = "neves"
		id = 16
		admin = false
	)
	println("ola", sobrenome, "registro", id, "Administrador: ", admin)
	println("Nome", name, "\nRegistro", id)
	println(c)

}