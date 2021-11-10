package trcnwords

import (
	"fmt"
)

//获取五笔全拼编码
func TrWubiFull(word string, mode SymbolMode, splitter ...string) (string, error) {
	if mode.HasFlag(RmTone) {
		return word, fmt.Errorf("not supported symbol mode : %d", mode)
	}

	return getFull(&WubiDict, WubiMode, word, mode, splitter...)
}

//获取五笔首字母编码
func TrWubiCode(word string) (string, error) {
	return getCode(&WubiCodeDict, word)
}
