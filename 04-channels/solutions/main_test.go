package main

import (
	"reflect"
	"runtime"
	"testing"
)

func TestAdd(t *testing.T) {
	resultCh := make(chan int)
	addTestData := GetAddTestData()
	go Add(addTestData.values, resultCh)
	result := <-resultCh
	if result != addTestData.expected {
		t.Errorf("Addition [4, 5, 10] expected: %d , actual: %d", addTestData.expected, result)
	}
}

func TestPartitionSlice(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	partitionSize := 4
	expectedPartitions := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19}}
	partitions := PartitionSlice(values, partitionSize)
	if !reflect.DeepEqual(partitions, expectedPartitions) {
		t.Errorf("PartitionSlice: expected: %v , actual: %v", expectedPartitions, partitions)
	}
}

func TestEmptyParitionSlice(t *testing.T) {
	values := []int{}
	partitionSize := 5
	partitions := PartitionSlice(values, partitionSize)
	expectedPartitions := [][]int{}
	if !reflect.DeepEqual(partitions, expectedPartitions) {
		t.Errorf("EmptyPartitionSlice: expected: %v , actual: %v", expectedPartitions, partitions)
	}
}

func TestAddResults(t *testing.T) {
	results := make(chan int, 3)
	results <- 1
	results <- 2
	results <- 6
	result := AddResults(results, 3)
	expected := 9
	if result != expected {
		t.Errorf("AddResults [1, 2, 6] expected: %d , actual: %d", expected, result)
	}
}

func TestAddPartitions(t *testing.T) {
	partitions := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}
	expected := 55
	result := AddPartitions(partitions)
	if result != expected {
		t.Errorf("AddPartitions expected: %d , actual: %d", expected, result)
	}
}

func TestParallelAdd(t *testing.T) {
	addTestData := GetAddTestData()
	partitionSize := 4
	result := ParallelAdd(addTestData.values, partitionSize)
	if result != addTestData.expected {
		t.Errorf("ParallelAdd [1...19] expected: %d , actual: %d", addTestData.expected, result)
	}
}

type AddTestData struct {
	values   []int
	expected int
}

func GetAddTestData() AddTestData {
	return AddTestData{
		values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		expected: 190,
	}
}

func BenchmarkAdd(b *testing.B) {
	values := getBenchmarkAddValues()
	for n := 0; n < b.N; n++ {
		result := make(chan int)
		go Add(values, result)
		<-result
	}
}

func BenchmarkParallelAdd(b *testing.B) {
	values := getBenchmarkAddValues()
	partitionSize := len(values) / runtime.NumCPU()
	for n := 0; n < b.N; n++ {
		ParallelAdd(values, partitionSize)
	}
}

func getBenchmarkAddValues() []int {
	var values []int
	for i := 0; i < 1000000; i++ {
		values = append(values, i)
	}
	return values
}
