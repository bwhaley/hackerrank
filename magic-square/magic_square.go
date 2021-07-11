package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func formingMagicSquare(s [][]int32) int32 {
	magicSquares := [][][]int32{
		[][]int32{[]int32{8, 3, 4}, []int32{1, 5, 9}, []int32{6, 7, 2}},
		[][]int32{[]int32{8, 1, 6}, []int32{3, 5, 7}, []int32{4, 9, 2}},
		[][]int32{[]int32{6, 7, 2}, []int32{1, 5, 9}, []int32{8, 3, 4}},
		[][]int32{[]int32{6, 1, 8}, []int32{7, 5, 3}, []int32{2, 9, 4}},
		[][]int32{[]int32{4, 3, 8}, []int32{9, 5, 1}, []int32{2, 7, 6}},
		[][]int32{[]int32{4, 9, 2}, []int32{3, 5, 7}, []int32{8, 1, 6}},
		[][]int32{[]int32{2, 9, 4}, []int32{7, 5, 3}, []int32{6, 1, 8}},
		[][]int32{[]int32{2, 7, 6}, []int32{9, 5, 1}, []int32{4, 3, 8}},
	}
	lowestCost := int32(81)
	for sq := 0; sq < len(magicSquares); sq++ {
		var thisCost int32
		for i, row := range s {
			for j, num := range row {
				thisCost += abs(num - magicSquares[sq][i][j])
			}
		}
		if thisCost < lowestCost {
			lowestCost = thisCost
		}
	}
	return lowestCost
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
