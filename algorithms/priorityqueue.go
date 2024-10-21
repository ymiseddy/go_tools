package algorithms

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type priorityQueueNode[T comparable, K constraints.Ordered] struct {
	priority K
	data     T
}

type PriorityQueue[T comparable, K constraints.Ordered] struct {
	data []*priorityQueueNode[T, K]
}

// Create a new priority queue.
func NewPriorityQueue[T comparable, K constraints.Ordered](capacity int) PriorityQueue[T, K] {
	return PriorityQueue[T, K]{
		data: make([]*priorityQueueNode[T, K], 0, capacity),
	}
}

// Push a new item onto the priority queue
func (q *PriorityQueue[T, K]) Push(value T, priority K) {
	node := priorityQueueNode[T, K]{
		priority: priority,
		data:     value,
	}
	q.data = append(q.data, &node)
	q.heapifyUp(len(q.data) - 1)
}

// Get the current size of the priority queue
func (q *PriorityQueue[T, K]) Size() int {
	return len(q.data)
}

// Retrieve the next item without popping it
func (q *PriorityQueue[T, K]) Peek() (*T, error) {
	if len(q.data) == 0 {
		return nil, errors.New("No items are in the queue.")
	}
	return &q.data[0].data, nil
}

// Pop the next item
func (q *PriorityQueue[T, K]) Pop() (*T, error) {
	if len(q.data) == 0 {
		return nil, errors.New("Queue is empty")
	}
	head := q.data[0]
	q.data[0] = q.data[len(q.data)-1]
	q.data[len(q.data)-1] = nil

	q.data = q.data[:len(q.data)-1]
	q.heapifyDown(0)
	return &head.data, nil
}

// Heapify the current index up
func (q *PriorityQueue[T, K]) heapifyUp(index int) {
	if index == 0 {
		return
	}
	currentNode := q.data[index]
	parentIndex := q.parent(index)
	parentNode := q.data[parentIndex]

	if parentNode.priority <= currentNode.priority {
		return
	}
	q.data[index] = parentNode
	q.data[parentIndex] = currentNode
	q.heapifyUp(parentIndex)
}

// Heapify the current index down
func (q *PriorityQueue[T, K]) heapifyDown(index int) {

	if index >= len(q.data) {
		return
	}

	leftChildIndex := q.leftChild(index)
	if leftChildIndex > len(q.data)-1 {
		return
	}

	rightChildIndex := q.rightChild(index)
	currentNode := q.data[index]
	leftNode := q.data[leftChildIndex]

	if rightChildIndex < len(q.data) {
		if leftNode.priority < currentNode.priority {
			q.data[index] = leftNode
			q.data[leftChildIndex] = currentNode
			q.heapifyDown(leftChildIndex)
		}
	}

	var rightNode *priorityQueueNode[T, K] = nil
	if rightChildIndex < len(q.data) {
		rightNode = q.data[rightChildIndex]
	}

	if (rightNode == nil || leftNode.priority < rightNode.priority) &&
		leftNode.priority < currentNode.priority {
		q.data[index] = leftNode
		q.data[leftChildIndex] = currentNode
		q.heapifyDown(leftChildIndex)
	} else if rightNode != nil && rightNode.priority < currentNode.priority {
		q.data[index] = rightNode
		q.data[rightChildIndex] = currentNode
		q.heapifyDown(rightChildIndex)
	}
}

func (q *PriorityQueue[T, K]) parent(index int) int {
	return (index - 1) / 2
}

func (q *PriorityQueue[T, K]) leftChild(index int) int {
	return 2*index + 1
}

func (q *PriorityQueue[T, K]) rightChild(index int) int {
	return 2*index + 2
}
