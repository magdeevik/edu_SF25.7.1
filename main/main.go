package main

import (
	"fmt"
	"log"
)

func main() {
	var s string
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %s\n", s)
}
