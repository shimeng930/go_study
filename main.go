package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_study/base"
	"go_study/img"
	"io/ioutil"
	"math"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Person struct {
	Name string
	Age int
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
	ResultCode 		string `json:"result_code"`
	ReportStatus 	string `json:"report_status"`
	ErrorCode 		string `json:"error_code"`
}

const currencyAmountMultiple  = 1000000
const runAtHour  = 52
const Interval  = 5

type Action struct {
	Name string
	obj  []int
}

func (a *Action) AddObj(objs ...int)  {
	a.obj = append(a.obj, objs...)
}

func GetMonthGapTime(curTime time.Time, gapNum int) time.Time {
	tempTime := curTime.AddDate(0, gapNum, 0)

	curYear, curMonth, _ := curTime.Date()
	tempYear, tempMonth, _ := tempTime.Date()
	if (tempYear-curYear) * 12 + int(tempMonth) != int(curMonth) + gapNum {
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
	defer fmt.Printf("defer print i=%d\n", i)
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

func httpTest()  {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["id"] = "1234"
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST","http://merchant-reports.apa-admin-th-dev.api.ingress",bytes.NewReader(bytesData))
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
	for i=0; i<=3; i++ {
		arr[i] = 0
		fmt.Printf("arr[i]=%d, i=%d, addr=%p, iaddr=%p\n", arr[i], i, &arr[i], &i)
	}
}

func linkListTest()  {
	node := base.NewListNode()
	head := node.InitList(10)
	node.Print(head)

	head = node.Reverse(head)
	node.Print(head)

	node.GetLastKNode(11, head)
}

func main()  {

	//httpTest()

	linkListTest()

	loopTest()

	var ta = []int{1,2,3,4,5,6,7,8,9,10,11}
	for i:=0;i<len(ta); {
		lastIndex := i+3
		if lastIndex > len(ta) {
			lastIndex = len(ta)
		}
		temp := ta[i:lastIndex]
		fmt.Println(temp)
		i = lastIndex
	}

	rule := &base.OutletCountInRule{
		MerchantRule: &base.MerchantRule{
			MerchantIdListExclude: []uint64{200,201},
			MerchantIdList: []uint64{200, 101},
			TypeListExclude: []uint32{2,3},
			TypeList: []uint32{2,3},
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

	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-31 12:00:00", loc)
	now = tmp
	fmt.Printf("Now:%s\n", now.Format("2006-01-02 15:04:05"))
	lastMonthTime := GetMonthGapTime(now, 6)
	//lastMonthTime := now.AddDate(0, 1, 0)
	fmt.Printf("lastMonthTime before Yesterday:%s\n", lastMonthTime.Format("2006-01-02 15:04:05"))
	lastMonthTime = now.AddDate(0, 2, 0)
	fmt.Printf("lastMonthTime before Yesterday:%s\n", lastMonthTime.Format("2006-01-02 15:04:05"))
	lastMonthTime = now.AddDate(0, 0, 0)
	fmt.Printf("lastMonthTime before Yesterday:%s\n", lastMonthTime.Format("2006-01-02 15:04:05"))

	for _, item := range groupAmountMap {
		if item[1] == 0 {
			fmt.Printf("error\n")
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

	var sett  = "{\"rollback\":false}"
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

	a1 := []uint64{8,5,3,2}
	a2 := []uint64{7,6,4,2,1}

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
	for ai < lenA && bi <lenB {
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
			fmt.Printf("continue time is %v\n", t)
			continue
		}
		fmt.Printf("time is %v\n", t)
	}
}