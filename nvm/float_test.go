// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42527

package nvm

import (
	"fmt"
	"math"
	"testing"
)

func Test_N_NaN_Inf(t *testing.T) {
	var print = func(s string, f float64) {
		fmt.Printf("%s\t%f\n", s, f)
		fmt.Println("IsNaN", IsNaN(f))
		fmt.Println("IsInf", IsInf(f))
		if IsInf(f) {
			fmt.Println(">0", f > 0)
			fmt.Println("<0", f < 0)
		}
	}

	print("1------", 2.1)
	print("2------", PosInf())
	print("3------", NegInf())
	print("4------", NaN())

	f0 := 0.0
	f1 := 1.0202004
	f2 := PosInf()
	f3 := NegInf()
	f4 := NaN()

	if !(math.IsNaN(NaN()) == true) {
		t.Errorf("NaN() error")
	}
	if !(IsNaN(f0) == false) {
		t.Errorf("IsNaN(%f) error", f0)
	}
	if !(IsNaN(f1) == false) {
		t.Errorf("IsNaN(%f) error", f1)
	}
	if !(IsNaN(f2) == false) {
		t.Errorf("IsNaN(%f) error", f2)
	}
	if !(IsNaN(f3) == false) {
		t.Errorf("IsNaN(%f) error", f3)
	}
	if !(IsNaN(f4) == true) {
		t.Errorf("IsNaN(%f) error", f4)
	}
	if !(IsInf(f2) == true) || !(f2 > 0) {
		t.Errorf("IsInf(%f) error", f2)
	}
	if !(IsInf(f3) == true) || !(f3 < 0) {
		t.Errorf("IsInf(%f) error", f3)
	}
}

func Test_N_StoN(t *testing.T) {
	var tests = []struct {
		s    string
		want float64
	}{
		{"12345.678908888", 12345.678908888},
		{"0.0000000002", 0.0000000002},
		{"1000000.0000000002", 1000000.0000000002},
		{"NaN", math.NaN()}, {"nan", math.NaN()}, {"NAN", math.NaN()},
		{"+Inf", math.Inf(1)}, {"-Inf", math.Inf(-1)}, {"Inf", math.Inf(1)}, {"INF", math.Inf(1)}, {"inf", math.Inf(1)},
		{"0", 0.0},
		{"-23.3e-2", -0.233},
		{"0.0004E200", 4E+196},
		{"xinb", math.NaN()},
	}
	for _, test := range tests {
		got := StoN(test.s)
		fmt.Println(test.s, "->", got)
		if math.IsNaN(got) {
			if !math.IsNaN(test.want) {
				t.Errorf("StoN(\"%s\")\n want:[%v]\n got:[%v]", test.s, test.want, got)
			}
		} else {
			if got != test.want {
				t.Errorf("StoN(\"%s\")\n want:[%v]\n got:[%v]", test.s, test.want, got)
			}
		}
	}
}

func Test_N_NtoS_NtoSci(t *testing.T) {
	fs := []float64{
		12345.678908888,
		0.0000000002,
		1000000.0000000002,
		NaN(),
		PosInf(),
		NegInf(),
	}
	for _, f := range fs {
		fmt.Println("========")
		fmt.Println(f, NtoS(f))
		fmt.Println("----")
		fmt.Println(f, NtoS2(f, -1))
		fmt.Println(f, NtoS2(f, 0))
		fmt.Println(f, NtoS2(f, 1))
		fmt.Println(f, NtoS2(f, 5))
		fmt.Println(f, NtoS2(f, 12))
		fmt.Println("----")
		fmt.Println(f, NtoSci(f))
		fmt.Println(f, NtoSci2(f, -1))
		fmt.Println(f, NtoSci2(f, 0))
		fmt.Println(f, NtoSci2(f, 1))
		fmt.Println(f, NtoSci2(f, 5))
		fmt.Println(f, NtoSci2(f, 15))
	}

	/*--------------------------*/

	var f float64
	var d int
	var s string
	var ws string

	f = 12345.678908888
	s = NtoS(f)
	ws = "12345.678908888"
	if !(s == ws) {
		t.Errorf("NtoS(%f) = %s , %s", f, s, ws)
	}

	f = 0.0000000002
	s = NtoS(f)
	ws = "2e-10"
	if !(s == ws) {
		t.Errorf("NtoS(%f) = %s , %s", f, s, ws)
	}

	f = -0.233
	s = NtoSci(f)
	ws = "-2.33e-01"
	if !(s == ws) {
		t.Errorf("NtoSci(%f) = %s , %s", f, s, ws)
	}

	f = NaN()
	s = NtoS(f)
	ws = "NaN"
	if !(s == ws) {
		t.Errorf("NtoS(%f) = %s , %s", f, s, ws)
	}

	f = NegInf()
	s = NtoSci(f)
	ws = "-Inf"
	if !(s == ws) {
		t.Errorf("NtoSci(%f) = %s , %s", f, s, ws)
	}

	f = PosInf()
	d = 3
	s = NtoS2(f, d)
	ws = "+Inf"
	if !(s == ws) {
		t.Errorf("NtoS2(%f, %d) = %s , %s", f, d, s, ws)
	}

	f = 12345.678908888
	d = -1
	s = NtoS2(f, d)
	ws = "12345.678908888"
	if !(s == ws) {
		t.Errorf("NtoS2(%f, %d) = %s , %s", f, d, s, ws)
	}

	f = 1000000.0000000007
	d = 9
	s = NtoS2(f, d)
	ws = "1000000.000000001"
	if !(s == ws) {
		t.Errorf("NtoS2(%f, %d) = %s , %s", f, d, s, ws)
	}

	f = 0
	d = 4
	s = NtoS2(f, d)
	ws = "0.0000"
	if !(s == ws) {
		t.Errorf("NtoS2(%f, %d) = %s , %s", f, d, s, ws)
	}

	f = 4E+196
	d = 2
	s = NtoSci2(f, d)
	ws = "4.00e+196"
	if !(s == ws) {
		t.Errorf("NtoSci2(%f, %d) = %s , %s", f, d, s, ws)
	}

}

func Test_N_NtoICeil(t *testing.T) {
	var datas = []struct {
		f  float64
		wf float64
	}{
		{NaN(), NaN()},
		{PosInf(), PosInf()},
		{NegInf(), NegInf()},
		{1.2, 2},
		{-1.2, -1},
		{0, 0},
		{0.8, 1},
		{-0.8, 0},
	}
	for _, data := range datas {
		gf := NtoICeil(data.f)
		fmt.Println(data.f, "->", gf)
		fmt := "NtoICeil(%f) = %f | %f"
		if IsNaN(gf) {
			if !IsNaN(data.wf) {
				t.Errorf(fmt, data.f, gf, data.wf)
			}
		} else {
			if !(gf == data.wf) {
				t.Errorf(fmt, data.f, gf, data.wf)
			}
		}
	}
}
func Test_N_NtoIFloor(t *testing.T) {
	var datas = []struct {
		f  float64
		wf float64
	}{
		{NaN(), NaN()},
		{PosInf(), PosInf()},
		{NegInf(), NegInf()},
		{1.2, 1},
		{-1.2, -2},
		{0, 0},
		{0.8, 0},
		{-0.8, -1},
	}
	for _, data := range datas {
		gf := NtoIFloor(data.f)
		fmt.Println(data.f, "->", gf)
		fmt := "NtoIFloor(%f) = %f | %f"
		if IsNaN(gf) {
			if !IsNaN(data.wf) {
				t.Errorf(fmt, data.f, gf, data.wf)
			}
		} else {
			if !(gf == data.wf) {
				t.Errorf(fmt, data.f, gf, data.wf)
			}
		}
	}
}
func Test_N_NtoI(t *testing.T) {
	var datas = []struct {
		f  float64
		wf float64
	}{
		{NaN(), NaN()},
		{PosInf(), PosInf()},
		{NegInf(), NegInf()},
		{1.2, 1},
		{-1.2, -1},
		{0, 0},
		{0.8, 0},
		{-0.8, 0},
	}
	for _, data := range datas {
		gf := NtoI(data.f)
		fmt.Println(data.f, "->", gf)
		fmt := "NtoI(%f) = %f | %f"
		if IsNaN(gf) {
			if !IsNaN(data.wf) {
				t.Errorf(fmt, data.f, gf, data.wf)
			}
		} else {
			if !(gf == data.wf) {
				t.Errorf(fmt, data.f, gf, data.wf)
			}
		}
	}
}
func Test_N_NtoIRound(t *testing.T) {
	var datas = []struct {
		f  float64
		wf float64
	}{
		{NaN(), NaN()},
		{PosInf(), PosInf()},
		{NegInf(), NegInf()},
		{1.2, 1},
		{-1.2, -1},
		{0, 0},
		{0.8, 1},
		{-0.8, -1},
		{0.4, 0},
		{0.5, 1},
		{0.49999999, 0},
		{-0.49999999, 0},
		{-9999.5, -10000},
		{-9999.499999, -9999},
	}
	for _, data := range datas {
		gf := NtoIRound(data.f)
		fmt.Println(data.f, "->", gf)
		fmt := "NtoIRound(%f) = %f | %f"
		if IsNaN(gf) {
			if !IsNaN(data.wf) {
				t.Errorf(fmt, data.f, gf, data.wf)
			}
		} else {
			if !(gf == data.wf) {
				t.Errorf(fmt, data.f, gf, data.wf)
			}
		}
	}
}

func Test_N_IsNEqual(t *testing.T) {
	var datas = []struct {
		x, y float64
		w    bool
	}{
		{NaN(), NaN(), false},
		{PosInf(), PosInf(), false},
		{NegInf(), NegInf(), false},
		{1.1, 1, false},
		{1e-4, 1e-5, false},
		{0, 0, true},
		{1e-6, 1e-5, false},
		{-1e-8, -1e-8, true},
	}
	for _, data := range datas {
		g := IsNEqual(data.x, data.y)
		fmt.Println(data.x, data.y, "->", g)
		fmt := "IsNEqual(%f, %f) = %t | %t"
		if !(g == data.w) {
			t.Errorf(fmt, data.x, data.y, g, data.w)
		}
	}
}
