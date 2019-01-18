/* Package implements simple priority queue. Using heap package from Go Standard library might
have been better idea with simpler implementation.
*/
package priority

// Item is an item that can be added to the priority queue.
type Item interface {
	// LessEqual is used to determine ordering in the priority queue.  Assuming the queue
	// is in ascending order, this should return true.
	Less(other Item) bool
}

type Queue []Item

func (items *Queue) swap(i, j int) {
	(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
}

func (items Queue) Peek() Item {
	if len(items) == 0 {
		return nil
	}
	return items[0]
}

func (items *Queue) Pop() Item {
	size := len(*items)

	// Move last leaf to root, and pop the last item.
	items.swap(size-1, 0)
	item := (*items)[size-1] // Item to return.
	(*items)[size-1], *items = nil, (*items)[:size-1]

	// Bubble down to restore heap property.
	index := 0
	childL, childR := 2*index+1, 2*index+2
	for len(*items) > childL {
		child := childL
		if len(*items) > childR && (*items)[childR].Less((*items)[childL]) {
			child = childR
		}

		if (*items)[child].Less((*items)[index]) {
			items.swap(index, child)

			index = child
			childL, childR = 2*index+1, 2*index+2
		} else {
			break
		}
	}

	return item
}

func (items *Queue) Push(item Item) {
	// Stick the item as the end of the last level.
	*items = append(*items, item)

	// Bubble up to restore heap property.
	index := len(*items) - 1
	parent := int((index - 1) / 2)
	for parent >= 0 && (*items)[parent].Less(item) == false {
		items.swap(index, parent)

		index = parent
		parent = int((index - 1) / 2)
		if parent == 0 {
			break
		}
	}
}
