package util

import (
	"reflect"
)

var CommonUtil util

type util struct {
}

func (u *util) ValueOrZero(obj interface{}) reflect.Value {
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return value
	}
	return reflect.Indirect(reflect.ValueOf(obj))
}

// EmptyThenOther 如a为0则返回other,不管other是否为0
func (u *util) EmptyThenOther(a uint64, other uint64) uint64 {
	if a == 0 {
		return other
	}
	return a
}

// target支持的类型array,slice,map
func (u *util) ContainKey(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	default:
		return false
	}

	return false
}

// 有序数组去重
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}