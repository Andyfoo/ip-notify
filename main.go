package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
	"net/smtp"
	"time"
)

func getIp() (string,error)  {
	resp ,err:=http.Get("http://api.ipify.org?format=json")
	defer resp.Body.Close()
	var result []byte
	if nil != err {
		return "" ,err
	}else {
		if result,err = ioutil.ReadAll(resp.Body);err != nil {
			return "",err
		}
	}
	resultMap := make(map[string]interface{})
	json.Unmarshal(result,&resultMap)
	v,ok := resultMap["ip"].(string)
	if ok {
		log.Printf("get ip success :%v\n",v)
		return v,nil
	}else {
		return "",err
	}
}

func sendMail(ip string,preIp string) error {

	auth := smtp.PlainAuth(
		"",
		"benkris1@126.com",
		"m522267128",
		"smtp.126.com",
	)
	err :=smtp.SendMail(
		"smtp.126.com:25",
		auth,
		"benkris1@126.com",
		[]string{"benkris1@126.com"},
		[]byte(fmt.Sprintf("To: benkris1@126.com \r\nFrom: 家里的树莓派<家里的树莓派> \r\nSubject: 外网IP地址变动 \r\nContent-Type: text/plain;charset=UTF-8\r\n\r\n 最新外网Ip为:%v。之前Ip为：%v",ip,preIp)),
	)
	 return err
}
func main()  {
	t := time.NewTicker(60 * time.Second)
	var preIp string
	for {
		select {
		 case <-t.C:
			 ip,err :=getIp()
			 if nil != err {
				 log.Fatal(err)
			 }
			 if preIp != ip{
			 	preIp = ip
				err := sendMail(ip,preIp)
				if nil != err {
					log.Println(err)
				}
			 }

		}
	}
}
