package trcnwords

import (
	"testing"
)

func TestTrWubiCode(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestCase_WubiCode_1", args: args{word: "生活就像海洋，只有意志坚强的人才能到达彼岸"}, want: "TIYWII，KEUFJXRWFCGDTM"},
		{name: "TestCase_WubiCode_2", args: args{word: "123456789"}, want: "123456789"},
		{name: "TestCase_WubiCode_3", args: args{word: "TIYWII，KEUFJXRWFCGDTM"}, want: "TIYWII，KEUFJXRWFCGDTM"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TrWubiCode(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrWubiCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TrWubiCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrWubiFull(t *testing.T) {
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
		{name: "TestCase_WubiFull_1", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     Split,
			splitter: []string{"-"},
		}, want: "tg-itd-yi-wqj-itx-iu-，-kw-e-ujn-fn-jcf-xk-r-w-ft-ce-gc-dp-thc-mdfj", wantErr: false},

		{name: "TestCase_WubiFull_2", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     CapFirst | Split,
			splitter: []string{"-"},
		}, want: "Tg-Itd-Yi-Wqj-Itx-Iu-，-Kw-E-Ujn-Fn-Jcf-Xk-R-W-Ft-Ce-Gc-Dp-Thc-Mdfj", wantErr: false},

		{name: "TestCase_WubiFull_3", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     Default,
			splitter: []string{""},
		}, want: "tgitdyiwqjitxiu，kweujnfnjcfxkrwftcegcdpthcmdfj", wantErr: false},

		{name: "TestCase_WubiFull_4", args: args{
			word:     "生活就像海洋，只有意志坚强的人才能到达彼岸",
			mode:     RmTone,
			splitter: []string{""},
		}, want: "生活就像海洋，只有意志坚强的人才能到达彼岸", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TrWubiFull(tt.args.word, tt.args.mode, tt.args.splitter...)
			if (err != nil) != tt.wantErr {
				t.Errorf("TrWubiFull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TrWubiFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
