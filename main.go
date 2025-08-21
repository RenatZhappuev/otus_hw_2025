package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Не получилось прочитать строку: ", err)
	}
	input = strings.TrimSpace(input)
	size, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Число введено неверно: ", err)
	}
	fmt.Println(createChessboard(size))
}

func createChessboard(size int) string {
	var chessboard strings.Builder
	for i := 0; i < size; i++{
		for j := 0; j < size; j++{
			if (i+j) % 2 == 0{
				chessboard.WriteString(" ")
			} else {
				chessboard.WriteString("#")
			}
		}
		chessboard.WriteString("\n")
	}
	return chessboard.String()
}