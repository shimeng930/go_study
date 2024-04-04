package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

var cnt = 8192

type Parse struct {
	programMap map[string]*vm.Program
	env        map[string]interface{}
}

func t(ctx context.Context) {
	go func(c context.Context) {
		//for i := 0; i < 10; i++ {
		//	fmt.Println(c)
		//	time.Sleep(1 * time.Second)
		//	fmt.Println(i)
		//}
		for {
			select {
			case <-ctx.Done():
				fmt.Println("退出携程")
				return
			default:
				fmt.Println("请求中..")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
}

func doTest(ctx context.Context, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	//time.Sleep(time.Millisecond * 3000)
	select {
	case <-ctx.Done():
		fmt.Println("handle")
	case <-time.After(3 * time.Millisecond):
		fmt.Println("process request with", 3)
	}
	fmt.Println("exec")
	if num == 10000 {
		fmt.Println(num)
	}
}

func watch(ctx context.Context) {
	//for {
	select {
	case <-ctx.Done():
		fmt.Println("watch get cancel and exit")

	case <-time.After(1 * time.Second):
		fmt.Println("process request with", 3)
		//default:
		//	fmt.Println("watching", 3)
		//	time.Sleep(time.Second)
	}
	//}
}

func main() {
	ctxp := context.Background()

	nc, cancel := context.WithCancel(ctxp)

	wg := &sync.WaitGroup{}
	wg.Add(10001)
	go watch(nc)
	cancel()
	fmt.Println("cancel")
	//for i := 0; i < 10001; i++ {
	//	go doTest(ctxp, i, wg)
	//}

	//ctx, _ := context.WithTimeout(ctxp, 3*time.Second)
	//
	//t(ctx)
	//cancel()

	//time.Sleep(10 * time.Second)
	//wg.Wait()
	time.Sleep(5000 * time.Millisecond)
	//cancel()
	fmt.Println("exit")
	//return
	//printMemStats()
	//
	//fmt.Println(4096 % 2048)
	//fmt.Println(4096 & 2047)
	//
	//env := map[string]interface{}{
	//	"params":  "params",
	//	"columns": "rowData",
	//	"nil": struct {
	//	}{},
	//}
	//p := &Parse{env: env}
	//initMap(p)
	//runtime.GC()
	//printMemStats()
	//
	//for i := 0; i < 20; i++ {
	//	initMap(p)
	//	printMemStats()
	//}

	//for i := 0; i < cnt; i++ {
	//	delete(intMap, i)
	//}
	//log.Println(len(intMap))

	//runtime.GC()
	//printMemStats()

	//intMap = nil
	//runtime.GC()
	//printMemStats()
}

func initMap(parse *Parse) {
	if parse.programMap == nil {
		parse.programMap = make(map[string]*vm.Program)
		parse.programMap["nil"] = nil
	}

	a, ok := parse.programMap["nil"]
	fmt.Println(ok)
	fmt.Println(a)

	for i := 0; i < cnt; i++ {
		key := strconv.FormatInt(int64(i), 10)

		//p, _ := expr.Compile(key, expr.Env(parse.env))
		//if p != nil {
		//}

		if parse.env["nil"] != struct{}{} {
			fmt.Println(false)
		}

		_, ok := parse.programMap[key]
		if !ok {
			parse.programMap[key], _ = expr.Compile(key, expr.Env(parse.env))
		} else {
		}
		//fmt.Printf("%p\n", parse.programMap[key])
	}
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc=%dKB, TotalAlloc=%dKB, Sys=%vKB, NumGC=%d\n", m.Alloc>>10, m.TotalAlloc>>10, m.Sys>>10, m.NumGC)
}
