package sum

import "testing"

func TestSumInts(t *testing.T) {
	type args struct {
		m map[string]int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test1",
			args{map[string]int64{
				"1": 1,
				"2": 2,
				"3": 3,
			},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumInts(tt.args.m); got != tt.want {
				t.Errorf("SumInts() = %v, want %v", got, tt.want)
			}
		})
	}

	type argsf struct {
		m map[string]float64
	}
	testsf := []struct {
		name string
		args argsf
		want float64
	}{
		{
			"test1",
			argsf{map[string]float64{
				"1": 1.1,
				"2": 2.2,
				"3": 3.3,
			},
			},
			6.6,
		},
	}
	for _, tt := range testsf {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumFloats(tt.args.m); got != tt.want {
				t.Errorf("SumInts() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumInt64OrFloat64(tt.args.m); got != tt.want {
				t.Errorf("SumInts() = %v, want %v", got, tt.want)
			}
		})
	}

	/*
		 	//this is to show that if the type parameter doens't match the type
			//we will see compile time error
			type argst struct {
				m map[string]int
			}
			testsft := []struct {
				name string
				args argst
				want int
			}{
				{
					"test1",
					argst{map[string]int{
						"1": 1,
						"2": 2,
						"3": 3,
					},
					},
					6,
				},
			}
			for _, tt := range testsft {
				t.Run(tt.name, func(t *testing.T) {
					if got := SumInt64OrFloat64(tt.args.m); got != tt.want {
						t.Errorf("SumInts() = %v, want %v", got, tt.want)
					}
				})
			}
	*/
}
