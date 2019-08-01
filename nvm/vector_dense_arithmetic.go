// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42618

package nvm

// dense equal dense
func vectorIsDenseEqual(x, y *V) bool {
	if !IsVSameShape(x, y) {
		return false
	}
	for i, e := range x.Data {
		if !IsNEqual(e, y.Data[i]) {
			return false
		}
	}
	return true
}

// dense scale
func vectorDenseScale(f float64, v *V) *V {
	if IsNaN(f) {
		panic(ErrNaN)
	}
	if v.IsNaV() {
		panic(ErrNaV)
	}
	r := NewV(v.Dim())
	for i, e := range v.Data {
		r.Data[i] = f * e
	}
	return r
}

// dense add dense
func vectorDenseAdd(x, y *V) *V {
	if !IsVSameShape(x, y) {
		panic(ErrShape)
	}
	v := NewV(x.Dim())
	for i, e := range x.Data {
		v.Data[i] = e + y.Data[i]
	}
	return v
}

// dense sub dense
func vectorDenseSub(x, y *V) *V {
	if !IsVSameShape(x, y) {
		panic(ErrShape)
	}
	v := NewV(x.Dim())
	for i, e := range x.Data {
		v.Data[i] = e - y.Data[i]
	}
	return v
}

// dense dot dense
func vectorDenseDot(x, y *V) float64 {
	if !IsVSameShape(x, y) {
		panic(ErrShape)
	}
	var f float64
	for i, e := range x.Data {
		f += e * y.Data[i]
	}
	return f
}
