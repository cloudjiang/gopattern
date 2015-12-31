package pattern

import "sync"

type Singleton struct {
	once sync.Once
	Inst interface{}
	newFunc func() interface{}
}

func (this *Ston)GetInstance() interface{} {
	this.once.Do(func() {
		this.Inst = this.newFunc()
	})
	return this.Inst
}