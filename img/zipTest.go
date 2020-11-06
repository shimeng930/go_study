package img

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"os"
	"time"
)

func Compress2File(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
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

func Compress(files []string) (*bytes.Buffer, error) {
	if files == nil {
		return nil, errors.New("no data")
	}

	var data bytes.Buffer
	w := zip.NewWriter(&data)
	defer w.Close()

	for _, path := range files {
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
	//writer.Write(ioutil.ReadAll(file))
	file.Close()
	if err != nil {
		return err
	}
	return nil
}

func TestCompress(p string)  {
	//path, _ := os.Getwd()
	//p1 := fmt.Sprintf("%s/static/bg_th.png", path)
	f1, err := os.Open("./static/bg_th.png")
	defer f1.Close()
	//p2 := fmt.Sprintf("%s/static/bg_vn.png", path)
	//f2, err := os.Open(p2)
	//defer f2.Close()
	//p3 := fmt.Sprintf("%s/static/test_qr.png", path)
	//f3, err := os.Open(p3)
	//defer f3.Close()
	//var files = []*os.File{f1, f2, f3}
	var files = []string{p, "./static/code_time.xlsx", "./static/code_time_1.xlsx"}
	//var files = []string{p}
	//dest := fmt.Sprintf("%s/static/test.zip", path)
	body, err := Compress(files)
	if err != nil {
		println(err, body)
	}

	up := NewUploadFile("ssss.zip", int(time.Now().Unix()))
	up.Upload(body, "bg_th.zip")

	//Upload(body, "bg_th.zip")
}
