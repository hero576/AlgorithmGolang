package queue

type Node struct {
	data  interface{}
	pNext *Node
}
type QueueLink struct {
	front *Node
	rear  *Node
}

func NewQueueLink() *QueueLink {
	qlk := new(QueueLink)
	return qlk
}

func (q *QueueLink) Size() int {
	pNext := q.front
	length := 0
	for pNext.pNext != nil {
		pNext = pNext.pNext
		length++
	}
	return length
}
func (q *QueueLink) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.front.data
}
func (q *QueueLink) End() interface{} {
	if q.IsEmpty() {
		return nil
	}
	pNext := q.front
	var value interface{}
	if pNext.pNext != nil {
		pNext = pNext.pNext
		value = pNext.data
	}
	return value
}
func (q *QueueLink) IsEmpty() bool {
	return q.front == nil
}
func (q *QueueLink) EnQueue(data interface{}) {
	newNode := &Node{data, nil}
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.pNext = newNode
		q.rear = q.rear.pNext
	}

}
func (q *QueueLink) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	newNode := q.front
	if q.front == q.rear {
		q.Clear()
	} else {
		q.front = q.front.pNext
	}
	return newNode.data
}
func (q *QueueLink) Clear() {
	q.front = nil
	q.rear = nil
}
