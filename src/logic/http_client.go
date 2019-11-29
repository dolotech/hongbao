package logic

import (
	"data"
	"github.com/golang/glog"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
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
