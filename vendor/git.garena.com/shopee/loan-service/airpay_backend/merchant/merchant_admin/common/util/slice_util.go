package util

import (
	"fmt"
	"strings"
)

var Uint64SliceUtil = (*uint64SliceUtil)(nil)
var Uint32SliceUtil = (*uint32SliceUtil)(nil)
var ObjectSliceUtil = (*objectSliceUtil)(nil)

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

func (r *uint64SliceUtil) ToObjectSlice(list []uint64) []interface{} {
	ret := make([]interface{}, 0, len(list))
	for _, e := range list {
		ret = append(ret, e)
	}
	return ret
}

func (r *uint64SliceUtil) Join(list []uint64) string {
	return strings.Replace(strings.Trim(fmt.Sprint(list), "[]"), " ", ",", -1)
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
