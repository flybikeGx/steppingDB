package steppingdb

type Commit interface {
	Base() Storage
	New() Commit
	Error() error
	Storage
}

func NewCommit(s Storage) Commit {
	return &CommitImpl{
		base:  s,
		patch: NewPatch(),
	}
}

type CommitImpl struct {
	base  Storage
	patch *Patch
	err   error
}

func (n *CommitImpl) Base() Storage {
	return n.base
}
func (n *CommitImpl) New() Commit {
	return &CommitImpl{
		base:  n,
		patch: NewPatch(),
	}
}
func (n *CommitImpl) Get(key string) Value { // 返回历次改变的结果
	val1 := n.base.Get(key)
	val2 := n.patch.Get(key)
	return merge(val1, val2)
}
func (n *CommitImpl) Set(key string, value Value) { // 直接写入patch，仅限基本类型
	n.patch.Set(key, value)
}
func (n *CommitImpl) HMGet(key, field string) Value { //
	val1 := n.base.HMGet(key, field)
	val2 := n.patch.HMGet(key, field)
	return merge(val1, val2)

}
func (n *CommitImpl) HMSet(key, field string, value Value) { // 在patch中写入mapdiff
	n.patch.HMSet(key, field, value)
}
func (n *CommitImpl) HMLen(key string) int { // 通过遍历mapdiff来计算最终的len
	b := n.base.HMLen(key)
	d := n.patch.HMLen(key)
	return b + d
}
func (n *CommitImpl) HMKeys(key string) []string { // 遍历mapDiff计算最终keys
	keys := n.base.HMKeys(key)
	dkeys := n.patch.HMKeys(key)
	return mergeKeys(keys, dkeys)
}
func (n *CommitImpl) ArrayGet(key string, i int) Value {
	val1 := n.base.ArrayGet(key, i)
	val2 := n.patch.ArrayGet(key, i)
	return merge(val1, val2)
}
func (n *CommitImpl) ArraySet(key string, i int, v Value) {
	n.patch.ArraySet(key, i, v)
}
func (n *CommitImpl) ArrayLen(key string) int {
	l := n.base.ArrayLen(key)
	d := n.patch.ArrayLen(key)
	return l + d
}
func (n *CommitImpl) Error() error {
	return n.err
}
func MergeBase(c *CommitImpl) Commit {
	// ...
	return c.base.(*CommitImpl)
}
