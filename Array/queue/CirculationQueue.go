package queue

type CircleQueue struct {
	data  [queueSize]interface{} // 最多存储queueSize-1个，最后一个表示队列已满
	front int
	rear  int
}

func NewCircleQueue() *CircleQueue {
	queue := new(CircleQueue)
	queue.front = 0
	queue.rear = 0
	return queue
}

func (c *CircleQueue) Size() int {
	return (c.rear - c.front + queueSize) % queueSize
}
func (c *CircleQueue) Front() interface{} {
	if c.IsEmpty() {
		return nil
	}
	return c.data[c.front]
}
func (c *CircleQueue) End() interface{} {
	if c.IsEmpty() {
		return nil
	}
	return c.data[c.rear]
}
func (c *CircleQueue) IsEmpty() bool {
	return c.rear == c.front
}
func (c *CircleQueue) EnQueue(data interface{}) {
	if (c.rear+1)%queueSize == c.front%queueSize {
		panic("queue is full")
	}
	c.data[c.rear] = data
	c.rear = (c.rear + 1) % queueSize
}
func (c *CircleQueue) DeQueue() interface{} {
	if c.IsEmpty() {
		return nil
	}
	res := c.data[c.front]
	c.data[c.front] = 0
	c.front = (c.front + 1) % queueSize
	return res
}
func (c *CircleQueue) Clear() {

}
