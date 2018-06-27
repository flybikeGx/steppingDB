package steppingdb

type DeleteType int

var Delete *DeleteType

type mapDiff struct {
	k      string
	v      Value
	new    bool
	delete bool
}

type sliceDiff struct {
	at int // -1 append
	v  Value
}
