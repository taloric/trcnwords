package trcnwords

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Pinyin     string = "data/pinyin.csv"
	Wubi       string = "data/wubi.csv"
	PinyinCode string = "data/pinyincode.csv"
	WubiCode   string = "data/wubicode.csv"
)

type SymbolMode byte

func (s SymbolMode) HasFlag(flag SymbolMode) bool {
	b1 := byte(s)
	b2 := byte(flag)
	return b1&b2 == b2
}

//output configurations
const (
	//default mode, output symbols with tone & no splitter, all lowercase
	Default SymbolMode = 0x1
	//capitalize first letter
	CapFirst SymbolMode = 0x2
	//remove pinyin tone
	RmTone SymbolMode = 0x4
	//return symbols with specific split char
	Split SymbolMode = 0x8
)

const (
	PinyinMode = 1
	WubiMode   = 2
)

type vowel int32

var (
	fir = []vowel{'ā', 'ē', 'ī', 'ō', 'ū', 'ǖ', 'Ā', 'Ē', 'Ī', 'Ō', 'Ū', 'Ǖ'} // 单韵母 一声
	sec = []vowel{'á', 'é', 'í', 'ó', 'ú', 'ǘ', 'Á', 'É', 'Í', 'Ó', 'Ú', 'Ǘ'} // 单韵母 二声
	thi = []vowel{'ǎ', 'ě', 'ǐ', 'ǒ', 'ǔ', 'ǚ', 'Ǎ', 'Ě', 'Ǐ', 'Ǒ', 'Ǔ', 'Ǚ'} // 单韵母 三声
	fou = []vowel{'à', 'è', 'ì', 'ò', 'ù', 'ǜ', 'À', 'È', 'Ì', 'Ò', 'Ù', 'Ǜ'} // 单韵母 四声
	non = []vowel{'a', 'e', 'i', 'o', 'u', 'v', 'A', 'E', 'I', 'O', 'U', 'V'} // 单韵母 无声调
)

//read a data copy from file
var (
	PinyinDict     CodeMap = make(CodeMap, 0)
	WubiDict       CodeMap = make(CodeMap, 0)
	PinyinCodeDict CodeMap = make(CodeMap, 0)
	WubiCodeDict   CodeMap = make(CodeMap, 0)
	toneDict       map[vowel]vowel
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
		//sort for search
		sort.Sort(*cm)
	}

	toneDict = make(map[vowel]vowel, 6*2*4)
	for _, vow := range [][]vowel{fir, sec, thi, fou} {
		for k, v := range vow {
			toneDict[v] = non[k]
		}
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

func isEmpty(src string) bool {
	if len(src) == 0 {
		return true
	}
	//u+0009=HT
	//u+000d=CR
	//u+00a0=NBSP
	//u+0085=NEL
	for i := 0; i < len(src); i++ {
		c := src[i]
		if !(c == ' ' || (c >= '\u0009' && c <= '\u000d') || c == '\u00a0' || c == '\u0085') {
			return false
		}
	}
	return true
}

func getFull(cm *CodeMap, fullMode int, word string, mode SymbolMode, splitter ...string) (string, error) {
	if isEmpty(word) {
		return "", nil
	}
	workRune := []rune(word)
	codes := make([]string, 0, len(workRune))

	for _, v := range workRune {
		if (v >= 48 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			//number,A-Z,a-z
			codes = append(codes, string(v))
			continue
		}

		val := cm.GetValue(v)

		if val == "" {
			codes = append(codes, string(v))
			continue
		}

		if mode.HasFlag(RmTone) && fullMode == PinyinMode {
			val = rmTone(val)
		}

		if mode.HasFlag(CapFirst) {
			//transfer to upper case val
			val = capitalizeFirstChar(val)
		}

		codes = append(codes, val)
	}

	def_splitter := ""
	if len(splitter) > 0 && mode.HasFlag(Split) {
		def_splitter = splitter[0]
	}

	return strings.Join(codes, def_splitter), nil
}

func capitalizeFirstChar(w string) string {
	w_arr := []byte(w)
	if w_arr[0] >= 97 && w_arr[0] <= 122 {
		w_arr[0] -= 32
	}
	return string(w_arr)
}

func rmTone(w string) string {
	w_arr := []int32(w)
	for k, v := range w_arr {
		idx := vowel(v)
		t, ok := toneDict[idx]
		if ok {
			w_arr[k] = int32(t)
		}
	}
	return string(w_arr)
}

//获取拼音首字母
func getCode(cm *CodeMap, word string) (string, error) {
	if isEmpty(word) {
		return "", nil
	}
	wordRune := []rune(word)
	codes := make([]string, 0, len(wordRune))
	for _, v := range wordRune {
		if (v >= 48 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			//number,A-Z,a-z
			codes = append(codes, string(v))
			continue
		}

		val := cm.GetValue(v)
		if val == "" {
			codes = append(codes, string(v))
		} else {
			codes = append(codes, val)
		}
	}
	return strings.Join(codes, ""), nil
}
