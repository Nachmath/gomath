// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42526|42618

package nvm

// Vector is a vector interface.
type Vector interface {
	// IsNaV reports whether `v` is "Not-a-Vector".
	IsNaV() bool

	// Dim returns the dimension of vector.
	Dim() int

	// At returns the element at position `i`. At will panic if `i` is out of bounds.
	At(i int) float64

	// SetAt sets `f` to the element at position `i`. SetAt will panic if `i` is out of bounds.
	SetAt(i int, f float64) Vector

	// NormL0 returns L0 norm. L0 norm equals the number of non-zero (include NaN/Inf) elements in the vector.
	NormL0() int

	// NormL1 returns L1 norm (Manhattan distance). L1 norm equals `\sum_{i=0}^n{|x_i|}`.
	NormL1() float64

	// NormL2 returns L2 norm (Euclidean distance). L2 norm equals `\sqrt{\sum_{i=0}^n{x_i^2}}`.
	NormL2() float64

	// Unit returns the unit vector.
	Unit() Vector

	// Scale scales the vector by `f`.
	Scale(f float64) Vector

	// IsZero reports whether the vector is zero vector, the every element is zero.
	IsZero() bool
}

// IsVSameShape reports whether the vectors `x` and `y` have the same dimension.
func IsVSameShape(x, y Vector) bool {
	if x.Dim() != y.Dim() || x.IsNaV() {
		return false
	}
	return true
}

// IsVEqual reports whether the vectors `x` and `y` have the same dimension, and their all elements have same value.
func IsVEqual(x, y Vector) bool {
	if x, ok := x.(*V); ok {
		if y, ok := y.(*V); ok {
			return vectorIsDenseEqual(x, y)
		}
	}
	panic(ErrImpType)
}

// VScale scales the vector `v` by `f`.
func VScale(f float64, v Vector) Vector {
	if v, ok := v.(*V); ok {
		return vectorDenseScale(f, v)
	}
	panic(ErrImpType)
}

// VAdd adds the vectors `x` and `y`.
func VAdd(x, y Vector) Vector {
	if x, ok := x.(*V); ok {
		if y, ok := y.(*V); ok {
			return vectorDenseAdd(x, y)
		}
	}
	panic(ErrImpType)
}

// VSub subtracts the vector `y` from `x`.
func VSub(x, y Vector) Vector {
	if x, ok := x.(*V); ok {
		if y, ok := y.(*V); ok {
			return vectorDenseSub(x, y)
		}
	}
	panic(ErrImpType)
}

// VDot returns the dot product of `x` and `y`, i.e. `\sum_{i=1}^N x[i]*y[i]`.
func VDot(x, y Vector) float64 {
	if x, ok := x.(*V); ok {
		if y, ok := y.(*V); ok {
			return vectorDenseDot(x, y)
		}
	}
	panic(ErrImpType)
}
