package fifo

import (
	"math/rand"
	"testing"
)

func testAssert(t *testing.T, b bool, objs ...interface{}) {
	if !b {
		t.Fatal(objs...)
	}
}

func TestBasic(t *testing.T) {
	q := NewQueue[int]()
	testAssert(t, q.Len() == 0, "Could not assert that new Queue has length zero (0).")
	q.Add(10)
	testAssert(t, q.Len() == 1, "Could not assert that Queue has lenght 1 after adding one item.")
	testAssert(t, q.Next() == 10, "Could not retrieve item from Queue correctly.")
	testAssert(t, q.Len() == 0, "Could not assert that Queue has length 0 after retrieving item.")
	q.Add(11)
	testAssert(t, q.Len() == 1, "Could not assert that Queue has length 1 after adding one item the second time.")
	testAssert(t, q.Next() == 11, "Could not retrieve item from Queue correctly the second time.")
	testAssert(t, q.Len() == 0, "Could not assert that Queue has length 0 after retrieving item the second time.")
}

func TestRandomized(t *testing.T) {
	var first, last int
	q := NewQueue[int]()
	for i := 0; i < 10000; i++ {
		if rand.Intn(2) == 0 {
			count := rand.Intn(128)
			for j := 0; j < count; j++ {
				q.Add(last)
				last++
			}
		} else {
			count := rand.Intn(128)
			if count > (last - first) {
				count = last - first
			}
			for i := 0; i < count; i++ {
				testAssert(t, q.Len() > 0, "len==0", q.Len())
				testAssert(t, q.Next() == first)
				first++
			}
		}
	}
}
