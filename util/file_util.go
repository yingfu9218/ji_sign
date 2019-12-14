package util

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"path/filepath"
)


var BasePath string
func init()  {
	BasePath=GetExecutePath()
	CheckRuntimePath()
}
func SetTestBasePath()  {
	BasePath=AppConfig.GetString("basepath")
}


func GetExecutePath() string {
	dir, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	exPath := filepath.Dir(dir)

	return exPath
}

func GetTmpFileName(prefix string) string  {

	rand:=uuid.NewV4().String()

	return BasePath+"/runtime/"+rand+"."+prefix

}

func CheckRuntimePath()  {
	_, err := os.Stat(BasePath+"/runtime")
	if err == nil {
		os.Mkdir(BasePath+"/runtime",os.ModePerm)
	}
}

func GetRuntimePath() string  {

	return  BasePath+"/runtime"
}

func CheckFileExit(f string) bool {

	_, err := os.Stat(f)
	if err == nil {
		return true
	}else {
		return false
	}

}