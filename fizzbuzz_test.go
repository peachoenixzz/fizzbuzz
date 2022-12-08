package fizzbuzz

import "testing"

type testFizz struct {
	name  string
	input int
	want  string
}

func TestFizzBuzz(t *testing.T) {

	tf := []testFizz{
		{name: "should 1 get 1", input: 1, want: "1"},
		{name: "should 2 get 2", input: 2, want: "2"},
		{name: "should 3 get Fizz", input: 3, want: "Fizz"},
		{name: "should 4 get 4", input: 4, want: "4"},
		{name: "should 5 get Buzz", input: 5, want: "Buzz"},
		{name: "should 6 get Fizz", input: 6, want: "Fizz"},
		{name: "should 7 get 7", input: 7, want: "7"},
		{name: "should 8 get 8", input: 8, want: "8"},
		{name: "should 9 get Fizz", input: 9, want: "Fizz"},
		{name: "should 10 get 10", input: 10, want: "Buzz"},
		{name: "should 11 get 11", input: 11, want: "11"},
		{name: "should 12 get Fizz", input: 12, want: "Fizz"},
		{name: "should 13 get 13", input: 13, want: "13"},
		{name: "should 14 get 14", input: 14, want: "14"},
		{name: "should 15 get FizzBuzz", input: 15, want: "FizzBuzz"},
	}

	for _, v := range tf {
		t.Run(v.name, func(t *testing.T) {
			get := FizzBuzz(v.input)
			if v.want != get {
				t.Errorf("Err want %v get %v", v.want, get)
			}
		})
	}
}
