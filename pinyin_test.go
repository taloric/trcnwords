package trcnwords

import (
	"testing"
)

func TestTrPinyinFull(t *testing.T) {
	type args struct {
		word     string
		mode     SymbolMode
		splitter []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestCase_PinyinFull_1", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     RmTone | Split,
			splitter: []string{"-"},
		}, want: "sheng-huo-jiu-xiang-hai-yang-，-zhi-you-yi-zhi-jian-qiang-de-ren-cai-neng-dao-da-bi-an", wantErr: false},

		{name: "TestCase_PinyinFull_2", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     RmTone | CapFirst | Split,
			splitter: []string{"-"},
		}, want: "Sheng-Huo-Jiu-Xiang-Hai-Yang-，-Zhi-You-Yi-Zhi-Jian-Qiang-De-Ren-Cai-Neng-Dao-Da-Bi-An", wantErr: false},

		{name: "TestCase_PinyinFull_2", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     Default,
			splitter: []string{""},
		}, want: "shēnghuójiùxiànghǎiyáng，zhǐyǒuyìzhìjiānqiángderéncáinéngdàodábǐàn", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TrPinyinFull(tt.args.word, tt.args.mode, tt.args.splitter...)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrPinyinFull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TrPinyinFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_TestPinyinFull(b *testing.B) {
	type args struct {
		word     string
		mode     SymbolMode
		splitter []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestCase_PinyinFull_1", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     RmTone | Split,
			splitter: []string{"-"},
		}, want: "sheng-huo-jiu-xiang-hai-yang-，-zhi-you-yi-zhi-jian-qiang-de-ren-cai-neng-dao-da-bi-an", wantErr: false},

		{name: "TestCase_PinyinFull_2", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     RmTone | CapFirst | Split,
			splitter: []string{"-"},
		}, want: "Sheng-Huo-Jiu-Xiang-Hai-Yang-，-Zhi-You-Yi-Zhi-Jian-Qiang-De-Ren-Cai-Neng-Dao-Da-Bi-An", wantErr: false},

		{name: "TestCase_PinyinFull_2", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     Default,
			splitter: []string{""},
		}, want: "shēnghuójiùxiànghǎiyáng，zhǐyǒuyìzhìjiānqiángderéncáinéngdàodábǐàn", wantErr: false},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			got, err := TrPinyinFull(tt.args.word, tt.args.mode, tt.args.splitter...)
			if (err != nil) != tt.wantErr {
				b.Errorf("TrPinyinFull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				b.Errorf("TrPinyinFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrPinyinCode(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestCase_PinyinCode_1", args: args{word: "生活就像海洋，只有意志坚强的人才能到达彼岸"}, want: "SHJXHY，ZYYZJQDRCNDDBA"},
		{name: "TestCase_PinyinCode_2", args: args{word: "123456789"}, want: "123456789"},
		{name: "TestCase_PinyinCode_3", args: args{word: "SHJXHY，ZYYZJQDRCNDDBA"}, want: "SHJXHY，ZYYZJQDRCNDDBA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TrPinyinCode(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrPinyinCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TrPinyinCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_TrPinyinCode(b *testing.B) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestCase_PinyinCode_1", args: args{word: "生活就像海洋，只有意志坚强的人才能到达彼岸"}, want: "SHJXHY，ZYYZJQDRCNDDBA"},
		{name: "TestCase_PinyinCode_2", args: args{word: "123456789"}, want: "123456789"},
		{name: "TestCase_PinyinCode_3", args: args{word: "SHJXHY，ZYYZJQDRCNDDBA"}, want: "SHJXHY，ZYYZJQDRCNDDBA"},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			got, err := TrPinyinCode(tt.args.word)
			if (err != nil) != tt.wantErr {
				b.Errorf("TrPinyinCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				b.Errorf("TrPinyinCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
