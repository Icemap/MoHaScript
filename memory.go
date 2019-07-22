package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	data []int64
	p int
}

func NewMemory () Memory {
	return Memory{
		data: make([]int64, 10),
		p: 0,
	}
}

func (memory *Memory) Get () int64 {
	return memory.data[memory.p]
}

func (memory *Memory) Add () {
	memory.data[memory.p] ++
}

func (memory *Memory) Minus () {
	memory.data[memory.p] --
}

// 自动扩充右边界
func (memory *Memory) ToRight () {
	memory.p ++
	if memory.p >= len(memory.data) {
		memory.data = append(memory.data, int64(0))
	}
}

// 左边触底，到达最右端
func (memory *Memory) ToLeft () {
	if memory.p == 0 {
		memory.p = len(memory.data) - 1
	} else {
		memory.p --
	}
}

func (memory *Memory) PrintChar () {
	fmt.Print(string(rune(memory.data[memory.p])))
}

func (memory *Memory) ReadNum () {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	fmt.Print(input)

	if err != nil {
		panic("[Error] 蛤？？？你的输入Too young，too simple...")
	}

	inputNum, err := strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		panic("[Error] 蛤？？？你的输入Sometimes naïve")
	}

	memory.data[memory.p] = int64(inputNum)
}