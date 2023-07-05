package main

import (
	"reflect"
	"testing"
)

var testNodes []*Node

func init() {
	b := new(Node)
	b.val = "b"

	a := new(Node)
	a.val = "a"
	a.next = b
	testNodes = append(testNodes, a, b)
}

func Test_buildList(t *testing.T) {
	type args struct {
		s        string
		d        string
		allNodes []*Node
	}
	tests := []struct {
		name string
		args args
		want []*Node
	}{
		{
			name: "test1",
			args: args{s: "a", d: "b"},
			want: testNodes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildList(tt.args.s, tt.args.d, tt.args.allNodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchNodes(t *testing.T) {
	type args struct {
		nodes []*Node
		val   string
	}

	c := new(Node)
	c.val = "c"
	testNodes[1].next = c
	testNodes = append(testNodes, c)

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "test",
			args: args{testNodes, "b"},
			want: testNodes[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchNodes(tt.args.nodes, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoLinkedListsIntersect(t *testing.T) {
	type args struct {
		nodes []*Node
		test  string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{testNodes, "b,c"},
			want: false,
		},
		{
			name: "test2",
			args: args{testNodes, "d,c"},
			want: false,
		},
		{
			name: "test2",
			args: args{testNodes, "a,b"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoLinkedListsIntersect(tt.args.nodes, tt.args.test); got != tt.want {
				t.Errorf("DoLinkedListsIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
