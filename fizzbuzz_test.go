package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestCorrectLength(t *testing.T) {
	want := 21
	act := fizzBuzz(21)
	got := len(act)
	if got != want {
		t.Errorf("wanted slice of length %v, got %v", want, got)
	}

	wantMore := 42
	actMore := fizzBuzz(42)
	gotMore := len(actMore)
	if gotMore != wantMore {
		t.Errorf("wanted slice of length %v, got %v", wantMore, gotMore)
	}
}

func TestCorrectType(t *testing.T) {
	want := reflect.TypeOf(make([]string, 1))
	got := reflect.TypeOf(fizzBuzz(1))
	if got != want {
		t.Errorf("wanted array of type %s, got %s", want, got)
	}
}

func TestFizz(t *testing.T) {
	want := []string{"1", "2", "fizz"}
	got := fizzBuzz(3)
	if reflect.DeepEqual(want, got) == false {
		t.Errorf("wanted %v, got %v", want, got)
		t.Errorf("%v", len(got))
	}
}

func TestBuzz(t *testing.T) {
	want := []string{"1", "2", "fizz", "4", "buzz"}
	got := fizzBuzz(5)
	if reflect.DeepEqual(want, got) == false {
		t.Errorf("wanted %v, got %v", want, got)
		t.Errorf("%v", len(got))
	}
}

func TestFizzBuzz(t *testing.T) {
	want := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}
	got := fizzBuzz(15)
	if reflect.DeepEqual(want, got) == false {
		t.Errorf("wanted %v, got %v", want, got)
		t.Errorf("%v", len(got))
	}
}
