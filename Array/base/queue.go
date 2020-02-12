package base

type Queue interface {
	Size() int
	Front() interface{}
	End() interface{}
	IsEmpty() bool
	EnQueue(data interface{})
	DeQueue() interface{}
	Clear()
}
