package main

// Write a function that adds the contents of the slice
// and sends the result to the result channel
func Add(values []int, result chan<- int) {
	total := 0
	for _, value := range values {
		total += value
	}
	result <- total
}

// Write a function that partitions the slice into
// partitions of size partitionSize
// Ex:
// 	values: [1, 2, 3, 4, 5, 6, 7]
// 	partitionSize: 3
//  result: [[1, 2, 3], [4, 5, 6], [7]]
func PartitionSlice(values []int, partitionSize int) [][]int {
	partitions := [][]int{}
	length := len(values)
	for i := 0; i < length; i += partitionSize {
		end := i + partitionSize
		if end >= length {
			end = length
		}
		partitions = append(partitions, values[i:end])
	}
	return partitions
}

func AddResults(results <-chan int, numResults int) int {
	var values []int
	for i := 0; i < numResults; i++ {
		values = append(values, <-results)
	}
	addResult := make(chan int)
	go Add(values, addResult)
	return <-addResult
}

func AddPartitions(partitions [][]int) int {
	results := make(chan int)
	for i := 0; i < len(partitions); i++ {
		go Add(partitions[i], results)
	}
	return AddResults(results, len(partitions))
}

func ParallelAdd(values []int, partitionSize int) int {
	partitions := PartitionSlice(values, partitionSize)
	return AddPartitions(partitions)
}
