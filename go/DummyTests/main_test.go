package main

import (
	"sync"
	"testing"
)

func TestVal_Get(t *testing.T) {
	type fields struct {
		mu  sync.Mutex
		val int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Val{
				mu:  tt.fields.mu,
				val: tt.fields.val,
			}
			if got := v.Get(); got != tt.want {
				t.Errorf("Val.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test",
			args{10},
			55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkFib(j int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(j)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1",
			args{121},
			true,
		},
		{"test2",
			args{-121},
			false,
		},
		{"test3",
			args{12345654321},
			true,
		},
		{"test4",
			args{11},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOdd(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{121},
			true,
		},
		{
			"test2",
			args{122},
			false,
		},
		{
			"test3",
			args{1},
			true,
		},
		{
			"test4",
			args{2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOdd(tt.args.i); got != tt.want {
				t.Errorf("isOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {

	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		/* {
			"test1",
			args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		}, */{
			"test2",
			args{
				nums1: []int{2, 5, 6},
				m:     3,
				nums2: []int{1, 2, 3, 0, 0, 0},
				n:     3,
			},
		},
		{
			"test3",
			args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
		{
			"test4",
			args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}

func Test_isPalindromeS(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		/* {
			"test1",
			args{"A man, a plan, a canal: Panama"},
			true,
		}, */
		{
			"test1",
			args{".,"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeS(tt.args.s); got != tt.want {
				t.Errorf("isPalindromeS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test",
			args{[]string{"flower", "flow", "flight"}},
			"fl",
		},
		{
			"test1",
			args{[]string{"dog", "racecar", "car"}},
			"",
		},
		{
			"test2",
			args{[]string{"ab", "a"}},
			"a",
		},
		{
			"test3",
			args{[]string{"a", "b"}},
			"",
		},
		{
			"test4",
			args{[]string{"aaa", "aa", ""}},
			"",
		},
		{
			"test5",
			args{[]string{"aaa"}},
			"aaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("test = %s longestCommonPrefix() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{"(){}[]"},
			true,
		},
		{
			"test2",
			args{"([)]"},
			false,
		},
		{
			"test3",
			args{"()["},
			false,
		},
		{
			"test4",
			args{"]"},
			false,
		},
		{
			"test5",
			args{"[{}]"},
			true,
		},
		{
			"test5",
			args{"({[)"},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("%s isValid() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test",
			args{"abab", "abab"},
			true,
		},
		{
			"test1",
			args{"abcd", "cbad"},
			true,
		},
		{
			"test2",
			args{"ab", "ab"},
			false,
		},
		{
			"test3",
			args{"aa", "aa"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("test= %s buddyStrings() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDefer1(t *testing.T) {
	type args struct {
		name *string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
		want2 int
	}{
		{
			name:  "test",
			want:  "HradCodedName-1",
			want1: 1,
			want2: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := Defer1(tt.args.name)
			if got != tt.want {
				t.Errorf("TestDefer() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TestDefer() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("TestDefer() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test",
			args{
				[]int{1, 4, 2, 3, 4, 5},
				4,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefer2(t *testing.T) {
	tests := []struct {
		name       string
		wantResult int
	}{
		{
			"test",
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Defer2(); gotResult != tt.wantResult {
				t.Errorf("Defer2() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

/* func Test_mergeTwoLists(t *testing.T) {

	var a *ListNode
	list1 := new(ListNode)
	list1.Val = 1
	list1.Next = new(ListNode)
	a = list1.Next
	a.Val = 2
	a.Next = new(ListNode)
	a = a.Next
	a.Val = 4

	list2 := new(ListNode)
	list2.Val = 1
	list2.Next = new(ListNode)
	a = list2.Next
	a.Val = 2
	a.Next = new(ListNode)
	a = a.Next
	a.Val = 3

	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test1",
			args{list1, list2},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
} */

func Test_scheduler(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheduler()
		})
	}
}

func Test_sortingInt(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortingInt()
		})
	}
}

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test",
			args{"aaaaaa", "xxaaaa"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		/* 		{
			"test1",
			args{
				[]int{-1, 0, 3, 5, 9, 12},
				9,
			},
			4,
		}, */
		{
			"test1",
			args{
				[]int{-1, 0, 3, 5, 9, 12},
				0,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
