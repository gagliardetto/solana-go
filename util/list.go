package util

type linkedListNode[T any] struct {
	value T
	next  *linkedListNode[T]
	prev  *linkedListNode[T]
}

type LinkedList[T any] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]
	Size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{Size: 0}
}

func (l *LinkedList[T]) Append(item T) {
	node := &linkedListNode[T]{value: item}
	l.Size++
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	} else {
		old := l.tail
		old.next = node
		node.prev = old
		l.tail = node
		return
	}
}

func (l *LinkedList[T]) Iterate(callback func(index int, value T) error) error {
	i := 0
	var err error
out:
	for node := l.head; node != nil; node = node.next {
		err = callback(i, node.value)
		if err != nil {
			break out
		}
		i++
	}
	return err
}

func (l *LinkedList[T]) Array() []T {
	ans := make([]T, l.Size)
	l.Iterate(func(index int, value T) error {
		ans[index] = value
		return nil
	})
	return ans
}
