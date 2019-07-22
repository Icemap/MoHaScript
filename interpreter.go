package main

func StartInterpreter (script string) {
	// 数据域
	memory := NewMemory()
	cmdList := []rune(string(script))
	var loopMark []int

	for cmdIndex := 0; cmdIndex < len(cmdList); cmdIndex ++ {
		cmdItem := cmdList[cmdIndex]

		if runeEquals(cmdItem, "⌚") {
			memory.ToRight()
		} else if runeEquals(cmdItem, "👓️") {
			memory.ToLeft()
		} else if runeEquals(cmdItem, "➕") {
			memory.Add()
		} else if runeEquals(cmdItem, "📹️") {
			memory.Minus()
		} else if runeEquals(cmdItem, "🎸") {
			memory.ReadNum()
		} else if runeEquals(cmdItem, "🤓") {
			memory.PrintChar()
		} else if runeEquals(cmdItem, "🧙") {
			if memory.Get() != 0 {
				loopMark = append(loopMark, cmdIndex)
			} else {
				nextCloseLoopIndex := findNextLoopClose(cmdList, cmdIndex, "🐶")
				if nextCloseLoopIndex < 0 {
					panic("[Error] 做膜法师🧙的人，必然需要言之有🐶啊，你缺🐶了知道吗？")
				} else {
					cmdIndex = nextCloseLoopIndex
				}
			}
		} else if runeEquals(cmdItem, "🐶") {
			if memory.Get() != 0 {
				cmdIndex = loopMark[len(loopMark) - 1]
			} else {
				loopMark = loopMark[:len(loopMark) - 1]
			}
		}
	}
}

func findNextLoopClose(cmdList []rune, pCurrent int, closeMark string) int {
	for i := pCurrent; i < len(cmdList); i++ {
		if runeEquals(cmdList[i], closeMark) {
			return i
		}
	}
	return -1
}

func runeEquals (runeChar rune, runeStr string) bool {
	return runeChar == []rune(runeStr)[0]
}
