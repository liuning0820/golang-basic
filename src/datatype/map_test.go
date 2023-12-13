package datatype

import (
	"testing"

)



func TestMap(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Access Map Item", func(t *testing.T) {
		got := Maps()
		want := 98
		assertCorrectMessage(t, got["GO"], want)
	})

	t.Run("Modify Map Item", func(t *testing.T) {
		got := Maps()
		got["GO"] = 100
		want := 100
		assertCorrectMessage(t, got["GO"], want)
	})


}

func assertValues(t *testing.T, got, want int) {
    t.Helper()

    if got != want {
        t.Errorf("got '%d' want '%d'", got, want)
    }
}

func assertError(t *testing.T, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error '%s' want '%s'", got, want)
    }
}

func TestSearch(t *testing.T)  {

	dictionary := map[string]int{"GO":98,"C#":90,"Python":100}

	got:= Search(dictionary,"Python")
	want:= 100
	assertValues(t,got, want)

	dictionary2 := Dictionary{"test": "this is just a test"}


	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary2.Search2("unknown")
	
		assertError(t, got, ErrNotFound)
	})

	
}
