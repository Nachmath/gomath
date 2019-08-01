// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42610

package nvm

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

// fs := make([]float64, 3471000000)
// 16G内存有时导致： fatal error: out of memory allocating heap arena metadata */

func matrixDenseTestPrintM(m Matrix) {
	if m.IsNaM() {
		fmt.Println("NaM")
		return
	}
	r, c := m.Dims()
	fmt.Println("Rows:", r, "\t", "Cols:", c)

	// for i := 0; i < r; i++ {
	// 	for j := 0; j < c; j++ {
	// 		fmt.Printf("%.1f ,\t", m.At(i, j))
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(mat.Formatted(m.(*M).O, mat.Prefix("")))
}

func Test_M_Gen(t *testing.T) {
	fmt.Println("GenMUnit")
	matrixDenseTestPrintM(GenMUnit(3))
	fmt.Println("GenMRandSquare")
	matrixDenseTestPrintM(GenMRandSquare(4))
	fmt.Println("GenMDiagonal")
	matrixDenseTestPrintM(GenMDiagonal(5, 8.2))
	fmt.Println("GenMRand")
	matrixDenseTestPrintM(GenMRand(4, 2))
}

func Test_M_Formula(t *testing.T) {
	m1 := NewMShare(4, 4, []float64{ // Det=266
		0, 7, 3, 5,
		5, 8, 9, 4,
		7, 0, 1, 2,
		3, 9, 8, 6,
	})
	matrixDenseTestPrintM(m1)

	m2 := m1.Inv()
	matrixDenseTestPrintM(m2)

	m3 := Mul(m1, m2)
	matrixDenseTestPrintM(m3)

	r, _ := m1.Dims()
	matrixDenseTestPrintM(GenMUnit(r))

	if !(IsEqual(m3, GenMUnit(r)) == true) {
		t.Fatal()
	}
}

func Test_M_NewM(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, m *M) {
		fmt.Println(s)
		fmt.Println("IsNaM:", m.IsNaM())
		matrixDenseTestPrintM(m)
	}
	var nam M
	print("1------", NewM(3, 2))
	print("2------", &nam)
	print("3------", NewM(2, 6))

	if r, c := NewM(5, 2).Dims(); !(r == 5 && c == 2) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	NewM(0, 2)
}

func Test_M_NewMShare(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, m *M) {
		fmt.Println(s)
		fmt.Println("IsNaM:", m.IsNaM())
		matrixDenseTestPrintM(m)
	}
	print("1------", NewMShare(2, 3, []float64{2.3, 0.4, 23, 5, 0, 1}))
	print("2------", NewMShare(4, 2, []float64{1, 2, 3, 4, 5, 6, 7, 8}))

	fmt.Println("===============")
	f64s := []float64{1, 2, 3, 4, 5, 6}
	fmt.Println("f64s:", f64s)
	mSharef64s := NewMShare(2, 3, f64s)
	print("3------", mSharef64s)
	f64s[3] = 999
	fmt.Println("f64s:", f64s)
	print("3.1----", mSharef64s)
	fmt.Println("===============")

	if !(NewMShare(2, 3, []float64{2.3, 0.4, 23, 5, 0, 1}).At(1, 2) == 1) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	NewMShare(2, 2, []float64{1, 2, 3, 4, 5, 6, 7, 8})
}

func Test_M_NewMCopy(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, m *M) {
		fmt.Println(s)
		fmt.Println("IsNaM:", m.IsNaM())
		matrixDenseTestPrintM(m)
	}
	print("1------", NewMCopy(2, 3, []float64{2.3, 0.4, 23, 5, 0, 1}))
	print("2------", NewMCopy(4, 2, []float64{1, 2, 3, 4, 5, 6, 7, 8}))

	fmt.Println("===============")
	f64s := []float64{1, 2, 3, 4, 5, 6}
	fmt.Println("f64s:", f64s)
	mSharef64s := NewMCopy(2, 3, f64s)
	print("3------", mSharef64s)
	f64s[3] = 999
	fmt.Println("f64s:", f64s)
	print("3.1----", mSharef64s)
	fmt.Println("===============")

	if !(NewMCopy(2, 3, []float64{2.3, 0.4, 23, 5, 0, 1}).At(1, 2) == 1) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	NewMCopy(2, 2, []float64{1, 2, 3, 4, 5, 6, 7, 8})
}

func Test_M_At_SetAt(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, m *M) {
		fmt.Println(s)
		fmt.Println("IsNaM:", m.IsNaM())
		matrixDenseTestPrintM(m)
	}

	m := NewMShare(2, 3, []float64{2.3, 0.4, 5, 678, 0, 1})
	print("1------", m)
	fmt.Println(m.At(0, 0))
	fmt.Println(m.At(1, 1))

	m.SetAt(1, 2, 999)
	print("2------", m)

	if !(m.SetAt(1, 2, 777).At(1, 2) == 777) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	m.SetAt(2, 2, 8888)
}

func Test_M_T(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, m Matrix) {
		fmt.Println(s)
		fmt.Println("IsNaM:", m.IsNaM())
		matrixDenseTestPrintM(m)
	}

	m := NewMShare(1, 1, []float64{777})
	print("1------", m)
	print("2------", m.T())

	m = NewMCopy(2, 3, []float64{2.3, 0.4, 5, 678, 0, 1})
	print("3------", m)
	tm := m.T()
	print("4------", tm)
	print("5------", tm.T())

	fmt.Println("---- goto panic ----")
	var nam M
	nam.T()
}

func Test_M_Det(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, m Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(m)
		fmt.Println("Det:", m.Det())
	}

	print("1------", NewMShare(4, 4, []float64{ // Det=266
		0, 7, 3, 5,
		5, 8, 9, 4,
		7, 0, 1, 2,
		3, 9, 8, 6,
	}))
	print("2------", NewMShare(1, 1, []float64{ // Det=5
		5,
	}))
	print("3------", NewMShare(2, 2, []float64{ // Det=-27
		0, 3,
		9, 6,
	}))
	print("4------", NewMShare(3, 3, []float64{
		1, 0, 0,
		0, 0, 0,
		0, 0, 1,
	}))
	print("5------", NewMShare(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}))

	if !(IsNEqual(NewMShare(3, 3, []float64{
		1, 0, 0,
		0, 0, 0,
		0, 0, 1,
	}).Det(), 0) == true) {
		t.Fatal()
	}

	if !(IsNEqual(NewMCopy(4, 4, []float64{
		0, 7, 3, 5,
		5, 8, 9, 4,
		7, 0, 1, 2,
		3, 9, 8, 6,
	}).Det(), 266) == true) {
		t.Fatal()
	}

}
func Benchmark_M_Det(b *testing.B) {
	fmt.Println("\n============= N", b.N)
	var rank = 300
	var M1 = GenMRandSquare(rank)
	fmt.Println("rank:", rank)
	// matrixDenseTestPrintM(M1)
	// fmt.Println("Det:", M1.Det())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		M1.Det()
	}

	// 3000 = 2280632900 ns/op	72284344 B/op	     126 allocs/op
	//  300 = 3656503 ns/op	  742251 B/op	      18 allocs/op
	//   30 = 22778 ns/op	    8631 B/op	       7 allocs/op
	//    3 = 2452 ns/op	     304 B/op	       7 allocs/op
}

func Test_M_Inv(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, m Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(m)
		fmt.Println("Inv:")
		matrixDenseTestPrintM(m.Inv())
	}

	print("1------", NewMShare(4, 4, []float64{ // Det=266
		0, 7, 3, 5,
		5, 8, 9, 4,
		7, 0, 1, 2,
		3, 9, 8, 6,
	}))
	print("2------", NewMShare(1, 1, []float64{ // Det=5
		5,
	}))
	print("3------", NewMShare(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}))

	if !(IsNEqual(NewMCopy(1, 1, []float64{5}).Inv().At(0, 0), 0.2) == true) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	NewMShare(3, 3, []float64{
		1, 0, 1,
		0, 0, 0,
		1, 0, 1,
	}).Inv()
}
func Benchmark_M_Inv(b *testing.B) {
	fmt.Println("\n============= N", b.N)
	var rank = 3
	var M1 = GenMUnit(rank)
	fmt.Println("rank:", rank)
	// matrixDenseTestPrintM(M1)
	// fmt.Println("Inv.Det:", M1.Inv().Det())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		M1.Inv()
	}

	// 3000 = 556263100 ns/op	74313720 B/op	     225 allocs/op
	//  300 = 2459567 ns/op	  930176 B/op	      26 allocs/op
	//   30 = 33239 ns/op	    8414 B/op	       6 allocs/op
	//    3 = 2803 ns/op	     256 B/op	       6 allocs/op
}

func Test_M_IsSameShape(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("IsSameShape:", IsSameShape(x, y))
	}

	print("1------",
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
	print("2------",
		NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
	print("3------",
		NewMCopy(2, 2, []float64{1, 2, 3, 5}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))

	if !(IsSameShape(
		NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})) == true) {
		t.Fatal()
	}

	var errm1, errm2 M
	if !(IsSameShape(&errm1, &errm2) == false) {
		t.Fatal()
	}
}

func Test_M_IsEqual(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("IsEqual:", IsEqual(x, y))
	}

	print("1------",
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
	print("2------",
		NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
	print("3------",
		NewMCopy(2, 2, []float64{1, 2, 3, 5}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))

	if !(IsEqual(
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})) == true) {
		t.Fatal()
	}
	if !(IsEqual(
		NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})) == false) {
		t.Fatal()
	}
}

func Test_M_Add(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("*M.Add:")
		matrixDenseTestPrintM(x.(*M).Add(x, y))
	}

	print("1------",
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))

	fmt.Println("2------")
	matrixDenseTestPrintM(NewM(2, 3).Add(
		NewMCopy(2, 3, []float64{6, 5, 4, 3, 2, 1}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})))

	if !(NewM(2, 3).Add(
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})).At(0, 0) == 2) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	NewM(2, 4).Add(
		NewMCopy(2, 3, []float64{1, 1, 1, 1, 1, 1}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
	NewM(2, 3).Add(
		NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
}
func Benchmark_M_Add(b *testing.B) {
	fmt.Println("\n============= N", b.N)
	var rows, cols int = 3000, 3000
	M1 := GenMRand(rows, cols)
	M2 := GenMRand(rows, cols)
	fmt.Println("rows:", rows, "cols", cols)
	// matrixDenseTestPrintM(M1)
	// matrixDenseTestPrintM(M2)
	// matrixDenseTestPrintM(M1.Add(M1, M2))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		M1.Add(M1, M2)
	}

	// 3000, 3000 = 11467497 ns/op	       0 B/op	       0 allocs/op
	//   300, 300 = 110161 ns/op	       0 B/op	       0 allocs/op
	//     30, 30 = 1229 ns/op	       0 B/op	       0 allocs/op
	//       3, 3 = 114 ns/op	       0 B/op	       0 allocs/op
}
func Test_M_Add2(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("Add:")
		matrixDenseTestPrintM(Add(x, y))
	}

	print("1------",
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))

	if !(Add(
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})).At(0, 0) == 2) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	Add(NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
}
func Benchmark_M_Add2(b *testing.B) {
	fmt.Println("\n============= N", b.N)
	var rows, cols int = 3, 3
	M1 := GenMRand(rows, cols)
	M2 := GenMRand(rows, cols)
	fmt.Println("rows:", rows, "cols", cols)
	// matrixDenseTestPrintM(M1)
	// matrixDenseTestPrintM(M2)
	// matrixDenseTestPrintM(Add(M1, M2))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add(M1, M2)
	}

	// 3000, 3000 = 18672117 ns/op	72007763 B/op	       3 allocs/op
	//   300, 300 = 216548 ns/op	  720976 B/op	       3 allocs/op
	//     30, 30 = 2890 ns/op	    8272 B/op	       3 allocs/op
	//       3, 3 = 281 ns/op	     160 B/op	       3 allocs/op
}

func Test_M_Sub(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("*M.Sub:")
		matrixDenseTestPrintM(x.(*M).Sub(x, y))
	}

	print("1------",
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))

	fmt.Println("2------")
	matrixDenseTestPrintM(NewM(2, 3).Sub(
		NewMCopy(2, 3, []float64{6, 5, 4, 3, 2, 1}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})))

	if !(NewM(2, 3).Sub(
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})).At(0, 0) == 0) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	NewM(2, 4).Sub(
		NewMCopy(2, 3, []float64{1, 1, 1, 1, 1, 1}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
	NewM(2, 3).Sub(
		NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
}
func Test_M_Sub2(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("Sub:")
		matrixDenseTestPrintM(Sub(x, y))
	}

	print("1------",
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))

	if !(Sub(
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})).At(0, 0) == 0) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	Sub(NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
}

func Test_M_Scale(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, f float64, x Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		fmt.Println("*M.Scale:")
		matrixDenseTestPrintM(x.(*M).Scale(f, x))
	}

	print("1------", 3,
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
	print("2------", 0.5,
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))

	if !(NewM(2, 2).Scale(3, NewMShare(2, 2, []float64{2, 3, 4, 5})).At(0, 0) == 6) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	NewM(2, 4).Scale(2, NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
}
func Test_M_Scale2(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, f float64, x Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		fmt.Println("Scale:")
		matrixDenseTestPrintM(Scale(f, x))
	}

	print("1------", 3,
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
	print("2------", 0.5,
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))

	if !(Scale(3, NewMShare(2, 2, []float64{2, 3, 4, 5})).At(0, 0) == 6) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	Scale(2, NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
}

func Test_M_DotMul(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("*M.DotMul:")
		matrixDenseTestPrintM(x.(*M).DotMul(x, y))
	}

	print("1------",
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))

	if !(NewM(2, 3).DotMul(
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})).At(1, 2) == 36) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	NewM(2, 4).DotMul(
		NewMCopy(2, 3, []float64{1, 1, 1, 1, 1, 1}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
	NewM(2, 3).DotMul(
		NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
}
func Test_M_DotMul2(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("DotMul:")
		matrixDenseTestPrintM(DotMul(x, y))
	}

	print("1------",
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))

	if !(DotMul(
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})).At(1, 2) == 36) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	DotMul(
		NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
}

func Test_M_Mul(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()

	{
		m1 := NewMCopy(2, 2, []float64{1, 2, 3, 4})
		matrixDenseTestPrintM(m1)
		m2 := NewMCopy(2, 2, []float64{1, 2, 3, 4})
		matrixDenseTestPrintM(m2)

		m3 := NewM(2, 2)
		m3.Mul(m1, m2)
		fmt.Println("m3")
		matrixDenseTestPrintM(m3)

		m2.Mul(m1, m2)
		fmt.Println("m2")
		matrixDenseTestPrintM(m2)
	}

	{
		m1 := NewMCopy(3, 2, []float64{1, 2, 3, 4, 5, 6})
		matrixDenseTestPrintM(m1)
		m2 := NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})
		matrixDenseTestPrintM(m2)

		m3 := NewM(3, 3)
		m3.Mul(m1, m2)
		fmt.Println("m3")
		matrixDenseTestPrintM(m3)

		if !(m3.At(1, 1) == 26) {
			t.Fatal()
		}
	}

	fmt.Println("---- goto panic ----")
	NewM(2, 4).Mul(
		NewMCopy(3, 2, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
}
func Test_M_Mul2(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("[recover]", p)
		}
	}()
	var print = func(s string, x, y Matrix) {
		fmt.Println(s)
		matrixDenseTestPrintM(x)
		matrixDenseTestPrintM(y)
		fmt.Println("Mul:")
		matrixDenseTestPrintM(Mul(x, y))
	}

	print("1------",
		NewMCopy(3, 2, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6}))
	print("2------",
		NewMCopy(1, 1, []float64{5}),
		NewMCopy(1, 1, []float64{6}))

	if !(Mul(
		NewMCopy(3, 2, []float64{1, 2, 3, 4, 5, 6}),
		NewMCopy(2, 3, []float64{1, 2, 3, 4, 5, 6})).At(0, 0) == 9) {
		t.Fatal()
	}

	fmt.Println("---- goto panic ----")
	Mul(NewMCopy(2, 3, []float64{1, 2, 0, 4, 5, 6}),
		NewMCopy(2, 2, []float64{1, 2, 3, 4}))
}
