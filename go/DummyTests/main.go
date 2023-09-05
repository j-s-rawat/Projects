package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"unicode"
)

/*
Race condition https://dave.cheney.net/2015/11/18/wednesday-pop-quiz-spot-the-race

	type RPC struct {
		result int
		done   chan struct{}
	}

	func (rpc *RPC) compute() {
		time.Sleep(time.Second) // strenuous computation intensifies
		rpc.result = 42
		close(rpc.done)
	}

	func (RPC) version() int {
		return 1 // never going to need to change this
	}

	func main() {
		rpc := &RPC{done: make(chan struct{})}

		go rpc.compute()         // kick off computation in the background
		version := rpc.version() // grab some other information while we're waiting
		<-rpc.done               // wait for computation to finish
		result := rpc.result

		fmt.Printf("RPC computation complete, result: %d, version: %d\n", result, version)
	}
*/
/*
func main() {
	var a int
	var b, c = &a, &a
	fmt.Println(b, c)   // 0x1040a124 0x1040a124
	fmt.Println(&b, &c) // 0x1040c108 0x1040c110
	var d int = 10
	b = &d
	fmt.Println(*b, *c) // 0x1040a124 0x1040a124
} */

/* func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		n := i // Create a unique variable for each closure.
		go func() {
			fmt.Print(n)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
} */

type Val struct {
	mu  sync.Mutex
	val int
}

func (v *Val) Get() int {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.val
}

func (v *Val) Add(n int) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.val += n
}

type Stats struct {
	a, b, c Val
}

func (s *Stats) Sum() int {
	return s.a.Get() + s.b.Get() + s.c.Get() // whoops
}

func main() {
	var t = Stats{
		a: Val{val: 1},
		b: Val{val: 2},
		c: Val{val: 3},
	}
	fmt.Println(t.Sum())

	var test Pill = 2
	fmt.Println(test.String())
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Pill
func isOdd(i int) bool {
	return i&1 != 0
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i := len(str)/2 - 1
	j := i + 1
	if isOdd(len(str)) {
		j++
	}
	for ; i >= 0; i, j = i-1, j+1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < j && j > 0; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func isPalindromeS(s string) bool {
	//  var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	//	s = strings.ToUpper(nonAlphanumericRegex.ReplaceAllString(s, ""))
	//	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < len(s) && !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
		}
		for j > i && !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
		}

		if i < j && unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			fmt.Printf("%c %c\n", s[i], s[j])
			return false
		}
	}
	return true
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, m+n)

	if m == 0 {
		fmt.Println(nums2[:n])

	} else if n == 0 {
		fmt.Println(nums1[:m])

	} else {
		t := m + n
		i, j := 0, 0
		for l := 0; l < t; l++ {
			if i == m {
				result = append(result[:i+1], nums2[j:n]...)
			} else if j == n {
				result = append(result[:j+1], nums1[i:m]...)
			} else if nums1[i] < nums2[j] {
				result[l] = nums1[i]
				i++
			} else {
				result[l] = nums2[j]
				j++
			}

		}
		fmt.Println(result)
	}

}

func isAnagram(s string, t string) bool {
	var m = make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	l := len(strs[0])
	var allMatched bool
	var result []byte
	for i := 0; i < l; i++ {
		v := strs[0][i]
		for _, u := range strs[1:] {
			if len(u) > i {
				if u[i] != v {
					allMatched = false
					break
				} else {
					allMatched = true
				}
			} else {
				allMatched = false
				break
			}
		}
		if allMatched {
			result = append(result, v)
		} else {
			break
		}
		allMatched = false
	}
	return string(result)
}

const (
	b1 = '('
	b2 = ')'
	b3 = '{'
	b4 = '}'
	b5 = '['
	b6 = ']'
)

// ******attempt a diff approach*****///
func isValid(s string) bool {
	var stack []rune
	for _, v := range s {
		if len(stack) > 0 {
			switch stack[len(stack)-1] {
			case b1:
				if v == rune(b2) {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}

			case b3:
				if v == rune(b4) {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}

			case b5:
				if v == rune(b6) {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}
			}
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	freqS := [26]int{}
	freqG := [26]int{}

	var diff int
	for i := 0; i < len(s); i++ {
		freqS[s[i]-'a']++
		freqG[goal[i]-'a']++
		if s[i] != goal[i] {
			diff++
		}
	}

	if freqS == freqG && diff == 2 {
		return true
	}
	if diff == 0 {
		for i := 0; i < 26; i++ {
			if freqS[i] >= 2 {
				return true
			}
		}
	}
	return false
}

var tempName = "HradCodedName"
var Task_id = 1

func Defer1(name *string) (s string, t, k int) {
	defer func() {
		Task_id++
		k = Task_id
	}()
	if name == nil {
		return fmt.Sprintf("%s-%d", tempName, Task_id), Task_id, k
	}
	return fmt.Sprintf("%s-%d", *name, Task_id), Task_id, k
}

func Defer2() (result int) {

	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	tail := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 == nil {
		tail.Next = list2
	} else if list2 == nil {
		tail.Next = list1
	}
	return head.Next
}

type gfun func(...any) any

// type with reflect
func myTest(arg ...any) any {
	for _, a := range arg {
		switch reflect.TypeOf(a).String() {
		case "string":
			fmt.Println(a, "is a string")
		case "int32":
			fmt.Println(a, "is a int32")
		default:
			fmt.Println(a, "is unknown type")
		}
	}
	return true
}

func scheduler() {
	var test gfun

	test = myTest
	test("hello", int32(1))

}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TypeWithInterface() {
	do(21)
	do("hello")
	do(true)
}

func sortingInt() {
	lst := []int{2, 5, 6, 7, 1, 9, 10}
	//sort.Ints(lst)
	sort.Sort(sort.IntSlice(lst))

	fmt.Println(lst)
}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
outer:
	for ; i < len(s); i++ {
		for ; j < len(t); j++ {
			if s[i] == t[j] {
				j++
				i++
				goto outer
			}
		}
		if j == len(t) {
			break
		}
	}
	return i == len(s)
}

func search(nums []int, target int) int {
	mI := len(nums) / 2
	if mI == 0 {
		return -1
	}
	if target == nums[mI] {
		return mI
	} else if target > nums[mI] {
		return mI + search(nums[mI:], target)
	} else {
		return mI - search(nums[:mI+1], target)
	}
}
