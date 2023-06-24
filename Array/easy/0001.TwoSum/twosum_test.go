package twosum

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	nums   []int = []int{2, 7, 11, 15}
	target int   = 9
)

func TestTwoSum(t *testing.T) {

	got := TwoSum(nums, target)
	expected := []int{0, 1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("want %v, but got %v ", expected, got)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}

func ExampleTwoSum() {
	indexSlice := TwoSum(nums, target)
	fmt.Println(indexSlice)
	// Output: [0 1]
}
