package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var StrUtil = (*stringUtil)(nil)

type stringUtil struct {
}

func (r *stringUtil) IsBlank(str string) bool {
	for _, e := range str {
		if !unicode.IsSpace(e) {
			return false
		}
	}
	return true
}

func (r *stringUtil) IsAnyBlank(list ...string) bool {
	for _, e := range list {
		if r.IsBlank(e) {
			return true
		}
	}
	return false
}

func (r *stringUtil) GetValidValue(s *string, defValue string) string {
	if s == nil {
		return defValue
	}
	return *s
}

func (r *stringUtil) StringOrDefault(s string, defaultV string) string {
	if strings.TrimSpace(s) == "" {
		return defaultV
	}
	return s
}

func (r *stringUtil) SafeStr(val interface{}) (str string) {
	defer func() {
		if r := recover(); r != nil {
			str = fmt.Sprintf("<panic>:%s", r)
		}
	}()
	if val == nil {
		return "<nil>"
	}
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)
	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "<nil>"
		}
		return fmt.Sprint(v.Elem())
	} else {
		return fmt.Sprint(v)
	}
}

func (r *stringUtil) SplitToUint64(str *string) ([]uint64, error) {
	ret, err := r.SplitToSliceType(*str, ",", []uint64{})
	if err != nil {
		return nil, err
	}

	result, ok := ret.([]uint64)
	if !ok {
		return nil, fmt.Errorf("unknow error")
	}
	return result, nil
}

func (r *stringUtil) SplitToString(str *string) ([]string, error) {
	result := strings.Split(*str, ",")
	return result, nil
}

func (r *stringUtil) In(str string, array []string) bool {
	for _, elem := range array {
		if str == elem {
			return true
		}
	}
	return false
}

func (r *stringUtil) SplitToSliceType(str string, sep string, obj interface{}) (interface{}, error) {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Slice {
		// obj must be a slice
		return nil, fmt.Errorf("obj just support slice")
	}

	objElemType := objType.Elem()
	covFunc, err := _getValueConvertFunc(objElemType)
	if err != nil {
		return nil, err
	}

	strElem := strings.Split(str, sep)
	result := reflect.MakeSlice(objType, 0, len(strElem))
	for _, elem := range strElem {
		if len(elem) == 0 {
			continue
		}
		// convert
		v, err := covFunc(elem)
		if err != nil {
			return nil, err
		}
		result = reflect.Append(result, reflect.ValueOf(v))
	}
	return result.Interface(), nil
}

func _getValueConvertFunc(objType reflect.Type) (func(string) (interface{}, error), error) {
	uint64ConvFunc := func(str string) (interface{}, error) {
		v, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		return uint64(v), nil
	}
	uint32ConvFunc := func(str string) (interface{}, error) {
		v, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		return uint32(v), nil
	}
	stringConvFunc := func(str string) (interface{}, error) {
		return str, nil
	}

	switch objType.Kind() {
	case reflect.Uint64:
		return uint64ConvFunc, nil
	case reflect.Uint32:
		return uint32ConvFunc, nil
	case reflect.String:
		return stringConvFunc, nil
		// 如需支持其他类型（必须可转），在此处添加
	default:
		// not support type
		return nil, fmt.Errorf("unsport elem type")
	}
}
