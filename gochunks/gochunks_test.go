package gochunks

import (
	"testing"
)

// Why am I even testing these 3 ?
// smart boy on stackoverflow gave a concise trick for testing panic: https://stackoverflow.com/questions/31595791/how-to-test-panics
func TestUnitNullParams1(t *testing.T) {
	defer func() { _ = recover() }()
	MakeConcurrencyChunks(0, 5)
	t.Fatal("code did not panic")
}

func TestUnitNullParams2(t *testing.T) {
	defer func() { _ = recover() }()
	MakeConcurrencyChunks(5, 0)
	t.Fatal("code did not panic")
}

func TestUnitNullParams3(t *testing.T) {
	defer func() { _ = recover() }()
	MakeConcurrencyChunks(-5, 1)
	t.Fatal("code did not panic")
}

func TestUnitChunks1(t *testing.T) {
	chunks := MakeConcurrencyChunks(10, 2)
	// should give 5 chunks of size 2 with unique ids
	if len(chunks) != 5 {
		t.Fatal("Wrong number of chunks returned")
	}
	idsMap := make(map[int]struct{}, 10)
	for i := range chunks {
		if len(chunks[i]) != 2 {
			t.Fatal("One of the chunk had unexpected size")
		}
		// test that chunks ids are unique an inbound
		for id := range chunks[i] {
			if chunks[i][id] < 0 || chunks[i][id] > 9 {
				t.Fatal("Chunks ids are not in bound")
			}
			if _, idAlreadyUsed := idsMap[chunks[i][id]]; idAlreadyUsed {
				t.Fatal("Some chunks have common or duplicated ids inside")
			}
			idsMap[chunks[i][id]] = struct{}{}
		}
	}
}

func TestUnitChunks2(t *testing.T) {
	chunks := MakeConcurrencyChunks(11, 2)
	// should give 5 chunks of size 2 with unique ids and one with size of 1
	if len(chunks) != 6 {
		t.Fatal("Wrong number of chunks returned")
	}
	idsMap := make(map[int]struct{}, 11)
	for i := range chunks {

		// first 5 chunks should be of size 2
		if i != 5 {
			if len(chunks[i]) != 2 {
				t.Fatal("One of the chunk had unexpected size")
			}
		} else {
			if len(chunks[i]) != 1 {
				t.Fatal("One of the chunk had unexpected size")
			}
		}

		// test that chunks ids are unique an inbound
		for id := range chunks[i] {
			if chunks[i][id] < 0 || chunks[i][id] > 10 {
				t.Fatal("Chunks ids are not in bound")
			}
			if _, idAlreadyUsed := idsMap[chunks[i][id]]; idAlreadyUsed {
				t.Fatal("Some chunks have common or duplicated ids inside")
			}
			idsMap[chunks[i][id]] = struct{}{}
		}
	}
}

func TestUnitChunks3(t *testing.T) {
	chunks := MakeConcurrencyChunks(7, 10)
	// should give single chunk of size seven
	if len(chunks) != 1 {
		t.Fatal("Wrong number of chunks returned")
	}
	idsMap := make(map[int]struct{}, 7)

	if len(chunks[0]) != 7 {
		t.Fatal("One of the chunk had unexpected size")
	}

	// test that chunks ids are unique an inbound
	for id := range chunks[0] {
		if chunks[0][id] < 0 || chunks[0][id] > 6 {
			t.Fatal("Chunks ids are not in bound")
		}
		if _, idAlreadyUsed := idsMap[chunks[0][id]]; idAlreadyUsed {
			t.Fatal("Some chunks have common or duplicated ids inside")
		}
		idsMap[chunks[0][id]] = struct{}{}
	}

}
