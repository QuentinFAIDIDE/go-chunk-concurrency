package gochunks

// MakeConcurrencyChunks will return chunks of ids. The basic use case
// is to fill in a list concurrently in a clear and concise way.
func MakeConcurrencyChunks(maxLength int, concurrency int) [][]int {
	var chunks [][]int
	if maxLength >= concurrency {
		chunks = makeChunks(makeRange(0, maxLength-1), concurrency)
	} else {
		chunks = [][]int{makeRange(0, maxLength-1)}
	}
	return chunks
}

// stolen from https://freshman.tech/snippets/go/split-slice-into-chunks/
func makeChunks(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

// stolen from https://stackoverflow.com/questions/39868029/how-to-generate-a-sequence-of-numbers
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
