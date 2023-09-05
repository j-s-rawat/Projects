package main

//test question https://gist.github.com/rumpl/c37f2373f3fd3705df5f1eff1e7f9522

import (
	"testing"
)

func TestRomanToIntString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{"III"},
			3,
		},
		{
			"test2",
			args{"IV"},
			4,
		},
		{
			"test2",
			args{"IVIII"},
			7,
		},
		{
			"test2",
			args{"VIIIX"},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToIntString(tt.args.input); got != tt.want {
				t.Errorf("RomanToIntString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToRoman(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"test1",
			args{23},
			"XXIII",
		},
		{
			"test1",
			args{243},
			"CCXLIII",
		},

		{
			"test1",
			args{943},
			"CMXLIII",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := IntToRoman(tt.args.input); gotOut != tt.wantOut {
				t.Errorf("IntToRoman() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
