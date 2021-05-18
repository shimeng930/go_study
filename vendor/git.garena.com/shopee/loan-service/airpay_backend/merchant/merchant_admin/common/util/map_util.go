package util

var StringMapUtil stringMapUtil

type stringMapUtil struct {
}

func (r *stringMapUtil) In(s string, smap map[string]interface{}) bool {
	if smap == nil {
		return false
	}
	_, exist := smap[s]
	return exist
}

func (r *stringMapUtil) NotIn(s string, smap map[string]interface{}) bool {
	return !r.In(s, smap)
}
