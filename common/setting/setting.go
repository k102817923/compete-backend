package setting

import (
	"flag"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	RunMode      string
	configPath   = flag.String("conf", "config/config.default.ini", "Specify the path of config file")
)

func init() {
	var err error

	flag.Parse()
	Cfg, err = ini.Load(*configPath)
	if err != nil {
		log.Fatalf("Fail to parse %v: %v", *configPath, err)
	}

	LoadBase()
	LoadServer()
}

// 初始化基础参数
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// 初始化获取server参数
func LoadServer() {
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HttpPort = server.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// 根据参数获取配置
func loadConfigParam(section string, keyName string, defaultValue interface{}) interface{} {
	c, err := Cfg.GetSection(section)
	if err != nil {
		log.Fatalf("Fail to get section '%s': %v", section, err)
	}

	key := c.Key(keyName)
	switch v := defaultValue.(type) {
	case string:
		return key.MustString(v)
	case int:
		return key.MustInt(v)
	case bool:
		return key.MustBool(v)
	default:
		log.Fatalf("Unsupported type for key '%s' in section '%s'", keyName, section)
		return nil
	}
}

// 根据参数获取配置(string)
func LoadStringParam(section string, keyName string) string {
	return loadConfigParam(section, keyName, "").(string)
}

// 根据参数获取配置(int)
func LoadIntParam(section string, keyName string) int {
	return loadConfigParam(section, keyName, 0).(int)
}

// 根据参数获取配置(bool)
func LoadBoolParam(section string, keyName string) bool {
	return loadConfigParam(section, keyName, false).(bool)
}
