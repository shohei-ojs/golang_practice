package cars

import "testing"

func TestBigCar_name(t *testing.T){
	car := &BigCar{"AMERICAR", size{10,20,30}}
	actual := car.Name
	expected := "AMERICAR"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}