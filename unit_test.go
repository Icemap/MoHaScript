package main

import (
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	int64Array := []int64{ 33503, 21033, 22269, 23478, 29983, 27515, 20197, 44, 20, 23682, 22240, 31096, 31119, 36991, 36235, 20043, 63 }
	var runeArray []rune
	for i := 0; i < len(int64Array); i++ {
		runeArray = append(runeArray, rune(int64Array[i]))
	}
	fmt.Println(string(runeArray))
}
