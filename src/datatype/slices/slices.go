package main

import (
	"fmt"
	"os"
	"regexp"
)

func Slices() {

	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	fmt.Println(len(names))
	fmt.Println(cap(names))

	// 切片就像数组的引用
	// 切片并不存储任何数据，它只是描述了底层数组中的一段。
	// 更改切片的元素会修改其底层数组中对应的元素。
	// 与它共享底层数组的切片都会观测到这些修改。

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("-----------")
	b[0] = "XXX"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(names)

	// A slice cannot be grown beyond its capacity.
	// names = names[:cap(names)+1]
	// fmt.Println(names)

}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

var digitRegexp = regexp.MustCompile("[0-9]+")

// FindDigits This code behaves as advertised, but the returned []byte points into an array containing the entire file. Since the slice references the original array, as long as the slice is kept around the garbage collector can’t release the array; the few useful bytes of the file keep the entire contents in memory.
func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	return digitRegexp.Find(b)
}

// CopyDigits To fix this problem one can copy the interesting data to a new slice before returning it:
func CopyDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func main() {
	Slices()
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	// p == []byte{2, 3, 5, 7, 11, 13}
	fmt.Println(p)
	// build in append
	p = append(p, 17, 23)
	fmt.Println(p)

}
