// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42528

package nvm

import (
	"fmt"
	"testing"
)

func Test_V_NewV(t *testing.T) {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("[recover]", p)
		}
	}()

	var print = func(s string, v *V) {
		fmt.Println(s, "\t", "Dim:", v.Dim(), "IsNaV:", v.IsNaV(), "\t", v)
	}

	var s = []float64{1, 2, 3, 4, 5}
	fmt.Println("s:\t", s)

	// print("1------", NewV(0))
	print("2------", NewV(7))
	print("3------", NewVShare(s[:3]))
	print("4------", NewVCopy(s[:]))
	// print("5------", NewVCopy(nil))

	{
		v := NewV(4)
		for i := 0; i < v.Dim(); i++ {
			if v.At(i) != 0 {
				t.Errorf("Not is zero vector")
			}
		}
	}

}

func Test_V_Dim_At(t *testing.T) {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("[recover]", p)
		}
	}()

	v := NewVCopy([]float64{1.0, 2.0, 3.0})
	fmt.Println("vCopy:\t", v, "dim:", v.Dim())

	for i := 0; i < v.Dim(); i++ {
		fmt.Println("At(", i, ")", v.At(i))
	}
	// fmt.Println("At(", v.Dim(), ")", v.At(v.Dim()))
	fmt.Println("At(", -1, ")", v.At(-1))

	// nav := NewV(0)
	// fmt.Println("nav:\t", nav, "dim:", nav.Dim())

	// fmt.Println("At(", nav.Dim(), ")", nav.At(nav.Dim()))
	// fmt.Println("At(", -1, ")", nav.At(-1))

}

func Test_V_IsNaV(t *testing.T) {

	v := NewVShare([]float64{1})
	fmt.Println("v:\t", v, "dim:", v.Dim())
	fmt.Println("v.IsNaV:", v.IsNaV())

	var v0 V
	fmt.Println("v0:\t", v0, "dim:", v0.Dim())
	fmt.Println("v0.IsNaV:", v0.IsNaV())

	vp := &v0
	fmt.Println("vp:\t", vp, "dim:", vp.Dim())
	fmt.Println("vp.IsNaV:", vp.IsNaV())

	if !(vp.IsNaV() == true) {
		t.Errorf("IsNaV is error")
	}
}

func Test_V_NewVCopy(t *testing.T) {
	var print = func(s string, v *V) {
		fmt.Println(s, "\t", "dim:", v.Dim(), "\t", v)
	}

	v := NewVCopy([]float64{1.0, 2.0, 3.0})
	print("v", v)

	v2 := v
	print("v2", v2)

	v3 := *v
	print("v3", &v3)

	v4 := NewVShare(v.Data)
	print("v4", v4)

	v6 := NewVCopy(v.Data)
	print("v6", v6)

	v.Data[2] = 999
	print("v", v)
	print("v2", v2)
	print("v3", &v3)
	print("v4", v4)

	print("v6", v6)

	if !(v6.At(2) == 3) {
		t.Errorf("NewVCopy is error")
	}

}

func Test_V_Norm(t *testing.T) {
	var print = func(s string, v *V) {
		fmt.Println(s, "\t", "dim:", v.Dim(), "\t", v)
		fmt.Println("NormL1=", v.NormL1())
		fmt.Println("NormL2=", v.NormL2())
	}

	print("ZeroV", NewV(4))
	print("V", NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	print("V2", NewVShare([]float64{0.2, NegInf(), -32, 9.772, 3.98}))

	if !(IsNEqual(NewVShare([]float64{0, 0, 0.001}).Unit().NormL2(), 1) == true) {
		t.Errorf("L2Norm is error")
	}
}

func Test_V_Unit(t *testing.T) {
	var print = func(s string, v *V) {
		fmt.Println(s, "\t", "dim:", v.Dim(), "\t", v)
		fmt.Println("IsZero=", v.IsZero())
		fmt.Println("NormL2=", v.NormL2())
		unit := v.Unit()
		vv := unit.(*V)
		fmt.Println("Unit", "\t\t", "dim:", unit.Dim(), "\t", unit)
		fmt.Println("NormL2=", vv.NormL2())
		fmt.Println("=1", IsNEqual(vv.NormL2(), 1))

	}

	print("1------", NewV(1))
	print("2------", NewV(4))
	print("3------", NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	print("4------", NewVShare([]float64{0.2, NegInf(), -32, 9.772, 3.98}))
	print("5------", NewVShare([]float64{1, 1, 1}))
	print("6------", NewVShare([]float64{0, 0, 0.001}))

	if !(IsNEqual(NewVShare([]float64{0, 0, 0.001}).Unit().NormL2(), 1) == true) {
		t.Errorf("Unit is error")
	}
}

func Test_V_Scale(t *testing.T) {
	var print = func(s string, v *V, f float64) {
		fmt.Println(s, "\t", "dim:", v.Dim(), "\t", v)
		fmt.Println("f=", f)
		result := VScale(f, v)
		fmt.Println("VScale", "\t\t", "dim:", result.Dim(), "\t", result)
	}

	print("1------", NewV(1), 1)
	print("2------", NewV(4), 22)
	print("3------", NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}), 3)
	print("4------", NewVShare([]float64{0.2, NegInf(), -32, 9.772, 3.98}), 0.5)

	var print2 = func(s string, v *V, f float64) {
		fmt.Println(s, "\t", "dim:", v.Dim(), "\t", v)
		fmt.Println("f=", f)
		result := v.Scale(f)
		fmt.Println("Scale", "\t\t", "dim:", result.Dim(), "\t", result)
	}

	print2("5------", NewV(1), 1)
	// print2("6------", NewV(4), NaN())
	print2("7------", NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}), 3)
	print2("8------", NewVShare([]float64{0.2, NegInf(), -32, 9.772, 3.98}), 4)

	if !(VScale(3, NewVShare([]float64{1, 2, 3, 4, 5})).At(3) == 12) {
		t.Errorf("VScale is error")
	}
}

func Test_V_Add(t *testing.T) {
	var print = func(s string, x, y *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("", "\t\t", "dim:", y.Dim(), "\t", y)
		result := VAdd(x, y)
		fmt.Println("result", "\t\t", "dim:", result.Dim(), "\t", result)
	}

	print("1------", NewV(1), NewV(1))
	// print("2------", NewV(4), NewV(0))
	print("3------", NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print("4------", NewVShare([]float64{0.2, NegInf(), 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print("5------", NewV(0), NewV(4))

	var print2 = func(s string, x, y *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("", "\t\t", "dim:", y.Dim(), "\t", y)
		result := x.Add(y)
		fmt.Println("result", "\t\t", "dim:", result.Dim(), "\t", result)
	}

	print2("6------", NewV(1), NewV(1))
	// print2("7------", NewV(4), NewV(0))
	print2("8------", NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print2("9------", NewVShare([]float64{0.2, NegInf(), 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print2("10------", NewV(0), NewV(4))

	if !(VAdd(NewVShare([]float64{3}), NewVShare([]float64{2})).At(0) == 5) {
		t.Errorf("VAdd is error")
	}
	if !(NewVShare([]float64{3}).Add(NewVShare([]float64{2})).At(0) == 5) {
		t.Errorf("Add is error")
	}
}

func Test_V_Sub(t *testing.T) {
	var print = func(s string, x, y *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("", "\t\t", "dim:", y.Dim(), "\t", y)
		result := VSub(x, y)
		fmt.Println("result", "\t\t", "dim:", result.Dim(), "\t", result)
	}

	print("1------", NewV(1), NewV(1))
	// print("2------", NewV(4), NewV(1))
	print("3------", NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print("4------", NewVShare([]float64{0.2, NegInf(), 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print("5------", NewV(1), NewV(4))

	var print2 = func(s string, x, y *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("", "\t\t", "dim:", y.Dim(), "\t", y)
		result := x.Sub(y)
		fmt.Println("result", "\t\t", "dim:", result.Dim(), "\t", result)
	}

	print2("6------", NewV(1), NewV(1))
	// print2("7------", NewV(4), NewV(1))
	print2("8------", NewVShare([]float64{2, 4, 6, 8}), NewVShare([]float64{1, 1, 1, 1}))
	// print2("9------", NewVShare([]float64{0.2, NegInf(), 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print2("10------", NewV(1), NewV(4))

	if !(VSub(NewVShare([]float64{3}), NewVShare([]float64{2})).At(0) == 1) {
		t.Errorf("VSub is error")
	}
	if !(NewVShare([]float64{3}).Sub(NewVShare([]float64{2})).At(0) == 1) {
		t.Errorf("Sub is error")
	}

}

func Test_V_Dot(t *testing.T) {
	var print = func(s string, x, y *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("", "\t\t", "dim:", y.Dim(), "\t", y)
		result := VDot(x, y)
		fmt.Println("VDot:", result)
	}

	print("1------", NewV(1), NewV(1))
	// print("2------", NewV(4), NewV(1))
	print("3------", NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print("4------", NewVShare([]float64{0.2, NegInf(), 9.772, 3.98}), NewVShare([]float64{0.2, 0.5, -32, 9.772, 3.98}))
	// print("5------", NewV(1), NewV(4))
	print("6------", NewVShare([]float64{1, 2, 3, 4, 5}), NewVShare([]float64{4, 7, 3, 8, 2}))

	if !(VDot(NewVShare([]float64{1, 2, 3, 4, 5}), NewVShare([]float64{4, 7, 3, 8, 2})) == 69) {
		t.Errorf("VDot is error")
	}
}

func Test_V_SetAt(t *testing.T) {
	var print = func(s string, x *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		x.SetAt(0, 9999)
		fmt.Println("\t", "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("result", x.SetAt(5, 8888))
	}

	// print("1------", NewV(1))
	// print("2------", NewV(4))
	print("3------", NewVShare([]float64{1, 2, 3, 4, 5, 6}))
	print("4------", NewVShare([]float64{1, 2, 3, 4, 5, 6, 7}))

	if !(NewVShare([]float64{1, 2, 3, 4, 5}).SetAt(1, 33).At(1) == 33) {
		t.Errorf("SetAt is error")
	}
}

func Test_V_Cross(t *testing.T) {
	var print = func(s string, x, y *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("", "\t\t", "dim:", y.Dim(), "\t", y)
		result := VCross(x, y)
		fmt.Println("result", "\t\t", "dim:", result.Dim(), "\t", result)
	}

	// print("1------", NewV(1), NewV(1))
	// print("2------", NewV(4), NewV(1))
	// print("3------", NewVShare([]float64{0.2, 0.5, -32, 9.772}), NewVShare([]float64{0.2, 0.5, -32, 9.772}))
	print("4------", NewVShare([]float64{0.2, NegInf(), 9.772}), NewVShare([]float64{0.2, 0.5, -32}))
	print("5------", NewVShare([]float64{PosInf(), 0, 0}), NewVShare([]float64{0, PosInf(), 0}))
	print("6------", NewVShare([]float64{-2, 2, 2}), NewVShare([]float64{2, 1, -3}))
	print("7------", NewVShare([]float64{1, 0, 0}), NewVShare([]float64{0, 1, 0}))
	print("8------", NewVShare([]float64{1, 1, 1}), NewVShare([]float64{0, 0, 0}))
	print("9------", NewV(3), NewV(3))

	v := VCross(NewVShare([]float64{-2, 2, 2}), NewVShare([]float64{2, 1, -3}))
	if !(IsNEqual(v.At(0), -8) && IsNEqual(v.At(1), -2) && IsNEqual(v.At(2), -6)) {
		t.Errorf("VCross is error!")
	}
}

func Test_V_IsZero(t *testing.T) {
	var print = func(s string, x *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("IsZeroV", x.IsZero())
	}

	print("1------", NewV(1))
	print("2------", NewV(4))
	print("3------", NewVShare([]float64{1, 2, 3, 4, 5, 6}))

	if !(NewV(3).IsZero() == true) {
		t.Errorf("IsZeroV is error")
	}
}

func Test_V_IsVSameShape(t *testing.T) {
	var print = func(s string, x, y *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("", "\t\t", "dim:", y.Dim(), "\t", y)
		fmt.Println("IsVSameShape:", IsVSameShape(x, y))
	}

	print("1------", NewV(1), NewV(1))
	print("2------", NewV(3), NewV(3))
	print("3------", NewV(4), NewV(1))
	print("4------", NewV(1), NewV(4))
	print("5------", NewVShare([]float64{0.2, 0.5, -32, 9.772}), NewVShare([]float64{0.2, 0.5, -32, 9.772}))
	print("6------", NewVShare([]float64{1, NaN(), 3}), NewVShare([]float64{1, NaN(), 3}))
	print("7------", NewVShare([]float64{-2, 2, 2}), NewVShare([]float64{2, 1, -3}))
	print("8------", NewVShare([]float64{PosInf(), 0, 0}), NewVShare([]float64{PosInf(), 0, 0}))

	if !(IsVSameShape(NewV(3), NewV(3)) == true) {
		t.Errorf("IsVSameShape is error")
	}
}

func Test_V_Max_Min_Abs(t *testing.T) {
	var print0 = func(x, y float64) {
		fmt.Println("------")
		fmt.Println(x, "<", y, x < y)
		fmt.Println(x, ">", y, x > y)
		fmt.Println(x, "=", y, x == y)
		fmt.Println(x, "equal", y, IsNEqual(x, y))
	}

	print0(NaN(), NaN())
	print0(NaN(), 0)
	print0(PosInf(), 0)
	print0(NegInf(), 0)
	print0(PosInf(), NaN())
	print0(NegInf(), NegInf())
	print0(PosInf(), PosInf())

	var print = func(s string, x *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		Max, index := x.Max()
		fmt.Println("Max:", Max, "index:", index)
		Min, index := x.Min()
		fmt.Println("Min:", Min, "index:", index)
		MaxAbs, index := x.MaxAbs()
		fmt.Println("MaxAbs:", MaxAbs, "index:", index)
		MinAbs, index := x.MinAbs()
		fmt.Println("MinAbs:", MinAbs, "index:", index)
	}

	print("1------", NewV(1))
	print("2------", NewV(4))
	print("3------", NewVShare([]float64{1, 2, 3, 4, 5, 6}))
	print("4------", NewVShare([]float64{NaN(), NegInf(), PosInf()}))
	print("5------", NewVShare([]float64{NaN(), 6}))
	print("5.1----", NewVShare([]float64{6, NaN()}))
	print("5.2----", NewVShare([]float64{NaN()}))
	print("5.3----", NewVShare([]float64{NaN(), NaN(), NaN()}))
	print("6------", NewVShare([]float64{PosInf(), 6}))
	print("7------", NewVShare([]float64{NegInf()}))
	print("7.2----", NewVShare([]float64{NegInf(), NegInf()}))
	print("8------", NewVShare([]float64{-8888, 222, 333}))

	f, i := NewVShare([]float64{222, 333, -8888}).MaxAbs()
	if !(f == 8888 && i == 2) {
		t.Errorf("MaxAbs is error")
	}
}

func Test_V_NormL0(t *testing.T) {
	var print = func(s string, x *V) {
		fmt.Println(s, "\t", "dim:", x.Dim(), "\t", x)
		fmt.Println("NormL0:", x.NormL0())
	}

	print("1------", NewV(1))
	print("2------", NewV(4))
	print("3------", NewVShare([]float64{1, 2, 0.0000000000000001, 4, 0, 6}))
	print("4------", NewVShare([]float64{NaN(), NegInf(), PosInf()}))
	print("5------", NewVShare([]float64{NaN(), 6}))
	print("5.1----", NewVShare([]float64{6, NaN()}))
	print("5.2----", NewVShare([]float64{NaN()}))
	print("5.3----", NewVShare([]float64{NaN(), NaN(), NaN()}))
	print("6------", NewVShare([]float64{PosInf(), 6}))
	print("7------", NewVShare([]float64{NegInf()}))
	print("7.2----", NewVShare([]float64{NegInf(), NegInf()}))

	if !(NewVShare([]float64{1, 2, 0.0000000000000001, 4, 0, 6}).NormL0() == 4) {
		t.Errorf("L0Norm is error")
	}
}

func Test_V_Vector(t *testing.T) {

	v := NewVShare([]float64{1, 2, 3, 4, 5, 6})
	fmt.Println("v", v)

	var vector Vector = v
	fmt.Println("vector", vector)
	fmt.Println(vector.Dim())

	fmt.Println(vector.(*V))

	var _ Vector = NewV(1)

}
