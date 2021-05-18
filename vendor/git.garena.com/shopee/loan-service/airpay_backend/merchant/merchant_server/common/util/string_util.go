package common_util

import (
	"fmt"
	"github.com/stretchr/objx"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var StringUtil = (*stringUtil)(nil)

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

func (r *stringUtil) TryConvert2UInt64(str string, tryFunc func(string) uint64) uint64 {
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return tryFunc(str)
	}
	return v
}

func SafeStr(val interface{}) (str string) {
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

func DefaultStr(s string, defaultV string) string {
	if strings.TrimSpace(s) == "" {
		return defaultV
	}
	return s
}

func GetStrValue(s *string, ignoreEmpty bool, defaultStr string) string {
	if s == nil {
		return defaultStr
	}
	if ignoreEmpty && *s == "" {
		return defaultStr
	}
	return *s
}

func FromJsonOrEmpty(json string) objx.Map {
	return objx.MustFromJSON(DefaultStr(json, "{}"))
}

func TrimAllSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}
