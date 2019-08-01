// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42610|42630

package nvm

// Matrix is a matrix interface.
type Matrix interface {
	// IsNaM reports whether `m` is "Not-a-Matrix".
	IsNaM() bool

	// Dims returns the rows `r` and cols `c` of the matrix.
	Dims() (r, c int)

	// At returns the element at position `i`th row and `j`th col.
	// At will panic if `i` or `j` is out of bounds.
	At(i, j int) float64

	// SetAt sets `f` to the element at position `i`th row and `j`th col.
	// SetAt will panic if `i` or `j` is out of bounds.
	SetAt(i, j int, f float64) Matrix

	// T returns the transpose of the matrix.
	T() Matrix

	// Det computes the determinant of the square matrix.
	Det() float64

	// Inv computes the inverse of the square matrix.
	// Inv will panic if the square matrix is not invertible.
	Inv() Matrix

	// // SwapRows swaps the `ri`th row and the `rj`th row.
	// SwapRows(ri, rj int) Matrix

	// // SwapCols swaps the `ci`th col and the `cj`th col.
	// SwapCols(ci, cj int) Matrix

}

// // TM represents a triangular matrix.
// // BM represents a band matrix.
// // SM represents a sparse matrix.

// IsSameShape reports whether the matrices `x` and `y` have the same shape.
// IsSameShape returns `false` if `x` or `y` is NaM.
func IsSameShape(x, y Matrix) bool {
	if x.IsNaM() || y.IsNaM() {
		return false
	}
	xr, xc := x.Dims()
	yr, yc := y.Dims()
	if xr != yr || xc != yc {
		return false
	}
	return true
}

// IsEqual reports whether the matrices `x` and `y` have the same shape,
// and are element-wise equal.
func IsEqual(x, y Matrix) bool {
	if !IsSameShape(x, y) {
		return false
	}
	r, c := x.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if !IsNEqual(x.At(i, j), y.At(i, j)) {
				return false
			}
		}
	}
	return true
}

// Add adds `x` and `y` element-wise, placing the result in the new matrix.
// Add will panic if the two matrices do not have the same shape.
func Add(x, y Matrix) Matrix {
	if !IsSameShape(x, y) {
		panic(ErrShape)
	}
	/// *M & *M
	if x, ok := x.(*M); ok {
		if y, ok := y.(*M); ok {
			return NewM(x.Dims()).Add(x, y)
		}
	}
	panic(ErrImpType)
}

// Sub subtracts `y` from `x` element-wise, placing the result in the new matrix.
// Sub will panic if the two matrices do not have the same shape.
func Sub(x, y Matrix) Matrix {
	if !IsSameShape(x, y) {
		panic(ErrShape)
	}
	/// *M & *M
	if x, ok := x.(*M); ok {
		if y, ok := y.(*M); ok {
			return NewM(x.Dims()).Sub(x, y)
		}
	}
	panic(ErrImpType)
}

// Scale multiplies the elements of `x` by `f`, placing the result in the new matrix.
func Scale(f float64, x Matrix) Matrix {
	if x.IsNaM() {
		panic(ErrNaM)
	}
	if x, ok := x.(*M); ok {
		return NewM(x.Dims()).Scale(f, x)
	}
	panic(ErrImpType)
}

// DotMul performs element-wise multiplication of `x` and `y`, placing the result in the new matrix.
// DotMul will panic if the two matrices do not have the same shape.
func DotMul(x, y Matrix) Matrix {
	if !IsSameShape(x, y) {
		panic(ErrShape)
	}
	if x, ok := x.(*M); ok {
		if y, ok := y.(*M); ok {
			return NewM(x.Dims()).DotMul(x, y)
		}
	}
	panic(ErrImpType)
}

// Mul computes the matrix product of `x` and `y`, placing the result in the new matrix.
// Mul will panic if the cols of `x` is not equal to the rows of `y`.
func Mul(x, y Matrix) Matrix {
	if x, ok := x.(*M); ok {
		if y, ok := y.(*M); ok {
			xr, _ := x.Dims()
			_, yc := y.Dims()
			return NewM(xr, yc).Mul(x, y)
		}
	}
	panic(ErrImpType)
}
