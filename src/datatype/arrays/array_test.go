package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	}

	t.Run("Sum Array Size 5", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})

	//compile fail: 数组有一个有趣的属性，它的大小也属于类型的一部分，如果你尝试将 [4]int 作为 [5]int 类型的参数传入函数，是不能通过编译的
	// t.Run("Sum Array Size 4", func(t *testing.T) {
	// 	numbers :=[4]int{1,2,3,4}
	// 	got :=Sum(numbers)
	// 	want :=10
	// 	assertCorrectMessage(t, got, want)
	// })

	// t.Run("collection of any size", func(t *testing.T) {
	//     numbers := []int{1, 2, 3}

	//     got := Sum(numbers)
	//     want := 6

	//     if got != want {
	//         t.Errorf("got %d want %d given, %v", got, want, numbers)
	//     }
	// })
}

// go test -cover
