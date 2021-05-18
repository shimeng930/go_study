package util

import (
	"reflect"
	"strings"
)

const (
	OrmName_Xorm = "xorm"
	OrmName_JSON = "json"
)

var SQLUtil sqlUtil

type sqlUtil struct {
}

// BuildColumnValueMap 可通过 model 构建反向构建列-值对map, 用于orm update
func (u *sqlUtil) BuildColumnValueMap(obj interface{}, ormName string, ignoreFileds []string) (map[string]interface{}, bool, string) {
	var tagutil tagUtil
	return tagutil.FieldValueMap(obj, ormName, nil, ignoreFileds)
}

func (u *sqlUtil) BuildUnIgnoredColumnList(obj interface{}, ormName string, ignoreFileds []string) ([]string, bool, string) {
	columnMap, got, msg := u.BuildColumnValueMap(obj, ormName, ignoreFileds)
	if !got {
		return nil, false, msg
	}

	list := make([]string, len(columnMap))
	for k, _ := range columnMap {
		list = append(list, k)
	}
	return list, true, ""
}

func (u *sqlUtil) BuildPointedColumnValueMap(obj interface{}, ormName string, pointedFieldAddress ...interface{}) (map[string]interface{}, bool, string) {
	var tagutil tagUtil
	return tagutil.FieldValueMap(obj, ormName, pointedFieldAddress, nil)
}

func (u *sqlUtil) BuildPointedColumnList(obj interface{}, ormName string, pointedFieldAddress ...interface{}) ([]string, bool, string) {
	columnMap, got, msg := u.BuildPointedColumnValueMap(obj, ormName, pointedFieldAddress...)
	if !got {
		return nil, false, msg
	}

	list := make([]string, len(columnMap))
	for k, _ := range columnMap {
		list = append(list, k)
	}
	return list, true, ""
}

type tagUtil struct {
}

// FieldValueMap 根据struct结构成员构建map，其中key为struct结构成员对应的列名(可通过tag指定，否则通过成员名)，值为成员值（不会排除零值成员）
// 如需要获取tag中tagName的子tag，则列名必须以单引号''括起
// excludeTagName,excludeFieldName 如需排除部分field，则field的tag中需要包含指定tag 指定field，空则无排除项
func (u *tagUtil) FieldValueMap(obj interface{}, tagName string, pointedFieldAddress []interface{}, ignoreFileds []string) ( /*tag-value*/ map[string]interface{} /*suc_done*/, bool /*description*/, string) {
	if obj == nil {
		return nil, false, "object is nil"
	}

	objValue := reflect.Indirect(reflect.ValueOf(obj))
	objType := objValue.Type()

	// only support structure
	if objType.Kind() != reflect.Struct {
		return nil, false, "object is not structure"
	}

	tagValueMap := make(map[string]interface{})
	num := objType.NumField()

	// build map
	var key string
	for i := 0; i < num; i++ {
		field := objType.Field(i)
		value := objValue.Field(i)

		// 排除非ignore的field
		// ignore if item is unexported
		if name := field.Name; !(name[0] >= 'A' && name[0] <= 'Z') {
			continue
		}

		// get name(key)
		if key = extractNameFromOrmTag(field.Tag, tagName); key == "" {
			key = extractNameFromFieldName(field.Name)
		}

		if len(pointedFieldAddress) > 0 {
			// 排除非指定的field
			var hit bool
			for _, pointed := range pointedFieldAddress {
				if value.CanAddr() && value.Addr().CanInterface() && value.Addr().Interface() == pointed {
					hit = true
					break // is pointed field
				}
			}
			if !hit {
				continue // ignore this field
			}
		} else {
			// check if is ignore field
			if StrUtil.In(key, ignoreFileds) {
				continue // ignore this field
			}
		}

		tagValueMap[key] = value.Interface()
	}
	return tagValueMap, true, "suc"
}

// 通过tab获取指定orm tag名称中的列名，列名必须以单引号''括起
func extractNameFromOrmTag(tag reflect.StructTag, tagName string) string {
	ormStr := tag.Get(tagName)
	tags := strings.Split(ormStr, " ")
	for _, key := range tags {
		// target name should be round by "'"
		if strings.HasPrefix(key, "'") && strings.HasSuffix(key, "'") {
			return key[1 : len(key)-1]
		}
	}
	return ""
}

// 根据名称大小写拆成"xx_xx_xx"形式的名称
func extractNameFromFieldName(name string) string {
	newstr := make([]rune, 0)
	for idx, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newstr = append(newstr, '_')
			}
			chr -= ('A' - 'a')
		}
		newstr = append(newstr, chr)
	}

	return string(newstr)
}
