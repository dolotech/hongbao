package logic

import (
	"data"
	"github.com/golang/glog"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"utils/cfg"
)

func Http(addr, hbid, cookie string, item data.HonbaoTb) error {
	/*resp, err := http.Post("http://"+addr+"/index/api/kai",
	"application/x-www-form-urlencoded",
		strings.NewReader("hbid="+hbid))
	if err != nil {
		return err
	}
	*/

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://"+addr+"/index/api/kai", strings.NewReader("hbid="+hbid))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "PHPSESSID="+cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	j := jsoniter.Get(body)

	if j.Get("code").ToString() == "ok" {
		glog.Info("抢包成功", item)
	} else {
		glog.Info(j.Get("message").ToString(), item)
	}

	return nil
}

func GetAll(cookies cfg.Cookies) error {
	client := &http.Client{}
	stamp := strconv.Itoa(int(time.Now().UnixNano() / 1000000))
	req, err := http.NewRequest("GET", "http://"+cookies.API+"/index/API/fabao?timestamp=", strings.NewReader("timestamp="+stamp))
	if err != nil {
		return err
	}
	r := rand.Int31n(int32(len(cookies.Cookie)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "PHPSESSID="+cookies.Cookie[r])

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ParseData(body, cookies.API, cookies.Cookie[r])
	return nil
}

//http://hb.yqsyyly.cn/index/API/fabao?timestamp=1575099030713
