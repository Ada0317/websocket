package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func main() {
	//open, err := gorm.Open(mysql.Open("szb02:iwala.netszb02@tcp(47.115.143.171:3306)/szb02?charset=utf8&loc=Asia%2FShanghai&parseTime=true&timeout=10s"))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(open)

	//return &Logger{
	//	Out:          os.Stderr,
	//	Formatter:    new(TextFormatter),
	//	Hooks:        make(LevelHooks),
	//	Level:        InfoLevel,
	//	ExitFunc:     os.Exit,
	//	ReportCaller: false,
	//}

	w1 := &bytes.Buffer{}
	w2 := os.Stdout
	w3, er := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if er != nil {
		log.Fatalf("create file failed %+v", er)
	}
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(io.MultiWriter(w1, w2, w3))
	//logrus.WithFields(logrus.Fields{"user_id": 1999910,
	//	"ip": "192.168.32.15",
	//})
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Panic("panic msg")
	logrus.Fatal("fatal msg")

}
