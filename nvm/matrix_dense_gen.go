// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42626|42629

package nvm

import "math/rand"

// GenMRand generates dense matrix,the elements are rand number in [0.0,1.0).
func GenMRand(r, c int) *M {
	data := make([]float64, r*c)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Float64()
	}
	return NewMShare(r, c, data)
}

// GenMRandSquare generates dense square matrix, the elements are rand number in [0.0,1.0).
func GenMRandSquare(rank int) *M {
	return GenMRand(rank, rank)
}

// GenMDiagonal generates dense diagonal matrix, the elements on the diagonal are `f`.
func GenMDiagonal(rank int, f float64) *M {
	data := make([]float64, rank*rank)
	m := NewMShare(rank, rank, data)
	for i := 0; i < rank; i++ {
		m.SetAt(i, i, f)
	}
	return m
}

// GenMUnit generates dense unit matrix, the elements on the diagonal are 1.
func GenMUnit(rank int) *M {
	return GenMDiagonal(rank, 1)
}
