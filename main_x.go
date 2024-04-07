package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"go_study/base"
	"go_study/excel"
	"go_study/img"
	"io/ioutil"
	"math"
	"net/http"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func getPerson(p *Person) {
	fmt.Println(p)
}

func (p Person) changeName() {
	p.Name = "shi"
}

func (p *Person) changeNamePtr() {
	p.Name = "shi"
}

type MerchantExtra struct {
	PushNotifyChoice uint32    `json:"push_notify_choice"`
	WhtChoice        FeeChoice `json:"wht_choice"`
	VatChoice        FeeChoice `json:"vat_choice"`
}

type FeeChoice struct {
	Current   uint32 `json:"current"`
	Last      uint32 `json:"last"`
	ValidTime uint64 `json:"valid_time"`
}

type SettlementExtra struct {
	ResultCode   string `json:"result_code"`
	ReportStatus string `json:"report_status"`
	ErrorCode    string `json:"error_code"`
}

const currencyAmountMultiple = 1000000
const runAtHour = 52
const Interval = 5

type Action struct {
	Name string
	obj  []int
}

func (a *Action) AddObj(objs ...int) {
	a.obj = append(a.obj, objs...)
}

func GetMonthGapTime(curTime time.Time, gapNum int) time.Time {
	tempTime := curTime.AddDate(0, gapNum, 0)

	curYear, curMonth, _ := curTime.Date()
	tempYear, tempMonth, _ := tempTime.Date()
	if (tempYear-curYear)*12+int(tempMonth) != int(curMonth)+gapNum {
		tempTime = tempTime.AddDate(0, 0, -1)
	}
	return tempTime
}

func sortDesc(ps []Person) {
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Age > ps[j].Age
	})
	return
}

func fmt_t(person *Person) {
	fmt.Println(person.Age, person.Name)
}

const THCurrencyAmountMultiple = 10000
const VNCurrencyAmountMultiple = 1000000

func parseAmount(amount int64) int64 {
	var currencyAmountMultiple float64
	currencyAmountMultiple = THCurrencyAmountMultiple
	actualAmount := float64(amount) / currencyAmountMultiple
	actualAmount = math.Floor(actualAmount + 0.5)
	return int64(actualAmount * currencyAmountMultiple)
}

func func3() {
	var i = 1
	defer fmt.Printf("defer print i=%d\\n", i)
}

func func1() {
	func3()
	func2()
}

func func2() {
	time.Sleep(time.Microsecond * 100)
	fmt.Println("func2 done")
}

func arrparam(arramt *[6]int64) {
	arramt[0] = 1
	arramt[1] = 0
	arramt[2] = 2
	arramt[3] = 3
	arramt[4] = 4
}

func httpTest() {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["id"] = "1234"
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "url", bytes.NewReader(bytesData))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func loopTest() {
	var arr [3]int
	var i int
	for i = 0; i <= 3; i++ {
		arr[i] = 0
		fmt.Printf("arr[i]=%d, i=%d, addr=%p, iaddr=%p\\n", arr[i], i, &arr[i], &i)
	}
}

func linkListTest() {
	node := base.NewListNode()
	head := node.InitList(5)
	node.Print(head)

	head = node.Reverse(head)
	node.Print(head)

	node.GetLastKNode(8, head)

	head = node.SortList(head)
	node.Print(head)

	head1 := node.InitList(5)
	node.Print(head1)
	head1 = node.SortList(head1)
	head1 = node.Merge2List(head1, head)
	node.Print(head1)

}

func personUpdate(ps []Person) {
	for _, item := range ps {
		item.Age = 15
	}
}

func CheckSigBytes(src, key []byte) string {
	checksum := fmt.Sprintf("%X", sha256.Sum256(append(src, key...)))
	return checksum
}

func divideSlice() {
	s := 10000

	fmt.Println(s << 2)
	var ta = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var cnt = 1
	for lastIndex := 0; lastIndex < len(ta); cnt += 1 {
		st := lastIndex
		lastIndex = st + 3
		if lastIndex > len(ta) {
			lastIndex = len(ta)
		}
		temp := ta[st:lastIndex]
		fmt.Println(cnt, temp)
	}
}

func compare() {
	now := time.Now().Unix()
	fmt.Println(now)

	var all = 10000000
	var a = 0
	for i := 0; i < all; i++ {
		a += 10000
	}
	fmt.Println(time.Now().Unix() - now)

	now = time.Now().Unix()
	a = 0
	for i := 1; i < all; i++ {
		a = i * 10000
	}
	fmt.Println(time.Now().Unix() - now)
}

func print(arr *Person) {
	fmt.Println(arr.Name)
}

//func lengthOfLongestSubstring(string s) {
//	tempMap := make(map[int]);//记录下窗口中出现过的字符
//	int longest = 0;
//	int i = 0, j = 0;
//	while(j < s.size()) {
//		if(!map[s[j]]) {//数组中没有当前的字符串，找到一个可能的解。
//		longest = max(longest, (j - i + 1));
//		} else {//数组中有记录当前的字符串，出现重复。
//		while(i < j && map[s[j]] > 0) {//将窗口左侧往前滑动，直到当前窗口内没有重复字符串为止。
//		map[s[i]]--;
//		i++;
//		}
//		}
//
//		map[s[j]] ++;
//		j++;
//
//	}
//
//	return longest;
//}

func jump(arr []int) bool {
	size := len(arr)

	maxr := 0
	for i := 0; i < size; i++ {
		if i <= maxr {
			if i+arr[i] > maxr {
				maxr = i + arr[i]
			}
			if maxr >= (size - 1) {
				return true
			}
		}
	}
	return false
}

func rec() error {
	if err := recover(); err != nil {
		if e, ok := err.(error); ok {
			return e
		}
		return fmt.Errorf("%v", err)
	}
	return nil
}

func doAtest() {
	defer func() {
		//if err := recover(); err != nil {
		//	if e, ok := err.(error); ok {
		//		fmt.Println(e)
		//	}
		//	fmt.Errorf("%v", err)
		//}
		fmt.Println(rec())
	}()
	fmt.Println("do a test")
	ThisPanic()
}

func ThisPanic() {
	panic("this is panic")
}

func main() {
	doAtest()
	base.TestReadCost()
	return

	//excel.ReadDir()
	excel.CheckOutletId()
	excel.ExcelDo()
	return

	var reg = "(([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})(((0[013578]|1[02])(0[0-9]|[12][0-9]|3[01]))|((0[469]|11)(0[1-9]|[12][0-9]|30))|(02(0[1-9]|[1][0-9]|2[0-8]))))|((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))0229)"
	r, _ := regexp.Compile(reg)
	fmt.Println(r.MatchString("20200000"))

	var parr []Person
	var pbrr []*Person
	parr = append(parr, Person{Name: "shi", Age: 10})
	parr = append(parr, Person{Name: "shim", Age: 12})
	parr = append(parr, Person{Name: "hh", Age: 20})
	for _, item := range parr {
		pbrr = append(pbrr, &item)
	}

	for _, item := range pbrr {
		print(item)
	}

	_, err := time.Parse("20060102", "20200101")
	_, err = time.Parse("2006-01-02", "20-01-10")
	_, err = time.Parse("2006/01/02", "2020/0101")
	fmt.Println(err)

	compare()
	divideSlice()

	type Merchant struct {
		MaxMerchantID int64 `json:"max_merchant_id"`
	}

	m := Merchant{MaxMerchantID: 0}
	key := "MxstMdeMUVboMAS54a6r5MtXEAUJPTEjrhqjfYgqtHFNcDYt4GlIA1jKOz90WKm2"

	body, _ := json.Marshal(m)

	fmt.Println(string(body))

	sign := CheckSigBytes(body, []byte(key))
	fmt.Println(sign)

	var ps []Person
	ps = append(ps, Person{Name: "shi", Age: 10})
	ps = append(ps, Person{Name: "li", Age: 11})
	ps = append(ps, Person{Name: "wang", Age: 12})
	personUpdate(ps)

	//httpTest()

	linkListTest()

	loopTest()

	rule := &base.OutletCountInRule{
		MerchantRule: &base.MerchantRule{
			MerchantIdListExclude: []uint64{200, 201},
			MerchantIdList:        []uint64{200, 101},
			TypeListExclude:       []uint32{2, 3},
			TypeList:              []uint32{2, 3},
		},
		OutletRule: &base.OutletRule{},
	}
	base.ConvertToPbReq(rule)

	img.GetAllFileTime("./static/new_test/")

	func1()

	//img.TestCompress("./static/bg_vn.png")

	//timeTick()

	var groupAmountMap = make(map[string][]int64)
	//var groupIdsMap = make(map[string][]string)
	now := time.Now()
	//unow := 1580471381

	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-31 12:00:00", loc)
	now = tmp
	fmt.Printf("Now:%s\\n", now.Format("2006-01-02 15:04:05"))
	lastMonthTime := GetMonthGapTime(now, 6)
	//lastMonthTime := now.AddDate(0, 1, 0)
	fmt.Printf("lastMonthTime before Yesterday:%s\\n", lastMonthTime.Format("2006-01-02 15:04:05"))
	lastMonthTime = now.AddDate(0, 2, 0)
	fmt.Printf("lastMonthTime before Yesterday:%s\\n", lastMonthTime.Format("2006-01-02 15:04:05"))
	lastMonthTime = now.AddDate(0, 0, 0)
	fmt.Printf("lastMonthTime before Yesterday:%s\\n", lastMonthTime.Format("2006-01-02 15:04:05"))

	for _, item := range groupAmountMap {
		if item[1] == 0 {
			fmt.Printf("error\\n")
		}
	}

	//base.RefRun()

	nowTime := time.Now()

	if nowTime.Hour() < runAtHour {
		fmt.Println("true")
	} else {
		fmt.Println("f")
	}

	ostr := "80868 , 80869 , 80870 "
	var strarr = strings.Split(ostr, ",")
	var apaStaffIDs []uint64
	for _, item := range strarr {
		_id, _ := strconv.Atoi(strings.TrimSpace(item))
		apaStaffIDs = append(apaStaffIDs, uint64(_id))
	}

	var strArr []string
	if strArr == nil {
		fmt.Println("strArr is nil")
	}
	fmt.Println(strings.Join(strArr, ","))

	var sett = "{\"rollback\":false}"
	var settl *SettlementExtra
	json.Unmarshal([]byte(sett), &settl)

	var s = "Mister Donut (Lotus Tha Tum Surin - โลตัสท่าตูม สุรินทร์)"
	//ls := strconv.QuoteToASCII(s)
	bytes := []rune(s)
	if len(bytes) <= 34 {
		fmt.Println(s, "")
	} else if len(bytes) <= 64 {
		fmt.Println(string(bytes[0:32]))
		fmt.Println(string(bytes[32:]))
	} else {
		fmt.Println(string(bytes[0:62])+"...", "")
	}

	coordinates := strings.Split("23 | 33", "|")
	if coordinates != nil && len(coordinates) > 0 {
		la := strings.Trim(coordinates[0], " ")
		val, err := strconv.ParseFloat(la, 64)
		if err == nil {
			fmt.Println(int64(val * 10000000))
		}
		if len(coordinates) == 2 {
			lo := strings.Trim(coordinates[1], " ")
			val, err := strconv.ParseFloat(lo, 64)
			if err == nil {
				fmt.Println(int64(val * 10000000))
			}
		}
	}

	str1 := "\"https://static.test.mitra.shopee.co.id/apa/下载.png\"\""
	//var newValue string
	bytesVal := []byte(str1)
	if string(bytesVal[0]) == "[" {
		str1 = strings.Replace(str1, "[", "", 1)
	}
	if string(bytesVal[len(bytesVal)-1]) == "]" {
		str1 = strings.Replace(str1, "]", "", 1)
	}

	str1 = strings.ReplaceAll(str1, "\"", "")

	a1 := []uint64{8, 5, 3, 2}
	a2 := []uint64{7, 6, 4, 2, 1}

	fmt.Println(mergeSortOrders(a1, a2))

	var a []string
	if a == nil {
		a = append(a, "1")
	}

	var b uint64
	//strconv.Itoa(int(b))
	fmt.Println(strconv.Itoa(int(b)))

	base.ParamsTest()
	base.ChannelTestOne()

	base.SelectCaseOne()
	//
	//base.PrintTest()

	//base.RefRun()
	//base.EscapeRun()
	//data := base.ForGoroutine(100)
	//fmt.Println(data)

	//qr := img.NewQrCode(102121, 332212, "abcdefg", "testdsdssasdcxsd")
	//
	//p, _ := qr.FileQrCodeWithBg("test_qr.png", "static/")
	//fmt.Println(p)

	//img.Upload()
}

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

func mergeSortOrders(ordersA, ordersB []uint64) []uint64 {
	if ordersA == nil || len(ordersA) == 0 {
		return ordersB
	}
	if ordersB == nil || len(ordersB) == 0 {
		return ordersA
	}
	lenA := len(ordersA)
	lenB := len(ordersB)
	ordersC := make([]uint64, lenA+lenB)

	ai := 0
	bi := 0
	ci := 0
	for ai < lenA && bi < lenB {
		if ordersA[ai] <= ordersB[bi] {
			ordersC[ci] = ordersA[ai]
			ai++
			ci++
		} else {
			ordersC[ci] = ordersB[bi]
			bi++
			ci++
		}
	}
	for ai < lenA {
		ordersC[ci] = ordersA[ai]
		ci++
		ai++
	}
	for bi < lenB {
		ordersC[ci] = ordersB[bi]
		ci++
		bi++
	}
	return ordersC
}

func timeTick() {
	ticker := time.NewTicker(time.Second * time.Duration(Interval))
	for t := range ticker.C {
		startTime := time.Now()
		if startTime.Minute() < 54 { // process the accounts at 3 a.m.
			fmt.Printf("continue time is %v\\n", t)
			continue
		}
		fmt.Printf("time is %v\\n", t)
	}
}
