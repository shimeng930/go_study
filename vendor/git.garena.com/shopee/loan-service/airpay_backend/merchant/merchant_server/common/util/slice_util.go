package common_util

var Uint64SliceUtil = (*uint64SliceUtil)(nil)
var Uint32SliceUtil = (*uint32SliceUtil)(nil)
var ObjectSliceUtil = (*objectSliceUtil)(nil)
var StringSliceUtil = (*stringSliceUtil)(nil)

type uint64SliceUtil struct {
}

func (r *uint64SliceUtil) In(e uint64, list ...uint64) bool {
	for _, v := range list {
		if e == v {
			return true
		}
	}
	return false
}

func (r *uint64SliceUtil) NotIn(e uint64, list ...uint64) bool {
	return !r.In(e, list...)
}

func (r *uint64SliceUtil) ToUint32(in []uint64) []uint32 {
	var res []uint32
	for _, e := range in {
		res = append(res, uint32(e))
	}
	return res
}

func (r *uint64SliceUtil) IsAnyZero(list ...uint64) bool {
	for _, e := range list {
		if e == 0 {
			return true
		}
	}
	return false
}

func  (r *uint64SliceUtil) FirstGreaterThanZero(list ...uint64) uint64 {
	for _, e := range list {
		if e > 0 {
			return e
		}
	}
	return list[len(list)-1]
}

func (r *uint64SliceUtil) ToObjectSlice(list []uint64) []interface{} {
	ret := make([]interface{}, 0, len(list))
	for _, e := range list {
		ret = append(ret, e)
	}
	return ret
}

type uint32SliceUtil struct {
}

func (r *uint32SliceUtil) In(e uint32, list ...uint32) bool {
	for _, v := range list {
		if e == v {
			return true
		}
	}
	return false
}

type objectSliceUtil struct {
}

func (r *objectSliceUtil) IsAnyNotNil(list ...interface{}) bool {
	for _, e := range list {
		if e != nil {
			return true
		}
	}
	return false
}

// 去重
func RemoveRepeatedIds(ids []uint64) []uint64 {
	idMap := make(map[uint64]bool)
	var result []uint64
	for _, v := range ids {
		if _, ok := idMap[v]; !ok {
			result = append(result, v)
			idMap[v] = true
		}
	}
	return result
}

type stringSliceUtil struct {
}

func (r *stringSliceUtil) In(e string, list ...string) bool {
	for _, v := range list {
		if e == v {
			return true
		}
	}
	return false
}

func (r *stringSliceUtil) NotIn(e string, list ...string) bool {
	return !r.In(e, list...)
}

// a is sub in b
func (r *stringSliceUtil) IsSub(a, b []string) bool {
	for _, e := range a {
		if r.NotIn(e, b...) {
			return false
		}
	}
	return true
}

// AnyNotIn is a any not b
func (r *stringSliceUtil) AnyNotIn(a, b []string) bool {
	for _, e := range a {
		if r.NotIn(e, b...) {
			return true
		}
	}
	return false
}

func (r *stringSliceUtil) Union(a, b []string) []string {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

	add := func(result, list []string) []string {
		for _, e := range a {
			if r.NotIn(e, result...) {
				result = append(result, e)
			}
		}
		return result
	}

	t := make([]string, 0, len(a)+len(b))
	t = add(t, a)
	t = add(t, b)

	return t
}

func (r *stringSliceUtil) Intersection(a, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return make([]string, 0, 0)
	}

	max := func(a1, b1 int) int {
		if a1 >= b1 {
			return a1
		}
		return b1
	}

	t := make([]string, 0, max(len(a), len(b)))

	for _, e := range a {
		if r.In(e, b...) {
			t = append(t, e)
		}
	}

	return t
}

func (r *stringSliceUtil) Filter(list []string, predict func(string) bool) []string {
	if len(list) == 0 {
		return make([]string, 0, 0)
	}

	ret := make([]string, 0)
	for i := range list {
		if predict(list[i]) {
			ret = append(ret, list[i])
		}
	}

	return ret
}

func (r *stringSliceUtil) PredictFuncOfIn(list []string) func(string) bool {
	return func(s string) bool {
		return r.In(s, list...)
	}
}

func (r *stringSliceUtil) PredictFuncOfNotIn(list []string) func(string) bool {
	return func(s string) bool {
		return r.NotIn(s, list...)
	}
}
