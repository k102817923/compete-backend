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

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HttpPort = server.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
