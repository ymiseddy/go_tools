package algorithms_test

import (
	"fmt"
	"testing"

	"github.com/ymiseddy/go_tools/algorithms"
)

func TestPriorityCache_RetrievesValuesByKey(t *testing.T) {
	lru := algorithms.NewLRUCache[int, string](10)

	first := "one"
	second := "two"

	lru.Put(1, first)
	lru.Put(2, second)

	val, found := lru.Get(1)
	if !found {
		t.Errorf("Failed to retrieve item from queue.")
	}

	if val != first {
		t.Errorf("Incorrect value retrieved from cache wanted '%s' got '%s'", first, val)
	}

}

func TestPriorityCache_ReturnsFalseIfMissing(t *testing.T) {
	lru := algorithms.NewLRUCache[int, string](10)

	first := "one"

	lru.Put(1, first)

	val, found := lru.Get(2)
	if found {
		t.Errorf("Cache should not have returned a value. Got: %s", val)
	}

	if val != "" {
		t.Errorf("Incorrect value retrieved from cache wanted '%s' got '%s'", "", val)
	}

}

func TestPriorityCache_RollsOffOldestValue(t *testing.T) {
	lru := algorithms.NewLRUCache[int, string](10)

	for x := 0; x < 11; x++ {
		lru.Put(x, fmt.Sprintf("Item %d", x))
	}

	// First item should have rolled off the cache.
	val, found := lru.Get(0)
	if found {
		t.Errorf("Cache should not have returned a value. Got: %s", val)
	}

	if val != "" {
		t.Errorf("Incorrect value retrieved from cache wanted '%s' got '%s'", "", val)
	}

	// Remaining items should still be in the cache.
	for x := 1; x < 11; x++ {
		expected := fmt.Sprintf("Item %d", x)
		val, found = lru.Get(x)

		if !found {
			t.Errorf("Failed to retrieve item from queue.")
		}
		if val != expected {
			t.Errorf("Incorrect value retrieved from cache wanted '%s' got '%s'", expected, val)
		}
	}
}
