package steppingdb

type Commit interface {
	Base() Storage
	New() Commit
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

}
func (n *CommitImpl) HMSet(key, field string, value Value) { // 在patch中写入mapdiff

}
func (n *CommitImpl) HMLen(key string) int { // 通过遍历mapdiff来计算最终的len

}
func (n *CommitImpl) HMKeys(key string) []string { // 遍历mapDiff计算最终keys

}
func (n *CommitImpl) ArrayGet(key string, i int) Value {

}
func (n *CommitImpl) ArraySet(key string, i int) Value {

}
func (n *CommitImpl) ArrayLen(key string) int {

}
