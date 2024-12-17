package algorithms_test

import (
	"fmt"
	"math"
	"math/rand/v2"
	"testing"

	"github.com/ymiseddy/go_tools/algorithms"
)

func PopExpectedMaximum(t *testing.T, q *algorithms.PriorityQueue[int], expected []int) {
	expectedMap := make(map[int]bool)
	for _, val := range expected {
		expectedMap[val] = false
	}

	prevVal := math.MaxInt64
	for q.Len() > 0 {
		val := q.PopItem()

		// Ensure val is in expected
		_, ok := expectedMap[val]
		if !ok {
			t.Fatalf("Popped unexpected value %d", val)
		}

		if prevVal < val {
			t.Fatalf(fmt.Sprintf("Item popped out of order %d > %d", prevVal, val))
		}
		expectedMap[val] = true

		prevVal = val
	}

	for _, val := range expected {
		if !expectedMap[val] {
			t.Fatalf("Expected %d but not found", val)
		}
	}
}

func PopExpectedMinimum(t *testing.T, q *algorithms.PriorityQueue[int], expected []int) {
	expectedMap := make(map[int]bool)
	for _, val := range expected {
		expectedMap[val] = false
	}

	prevVal := -1
	for q.Len() > 0 {
		val := q.PopItem()

		// Ensure val is in expected
		_, ok := expectedMap[val]
		if !ok {
			t.Fatalf("Popped unexpected value %d", val)
		}

		if prevVal > val {
			t.Fatalf(fmt.Sprintf("Item popped out of order %d > %d", prevVal, val))
		}
		expectedMap[val] = true

		prevVal = val
	}

	for _, val := range expected {
		if !expectedMap[val] {
			t.Fatalf("Expected %d but not found", val)
		}
	}
}

func TestPriorityQueue_PushInOrder_ShouldPopInOrder(t *testing.T) {
	q := algorithms.NewMinCostPriorityQueue[int](100)
	expectedItems := []int{}

	for x := 0; x < 100; x++ {
		expectedItems = append(expectedItems, x)
		q.PushItem(x, x)
	}
	PopExpectedMinimum(t, q, expectedItems)
}

func TestPriorityQueueMaximumCost_PushInOrder_ShouldPopInOrder(t *testing.T) {
	q := algorithms.NewMaxCostPriorityQueue[int](100)
	expectedItems := []int{}

	for x := 0; x < 100; x++ {
		expectedItems = append(expectedItems, x)
		q.PushItem(x, x)
	}
	PopExpectedMaximum(t, q, expectedItems)
}

func TestPriorityQueue_PushOutOfOrder_ShouldPopInOrder(t *testing.T) {
	q := algorithms.NewMinCostPriorityQueue[int](100)
	expectedItems := []int{}
	for x := 99; x >= 0; x-- {
		expectedItems = append(expectedItems, x)
		q.PushItem(x, x)
	}

	PopExpectedMinimum(t, q, expectedItems)
}

func TestPriorityQueue_PushOutOfOrder_PeekShouldShowValueBeforePopping(t *testing.T) {
	q := algorithms.NewMinCostPriorityQueue[int](10)
	for x := 9; x >= 0; x-- {
		q.PushItem(x, x)
		t.Logf("Pushing %v", q.String())
	}

	for q.Len() > 0 {
		peekVal, err := q.Peek()
		if err != nil {
			panic(err)
		}
		popVal := q.PopItem()

		if *peekVal != popVal {
			t.Fatalf("Peek value mismatch expected %d got %d", peekVal, popVal)
		}
	}
}

func TestPriorityQueue_Peek_ShouldFailOnEmptyQueue(t *testing.T) {
	q := algorithms.NewMinCostPriorityQueue[int](100)
	_, err := q.Peek()
	if err == nil {
		t.Fatalf("Expected error.")
	}
}

func TestPriorityQueue_FuzzTesting_ShouldPopInOrder(t *testing.T) {
	q := algorithms.NewMinCostPriorityQueue[int](100)
	expectedItems := make([]int, 0, 100)
	for x := 0; x < 100; x++ {
		val := rand.IntN(1000)
		expectedItems = append(expectedItems, val)
		q.PushItem(val, val)
	}
	PopExpectedMinimum(t, q, expectedItems)
}

func PriorityQueueRandom(count int) error {
	q := algorithms.NewMinCostPriorityQueue[int](count)
	for x := 0; x < count; x++ {
		val := rand.IntN(1000)
		q.PushItem(val, val)
	}

	for q.Len() > 0 {
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
