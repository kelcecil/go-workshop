package main

// Write a function that adds the contents of the slice
// and sends the result to the result channel
func Add(values []int, result chan<- int) {
	result <- 0
}

// Write a function that partitions the slice into
// partitions of size partitionSize
// Ex:
// 	values: [1, 2, 3, 4, 5, 6, 7]
// 	partitionSize: 3
//  result: [[1, 2, 3], [4, 5, 6], [7]]
func PartitionSlice(values []int, partitionSize int) [][]int {
	return [][]int{}
}

// Write a function that reads numResults from the results channel
// and then adds the result together.
func AddResults(results <-chan int, numResults int) int {
	return 0
}

// Write a function to take partitions and add all of the values
// in the partitions together. The Add and AddResults functions
// will help here.
func AddPartitions(partitions [][]int) int {
	return 0
}

// This function does parallel addition, and is already implemented.
func ParallelAdd(values []int, partitionSize int) int {
	partitions := PartitionSlice(values, partitionSize)
	return AddPartitions(partitions)
}
