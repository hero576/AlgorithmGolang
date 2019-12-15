package base

type Iter interface {
	Next() (interface{}, error)
	Remove() error
	GetIndex() int
	HasNext() bool
}

type Iterable interface {
	Iterator() Iter //构造初始化接口
}
