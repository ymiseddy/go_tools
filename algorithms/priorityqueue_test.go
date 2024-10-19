package algorithms_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/ymiseddy/go_tools/algorithms"
)

func popQueue(t *testing.T, q *algorithms.PriorityQueue[int, int]) {
	prev := -1
	for q.Size() > 0 {
		val, err := q.Pop()
		if err != nil {
			panic(err)
		}
		if prev > *val {
			t.Fatalf(fmt.Sprintf("Item popped out of order %d > %d", prev, *val))
		}

		prev = *val
	}
}

func TestPriorityQueue_PushInOrder_ShouldPopInOrder(t *testing.T) {
	q := algorithms.NewPriorityQueue[int, int](100)
	for x := 0; x < 100; x++ {
		q.Push(x, x)
	}
	popQueue(t, &q)
}

func TestPriorityQueue_PushOutOfOrder_ShouldPopInOrder(t *testing.T) {
	q := algorithms.NewPriorityQueue[int, int](100)
	for x := 0; x < 100; x++ {
		q.Push(x, x)
	}

	popQueue(t, &q)
}

func TestPriorityQueue_PushOutOfOrder_PeekShouldShowValueBeforePopping(t *testing.T) {
	q := algorithms.NewPriorityQueue[int, int](100)
	for x := 0; x < 100; x++ {
		q.Push(x, x)
	}

	for q.Size() > 0 {
		peekVal, err := q.Peek()
		if err != nil {
			panic(err)
		}
		popVal, popErr := q.Pop()
		if popErr != nil {
			panic(popErr)
		}

		if *peekVal != *popVal {
			t.Fatalf("Peek value mismatch expected %d got %d", peekVal, popVal)
		}
	}
}

func TestPriorityQueue_Pop_ShouldFailOnEmptyQueue(t *testing.T) {
	q := algorithms.NewPriorityQueue[int, int](100)
	_, err := q.Pop()
	if err == nil {
		t.Fatalf("Expected error.")
	}
}

func TestPriorityQueue_Peek_ShouldFailOnEmptyQueue(t *testing.T) {
	q := algorithms.NewPriorityQueue[int, int](100)
	_, err := q.Peek()
	if err == nil {
		t.Fatalf("Expected error.")
	}
}

func TestPriorityQueue_FuzzTesting_ShouldPopInOrder(t *testing.T) {
	q := algorithms.NewPriorityQueue[int, int](100)

	for x := 0; x < 100; x++ {
		val := rand.IntN(1000)
		q.Push(val, val)
	}
	popQueue(t, &q)
}

func PriorityQueueRandom(count int) error {
	q := algorithms.NewPriorityQueue[int, int](count)
	for x := 0; x < count; x++ {
		val := rand.IntN(1000)
		q.Push(val, val)
	}

	for q.Size() > 0 {
		q.Pop()
	}
	return nil
}

func BenchmarkPriorityQueue10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PriorityQueueRandom(10)
	}
}

func BenchmarkPriorityQueue100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PriorityQueueRandom(100)
	}
}

func BenchmarkPriorityQueue1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PriorityQueueRandom(1000)
	}
}

func BenchmarkPriorityQueue10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PriorityQueueRandom(10000)
	}
}

func BenchmarkPriorityQueue1000000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PriorityQueueRandom(1000000)
	}
}
