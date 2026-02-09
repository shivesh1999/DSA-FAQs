package main

type DLL struct {
	data int
	prev *DLL
	next *DLL
}

func (d *DLL) add(value int) {
	newNode := &DLL{data: value, prev: nil, next: nil}
	if d == nil {
		d = newNode
		return
	}
	current := d
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current
}

func convertArraytoDLL(arr []int) *DLL {
	if len(arr) == 0 {
		return nil
	}
	head := &DLL{data: arr[0], prev: nil, next: nil}
	current := head
	for i := 1; i < len(arr); i++ {
		newNode := &DLL{data: arr[i], prev: nil, next: nil}
		current.next = newNode
		newNode.prev = current
		current = newNode
	}
	return head
}

func main() {
	dll := &DLL{data: 1, prev: nil, next: nil}
	dll.add(2)
	dll.add(3)
	dll.add(4)
	dll.add(5)

	// Traverse the DLL
	current := dll
	for current != nil {
		println(current.data)
		current = current.next
	}

	arr := []int{1, 7, 11, 4, 9}
	head := convertArraytoDLL(arr)
	// Traverse the DLL
	current = head
	for current != nil {
		println(current.data)
		current = current.next
	}
}
