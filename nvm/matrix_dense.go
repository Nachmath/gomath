// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42611|42630

package nvm

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

var (
	_m *M
	_  Matrix = _m
)

// M represents a dense matrix.
type M struct {
	// O can be *mat.Dense or mat.Transpose.
	O mat.Matrix
}

// NewM creates a new matrix of rows `r` and cols `c`,
// NewM will panic if `r <= 0` or `c <= 0`.
func NewM(r, c int) *M {
	if r <= 0 {
		panic(ErrRows)
	}
	if c <= 0 {
		panic(ErrCols)
	}
	return &M{O: mat.NewDense(r, c, nil)}
}

// NewMShare creates a new matrix of rows `r` and cols `c` with `data`,
// NewMShare will panic if `r <= 0` or `c <= 0` or `len(data) != r*c`.
// Note: Slice both `M.Data` and `data` share a same underlying array.
func NewMShare(r, c int, data []float64) *M {
	if r <= 0 {
		panic(ErrRows)
	}
	if c <= 0 {
		panic(ErrCols)
	}
	if len(data) != r*c {
		panic(ErrShape)
	}
	return &M{O: mat.NewDense(r, c, data)}
}

// NewMCopy creates a new matrix of rows `r` and cols `c` with `data`,
// NewMCopy will panic if `r <= 0` or `c <= 0` or `len(data) != r*c`.
func NewMCopy(r, c int, data []float64) *M {
	l := len(data)
	if r <= 0 {
		panic(ErrRows)
	}
	if c <= 0 {
		panic(ErrCols)
	}
	if l != r*c {
		panic(ErrShape)
	}
	newData := make([]float64, l)
	copy(newData, data)
	return &M{O: mat.NewDense(r, c, newData)}
}

// IsNaM reports whether `m` is "Not-a-Matrix".
func (m *M) IsNaM() bool {
	if m.O == nil {
		return true
	}
	return false
}

// Dims returns the rows `r` and cols `c` of the matrix.
func (m *M) Dims() (r, c int) {
	if m.IsNaM() {
		return 0, 0
	}
	return m.O.Dims()
}

// At returns the element at position `i`th row and `j`th col.
// At will panic if `i` or `j` is out of bounds.
func (m *M) At(i, j int) float64 {
	if m.IsNaM() {
		panic(ErrNaM)
	}
	r, c := m.Dims()
	if i < 0 || i >= r {
		panic(ErrRowIndex)
	}
	if j < 0 || j >= c {
		panic(ErrColIndex)
	}
	return m.O.At(i, j)
}

// SetAt sets `f` to the element at position `i`th row and `j`th col.
// SetAt will panic if `i` or `j` is out of bounds.
func (m *M) SetAt(i, j int, f float64) Matrix {
	if m.IsNaM() {
		panic(ErrNaM)
	}
	r, c := m.Dims()
	if i < 0 || i >= r {
		panic(ErrRowIndex)
	}
	if j < 0 || j >= c {
		panic(ErrColIndex)
	}
	m.O.(*mat.Dense).Set(i, j, f)
	return m
}

// T returns the transpose of the matrix.
func (m *M) T() Matrix {
	if m.IsNaM() {
		panic(ErrNaM)
	}
	return &M{O: m.O.T()}
}

// Det computes the determinant of the square matrix.
func (m *M) Det() float64 {
	if m.IsNaM() {
		panic(ErrNaM)
	}
	if r, c := m.Dims(); r != c {
		panic(ErrShape)
	}
	return mat.Det(m.O)
}

// Inv computes the inverse of the square matrix.
// Inv will panic if the square matrix is not invertible.
func (m *M) Inv() Matrix {
	if m.IsNaM() {
		panic(ErrNaM)
	}
	r, c := m.Dims()
	if r != c {
		panic(ErrShape)
	}
	inv := mat.NewDense(r, c, nil)
	err := inv.Inverse(m.O)
	if err != nil {
		panic(errors.New(ErrInverse.Error() + " > " + err.Error()))
	}
	return &M{O: inv}
	// https://en.wikipedia.org/wiki/Invertible_matrix
}

// Add adds the matrices `x` and `y` element-wise, placing the result in the receiver.
// Add will panic if the two matrices do not have the same shape.
func (m *M) Add(x, y Matrix) Matrix {
	if !IsSameShape(x, y) || !IsSameShape(m, x) {
		panic(ErrShape)
	}
	/// *M & *M
	if x, ok := x.(*M); ok {
		if y, ok := y.(*M); ok {
			m.O.(*mat.Dense).Add(x.O, y.O)
			return m
		}
	}
	panic(ErrImpType)
}

// Sub subtracts the matrices `y` from `x` element-wise, placing the result in the receiver.
// Sub will panic if the two matrices do not have the same shape.
func (m *M) Sub(x, y Matrix) Matrix {
	if !IsSameShape(x, y) || !IsSameShape(m, x) {
		panic(ErrShape)
	}
	/// *M & *M
	if x, ok := x.(*M); ok {
		if y, ok := y.(*M); ok {
			m.O.(*mat.Dense).Sub(x.O, y.O)
			return m
		}
	}
	panic(ErrImpType)
}

// Scale multiplies the elements of `x` by `f`, placing the result in the receiver.
func (m *M) Scale(f float64, x Matrix) Matrix {
	if !IsSameShape(m, x) {
		panic(ErrShape)
	}
	/// *M
	if x, ok := x.(*M); ok {
		m.O.(*mat.Dense).Scale(f, x.O)
		return m
	}
	panic(ErrImpType)
}

// DotMul performs element-wise multiplication of `x` and `y`, placing the result in the receiver.
// DotMul will panic if the two matrices do not have the same shape.
func (m *M) DotMul(x, y Matrix) Matrix {
	if !IsSameShape(x, y) || !IsSameShape(m, x) {
		panic(ErrShape)
	}
	/// *M & *M
	if x, ok := x.(*M); ok {
		if y, ok := y.(*M); ok {
			m.O.(*mat.Dense).MulElem(x.O, y.O)
			return m
		}
	}
	panic(ErrImpType)
}

// Mul computes the matrix product of `x` and `y`, placing the result in the receiver.
// Mul will panic if the cols of `x` is not equal to the rows of `y`.
// Note: Mul will panic if the receiver and the product do not have the same shape.
func (m *M) Mul(x, y Matrix) Matrix {
	xr, xc := x.Dims()
	yr, yc := y.Dims()
	if xc != yr || x.IsNaM() || y.IsNaM() {
		panic(ErrShape)
	}
	mr, mc := m.Dims()
	if mr != xr || mc != yc {
		panic(ErrShape)
	}
	/// *M & *M
	if x, ok := x.(*M); ok {
		if y, ok := y.(*M); ok {
			m.O.(*mat.Dense).Mul(x.O, y.O)
			return m
		}
	}
	panic(ErrImpType)
}
