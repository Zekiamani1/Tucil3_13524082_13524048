package core

import(
	"io"
	"bufio"
	"fmt"
)

func ParseInput(r io.Reader) (int, int, []string, [][]int, error) {
	reader := bufio.NewReader(r)

	var X, Y int
	if _, err := fmt.Fscan(reader, &X, &Y); err != nil {
		return 0, 0, nil, nil, err
	}

	matrix := make([]string, X)
	for i := 0; i < X; i++ {
		if _, err := fmt.Fscan(reader, &matrix[i]); err != nil {
			return 0, 0, nil, nil, err
		}
	}

	costMatrix := make([][]int, X)
	for i := 0; i < X; i++ {
		costMatrix[i] = make([]int, Y)
		for j := 0; j < Y; j++ {
			if _, err := fmt.Fscan(reader, &costMatrix[i][j]); err != nil {
				return 0, 0, nil, nil, err
			}
		}
	}

	return X, Y, matrix, costMatrix, nil
}