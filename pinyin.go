package trcnwords

//获取全拼拼音
func TrPinyinFull(word string, mode SymbolMode, splitter ...string) (string, error) {
	return getFull(&PinyinDict, PinyinMode, word, mode, splitter...)
}

//获取拼音首字母编码
func TrPinyinCode(word string) (string, error) {
	return getCode(&PinyinCodeDict, word)
}
