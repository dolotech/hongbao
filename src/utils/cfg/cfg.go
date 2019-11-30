package cfg

import (
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/kr/pretty"
	"os"
)

type Cookies struct {
	Cookie  []string
	Address string
	API     string
	Mode    int
}

type MysqlDB struct {
	Host     string
	User     string
	Password string
	DBName   string
}

// Config 配置类型
type Config struct {
	Cookies []Cookies
	MysqlDB MysqlDB
}

// Opts Config 默认配置
var opts *Config

// ParseToml 解析配置文件
var file string

func Reload() error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		glog.Errorln("没有找到配置文件 ...")
		return nil
	}
	opts = &Config{}
	_, err := toml.DecodeFile(file, opts)
	if err != nil {
		glog.Errorln("配置文件解析错误：", err)
		return err
	}
	glog.Infof("cfg is %v", pretty.Formatter(opts))
	return nil
}
func ParseToml(f string) error {
	file = f
	return Reload()
}

// Opts 获取配置
func Opts() *Config {
	return opts
}
