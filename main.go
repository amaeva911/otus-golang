package main

import (
	"fmt"
	"log"
)

func scanNumbersInput() int {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Неверный формат вводимых данных. Ожидается целое число.")
		log.Fatal(err)
	}
	return number
}

func drawBoard(columns int, lines int) {
	for y := 0; y < lines; y++ {
		for x := 0; x < columns; x++ {
			print("#  ")
		}
		print("\n")
	}

}

func main() {
	fmt.Println("Введите число строк: ")
	numberOfColumns := scanNumbersInput()
	fmt.Println("Введите число столбцов: ")
	numberOfLines := scanNumbersInput()

	fmt.Printf("Рисуем шахматную доску %v на %v.\n", numberOfColumns, numberOfLines)

	drawBoard(numberOfLines, numberOfColumns)
}
