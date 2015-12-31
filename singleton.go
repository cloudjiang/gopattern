package pattern

import "sync"

type Singleton struct {
	once sync.Once
	Inst interface{}
	NewFunc func() interface{}
}

func (this *Singleton)GetInstance() interface{} {
	this.once.Do(func() {
		this.Inst = this.NewFunc()
	})
	return this.Inst
}
