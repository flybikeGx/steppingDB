package steppingdb

import (
	"fmt"
	"sync"
)

type Patch struct {
	rw   sync.RWMutex
	data map[string]Value
	err  error
}

func NewPatch() *Patch {

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
	p.err = checkBase(value)
	if p.err != nil {
		return
	}

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

}
func (p *Patch) HMLen(key string) int {
	c := 0
	for _, v := range p.data[key].(map[string]*mapDiff) {
		if v.new {
			c++
		} else if v.delete {
			c--
		}
	}
	return c

}
func (p *Patch) HMKeys(key string) []*mapDiff {
	rtn := make([]*mapDiff, 0)
	for _, v := range p.data[key].(map[string]*mapDiff) {
		rtn = append(rtn, v)
	}
}
func (p *Patch) ArrayGet(key string, i int) Value {
	arr, ok := p.data[key]
	if !ok {
		return nil
	}
	return arr.(map[int]*sliceDiff)[i]
}
func (p *Patch) ArraySet(key string, i int, value Value) {
	p.err = checkBase(value)
	if p.err != nil {
		return
	}

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
	p.data[key].(map[int]*sliceDiff)[i] = dv
}
func (p *Patch) ArrayLen(key string) int {
	c := 0
	for _, v := range p.data[key].(map[string]*sliceDiff) {
		if v.at == -1 {
			c++
		}
	}
	return c
}
func (p *Patch) Error() error {
	return p.err
}
