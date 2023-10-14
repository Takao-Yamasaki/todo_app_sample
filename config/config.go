package config

import (
	"log"
	"todo_app_sample/utils"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

// NOTE: 外部のパッケージから呼び出せるように
var Config ConfigList

// main関数より前に読み込む
func init() {
	// configの読み込み
	LoadConfig()
	// ログファイルの設定
	// NOTE: Config.LogFileはLoadConfigで設定
	utils.LoggingSettings(Config.LogFile)
}

// iniファイルを設定の読み込んで、ConfigListを作成する
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
