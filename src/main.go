package main

import (
	"flag"
	"github.com/golang/glog"
	"logic"
	"os"
	"os/signal"
	"syscall"
	"utils/cfg"
	"utils/db"
)

func SIGINT() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT:
			os.Exit(0)
		case syscall.SIGHUP, syscall.SIGTERM:
			cfg.Reload()
		}
	}
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "conf", "cfg.toml", "Configuration file to start game")
	flag.Parse()
	glog.Infoln("Configuration is", fileName)
	err := cfg.ParseToml(fileName)
	if err != nil {
		glog.Errorln("配置文件.toml出错")
		glog.Fatal(err)
	}

	go SIGINT()
	db.InitMysql(cfg.Opts().MysqlDB.User, cfg.Opts().MysqlDB.Password, cfg.Opts().MysqlDB.Host, cfg.Opts().MysqlDB.DBName)

	list := cfg.Opts().Cookies

	//db.Get().CreateTable(&data.HonbaosTb{})
	//db.Get().CreateTable(&data.HonbaoTb{})
	s:=logic.Float(10.13)

	glog.Info(s.String())
	glog.Info(s.Last())

	for _, value := range list {
		go logic.Websocet(value)
	}
	select {}
}
