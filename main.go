package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(createChessboard(8, 8))
}

func createChessboard(lenn int, wid int) string {
	var chessboard strings.Builder
	for i := 0; i < lenn; i++{
		for j := 0; j < wid; j++{
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

