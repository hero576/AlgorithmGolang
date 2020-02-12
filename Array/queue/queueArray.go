package queue

type QueueArray struct {
	dataStore    []interface{}
	currentIndex int
}

func NewQueueArray() *QueueArray {
	queue := new(QueueArray)
	queue.dataStore = make([]interface{}, 0)
	queue.currentIndex = 0
	return queue
}

func (q *QueueArray) Size() int {
	return q.currentIndex
}
func (q *QueueArray) Front() interface{} {
	if q.IsEmpty() {
		return nil
	} else {
		return q.dataStore[0]
	}
}
func (q *QueueArray) End() interface{} {
	if q.IsEmpty() {
		return nil
	} else {
		return q.dataStore[q.currentIndex-1]
	}
}
func (q *QueueArray) IsEmpty() bool {
	return q.currentIndex == 0
}
func (q *QueueArray) EnQueue(data interface{}) {
	q.dataStore = append(q.dataStore, data)
	q.currentIndex++
}
func (q *QueueArray) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.dataStore[0]
	q.dataStore = q.dataStore[1:q.currentIndex]
	q.currentIndex--
	return data
}
func (q *QueueArray) Clear() {
	q.dataStore = make([]interface{}, 0)
	q.currentIndex = 0
}
