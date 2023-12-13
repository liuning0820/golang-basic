package conditioncontrol

import "testing"

func TestSwitch(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Switch With Condition", func(t *testing.T) {
		got := SwitchWithCondition()
		want := "darwin"
		assertCorrectMessage(t, got, want)
	})

}
