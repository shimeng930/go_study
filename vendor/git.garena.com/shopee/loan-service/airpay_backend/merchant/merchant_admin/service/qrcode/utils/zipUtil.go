package utils

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"os"
)

func Compress(filePaths []string) (*bytes.Buffer, error) {
	if filePaths == nil {
		return nil, errors.New("no data")
	}

	var data bytes.Buffer
	w := zip.NewWriter(&data)
	defer w.Close()

	for _, path := range filePaths {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		err = compress(f, "", w)
		if err != nil {
			return nil, err
		}
	}

	return &data, nil
}

//func Compress(files []*os.File) (*bytes.Buffer, error) {
//	var data bytes.Buffer
//	w := zip.NewWriter(&data)
//	defer w.Close()
//	for _, file := range files {
//		err := compress(file, "", w)
//		if err != nil {
//			return nil, err
//		}
//	}
//	return &data, nil
//}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	header.Name = prefix + "/" + header.Name
	if err != nil {
		return err
	}
	writer, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, file)
	file.Close()
	if err != nil {
		return err
	}
	return nil
}