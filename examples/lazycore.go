package main

import (
	"fmt"
	"github.com/orkunkaraduman/lazygo/lazycore"
)

func DuplicateTest() {
	var pSrc, pDst *int64
	var src, dst int64
	pSrc = &src
	src = -3 * 1024 * 1024 * 1024 * 1024
	pDst = lazycore.Duplicate(src).(*int64)
	dst = *pDst
	fmt.Printf("pDst := Duplicate(src).(*int64) -> src=%d(0x%x) dst=%d(0x%x)\n", src, pSrc, dst, pDst)
}

func IndexTest() {
	s := []int{3, 4, 5, 6}
	fmt.Printf("Slice is %v\n", s)
	fmt.Printf("Index of 5 in slice: %d\n", lazycore.Index(s, 5))
	fmt.Printf("Index of 1 in slice: %d\n", lazycore.Index(s, 1))
}

func EachTest() {
	m := map[string]int{"a": 3, "b": 4, "c": 5, "d": 6}
	fmt.Printf("Map is %v\n", m)
	k, v := lazycore.Each(m)
	fmt.Printf("Keys and values in map with Each function (ordered): %v %v\n", k, v)
	k, v = lazycore.Keys(m), lazycore.Values(m)
	fmt.Printf("Keys and values in map with Keys, Values functions: %v %v\n", k, v)
}

func main() {
	DuplicateTest()
	IndexTest()
	EachTest()
}
