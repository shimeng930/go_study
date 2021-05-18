package utils

import (
	"path"
	"runtime"
)

var FileUtil *fileUtil

type fileUtil struct {
}

func (f *fileUtil) GetCurDirPth() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
