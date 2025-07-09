package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("***Matrix Multiplication CLI***\n")

	// Get dimensions for the first matrix
	fmt.Println("---First matrix size---")
	rows1, cols1 := getMatrixSize()
	fmt.Print("\n")

	// Get dimensions for the second matrix
	fmt.Println("---Second matrix size---")
	rows2, cols2 := getMatrixSize()
	fmt.Print("\n")

	// Check if matrix multiplication is possible with given dimensions
	err := checkValid(cols1, rows2)
	if err != nil {
		log.Fatal(err)
	}

	// Create and fill the first matrix
	fmt.Println("---Fill in first matrix---\n")
	matrix1 := createMatrix(rows1, cols1)
	fmt.Print("\n")

	// Create and fill the second matrix
	fmt.Println("---Fill in second matrix---\n")
	matrix2 := createMatrix(rows2, cols2)

	// Multiply the matrices and display the result
	result := multiplyMatrices(matrix1, matrix2)
	fmt.Println("\n---RESULT---")
	for i := range result.Rows {
		fmt.Println(result.Data[i])
	}
}

// Matrix struct represents a mathematical matrix with rows, columns, and data
type Matrix struct {
	Rows int         // Number of rows in the matrix
	Cols int         // Number of columns in the matrix
	Data [][]float64 // 2D slice containing the matrix elements
}

// getMatrixSize prompts the user to input matrix dimensions and returns them
func getMatrixSize() (int, int) {
	var rows, cols int

	fmt.Print("Rows and columns: ")

	// Read two integers from standard input
	_, err := fmt.Scanln(&rows, &cols)
	if err != nil {
		log.Fatalf("Failed to get rows and columns: %v", err)
	}

	return rows, cols
}

// checkValid verifies if matrix multiplication is possible (cols1 must equal rows2)
func checkValid(cols1, rows2 int) error {
	if cols1 != rows2 {
		return errors.New("multiplication impossible for given sizes")
	} else {
		return nil
	}
}

// createMatrix initializes and fills a matrix with user-provided values
func createMatrix(rows, cols int) Matrix {
	var matrix Matrix

	matrix.Rows = rows
	matrix.Cols = cols

	// Initialize the 2D slice for matrix data
	matrix.Data = make([][]float64, rows)
	for i := range matrix.Data {
		matrix.Data[i] = make([]float64, cols)
	}

	// Fill each element of the matrix with user input
	for i := range rows {
		// Print row number with proper ordinal suffix
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

		// Get values for each column in the current row
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

// multiplyMatrices performs matrix multiplication on two matrices
func multiplyMatrices(m1, m2 Matrix) Matrix {
	var matrix Matrix

	// Result matrix has rows from first matrix and columns from second matrix
	matrix.Rows = m1.Rows
	matrix.Cols = m2.Cols

	// Initialize the result matrix data structure
	matrix.Data = make([][]float64, matrix.Rows)
	for i := range matrix.Data {
		matrix.Data[i] = make([]float64, matrix.Cols)
	}

	// Perform matrix multiplication using triple nested loops
	for i := 0; i < matrix.Rows; i++ {
		for j := 0; j < matrix.Cols; j++ {
			for k := 0; k < m2.Rows; k++ {
				matrix.Data[i][j] += m1.Data[i][k] * m2.Data[k][j]
			}
		}
	}
	return matrix
}
