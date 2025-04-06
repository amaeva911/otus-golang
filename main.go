package main

import (
	"fmt"
	"log"
)

func scanNumbersInput() int {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func drawBoard(lines int, columns int) {
	for y := 0; y < columns; y++ {
		for x := 0; x < lines; x++ {
			print("#  ")
		}
		print("\n")
	}

}

func main() {
	fmt.Println("Введите число строк: ")
	numberOfLines := scanNumbersInput()
	fmt.Println("Введите число столбцов: ")
	numberOfColumns := scanNumbersInput()

	fmt.Printf("Рисуем шахматную доску %v на %v.\n", numberOfLines, numberOfColumns)

	drawBoard(numberOfLines, numberOfColumns)
}
