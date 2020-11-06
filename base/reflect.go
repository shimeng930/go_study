package base

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Foo struct {
	A int
	B string
	C int
}

type Foo1 struct {
	A int
	B string
	C int
	D []int
	E bool
}

type Foo2 struct {
	A string
	B string
	C int
	D []uint16
	E bool
}

func RefRun() {
	greeting := "hello"
	f := Foo{A: 10, B: "Salutations", C:0}
	//source := &Foo1{A: 2, B:"new B", D: []int{1,23,4}}

	dest := &Foo2{A: "2", B:"new B", D: []uint16{1,23,4}}
	//reflect.Copy(reflect.ValueOf(dest), reflect.ValueOf(f))

	//err := CopyProperties(source, &f, []string{"A", "B", "C", "D", "E"})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	gVal := reflect.ValueOf(greeting)
	// not a pointer so all we can do is read it
	fmt.Println(gVal.Interface())

	gpVal := reflect.ValueOf(&greeting)
	// it’s a pointer, so we can change it, and it changes the underlying variable
	gpVal.Elem().SetString("goodbye")
	fmt.Println(greeting)

	//fType := reflect.TypeOf(f)
	//fVal := reflect.New(fType)
	//fVal.Elem().Field(0).SetInt(20)
	//fVal.Elem().Field(1).SetString("Greetings")
	//f2 := fVal.Elem().Interface().(Foo)
	//fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)

	fv := reflect.ValueOf(dest)
	if fv.Kind() != reflect.Ptr {
		fmt.Printf("err kind")
	}

	tv, ok := reflect.TypeOf(f).FieldByName("AA")
	fmt.Println(tv, ok)

	sv := fv.Elem()

	bv := sv.FieldByName("D")
	_convertSlice2String(bv)
	if bv.IsValid() {
		fmt.Println(bv.Interface())
		bv.Set(reflect.ValueOf("new v"))
		fmt.Println(f)
	}

}

func CopyProperties(sourceData, targetData interface{}, specificFields []string) error {
	source := reflect.ValueOf(sourceData)
	if source.Kind() != reflect.Ptr {
		return errors.New("[source]pls use ptr kind")
	}
	target := reflect.ValueOf(targetData)
	if target.Kind() != reflect.Ptr {
		return errors.New("[target]pls use ptr kind")
	}

	sourceValue := source.Elem()
	targetValue := target.Elem()

	sourceType := sourceValue.Type()
	if sourceType.Kind() != reflect.Struct {
		return errors.New("source must be struct type")
	}
	targetType := targetValue.Type()
	if targetType.Kind() != reflect.Struct {
		return errors.New("target must be struct type")
	}

	for _, item := range specificFields {
		if _, ok := sourceType.FieldByName(item); !ok {
			continue
		}
		if _, ok := targetType.FieldByName(item); !ok {
			continue
		}

		sourceFieldValue := sourceValue.FieldByName(item)
		targetFieldValue := targetValue.FieldByName(item)
		switch targetFieldValue.Kind() {
		case reflect.String:
			targetFieldValue.SetString(sourceFieldValue.String())
		case reflect.Int, reflect.Int32, reflect.Int64:
			targetFieldValue.SetInt(sourceFieldValue.Int())
		case reflect.Uint32, reflect.Uint64:
			targetFieldValue.SetUint(sourceFieldValue.Uint())
		default:
			return errors.New("not support fieldType")
		}
	}

	return nil
}

func _convertSlice2String(sourceFieldValue reflect.Value) error {
	var strs []string
	for i := 0; i < sourceFieldValue.Len(); i++ {
		value := sourceFieldValue.Index(i)
		switch value.Kind() {
		case reflect.String:
			strs = append(strs, value.String())
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			strs = append(strs, strconv.Itoa(int(value.Int())))
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			strs = append(strs, strconv.Itoa(int(value.Uint())))
		default:
			return errors.New("not support array itemType")
		}
	}
	fmt.Println(strs)
	return nil
}

func ConvertOpAndValue(inList, excludeList interface{}) (int, interface{}) {
	var valueKind = reflect.TypeOf(inList).Kind()
	if valueKind != reflect.Slice && valueKind != reflect.Array {
		return 0, nil
	}

	if inList == nil && excludeList == nil {
		return 0, nil
	}
	if inList != nil && excludeList == nil {
		return 1, inList
	}
	if excludeList != nil && excludeList == nil {
		return 2, excludeList
	}

	inValue := reflect.ValueOf(inList)
	excludeValue := reflect.ValueOf(excludeList)
	var inLen = inValue.Len()
	var valueMap = make(map[interface{}]bool, inLen)
	for i := 0; i < inLen; i++ {
		valueMap[inValue.Index(i).Interface()] = true
	}
	for i := 0; i < excludeValue.Len(); i++ {
		if valueMap[excludeValue.Index(i).Interface()] {
			valueMap[excludeValue.Index(i).Interface()] = false
		}
	}

	var actualList []interface{}
	for k, v := range valueMap {
		if v {
			actualList = append(actualList, k)
		}
	}
	if actualList != nil {
		return 1, actualList
	} else {
		return 0, nil
	}
}

type OutletCountInRule struct {
	MerchantRule 	*MerchantRule `json:"merchant_rule,omitempty"`
	OutletRule 		*OutletRule `json:"outlet_rule,omitempty"`
}

type MerchantRule struct {
	MerchantIdList 			[]uint64 `json:"merchant_id_list,omitempty"`
	MerchantIdListExclude 	[]uint64 `json:"merchant_id_list_exclude,omitempty"`
	TypeList 				[]uint32 `json:"type_list,omitempty"`
	TypeListExclude 		[]uint32 `json:"type_list_exclude,omitempty"`
	StateList 				[]string `json:"state_list,omitempty"`
	StateListExclude 		[]string `json:"state_list_exclude,omitempty"`
	DistrictList 			[]string `json:"district_list,omitempty"`
	DistrictListExclude 	[]string `json:"district_list_exclude,omitempty"`
	VatChoiceList 			[]uint32 `json:"vat_choice_list,omitempty"`
	VatListExclude 			[]uint32 `json:"vat_choice_list_exclude,omitempty"`
	WhtChoiceList 			[]uint32 `json:"wht_choice_list,omitempty"`
	WhtListExclude 			[]uint32 `json:"wht_choice_list_exclude,omitempty"`
	OnlyActive 				bool `json:"only_active,omitempty"`
}

type OutletRule struct {
	OutletIdList 		[]uint64 `json:"outlet_id_list,omitempty"`
	OutletIdListExclude []uint64 `json:"outlet_id_list_exclude,omitempty"`
	CategoryList 		[]uint32 `json:"category_list,omitempty"`
	CategoryListExclude []uint32 `json:"category_list_exclude,omitempty"`
	StateList 			[]string `json:"state_list,omitempty"`
	StateListExclude 	[]string `json:"state_list_exclude,omitempty"`
	DistrictList 		[]string `json:"district_list,omitempty"`
	DistrictListExclude []string `json:"district_list_exclude,omitempty"`
	OnlyActive 			bool `json:"only_active,omitempty"`
}

func ConvertToPbReq(rulr *OutletCountInRule) (bool) {
	merchantRule := rulr.MerchantRule
	if merchantRule != nil {
		if merchantRule.MerchantIdList != nil || merchantRule.MerchantIdListExclude != nil {
			op, value := convertOpAndValue(merchantRule.MerchantIdList, merchantRule.MerchantIdListExclude, reflect.Uint64)
			merchantRule.MerchantIdList = value.([]uint64)
			fmt.Println(op, value.([]uint64))
		}
		if merchantRule.TypeList != nil || merchantRule.TypeListExclude != nil {
			op, value := convertOpAndValue(merchantRule.TypeList, merchantRule.TypeListExclude, reflect.Uint32)
			fmt.Println(op, value)
		}
		if merchantRule.StateList != nil || merchantRule.StateListExclude != nil {
			op, value := convertOpAndValue(merchantRule.StateList, merchantRule.StateListExclude, reflect.String)
			fmt.Println(op, value)
		}
		if merchantRule.DistrictList != nil || merchantRule.DistrictListExclude != nil {
			op, value := convertOpAndValue(merchantRule.DistrictList, merchantRule.DistrictListExclude, reflect.String)
			fmt.Println(op, value)
		}
		if merchantRule.VatChoiceList != nil || merchantRule.VatListExclude != nil {
			op, value := convertOpAndValue(merchantRule.VatChoiceList, merchantRule.VatListExclude, reflect.Uint32)
			fmt.Println(op, value)
		}
		if merchantRule.WhtChoiceList != nil || merchantRule.WhtListExclude != nil {
			op, value := convertOpAndValue(merchantRule.WhtChoiceList, merchantRule.WhtListExclude, reflect.Uint32)
			fmt.Println(op, value)
		}
	}

	outletRule := rulr.OutletRule
	if outletRule != nil {
		if outletRule.OutletIdList != nil || outletRule.OutletIdListExclude != nil {
			op, value := convertOpAndValue(outletRule.OutletIdList, outletRule.OutletIdListExclude, reflect.Uint64)
			fmt.Println(op, value)
		}
		if outletRule.CategoryList != nil || outletRule.CategoryListExclude != nil {
			op, value := convertOpAndValue(outletRule.CategoryList, outletRule.CategoryListExclude,reflect.Uint32)
			fmt.Println(op, value)
		}
		if outletRule.StateList != nil || outletRule.StateListExclude != nil {
			op, value := convertOpAndValue(outletRule.StateList, outletRule.StateListExclude, reflect.String)
			fmt.Println(op, value)
		}
		if outletRule.DistrictList != nil || outletRule.DistrictListExclude != nil {
			op, value := convertOpAndValue(outletRule.DistrictList, outletRule.DistrictListExclude, reflect.String)
			fmt.Println(op, value)
		}
	}

	return true
}

// 转换参数里的inList 和 excludeList
func convertOpAndValue(inList, excludeList interface{}, vKind reflect.Kind) (int, interface{}) {
	var valueKind = reflect.TypeOf(inList).Kind()
	if valueKind != reflect.Slice && valueKind != reflect.Array {
		return 0, nil
	}

	inValue := reflect.ValueOf(inList)
	excludeValue := reflect.ValueOf(excludeList)
	var inLen = inValue.Len()
	var excludeLen = excludeValue.Len()
	if inLen == 0 && excludeLen == 0 {
		return 0, nil
	}
	if inLen != 0 && excludeLen == 0 {
		return 1, inList
	}
	if inLen == 0 && excludeLen != 0 {
		return 2, excludeList
	}

	var valueMap = make(map[interface{}]bool, excludeLen)
	for i := 0; i < excludeLen; i++ {
		valueMap[excludeValue.Index(i).Interface()] = true
	}
	for i := 0; i < inLen; i++ {
		if valueMap[inValue.Index(i).Interface()] {
			valueMap[inValue.Index(i).Interface()] = false
		}
	}

	var actualList interface{}
	switch vKind {
	case reflect.String:
		actualList = toStringList(valueMap)
	case reflect.Uint32:
		actualList = toUint32List(valueMap)
	case reflect.Uint64:
		actualList = toUint64List(valueMap)
	case reflect.Int32:
		actualList = toInt32List(valueMap)
	case reflect.Int64:
		actualList = toInt64List(valueMap)
	}
	if reflect.ValueOf(actualList).IsNil() {
		return 0, nil
	} else {
		return 2, actualList
	}
}

func toStringList(valueMap map[interface{}]bool) []string {
	var rs []string
	for k, v := range valueMap {
		if v {
			rs = append(rs, k.(string))
		}
	}
	return rs
}

func toUint64List(valueMap map[interface{}]bool) []uint64 {
	var rs []uint64
	for k, v := range valueMap {
		if v {
			rs = append(rs, k.(uint64))
		}
	}
	return rs
}

func toUint32List(valueMap map[interface{}]bool) []uint32 {
	var rs []uint32
	for k, v := range valueMap {
		if v {
			rs = append(rs, k.(uint32))
		}
	}
	return rs
}

func toInt64List(valueMap map[interface{}]bool) []int64 {
	var rs []int64
	for k, v := range valueMap {
		if v {
			rs = append(rs, k.(int64))
		}
	}
	return rs
}

func toInt32List(valueMap map[interface{}]bool) []int32 {
	var rs []int32
	for k, v := range valueMap {
		if v {
			rs = append(rs, k.(int32))
		}
	}
	return rs
}