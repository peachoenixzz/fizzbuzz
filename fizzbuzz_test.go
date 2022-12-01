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
		 {name : "should 2 get 2",input: 2,want : "2"},
		 {name : "should 3 get Fizz",input: 3,want : "Fizz"},
		 {name : "should 4 get 4",input: 4,want : "4"},
		 {name : "should 5 get Buzz",input: 5,want : "Buzz"},
		 {name : "should 6 get Fizz",input: 6,want : "Fizz"},

	 }

	for _ , v := range tf{
		get := FizzBuzz(v.input)
		if v.want != get {
			t.Errorf("Err want %v get %v",v.want,get)
		}
	}
}
