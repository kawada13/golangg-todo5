package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// 読み書き、ファイルがなければ作成、追記ができる
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln(err)
	}

	// ログの出力先を代入
	multiLogFile := io.MultiWriter(os.Stdout, logfile)

	// ログのフォーマットを作成
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// 実際にログの出力先を設定
	log.SetOutput(multiLogFile)

}
