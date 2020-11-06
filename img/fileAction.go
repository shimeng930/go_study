package img

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

func getFileTime(file os.FileInfo) int64 {
	statTime := file.Sys().(*syscall.Stat_t)
	return int64(statTime.Ctimespec.Sec)
}

func GetAllFileTime(path string) {
	rd, _ := ioutil.ReadDir(path)
	for _, fi := range rd {
		if fi.IsDir() {
			GetAllFileTime(path + fi.Name())
		} else {
			fmt.Printf("[%s:%d]\n", path+fi.Name(), getFileTime(fi))
			os.Remove(path+fi.Name())
		}
	}
}
