package conditioncontrol

import (
	"fmt"
	"runtime"
	"time"
)

func SwitchWithCondition() string {

	currentOS := runtime.GOOS

	switch currentOS {
	case "darwin":
		fmt.Println("Go runs on OS X.")

	case "linux":
		fmt.Println("Go runs on Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("Go runs on %s.\n", currentOS)

	}

	return currentOS

}

func Switch() {

	// 没有条件的 switch 同 switch true 一样。
	// 这种形式能将一长串 if-then-else 写得更加清晰
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
