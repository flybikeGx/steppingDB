package steppingdb

import "sync"

type Patch struct {
	rw   sync.RWMutex
	data map[string]Value
}

func NewPatch() *Patch {

}
func (p *Patch) Get(key string) Value {
	return p.data[key]
}
func (p *Patch) Set(key string, value Value) {

}
func (p *Patch) HMGet(key, field string) Value {

}
func (p *Patch) HMSet(key, field string, value Value) {

}
func (p *Patch) HMLen(key string) int {

}
func (p *Patch) HMKeys(key string) []string {

}
func (p *Patch) ArrayGet(key string, i int) Value {

}
func (p *Patch) ArraySet(key string, i int) Value {

}
func (p *Patch) ArrayLen(key string) int {

}
