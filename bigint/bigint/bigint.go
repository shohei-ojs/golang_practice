package bigint


import (
	"strconv"
	"math"
)

type BigInt []uint8

func (b BigInt)String() string{
	var str string
	for _,v := range arrayReverse(b) {
		str += strconv.Itoa(int(v))
	}
	return str
}

func FromString(s string) (bigint BigInt) {
	for _, c := range stringReverse(s) {
		bigint = append(bigint, uint8(c)-48)
	}
	return
}

func New(v []uint8) BigInt {
	return BigInt(v)
}

func Zero() BigInt {
	return BigInt([]uint8{0})
}

func (b BigInt)add(a BigInt) BigInt {
	diff := int(math.Abs(float64(len(a))-float64(len(b))))
	if len(a) > len(b) {
		for i := 0; i < diff; i++ {
			b = append(b, 0)
		} 
	} else if len(a) < len(b) {
		for i := 0; i < diff; i++ {
			a = append(a, 0)
		} 
	}
	bigint := make(BigInt, len(a))
	for i := range b {
		d := b[i]+a[i]
		if d < 10 {
			bigint[i] = d
		} else if i == len(b)-1 {
			bigint[i] = d%10
			bigint = append(bigint, 1)
		} else {
			bigint[i] = d%10
			b[i+1]++
		}
	}
	return bigint
}

func (b BigInt)Mul(a BigInt) BigInt {
	result := make(BigInt, 0)
	
	for i, v := range a {
		result = result.add(b.mul(v).shift(uint8(i)))
	}
	return result
}

func (b BigInt)mul(a uint8) BigInt {
	result := make(BigInt, len(b))
	for i,v := range b {
		d := v*a
		if d >= 10 {
			if i == len(b)-1 {
				result = append(result, 0)
			}
			result[i+1] += d/10
			result[i] += d%10
		} else {
			result[i] += d
		}
	}
	return result
}

func (b BigInt)shift(n uint8) BigInt {
	for i := 0; int(n) > i; i++ {
		b, b[0] = append(b[0:1], b[0:]...), 0
		b[i] = 0
	}
	return b
}

func stringReverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}

func arrayReverse(arr []uint8) (result []uint8) {
	result = make([]uint8, len(arr))
	n := len(arr)
	for i, v := range arr {
		result[n-(i+1)] = v
	}
	return
}