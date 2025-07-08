package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("***Matrix Multiplication CLI***\n")

	fmt.Println("---First matrix size---")
	rows1, cols1 := getMatrixSize()
	fmt.Print("\n")
	fmt.Println("---Second matrix size---")
	rows2, cols2 := getMatrixSize()
	fmt.Print("\n")

	err := checkValid(cols1, rows2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---Fill in first matrix---\n")
	matrix1 := createMatrix(rows1, cols1)
	fmt.Print("\n")
	fmt.Println("---Fill in second matrix---\n")
	matrix2 := createMatrix(rows2, cols2)

	result := multiplyMatrices(matrix1, matrix2)
	fmt.Println("\n---RESULT---")
	for i := range result.Rows {
		fmt.Println(result.Data[i])
	}
}

type Matrix struct {
	Rows int
	Cols int
	Data [][]float64
}

func getMatrixSize() (int, int) {
	var rows, cols int

	fmt.Print("Rows and columns: ")

	_, err := fmt.Scanln(&rows, &cols)
	if err != nil {
		log.Fatalf("Failed to get rows and columns: %v", err)
	}

	return rows, cols
}

func checkValid(cols1, rows2 int) error {
	if cols1 != rows2 {
		return errors.New("multiplication impossible for given sizes")
	} else {
		return nil
	}
}

func createMatrix(rows, cols int) Matrix {
	var matrix Matrix

	matrix.Rows = rows
	matrix.Cols = cols

	matrix.Data = make([][]float64, rows)
	for i := range matrix.Data {
		matrix.Data[i] = make([]float64, cols)
	}

	for i := range rows {
		switch i {
		case 0:
			fmt.Printf("%dST ROW\n", i+1)
		case 1:
			fmt.Printf("%dND ROW\n", i+1)
		case 2:
			fmt.Printf("%dRD ROW\n", i+1)
		default:
			fmt.Printf("%dTH ROW\n", i+1)
		}

		for j := range cols {
			fmt.Printf("Column %d: ", j+1)

			_, err := fmt.Scanln(&matrix.Data[i][j])
			if err != nil {
				log.Fatalf("could not get data: %v", err)
			}
		}
	}

	return matrix
}

func multiplyMatrices(m1, m2 Matrix) Matrix {
	var matrix Matrix

	matrix.Rows = m1.Rows
	matrix.Cols = m2.Cols

	matrix.Data = make([][]float64, matrix.Rows)
	for i := range matrix.Data {
		matrix.Data[i] = make([]float64, matrix.Cols)
	}

	for i := 0; i < matrix.Rows; i++ {
		for j := 0; j < matrix.Cols; j++ {
			for k := 0; k < m2.Rows; k++ {
				matrix.Data[i][j] += m1.Data[i][k] * m2.Data[k][j]
			}
		}
	}
	return matrix
}
