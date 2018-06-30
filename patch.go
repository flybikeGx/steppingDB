package steppingdb

import (
	"fmt"
	"sync"
)

type Patch struct {
	rw        sync.RWMutex
	data      map[string]Value
	hmKeys    map[string][]*mapDiff
	hmLens    map[string]int
	arrayLens map[string]int
	err       error
}

func NewPatch() *Patch {
	return &Patch{
		data:      make(map[string]Value),
		hmKeys:    make(map[string][]*mapDiff),
		hmLens:    make(map[string]int),
		arrayLens: make(map[string]int),
		err:       nil,
	}

}
func (p *Patch) Get(key string) Value {
	return p.data[key]
}
func (p *Patch) Set(key string, value Value) {
	p.err = checkBase(value)
	if p.err != nil {
		return
	}
	p.data[key] = value
}
func (p *Patch) HMGet(key, field string) Value {
	diffmap, ok := p.data[key]
	if !ok {
		return nil
	}
	return diffmap.(map[string]*mapDiff)[field]
}
func (p *Patch) HMSet(key, field string, value Value) {
	dv, vok := value.(*mapDiff)
	if !vok {
		p.err = fmt.Errorf("hm set not mapdiff: %v", value)
	}
	diffmap0, ok := p.data[key]
	if !ok {
		p.data[key] = make(map[string]*mapDiff)
	} else {
		_, typeOK := diffmap0.(map[string]*mapDiff)
		if !typeOK {
			p.data[key] = make(map[string]*mapDiff)
		}
	}
	p.data[key].(map[string]*mapDiff)[field] = dv
	p.data[key] = append(p.hmKeys[key], dv)
	if dv.new {
		p.hmLens[key] = p.hmLens[key] + 1
	} else {
		p.hmLens[key] = p.hmLens[key] - 1
	}

}
func (p *Patch) HMLen(key string) int {
	return p.hmLens[key]
}
func (p *Patch) HMKeys(key string) []*mapDiff {
	return p.hmKeys[key]
}
func (p *Patch) ArrayGet(key string, i int) Value {
	arr, ok := p.data[key]
	if !ok {
		return nil
	}
	return arr.(map[int]*sliceDiff)[i]
}
func (p *Patch) ArraySet(key string, i int, value Value) {

	dv, vok := value.(*sliceDiff)
	if !vok {
		p.err = fmt.Errorf("array set not slicediff: %v", value)
	}
	diffmap0, ok := p.data[key]
	if !ok {
		p.data[key] = make(map[int]*sliceDiff)
	} else {
		_, typeOK := diffmap0.(map[int]*sliceDiff)
		if !typeOK {
			p.data[key] = make(map[int]*sliceDiff)
		}
	}
	if i == -1 {
		p.arrayLens[key] = dv.at
		return
	}
	p.data[key].(map[int]*sliceDiff)[i] = dv
}
func (p *Patch) ArrayLen(key string) int {
	resize, ok := p.arrayLens[key]
	if !ok {
		return 0
	}
	return resize
}
func (p *Patch) Error() error {
	return p.err
}
