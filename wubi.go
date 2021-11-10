package trcnwords

//获取五笔全拼编码
func TrWubiFull(word string, mode SymbolMode, splitter ...string) string {
	return getFull(WubiDict, WubiMode, word, mode, splitter...)
}

//获取五笔首字母编码
func TrWubiCode(word string) string {
	return getCode(WubiCodeDict, word)
}
