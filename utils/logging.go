package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// ログファイルを開く（書き込み・読み込み権限/ファイルがなければ作成/追記）
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// 標準出力とログファイルを指定したうえで、ログの出力先を指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// ログのフォーマット指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// ログの出力先を指定
	log.SetOutput(multiLogFile)
}
