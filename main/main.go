package main

import (
	"fmt"
	"log"
)

func main() {

	// Комментарий для релиза 1.0
	var s string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели данные: %s\n", s)
}
