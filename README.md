# Go Chunk Concurrency Utils
A go package to break down a slice into chunk of indexes to help implementing
chunk based concurrency.

Chunk based concurreny may be less verbose than the usual channel/worker concurrency pattern,
albeit having the disadvantage of synchronuously waiting for every batch/chunk to complete,
and recreating one goroutine per item.

## Typical use case
```go
package main

import (
    "github.com/QuentinFAIDIDE/go-chunk-concurrency/gochunks"
)

numbersToHalves := []int{4, 10, 30, 200, 3090}
halvedNumbers := make([]int, 5)

concurrentIdsChunks := gochunks.MakeConcurrencyChunks(5, 2)

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
