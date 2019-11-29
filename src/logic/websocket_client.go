package logic

import (
	"data"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/url"
	"strconv"
	"sync"
	"time"
	"utils/cfg"
)

const (
	pongWait = 3 * time.Second
)

type Resp struct {
	Code    string          `json:"code"`
	Honbao  data.HonbaoTbs  `json:"honbao"`
	Honbaos data.HonbaosTbs `json:"honbaos"`
}

type Float float64

func (this Float) Last() int {
	s := this.String()
	l := s[len(s)-1:]
	i, _ := strconv.Atoi(l)
	return i
}

func (this Float) String() string {
	return fmt.Sprintf("%.2f", this)
}

var syc sync.Map

func Websocet(cookies cfg.Cookies) {
	u := url.URL{Scheme: "ws", Host: cookies.Address, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		c.SetReadDeadline(time.Now().Add(pongWait))
		_, message, err := c.ReadMessage()
		if err != nil {
			glog.Info("read:", err)
			return
		}
		any := jsoniter.Get(message)

		if any.Get("code").ToString() == "ok" {
			d := &Resp{}
			any.ToVal(d)

			honbao := d.Honbao
			honbaos := d.Honbaos
			go honbao.Save()
			go honbaos.Save()

			for i := 0; i < len(honbao); i++ {
				//v := honbao[i]
				//抢包处理
				/*if v.Number == "6" && v.GetMoney()<=20{
					if _, ok := syc.Load(v.Id); !ok {
						syc.Store(v.Id, struct{}{})

						money := v.GetMoney()

						var moneyLeft float64
						for _, value := range honbaos {
							if value.HbId == v.Id {
								moneyLeft += value.GetMoney()
							}
						}
						if Float(money - moneyLeft).Last() != v.Lei {
							r := rand.Int31n(int32(len(cookies.Cookie)))
							//time.AfterFunc(time.Millisecond*200, func() {
								go Http(cookies.API, strconv.Itoa(v.Id), cookies.Cookie[r],v)
							//})
						}
					}
				}*/
			}
		}
	}
}
