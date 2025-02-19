package stack

import (
	"sync"
)

const chunkSize = 64

// chunks are used to make a queue auto resizeable.
type chunk[t any] struct {
	items       [chunkSize]t // list of queue'ed items
	first, last int          // positions for the first and list item in this chunk
	next        *chunk[t]    // pointer to the next chunk (if any)
}

// fifo queue
type FIFO[t any] struct {
	head, tail *chunk[t]  // chunk head and tail
	count      int        // total amount of items in the queue
	lock       sync.Mutex // synchronisation lock
}

// NewQueue creates a new and empty *fifo.Queue
func NewQueue[t any]() (q *FIFO[t]) {
	initChunk := new(chunk[t])
	q = &FIFO[t]{
		head: initChunk,
		tail: initChunk,
	}
	return q
}

// Return the number of items in the queue
func (q *FIFO[t]) Len() (length int) {
	// locking to make Queue thread-safe
	q.lock.Lock()
	defer q.lock.Unlock()

	// copy q.count and return length
	length = q.count
	return length
}

// Add an item to the end of the queue
func (q *FIFO[t]) Add(item t) {
	// locking to make Queue thread-safe
	q.lock.Lock()
	defer q.lock.Unlock()

	// if the tail chunk is full, create a new one and add it to the queue.
	if q.tail.last >= chunkSize {
		q.tail.next = new(chunk[t])
		q.tail = q.tail.next
	}

	// add item to the tail chunk at the last position
	q.tail.items[q.tail.last] = item
	q.tail.last++
	q.count++
}

// Remove the item at the head of the queue and return it.
// Returns nil when there are no items left in queue.
func (q *FIFO[t]) Next() (item t) {
	// locking to make Queue thread-safe
	q.lock.Lock()
	defer q.lock.Unlock()

	// Return nil if there are no items to return
	if q.count == 0 {
		return item
	}
	// FIXME: why would this check be required?
	if q.head.first >= q.head.last {
		return item
	}

	// Get item from queue
	item = q.head.items[q.head.first]

	// increment first position and decrement queue item count
	q.head.first++
	q.count--

	if q.head.first >= q.head.last {
		// we're at the end of this chunk and we should do some maintainance
		// if there are no follow up chunks then reset the current one so it can be used again.
		if q.count == 0 {
			q.head.first = 0
			q.head.last = 0
			q.head.next = nil
		} else {
			// set queue's head chunk to the next chunk
			// old head will fall out of scope and be GC-ed
			q.head = q.head.next
		}
	}

	// return the retrieved item
	return item
}
