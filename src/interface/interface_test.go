package interfaces

import (
	"testing"

	"fmt"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Test Speed", func(t *testing.T) {
		car := Car("Tesla")
		got := car.Speed()
		fmt.Println(got)
		want := "200 Km/Hrs"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test Structure", func(t *testing.T) {
		car := Car("Tesla")
		got := car.Structure()[0]
		want := "ECU"
		assertCorrectMessage(t, got, want)
	})

}
