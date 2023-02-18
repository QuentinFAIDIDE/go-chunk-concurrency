# Go Chunk Concurrency Utils
A tiny package to break down a slice into index chunks to help implementing
chunk based concurrency.

## Typical use case
```go
numbersToHalves := []int{4, 10, 30, 200, 3090}
halvedNumbers := make([]int, 5)

concurrentIdsChunks := MakeConcurrencyChunks(5, 2)

// iterate over chunks with their own waitgroup synchronously
for _, cChunk := range concurrentIdsChunks {
    var wg sync.WaitGroup
    wg.Add(len(cChunk))
    // for each chunk conrrent task pop a worker and perform task
    for _, taskId := range cChunk {
        go func(i int) {
            halvedNumbers[i] = numbersToHalves[i] >> 1
            wg.Done()
        }(taskId)
    }
    wg.Wait()
}
fmt.Println(halvedNumbers)
// Output: {2, 5, 15, 100, 1545}
```