package singleton

import "sync"

type singleton struct {
	count int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		if instance == nil {
			instance = &singleton{}
		}
	})

	return instance
}

func (inst *singleton) AddOne() {
	inst.count++
}

func (inst *singleton) GetCount() int {
	return inst.count
}
