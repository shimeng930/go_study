package base

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"sync"
	"time"
)

func TestReadCost() {
	start := time.Now()
	ReadFileAll("/Users/mengshi/Downloads/ubuntu-dev.box")
	fmt.Printf("read cost %d\n", time.Now().Sub(start)/1e6)
}

// 文件直接读取到[]byte
func ReadDirect() error {
	_, err := ioutil.ReadFile("file/test")
	if err != nil {
		fmt.Println("read fail", err)
	}
	return err
}

// 读取到file中，再利用ioutil将file直接读取到[]byte
func ReadFileAll(path string) error {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("cannot able to read the file", err)
		return err
	}
	defer f.Close()

	_, err = ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return err
	}

	// fmt.Println(string(fd))
	return nil
}

//
func ReadFileByLine(path string) error {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("cannot able to read the file", err)
		return err
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		_, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
	}
	return nil
}

// file->buf->[]byte
func Read1(path string) error {
	//获得一个file
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("read fail")
		return err
	}

	//把file读取到缓冲区中
	defer f.Close()
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return err
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}

	//fmt.Println(string(chunk))
	return nil
}

// file->reader->buf->[]byte
func Read2(path string) error {
	//获得一个file
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("read fail")
		return err
	}

	//把file读取到缓冲区中
	defer f.Close()
	reader := bufio.NewReader(f)
	var chunk []byte
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return err
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}

	//fmt.Println(string(chunk))
	return nil
}

// sync.Pool是一个强大的对象池，可以重用对象来减轻垃圾收集器的压力。我们将重用各个分片的内存，以减少内存消耗，大大加快我们的工作
// Go Routines帮助我们同时处理缓冲区块，这大大提高了处理速度
func Process(f *os.File, start time.Time, end time.Time) error {
	//sync pools to reuse the memory and decrease the preassure on Garbage Collector
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 500*1024)
		return lines
	}}
	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}
	//slicePool := sync.Pool{New: func() interface{} {
	//	lines := make([]string, 100)
	//	return lines
	//}}
	r := bufio.NewReader(f)
	var wg sync.WaitGroup //wait group to keep track off all threads
	for {

		buf := linesPool.Get().([]byte)
		n, err := r.Read(buf)
		buf = buf[:n]
		if n == 0 {
			if err != nil {
				fmt.Println(err)
				break
			}
			if err == io.EOF {
				break
			}
			return err
		}
		nextUntillNewline, err := r.ReadBytes('\n') //read entire line

		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}

		wg.Add(1)
		go func() {

			//process each chunk concurrently
			//start -> log start time, end -> log end time

			ProcessChunk(buf, &linesPool, &stringPool, start, end)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, start time.Time, end time.Time) {

	var wg2 sync.WaitGroup

	logs := stringPool.Get().(string)
	logs = string(chunk)

	linesPool.Put(chunk)

	logsSlice := strings.Split(logs, "\n")

	stringPool.Put(logs)

	chunkSize := 300
	n := len(logsSlice)
	noOfThread := n / chunkSize

	if n%chunkSize != 0 {
		noOfThread++
	}

	for i := 0; i < (noOfThread); i++ {

		wg2.Add(1)
		go func(s int, e int) {
			defer wg2.Done() //to avaoid deadlocks
			for i := s; i < e; i++ {
				text := logsSlice[i]
				if len(text) == 0 {
					continue
				}
				logSlice := strings.SplitN(text, ",", 2)
				logCreationTimeString := logSlice[0]

				logCreationTime, err := time.Parse("2006-01-02T15:04:05.0000Z", logCreationTimeString)
				if err != nil {
					fmt.Printf("\n Could not able to parse the time :%s for log : %v", logCreationTimeString, text)
					return
				}

				if logCreationTime.After(start) && logCreationTime.Before(end) {
					//fmt.Println(text)
				}
			}

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
	}

	wg2.Wait()
	logsSlice = nil
}
