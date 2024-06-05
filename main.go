package main

import (
	"TaskGo/usecase"
	"fmt"
)

func main() {
	// Membuat instance dari login melalui fungsi NewLogin
	loginInstance := usecase.NewLogin()

	Username := "Admin"
	Password := "admin123"

	if loginInstance.Authenticate(Username, Password) {
		fmt.Println("Berhasil login")
	} else {
		fmt.Println("Gagal Login")
	}

	fmt.Println(loginInstance)

}
