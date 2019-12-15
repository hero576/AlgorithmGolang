package base

type List interface {
	Size() int
	Get(int) (interface{}, error)
	Set(int, interface{}) error
	Insert(int, interface{}) error
	Append(interface{})
	Clear()
	Delete(int) error
	String() string
}

type ArrayListIter interface {
	Iterator() Iter
}
