package main

import (
	"fmt"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getMatrixFromUser() (matrix [][]int) {
	fmt.Print("No. of rows for matrix: ")
	var noOfRows int
	_, err1 := fmt.Scan(&noOfRows)
	checkError(err1)

	fmt.Print("No. of columns for matrix: ")
	var noOfColumns int
	_, err2 := fmt.Scan(&noOfColumns)
	checkError(err2)

	fmt.Println("Enter Matrix Elements: ")
	matrix = make([][]int, noOfRows)
	for i := 0; i < noOfRows; i++ {
		matrix[i] = make([]int, noOfColumns)
		for j := 0; j < noOfColumns; j++ {
			_, err := fmt.Scan(&matrix[i][j])
			checkError(err)
		}
	}

	return
}
func matrixMultiplication(m1, m2 [][]int) (m3 [][]int) {
	if len(m1[0]) != len(m2) {
		fmt.Println("Matrix multiplication not possible!!!")
		return
	}
	m3 = make([][]int, len(m1))
	for i := 0; i < len(m3); i++ {
		m3[i] = make([]int, len(m2[0]))
		for j := 0; j < len(m3[i]); j++ {
			m3[i][j] = 0
			for k := 0; k < len(m2); k++ {
				m3[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	return
}

func main() {
	fmt.Println("Matrix 1: ")
	matrix1 := getMatrixFromUser()
	fmt.Println("Matrix 2: ")
	matrix2 := getMatrixFromUser()
	fmt.Printf("Matrix Mulitiplication : %v", matrixMultiplication(matrix1, matrix2))
}
