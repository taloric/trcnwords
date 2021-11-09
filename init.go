package trcnwords

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	Pinyin     string = "data/pinyin.csv"
	Wubi       string = "data/wubi.csv"
	PinyinCode string = "data/pinyincode.csv"
	WubiCode   string = "data/wubicode.csv"
)

type Symbol int

//output configurations
const (
	Default  Symbol = 1 //default mode
	CapFirst Symbol = 2 //capitalize first letter
	Tone     Symbol = 4 //all pinyin code with tone
	Split    Symbol = 8 //return with split char
)

type vowel int32

var (
	fir = []vowel{'ā', 'ē', 'ī', 'ō', 'ū', 'ǖ', 'Ā', 'Ē', 'Ī', 'Ō', 'Ū', 'Ǖ'} // 单韵母 一声
	sec = []vowel{'á', 'é', 'í', 'ó', 'ú', 'ǘ', 'Á', 'É', 'Í', 'Ó', 'Ú', 'Ǘ'} // 单韵母 二声
	thi = []vowel{'ǎ', 'ě', 'ǐ', 'ǒ', 'ǔ', 'ǚ', 'Ǎ', 'Ě', 'Ǐ', 'Ǒ', 'Ǔ', 'Ǚ'} // 单韵母 三声
	fou = []vowel{'à', 'è', 'ì', 'ò', 'ù', 'ǜ', 'À', 'È', 'Ì', 'Ò', 'Ù', 'Ǜ'} // 单韵母 四声
	non = []vowel{'a', 'e', 'i', 'o', 'u', 'v', 'A', 'E', 'I', 'O', 'U', 'V'} // 单韵母 无声调
)

type CodeItem struct {
	Key   rune
	Value string
}

type CodeMap []CodeItem

func (a CodeMap) Len() int           { return len(a) }
func (a CodeMap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CodeMap) Less(i, j int) bool { return a[i].Key < a[j].Key }

//read a data copy from file
var (
	PinyinDict     CodeMap = make(CodeMap, 0)
	WubiDict       CodeMap = make(CodeMap, 0)
	PinyinCodeDict CodeMap = make(CodeMap, 0)
	WubiCodeDict   CodeMap = make(CodeMap, 0)
)

func init() {
	dataHandle := map[string](*CodeMap){
		Pinyin:     &PinyinDict,
		Wubi:       &WubiDict,
		PinyinCode: &PinyinCodeDict,
		WubiCode:   &WubiCodeDict,
	}

	for k := range dataHandle {
		f, err := os.Open(k)
		if err != nil {
			fmt.Printf("read data from file %s throws error : %s", k, err.Error())
			continue
		}
		defer f.Close()

		cm := dataHandle[k]
		loadData(f, cm)
		sort.Sort(*cm)
	}
}

func loadData(f *os.File, rst *CodeMap) {
	csvReader := csv.NewReader(f)
	for {
		row, err := csvReader.Read()
		if err != nil || len(row) == 0 {
			break
		}
		idx, err := strconv.ParseInt(row[0], 16, 64)
		if err != nil {
			//ignore this error, just don't add it to array
			continue
		}
		(*rst) = append((*rst), CodeItem{
			Key:   rune(idx),
			Value: row[1],
		})
	}
}
