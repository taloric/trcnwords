package trcnwords

type CodeItem struct {
	Key   rune
	Value string
}

type CodeMap []CodeItem

func (a CodeMap) Len() int           { return len(a) }
func (a CodeMap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CodeMap) Less(i, j int) bool { return a[i].Key < a[j].Key }

func (a CodeMap) Get(key rune) CodeItem {
	return *(a.binarySearch(key))
}

func (a CodeMap) GetValue(key rune) string {
	r := a.binarySearch(key)
	if r != nil {
		return r.Value
	} else {
		return string(key)
	}
}

//查找指定的key
func (a CodeMap) binarySearch(key rune) *CodeItem {
	left := 0
	right := len(a) - 1
	for {
		//exit
		if left > right {
			break
		}

		mid := (left + right) / 2
		if a[mid].Key == key {
			return &a[mid]
		} else if a[mid].Key > key {
			right = mid - 1
		} else if a[mid].Key < key {
			left = mid + 1
		}
	}
	return nil
}
