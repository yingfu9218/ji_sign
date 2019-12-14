package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/robfig/cron"
	"ji_sign/util"
	"log"
	"time"
)

func init()  {
	//获取执行文件路径
	util.GetExecutePath()
	//加载配置文件
	util.LoadConfig()
	util.OpenLogFile()

}

func main() {
	c := cron.New()
	c.AddFunc("0 * * * * ?", func() {
		timeNow:=time.Now()
		timeNowStr:=timeNow.Format("2006-01-02 15:04:05")
		fmt.Print(timeNowStr +": test cron ,every hour \n")
	})
	c.AddFunc("0 0 9 * * ?", func() {
		fmt.Print("每天9:00签到")
		sign()
	})
	c.Start()
	select {}
}
//登录并签到
func sign()  {
	// create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("ji-bt.pw"),
	)

	// authenticate
	err := c.Post("https://ji-bt.pw/signin", map[string]string{"email": util.AppConfig.GetString("email"), "passwd": util.AppConfig.GetString("passwd")})
	if err != nil {
		log.Fatal(err)
		util.Log(err.Error())
	}

	c.OnResponse(func(r *colly.Response) {
		log.Println("response revice", string(r.Body))
		util.Log("response revice :"+string(r.Body))

	})
	c.Visit("https://ji-bt.pw/xiaoma/get_user")
	//签到
	err = c.Post("https://ji-bt.pw/user/checkin", map[string]string{})
	if err != nil {
		log.Fatal(err)
		util.Log(err.Error())
	}
}
