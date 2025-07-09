# Matrix Multiplication CLI

A simple command-line interface program written in Go that performs matrix multiplication based on user input. It checks dimension compatibility, takes in matrix values, multiplies them, and prints the result.

---

## Project Link

[https://github.com/kleinebossie/matrix-multiplication](https://github.com/kleinebossie/matrix-multiplication)

---

## Features

- Prompt-based CLI for entering matrix sizes and values
- Dimension validation for multiplication
- Dynamically builds and multiplies matrices
- Displays the resulting matrix after computation

---

## Requirements

- Go 1.20 or higher (or any version with support for modules and `go run`)
- Terminal/command prompt access

---

## Usage

1. **Clone the repo**:
   ```bash
   git clone https://github.com/kleinebossie/matrix-multiplication.git
   cd matrix-multiplication
   ```
2. **Run the program**:
   ```bash
   go run main.go
   ```
3. **Follow the prompts**:
   - Enter the dimensions (rows and columns) for the first and second matrices.
   - Input all values row by row.
   - The program multiplies the matrices and prints the result.

## Example

```bash
***Matrix Multiplication CLI***

---First matrix size---
Rows and columns: 2 3

---Second matrix size---
Rows and columns: 3 2

---Fill in first matrix---

1ST ROW
Column 1: 1
Column 2: 2
Column 3: 3
2ND ROW
Column 1: 4
Column 2: 5
Column 3: 6

---Fill in second matrix---

1ST ROW
Column 1: 7
Column 2: 8
2ND ROW
Column 1: 9
Column 2: 10
3RD ROW
Column 1: 11
Column 2: 12

---RESULT---
[58 64]
[139 154]
```
---
## Structure

Main file: main.go

**Key functions**:
- getMatrixSize() – prompts for dimensions
- createMatrix() – collects matrix values
- checkValid() – checks multiplication compatibility
- multiplyMatrices() – multiplies two matrices and returns the result

## Notes

Input must be valid numbers; bad input will crash the program (no input sanitization).
Only works with float64 values (entered as numbers, not fractions).
Strictly synchronous and linear—no concurrency or advanced optimizations.

## License

**MIT**
