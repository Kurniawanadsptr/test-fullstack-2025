package main

import (
	"fmt"
)

func main() {
	// if status == "active" {
	// fmt.Println("User is active")
	// } else if status == "inactive" {
	// fmt.Println("User is inactive")
	// } else if status == "banned" {
	// fmt.Println("User is banned")
	// } else {
	// fmt.Println("Unknown status")
	// }

	status := "playing game"

	switch status {
	case "active":
		fmt.Println("User is active")
	case "inactive":
		fmt.Println("User is inactive")
	case "banned":
		fmt.Println("User is banned")
	default:
		fmt.Println("Unknown status")
	}

	//Jawaban B
	//Alasan saya memilih jawaban B karena daripada menggunakan if dan else itu lebih baik menggunakan switch yang dimana lebih singkat dan rapi 
	//Daripada if dan else selain itu juga sangat gampang untuk menambahkan status yang baru 
	//Dan saya akan mengirimkan hasil juga yang saya running tiap dari status
}