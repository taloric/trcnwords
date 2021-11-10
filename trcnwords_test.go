package trcnwords

import (
	"testing"
)

func TestMain(m *testing.M) {
	_ = LoadFiles("data/pinyin.csv", "data/wubi.csv", "data/pinyincode.csv", "data/wubicode.csv")
	m.Run()
}

func TestSymbolMode_HasFlag(t *testing.T) {
	type args struct {
		flag SymbolMode
	}
	tests := []struct {
		name string
		s    SymbolMode
		args args
		want bool
	}{
		{name: "TestCase_HasFlag_1", s: Default | Split, args: args{Split}, want: true},
		{name: "TestCase_HasFlag_2", s: Default, args: args{Split}, want: false},
		{name: "TestCase_HasFlag_3", s: Default | Split | RmTone | CapFirst, args: args{Split}, want: true},
		{name: "TestCase_HasFlag_4", s: RmTone | CapFirst, args: args{Default}, want: false},
		{name: "TestCase_HasFlag_5", s: 15, args: args{Split}, want: true},
		{name: "TestCase_HasFlag_6", s: 14, args: args{Default}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HasFlag(tt.args.flag); got != tt.want {
				t.Errorf("SymbolMode.HasFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEmpty(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestCase_isEmpty_1", args: args{src: ""}, want: true},
		{name: "TestCase_isEmpty_2", args: args{src: "test"}, want: false},
		{name: "TestCase_isEmpty_3", args: args{src: "	"}, want: true},
		{name: "TestCase_isEmpty_4", args: args{src: " "}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmpty(tt.args.src); got != tt.want {
				t.Errorf("isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rmTone(t *testing.T) {
	type args struct {
		w string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestCase_rmTone_1", args: args{w: "āáǎàīíǐìōóǒòūúǔùǖǘǚǜĀÁǍÀĒÉĚÈĪÍǏÌŌÓǑÒŪÚǓÙǕǗǙǛ"},
			want: "aaaaiiiioooouuuuvvvvAAAAEEEEIIIIOOOOUUUUVVVV"},
		{name: "TestCase_rmTone_1", args: args{w: "TestCase123"},
			want: "TestCase123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rmTone(tt.args.w); got != tt.want {
				t.Errorf("rmTone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_capitalizeFirstChar(t *testing.T) {
	type args struct {
		w string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestCase_capitalizeFirstChar_1", args: args{w: "test"}, want: "Test"},
		{name: "TestCase_capitalizeFirstChar_2", args: args{w: "TEST"}, want: "TEST"},
		{name: "TestCase_capitalizeFirstChar_3", args: args{w: "tEST"}, want: "TEST"},
		{name: "TestCase_capitalizeFirstChar_4", args: args{w: "123123"}, want: "123123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capitalizeFirstChar(tt.args.w); got != tt.want {
				t.Errorf("capitalizeFirstChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
