# trcnwords

用于获取中文字对应的五笔码、拼音码的Go实现


### 为什么要重复造轮子
因为能找到的实现要不就是只能获取拼音码，要不就是只能获取每个字的头字母，不能获取全拼，所以只能自己动手了


### usage

#### init before everything

``` go 
err := LoadFiles("data/pinyin.csv", "data/wubi.csv", "data/pinyincode.csv", "data/wubicode.csv")
```

#### call and get result
``` go
	word := "我是一个中国人"
	pyCode := TrPinyinCode(word)
    //WSYGZGR
	wbCode := TrWubiCode(word)
    //QJGWKLW
	fullPy := TrPinyinFull(word, Split|RmTone|CapFirst, "-")
    //Wo-Shi-Yi-Ge-Zhong-Guo-Ren
	fullWb := TrWubiFull(word, CapFirst|Split, "-")
    //Q-J-G-Wh-K-L-W
```
