package datatypeV1

import (
	"testing"
)

func TestMap(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	}

	t.Run("BANNA Happen 1 time", func(t *testing.T) {
		s := "NAABXXAN"
		got := Solution(s)
		want := 1
		assertCorrectMessage(t, got, want)
	})

	t.Run("BANNA Happen 2 time", func(t *testing.T) {
		s := "NAANAAXNABABYNNBZ"
		got := Solution(s)
		want := 2
		assertCorrectMessage(t, got, want)
	})

}

// go test -cover
