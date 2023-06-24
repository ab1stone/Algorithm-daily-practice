package twosum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	got := TwoSum(nums, target)
	expected := []int{0, 1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("want %v, but got %v ", expected, got)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum([]int{2, 7, 11, 15}, 9)
	}
}

func ExampleTwoSum() {
	indexSlice := TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(indexSlice)
	// Output: [0 1]
}
