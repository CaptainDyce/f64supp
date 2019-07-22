package f64supp

import (
	"fmt"
	is "github.com/CaptainDyce/intsupp"
	"math"
)

///////////////////////////
// interface...
///////////////////////////

type IndexedFunc func(index int) float64
type Operator func(value float64) float64
type IndexedOperator func(index int, value float64) float64
type Consumer func(value float64)
type IndexedConsumer func(index int, value float64)
type ReduceOperator func(vleft float64, vright float64) float64

///////////////////////////
// building...
///////////////////////////

// -> f(i) = c
func Constant(constant float64) IndexedFunc {
	return func(i int) float64 {
		return constant
	}
}

func CoerceInt(i int) float64 {
	return float64(i)
}

func CoerceInts(values []int) IndexedFunc {
	return func(i int) float64 {
		return float64(values[i])
	}
}

// -> f(i) = values[i]
func Get(values []float64) IndexedFunc {
	return func(i int) float64 {
		return values[i]
	}
}

func Plus(vleft float64, vright float64) float64 {
	return vleft + vright
}

func Minus(vleft float64, vright float64) float64 {
	return vleft - vright
}

func Times(vleft float64, vright float64) float64 {
	return vleft * vright
}

func Div(vleft float64, vright float64) float64 {
	return vleft / vright
}

func Neg(value float64) float64 {
	return -value
}

///////////////////////////
// operations...
///////////////////////////

func accept(v []float64, v1 []float64) {
	if len(v1) < len(v) {
		panic(fmt.Sprintf("invalid array size %d (out of bounds for %d-element vector)", len(v1), len(v)))
	}
}

func Apply(s []float64, f IndexedFunc) []float64 {
	for i, _ := range s {
		s[i] = f(i)
	}
	return s
}

func ApplyOp(s []float64, f Operator) []float64 {
	for i, val := range s {
		s[i] = f(val)
	}
	return s
}

func ApplyOpi(s []float64, f IndexedOperator) []float64 {
	for i, val := range s {
		s[i] = f(i, val)
	}
	return s
}

func Ident(s []float64) []float64 {
	return Apply(s, CoerceInt)
}

func Setl(s []float64, value float64) []float64 {
	for i, _ := range s {
		s[i] = value
	}
	return s
}

func Setv(s []float64, v []float64) []float64 {
	accept(s, v)
	for i, _ := range s {
		s[i] = v[i]
	}
	return s
}

func SetMaskl(s []float64, value float64, p is.Predicate) []float64 {
	for i, _ := range s {
		if p(i) {
			s[i] = value
		}
	}
	return s
}

func SetMaskv(s []float64, v []float64, p is.Predicate) []float64 {
	accept(s, v)
	for i, _ := range s {
		if p(i) {
			s[i] = v[i]
		}
	}
	return s
}

/////////////////
// Plus
/////////////////
// -> s'[i] = s[i] + v[i]
func Plusv(s []float64, v1 []float64) []float64 {
	accept(s, v1)
	for i, _ := range s {
		s[i] += v1[i]
	}
	return s
}

// -> s'[i] = s[i] + value
func Plusl(s []float64, value float64) []float64 {
	for i, _ := range s {
		s[i] += value
	}
	return s
}

// -> s'[i] = s[i] + o(i)
func PlusOp(s []float64, o IndexedFunc) []float64 {
	for i, _ := range s {
		s[i] += o(i)
	}
	return s
}

// -> s'[i] = s[i] + o(i, s[i])
func PlusOpi(s []float64, o IndexedOperator) []float64 {
	for i, val := range s {
		s[i] += o(i, val)
	}
	return s
}

/////////////////
// Minus
/////////////////
// -> s'[i] = s[i] - v[i]
func Minusv(s []float64, v1 []float64) []float64 {
	accept(s, v1)
	for i, _ := range s {
		s[i] -= v1[i]
	}
	return s
}

// -> s'[i] = s[i] - value
func Minusl(s []float64, value float64) []float64 {
	for i, _ := range s {
		s[i] -= value
	}
	return s
}

// -> s'[i] = s[i] - o(i)
func MinusOp(s []float64, o IndexedFunc) []float64 {
	for i, _ := range s {
		s[i] -= o(i)
	}
	return s
}

// -> s'[i] = s[i] - o(i, s[i])
func MinusOpi(s []float64, o IndexedOperator) []float64 {
	for i, val := range s {
		s[i] -= o(i, val)
	}
	return s
}

/////////////////
// Times
/////////////////
// -> s'[i] = s[i] * v[i]
func Timesv(s []float64, v1 []float64) []float64 {
	accept(s, v1)
	for i, _ := range s {
		s[i] *= v1[i]
	}
	return s
}

// -> s'[i] = s[i] * value
func Timesl(s []float64, value float64) []float64 {
	for i, _ := range s {
		s[i] *= value
	}
	return s
}

// -> s'[i] = s[i] * o(i)
func TimesOp(s []float64, o IndexedFunc) []float64 {
	for i, _ := range s {
		s[i] *= o(i)
	}
	return s
}

// -> s'[i] = s[i] * o(i, s[i])
func TimesOpi(s []float64, o IndexedOperator) []float64 {
	for i, val := range s {
		s[i] *= o(i, val)
	}
	return s
}

/////////////////
// Div
/////////////////
// -> s'[i] = s[i] / v[i]
func Divv(s []float64, v1 []float64) []float64 {
	accept(s, v1)
	for i, _ := range s {
		s[i] /= v1[i]
	}
	return s
}

// -> s'[i] = s[i] / value
func Divl(s []float64, value float64) []float64 {
	for i, _ := range s {
		s[i] /= value
	}
	return s
}

// -> s'[i] = s[i] / o(i)
func DivOp(s []float64, o IndexedFunc) []float64 {
	for i, _ := range s {
		s[i] /= o(i)
	}
	return s
}

// -> s'[i] = s[i] / o(i, s[i])
func DivOpi(s []float64, o IndexedOperator) []float64 {
	for i, val := range s {
		s[i] /= o(i, val)
	}
	return s
}

/////////////////
// Pow
/////////////////
// -> s'[i] = s[i] ^ v[i] (as in e.g. 2^3...)
func Powv(s []float64, v1 []float64) []float64 {
	accept(s, v1)
	for i, val := range s {
		s[i] = math.Pow(val, v1[i])
	}
	return s
}

// -> s'[i] = s[i] ^ value
func Powl(s []float64, value float64) []float64 {
	for i, val := range s {
		s[i] = math.Pow(val, value)
	}
	return s
}

// -> s'[i] = s[i] ^ o(i)
func PowOp(s []float64, o IndexedFunc) []float64 {
	for i, val := range s {
		s[i] = math.Pow(val, o(i))
	}
	return s
}

// -> s'[i] = s[i] ^ o(i, s[i])
func PowOpi(s []float64, o IndexedOperator) []float64 {
	for i, val := range s {
		s[i] = math.Pow(val, o(i, val))
	}
	return s
}

/////////////////
// Max
/////////////////
// -> s'[i] = max(s[i], v[i])
func Maxv(s []float64, v1 []float64) []float64 {
	accept(s, v1)
	for i, val := range s {
		s[i] = math.Max(val, v1[i])
	}
	return s
}

// -> s'[i] = max(s[i], value)
func Maxl(s []float64, value float64) []float64 {
	for i, val := range s {
		s[i] = math.Max(val, value)
	}
	return s
}

// -> s'[i] = max(s[i], o(i))
func MaxOp(s []float64, o IndexedFunc) []float64 {
	for i, val := range s {
		s[i] = math.Max(val, o(i))
	}
	return s
}

// -> s'[i] = max(s[i], o(i, s[i]))
func MaxOpi(s []float64, o IndexedOperator) []float64 {
	for i, val := range s {
		s[i] = math.Max(val, o(i, val))
	}
	return s
}

/////////////////
// Min
/////////////////
// -> s'[i] = max(s[i], v[i])
func Minv(s []float64, v1 []float64) []float64 {
	accept(s, v1)
	for i, val := range s {
		s[i] = math.Min(val, v1[i])
	}
	return s
}

// -> s'[i] = max(s[i], value)
func Minl(s []float64, value float64) []float64 {
	for i, val := range s {
		s[i] = math.Min(val, value)
	}
	return s
}

// -> s'[i] = max(s[i], o(i))
func MinOp(s []float64, o IndexedFunc) []float64 {
	for i, val := range s {
		s[i] = math.Min(val, o(i))
	}
	return s
}

// -> s'[i] = max(s[i], o(i, s[i]))
func MinOpi(s []float64, o IndexedOperator) []float64 {
	for i, val := range s {
		s[i] = math.Min(val, o(i, val))
	}
	return s
}

// reverse slice
func Rev(s []float64) []float64 {
	size := len(s)
	for i := 0; i < size/2; i++ {
		tmp := s[i]
		s[i] = s[size-i-1]
		s[size-i-1] = tmp
	}
	return s
}

// -> s'[i] = -s[i]
func Negv(s []float64) []float64 {
	for i, val := range s {
		s[i] = -val
	}
	return s
}

// -> s'[i] = abs(s[i])
func Abs(s []float64) []float64 {
	for i, val := range s {
		s[i] = math.Abs(val)
	}
	return s
}

// -> s'[i] = value / s[i]
func Idivl(s []float64, value float64) []float64 {
	for i, val := range s {
		s[i] = value / val
	}
	return s
}

// -> s'[i] = log(s[i])
func Log(s []float64) []float64 {
	for i, val := range s {
		s[i] = math.Log(val)
	}
	return s
}

// -> s'[i] = exp(s[i])
func Exp(s []float64) []float64 {
	for i, val := range s {
		s[i] = math.Exp(val)
	}
	return s
}

// -> s'[i] = value ^ s[i]
func Expl(s []float64, value float64) []float64 {
	for i, val := range s {
		s[i] = math.Pow(value, val)
	}
	return s
}
