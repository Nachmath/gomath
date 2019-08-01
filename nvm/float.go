// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42526|42605

package nvm

import (
	"math"
	"strconv"
)

// FloatEpsilon represents difference between 1 and the least value greater than 1 that is representable.
var FloatEpsilon = 1e-8

// IsNEqual reports whether floats `x` and `y` are equal.
// `NaN != NaN`, `(-)Inf != (-)Inf`
func IsNEqual(x, y float64) bool {
	if math.Abs(x-y) <= FloatEpsilon {
		return true
	}
	return false
}

// NaN returns "Not-a-Number" value.
func NaN() float64 {
	return math.NaN()
}

// IsNaN reports whether `f` is a "Not-a-Number" value.
func IsNaN(f float64) bool {
	return math.IsNaN(f)
}

// PosInf returns positive infinity.
func PosInf() float64 {
	return math.Inf(1)
}

// NegInf returns negative infinity.
func NegInf() float64 {
	return math.Inf(-1)
}

// IsInf reports whether `f` is positive or negative infinity.
func IsInf(f float64) bool {
	return math.IsInf(f, 0)
}

// NtoS converts the floating-point number `f` to a string,
// it looks like (-ddd.dddd, no exponent) or (-d.dddde±dd, a decimal exponent, for large exponents).
func NtoS(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

// NtoS2 converts the floating-point number `f` to a string,
// it looks like (-ddd.dddd, no exponent),
// `d` is the number of digits after the decimal point.
func NtoS2(f float64, d int) string {
	return strconv.FormatFloat(f, 'f', d, 64)
}

// NtoSci converts the floating-point number `f` to a scientific notation,
// it looks like (-d.dddde±dd, a decimal exponent).
func NtoSci(f float64) string {
	return strconv.FormatFloat(f, 'e', -1, 64)
}

// NtoSci2 converts the floating-point number `f` to a scientific notation,
// it looks like (-d.dddde±dd, a decimal exponent),
// `d` is the number of digits after the decimal point.
func NtoSci2(f float64, d int) string {
	return strconv.FormatFloat(f, 'e', d, 64)
}

// StoN converts the string `s` to a floating-point number,
// the string can be -ddd.dddd, -d.dddde±dd, Inf, -Inf, NaN.
func StoN(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return NaN()
	}
	return f
}

// NtoICeil returns the integer, toward +Inf.
func NtoICeil(f float64) float64 {
	return math.Ceil(f)
}

// NtoIFloor returns the integer, toward -Inf.
func NtoIFloor(f float64) float64 {
	return math.Floor(f)
}

// NtoI returns the integer, toward zero.
func NtoI(f float64) float64 {
	if f < 0 {
		return NtoICeil(f)
	}
	return NtoIFloor(f)
}

// NtoIRound returns the rounding nearest integer, toward zero.
func NtoIRound(f float64) float64 {
	return math.Round(f)
}
