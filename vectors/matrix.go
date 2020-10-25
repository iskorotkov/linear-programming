package vectors

type Matrix []Vector

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Columns() int {
	return len(m[0])
}

func (m Matrix) SubtractRow(rowIndex int) {
	mainCol := rowIndex
	for row := rowIndex + 1; row < m.Rows(); row++ {
		multiplier := m[row][mainCol] / m[rowIndex][mainCol]
		for col := mainCol; col < m.Columns(); col++ {
			m[row][col] -= multiplier * m[rowIndex][col]
		}
	}
}

func (m Matrix) GetVariable(index int) Vector {
	return m[index].GetVariable(index)
}

func (m Matrix) Slice(rows, columns Range) Matrix {
	totalRows := rows.Total()
	matrix := make([]Vector, totalRows)
	for i := 0; i < totalRows; i++ {
		matrix[i] = make([]float64, columns.Total())
		copy(matrix[i], m[rows.Start+i][columns.Start:columns.End])
	}
	return matrix
}

func (m Matrix) Determinant() float64 {
	for i := 0; i < m.Rows()-1; i++ {
		m.SubtractRow(i)
	}

	res := 1.0
	for i := 0; i < m.Rows(); i++ {
		res *= m[i][i]
	}

	return res
}

func (m Matrix) String() string {
	s := ""
	for _, row := range m {
		s += row.String() + "\n"
	}
	return s[:len(s)-1]
}
