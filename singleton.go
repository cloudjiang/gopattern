package pattern

import "sync"

type Singleton struct {
	once sync.Once
	inst interface{}
	NewFunc func() interface{}
}

func (this *Singleton)GetInstance() interface{} {
	this.once.Do(func() {
		this.inst = this.NewFunc()
	})
	return this.inst
}
