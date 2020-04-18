package leetcode

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTwo3Sum(t *testing.T) {
	type test struct {
		input  []int
		target int
		want   []int
	}

	tests := []test{
		{input: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{input: []int{2, 7, 11, 15}, target: 20, want: []int{0, 0}},
		{input: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
	}

	for _, tc := range tests {
		got := TwoSum2(tc.input, tc.target)
		if !reflect.DeepEqual(got, tc.want) {
			t.Log(cmp.Diff(tc.want, got))
		}
	}
}

func TestFindRepeatNumber(t *testing.T) {
	type test struct {
		input []int
		want  int
	}
	tests := []test{
		{input: []int{2, 3, 1, 0, 2, 5, 3}, want: 2},
		{input: []int{2, 3, 1, 0}, want: -1},
	}
	for _, tc := range tests {
		got := FindRepeatNumber2(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Log(cmp.Diff(tc.want, got))
		}
	}
}

func TestMerge(t *testing.T) {
	A := []int{1, 2, 3, 0, 0, 0}
	m := 3
	B := []int{2, 5, 6}
	n := 3
	Merge(A, m, B, n)
}

func TestDistributeCandies(t *testing.T) {
	type test struct {
		candies int
		people  int
		want    []int
	}
	tests := []test{
		{candies: 7, people: 4, want: []int{1, 2, 3, 1}},
		{candies: 10, people: 3, want: []int{5, 2, 3}},
	}
	for _, tc := range tests {
		got := DistributeCandies(tc.candies, tc.people)
		if !reflect.DeepEqual(got, tc.want) {
			t.Log(cmp.Diff(tc.want, got))
		}
	}
}

func TestExange(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	t.Log(Exchange(nums))
}

func TestPivotIndex(t *testing.T) {
	type test struct {
		input []int
		want  int
	}
	tests := []test{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{}, -1},
	}
	for _, tc := range tests {
		got := PivotIndex(tc.input)
		if got != tc.want {
			t.Log(tc.want, got)
		}
	}
}

func TestDominantIndex(t *testing.T) {
	nums := []int{0, 0, 0, 1}
	t.Log(DominantIndex2(nums))
}

func TestPlusOne(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}
	tests := []test{
		{[]int{9}, []int{1, 0}}, // 注意这种最高位有进位的情况
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{}, []int{}},
		{[]int{1, 2, 9}, []int{1, 3, 0}},
	}
	for _, tc := range tests {
		got := PlusOne(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Log(cmp.Diff(tc.want, got))
		}
	}
}

func TestSingleNumber(t *testing.T) {
	type test struct {
		input []int
		want  int
	}
	tests := []test{
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{2, 2, 1}, 1},
	}
	for _, tc := range tests {
		got := SingleNumber(tc.input)
		if got != tc.want {
			t.Log(cmp.Diff(tc.want, got))
		}
	}
}

func TestSingleNumber2(t *testing.T) {
	type test struct {
		input []int
		want  int
	}
	tests := []test{
		{[]int{2, 2, 3, 2}, 3},
	}
	for _, tc := range tests {
		got := SingleNumber22(tc.input)
		if got != tc.want {
			t.Log(cmp.Diff(tc.want, got))
		}
	}
}

func TestMajorityElement(t *testing.T) {
	type test struct {
		input []int
		want  int
	}
	tests := []test{
		{[]int{1}, 1},
		{[]int{3, 3, 1}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tc := range tests {
		got := MajorityElement3(tc.input)
		if got != tc.want {
			t.Log(cmp.Diff(tc.want, got))
		}
	}
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{"case one num", args{[]int{1}, 1}, 0},
		{"case repeat num", args{[]int{1, 2, 2, 3, 4}, 2}, 2},
		{"case  normal", args{[]int{1, 2, 2, 3, 4}, 3}, 3},
		{"case no target", args{[]int{1, 2, 2, 3, 4}, 5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := BinarySearch(tt.args.nums, tt.args.target); gotIndex != tt.wantIndex {
				t.Errorf("BinarySearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestRemoveDuplicateNums(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3, 5, 5, 6, 7}
	t.Log(RemoveDuplicateNums(nums))
}
