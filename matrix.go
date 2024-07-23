package mat

import "fmt"

type Matrix struct {
	data [][]float64
	rows, columns int
}

func CreateMatrix(rows, columns int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, columns)
	}

	return &Matrix{data: data, rows: rows,columns:columns}
}

func (m *Matrix) Add(n *Matrix) (*Matrix, error) {
	if m.rows != n.rows || m.columns != n.columns {
		return nil, fmt.Errorf("Matrix dimensions must match")
	}
	result := CreateMatrix(m.rows, m.columns)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
            result.data[i][j] = m.data[i][j] + n.data[i][j]
        }
	}
	return result, nil

}

func (m *Matrix) Multiply(n *Matrix) (*Matrix, error) {
    if m.columns != n.rows {
        return nil, fmt.Errorf("matrix dimensions must match for multiplication")
    }

    result := CreateMatrix(m.rows, n.columns)
    for i := 0; i < m.rows; i++ {
        for j := 0; j < n.columns; j++ {
            sum := 0.0
            for k := 0; k < m.columns; k++ {
                sum += m.data[i][k] * n.data[k][j]
            }
            result.data[i][j] = sum
        }
    }
    return result, nil
}



func (m *Matrix) Transpose () *Matrix {
	result := CreateMatrix(m.columns, m.rows)
	for i:= 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			result.data[j][i] = m.data[i][j]
		}
	}
	return result
}

func (m *Matrix) String() string {
	result := ""
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			result += fmt.Sprintf("%f ", m.data[i][j])
		}
		result += "\n"
	}
	return result
}