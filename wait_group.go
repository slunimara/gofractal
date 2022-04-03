package gofractal

import (
	"sync"
)

type WaitGroup struct {
	ctx    sync.WaitGroup
	length int64
}

// TODO: Documenation
func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		ctx:    sync.WaitGroup{},
		length: 0,
	}
}

// TODO: Documenation
func (wg WaitGroup) Length() int64 {
	return wg.length
}

// TODO: Documenation
func (wg *WaitGroup) Add(delta int) {
	wg.length++
	wg.ctx.Add(delta)
}

// TODO: Documenation
func (wg *WaitGroup) Done() {
	wg.ctx.Done()
	wg.length--
}

// TODO: Documenation
func (wg *WaitGroup) Wait() {
	wg.ctx.Wait()
}
