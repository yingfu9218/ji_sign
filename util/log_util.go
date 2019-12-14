package util

import (
	"fmt"
	"os"
	"time"
)



var logPath string

func OpenLogFile()  {
	logPath=BasePath+"/log.txt"
	_, err := os.Stat(logPath)
	if err != nil {
		os.Create(logPath)
	}


}


func Log(str string)  {
	// 以追加模式打开文件
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	defer  logFile.Close()
	if err !=nil{
		fmt.Print(err.Error())
	}
	timeNow:=time.Now()
	timeNowStr:=timeNow.Format("2006-01-02 15:04:05")
	str=timeNowStr+"  "+str+"\n"
	fmt.Println(str)
	logFile.WriteString(str)
}
