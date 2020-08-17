package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//el paquete bcrypt no esta cargado por default por eso antes que nada en la terminal hay que correr:
// go get golang.org/x/crypto/bcrypt
//si el paquete esta descargado y se lo quiere actualizar se usa go get -u golang.org/x/crypto/bcrypt
func main() {

	//la idea al almacenar datos sensibles es hacerlo encriptando los datos, para eso existe el paquete bcrypt
	pass := "password1234"
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pass)
	fmt.Println(string(hash))

	error := bcrypt.CompareHashAndPassword(hash, []byte(pass))

	if error != nil {
		fmt.Println("password incorrecto")
	} else {
		fmt.Println("password correcto")
	}

	loginPassError := "lalalalal"
	hashTwo, errTwo := bcrypt.GenerateFromPassword([]byte(loginPassError), bcrypt.MinCost)
	fmt.Println(string(hashTwo))
	if errTwo != nil {
		fmt.Println(errTwo)
	}
	ok := bcrypt.CompareHashAndPassword(hashTwo, []byte(pass))
	if ok != nil {
		fmt.Println("la contraseña no es valida")
	} else {
		fmt.Println("la contraseña es valida")
	}

}
