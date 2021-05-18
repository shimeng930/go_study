package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	postUrl = "http://competition.jn-rencai.com:8080/register/saveUser"
	method  = "POST"
)
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Resp struct {
	Msg string
	Code int
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func MakeAccount(count int) {

	//payload := strings.NewReader("name=shim_t3&username=shimt3&password=shim_t2&confirmPassword=shim_t2&email=123456%40qq.com&mobile=13456721345")
	//req, err := http.NewRequest(method, postUrl, payload)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//
	//client := &http.Client {
	//}
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
	//
	//var resp *Resp
	//err = json.Unmarshal(body, &resp)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if resp.Code != 0 {
	//	fmt.Printf("ERROR|business error|code=%d, msg=%s\n", resp.Code, resp.Msg)
	//}

	var (
		pwd 	= "jn123456"
		email 	= "jn123456@qq.com"
		mobile 	= "13325658765"
		playlod = fmt.Sprintf("password=%s&confirmPassword=%s&email=%s&mobile=%s", pwd, pwd, email, mobile)
	)

	client := &http.Client {}

	for i:=0; i<count; i++ {
		name := fmt.Sprintf("jn_%s", randStringRunes(5))
		playlod = fmt.Sprintf("name=%s&username=%s&%s", name, name, playlod)
		body, err := doRequest(client, playlod)
		if err != nil {
			fmt.Printf("ERROR|doRequest error|i=%d, err=%v", i, err)
			continue
		}
		var resp *Resp
		err = json.Unmarshal(body, &resp)
		if err != nil {
			fmt.Println(err)
			return
		}
		if resp.Code != 0 {
			fmt.Printf("ERROR|business error|code=%d, msg=%s\n", resp.Code, resp.Msg)
		}
		fmt.Printf("SUCCESS|name=username=%s\n", name)
		time.Sleep(10*time.Duration(time.Millisecond))
	}
}

func doRequest(client *http.Client, params string) ([]byte, error) {
	payload := strings.NewReader(params)
	req, err := http.NewRequest(method, postUrl, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}