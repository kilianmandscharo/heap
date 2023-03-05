package heap

func New(a []int) []int {
	if a == nil {
		return []int{}
	}
	if len(a) < 2 {
		return a
	}
	for i := len(a)/2 - 1; i >= 0; i-- {
		largest := largestChild(a, i, len(a))
		if largest != i {
			a[largest], a[i] = a[i], a[largest]
			downheap(a, largest, len(a))
		}
	}
	return a
}

func Insert(heap []int, val int) []int {
	heap = append(heap, val)
	upheap(heap)
	return heap
}

func Extract(heap []int) (int, bool) {
	if len(heap) == 0 {
		return 0, false
	}
	val := heap[0]
	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]
	downheap(heap, 0, len(heap))
	return val, true
}

func upheap(heap []int) {
	i := len(heap) - 1
	parent := (i - 1) / 2
	for i > 0 && heap[parent] < heap[i] {
		heap[parent], heap[i] = heap[i], heap[parent]
		i = parent
		parent = (i - 1) / 2
	}
}

func downheap(heap []int, i, n int) {
	largest := largestChild(heap, i, n)
	for largest != i {
		heap[largest], heap[i] = heap[i], heap[largest]
		i = largest
		largest = largestChild(heap, i, n)
	}
}

func largestChild(heap []int, i, n int) int {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	if left < n && heap[left] > heap[largest] {
		largest = left
	}
	if right < n && heap[right] > heap[largest] {
		largest = right
	}
	return largest
}

func isHeap(heap []int, i int) bool {
	left := i*2 + 1
	right := i*2 + 2
	if left >= len(heap) {
		return true
	}
	if right >= len(heap) {
		if heap[i] < heap[left] {
			return false
		}
		return isHeap(heap, left)
	}
	if heap[i] >= heap[left] && heap[i] >= heap[right] {
		return isHeap(heap, left) && isHeap(heap, right)
	}
	return false
}
