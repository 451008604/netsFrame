package config

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/451008604/nets/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"netsFrame/logs"
	"os"
	"time"
)

var jsonsPath = "./config/jsons/"

type GlobalConf struct {
	config.ServerConf
	Debug   bool
	Version string
	Redis   Database
	Mysql   Database
}

type Server struct {
	Address     string
	Port        string
	TLSCertPath string
	TLSKeyPath  string
}

type Database struct {
	Address  string
	Username string
	Password string
}

var conf *GlobalConf

func initServerConfig(remoteAddress string) {
	viper.SetConfigType("toml")
	// 注册需要监控的配置文件
	viper.SetConfigFile("./config.toml")
	viper.WatchConfig()
	// 开启监控回调，限制每秒最多执行1次
	t := time.Now().Unix()
	viper.OnConfigChange(func(in fsnotify.Event) {
		now := time.Now().Unix()
		if in.Has(fsnotify.Write) && now-1 > t {
			t = now
			loadServerConfig()
		}
	})

	// 初始化配置内容
	var configByte []byte
	var err error
	if remoteAddress != "" {
		configByte, err = GetRemoteConfigData(remoteAddress, "config.toml")
	} else {
		configByte, err = os.ReadFile("./config.toml")
	}
	logs.PrintLogPanic(err)
	logs.PrintLogPanic(viper.ReadConfig(bytes.NewBuffer(configByte)))
	loadServerConfig()
}

// 解析配置内容到结构体
func loadServerConfig() {
	logs.PrintLogErr(viper.Unmarshal(&conf))
	logs.PrintLogErr(viper.Unmarshal(&conf.ServerConf))
	logs.SetPrintMode(conf.Debug)
	config.SetServerConf(conf.ServerConf)
	logs.PrintLogInfo(fmt.Sprintf("服务配置参数：%v", viper.AllSettings()))
}

// GetGlobalObject 获取全局配置对象
func GetGlobalObject() GlobalConf {
	if conf == nil {
		serverUrl := os.Getenv("REMOTE_CONFIG_SERVER_URL")
		initServerConfig(serverUrl)
	}
	return *conf
}

func GetRemoteConfigData(remoteAddress, fileName string) ([]byte, error) {
	resp, err := http.Get(remoteAddress + "/configFile?fileName=" + fileName)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if body, _ := io.ReadAll(resp.Body); len(body) != 0 {
		return body, nil
	}
	return nil, errors.New("文件 " + fileName + " 获取失败")
}
