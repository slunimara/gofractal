package gofractal

import (
	"sync"
)

type WaitGroup struct {
	ctx    sync.WaitGroup
	length int64
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		ctx:    sync.WaitGroup{},
		length: 0,
	}
}

func (wg WaitGroup) Length() int64 {
	return wg.length
}

func (wg *WaitGroup) Add(delta int) {
	wg.length++
	wg.ctx.Add(delta)
}

func (wg *WaitGroup) Done() {
	wg.ctx.Done()
	wg.length--
}

func (wg *WaitGroup) Wait() {
	wg.ctx.Wait()
}
