package bigint

import (
    "testing" 
)

func TestAdd(t *testing.T) {
		a := FromString("1234")
		
		a.add(FromString("34567"))
}
func TestMul(t *testing.T) {
	a := FromString("1234")
		
	a.mul(uint8(3))
}

func TestShift(t *testing.T) {
	a := FromString("1234")
	a.shift(uint8(3))
}
