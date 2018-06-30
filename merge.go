package steppingdb

func merge(v1, v2 Value) Value {
	switch {
	case isDelete(v2):
		return nil
	case isArray(v1) && isArray(v2):
	case isMap(v1) && isMap(v2):
	default:
		return v2
	}
}

func isMap(value Value) bool {
	_, ok := value.(*mapDiff)
	return ok
}

func isArray(value Value) bool {
	_, ok := value.(*sliceDiff)
	return ok
}

func isDelete(value Value) bool {
	_, ok := value.(*DeleteType)
	return ok
}

func mergeKeys(v1 []string, v2 []*mapDiff) []string {

}
