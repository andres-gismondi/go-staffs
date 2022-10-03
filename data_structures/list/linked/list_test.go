package linked_test

import (
	"testing"

	"go-staffs/data_structures/list/linked"
)

func TestList(t *testing.T) {
	list := linked.NewLinkedList(1)

	list.Enqueue(2)
	list.Enqueue(3)
	list.ToString()

	list.Dequeue()
	list.ToString()

	list.Dequeue()
	list.ToString()

	list.Enqueue(2)
	list.Enqueue(4)
	list.ToString()

	list.EnqueueMiddle(3, 2)
	list.ToString()

	list.EnqueueMiddle(3, 2)
	list.ToString()

	list.EnqueueFront(-1)
	list.ToString()

	list.DequeueHead()
	list.ToString()
}
