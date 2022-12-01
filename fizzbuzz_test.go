package fizzbuzz

import "testing"

type testFizz struct {
	name string
	input int
	want string
}

func TestFizzBuzz(t *testing.T) {

	 tf := []testFizz{
		 {name : "should 1 get 1",input: 1,want : "1"},
	 }

	for _ , v := range tf{
		get := FizzBuzz(v.input)
		if v.want != get {
			t.Errorf("Err want %v get %v",v.want,get)
		}
	}
}
