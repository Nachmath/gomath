// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42526|42617

package nvm

import "math"

var (
	_v *V
	_  Vector = _v
)

// V represents a vector.
type V struct {
	Data []float64
}

// NewV creates a new vector of dimension `d`. NewV will panic if `d <= 0`.
func NewV(d int) *V {
	if d <= 0 {
		panic(ErrDim)
	}
	return &V{Data: make([]float64, d)}
}

// NewVShare creates a new vector of dimension `len(data)`. NewVShare will panic if `len(data) == 0`.
// Note: Slice both `V.Data` and 'data' share a same underlying array.
func NewVShare(data []float64) *V {
	if len(data) == 0 {
		panic(ErrDim)
	}
	return &V{Data: data}
}

// NewVCopy creates a new vector of dimension `len(data)`. NewVCopy will panic if `len(data) == 0`.
func NewVCopy(data []float64) *V {
	d := len(data)
	if d == 0 {
		panic(ErrDim)
	}
	newData := make([]float64, d)
	copy(newData, data)
	return &V{Data: newData}
}

// IsNaV reports whether `v` is "Not-a-Vector".
func (v *V) IsNaV() bool {
	if len(v.Data) == 0 {
		return true
	}
	return false
}

// Dim returns the dimension of vector.
func (v *V) Dim() int {
	return len(v.Data)
}

// At returns the element at position `i`. At will panic if `i` is out of bounds.
func (v *V) At(i int) float64 {
	if i < 0 || i >= v.Dim() {
		panic(ErrIndex)
	}
	return v.Data[i]
}

// SetAt sets `f` to the element at position `i`. SetAt will panic if `i` is out of bounds.
func (v *V) SetAt(i int, f float64) Vector {
	if i < 0 || i >= v.Dim() {
		panic(ErrIndex)
	}
	v.Data[i] = f
	return v
}

// NormL0 returns L0 norm. L0 norm equals the number of non-zero (include NaN/Inf) elements in the vector.
func (v *V) NormL0() int {
	var c int
	for _, e := range v.Data {
		if !IsNEqual(e, 0) {
			c++
		}
	}
	return c
}

// NormL1 returns L1 norm (Manhattan distance). i.e. `\sum_{i=0}^n{|x_i|}`.
func (v *V) NormL1() float64 {
	if v.IsNaV() {
		panic(ErrNaV)
	}
	var n float64
	for _, e := range v.Data {
		n += math.Abs(e)
	}
	return n
}

// NormL2 returns L2 norm (Euclidean distance). i.e. `\sqrt{\sum_{i=0}^n{x_i^2}}`.
func (v *V) NormL2() float64 {
	if v.IsNaV() {
		panic(ErrNaV)
	}
	var n float64
	for _, e := range v.Data {
		n += e * e
	}
	return math.Sqrt(n)
}

// Unit returns the unit vector.
func (v *V) Unit() Vector {
	if v.IsNaV() {
		panic(ErrNaV)
	}
	norm := v.NormL2()
	rv := NewV(v.Dim())
	for i, e := range v.Data {
		rv.Data[i] = e / norm
	}
	return rv
}

// Scale scales the vector by `f`.
func (v *V) Scale(f float64) Vector {
	if IsNaN(f) {
		panic(ErrNaN)
	}
	for i, e := range v.Data {
		v.Data[i] = f * e
	}
	return v
}

// IsZero reports whether the vector is zero vector, the every element is zero.
func (v *V) IsZero() bool {
	if v.IsNaV() {
		panic(ErrNaV)
	}
	for _, e := range v.Data {
		if !IsNEqual(e, 0) {
			return false
		}
	}
	return true
}

// Max returns the maximum value `f` and the first found index `idx` in the vector.
// `NaN, -1` will be returned if the vector is not normal.
func (v *V) Max() (f float64, idx int) {
	if v.IsNaV() {
		panic(ErrNaV)
	}
	f = NaN()
	idx = -1
	for i, e := range v.Data {
		if e > f || IsNaN(f) { /*NaN与任何float64比较大小的结果是false*/
			f = e
			idx = i
		}
	}
	if IsNaN(f) {
		idx = -1
		return
	}
	return f, idx
}

// Min returns the minimum value `f` and the first found index `idx` in the vector `v`.
// `NaN, -1` will be returned if the vector is not normal.
func (v *V) Min() (f float64, idx int) {
	if v.IsNaV() {
		panic(ErrNaV)
	}
	f = NaN()
	idx = -1
	for i, e := range v.Data {
		if e < f || IsNaN(f) {
			f = e
			idx = i
		}
	}
	if IsNaN(f) {
		idx = -1
		return
	}
	return f, idx
}

// MaxAbs returns the maximum absolute value `f` and the first found index `idx` in the vector `v`.
// It is also known as special case of infinity norm, i.e. `\max(|x_1|,|x_2|,...,|x_n|)`.
// `NaN, -1` will be returned if the vector is not normal.
func (v *V) MaxAbs() (f float64, idx int) {
	if v.IsNaV() {
		panic(ErrNaV)
	}
	f = NaN()
	idx = -1
	for i, e := range v.Data {
		abs := math.Abs(e)
		if abs > f || IsNaN(f) {
			f = abs
			idx = i
		}
	}
	if IsNaN(f) {
		idx = -1
		return
	}
	return f, idx
}

// MinAbs returns the minimum absolute value `f` and the first found index `idx` in the vector `v`.
// i.e. `\min(|x_1|,|x_2|,...,|x_n|)`.
// `NaN, -1` will be returned if the vector is not normal.
func (v *V) MinAbs() (f float64, idx int) {
	if v.IsNaV() {
		panic(ErrNaV)
	}
	f = NaN()
	idx = -1
	for i, e := range v.Data {
		abs := math.Abs(e)
		if abs < f || IsNaN(f) {
			f = abs
			idx = i
		}
	}
	if IsNaN(f) {
		idx = -1
		return
	}
	return f, idx
}

// Add adds the vector `x` to `v`.
func (v *V) Add(x *V) *V {
	if !IsVSameShape(x, v) {
		panic(ErrShape)
	}
	for i, e := range x.Data {
		v.Data[i] += e
	}
	return v
}

// Sub subtracts the vector `x` from `v`.
func (v *V) Sub(x *V) *V {
	if !IsVSameShape(x, v) {
		panic(ErrShape)
	}
	for i, e := range x.Data {
		v.Data[i] -= e
	}
	return v
}

// VCross returns the cross product of `x` cross `y` in three-dimensional space.
func VCross(x, y *V) *V {
	if x.Dim() != 3 || y.Dim() != 3 {
		panic(ErrShape)
	}
	v := NewV(3)
	v.SetAt(0, x.At(1)*y.At(2)-x.At(2)*y.At(1))
	v.SetAt(1, x.At(2)*y.At(0)-x.At(0)*y.At(2))
	v.SetAt(2, x.At(0)*y.At(1)-x.At(1)*y.At(0))
	return v
}
