package setting

//负责调用配置

import (
"log"
"time"

"github.com/go-ini/ini"
)

type Server struct {
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type App struct {
	PageSize int
	JWTSecret string

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogFileExt string
	TimeFormat string
}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
}

var (
	RunMode string

	ServerSetting = &Server{}
	AppSetting = &App{}
	DatabaseSetting  = &Database{}
)

func Setup() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	RunMode = cfg.Section("").Key("RUN_MODE").MustString("debug")

	if err := cfg.Section("server").MapTo(ServerSetting); err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	if err := cfg.Section("app").MapTo(AppSetting); err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	AppSetting.ImageMaxSize *=  1024 * 1024

	if err := cfg.Section("database").MapTo(DatabaseSetting); err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}
