package main

import "testing"

func Test_byteArrayToInt(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{[]byte{0xFF, 0xFF, 0xFF, 0xFE}},
			-2,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := byteArrayToInt(tt.args.data); got != tt.want {
				t.Errorf("byteArrayToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
