package main

import (
	"github.com/zhyea/vibe"
	"log"
)

var W = new(Workshop)

func init() {
	vibe.AddConfigFiles("conf/config.yml")

	if err := vibe.ReadConfig(); nil != err {
		log.Fatalf("读取配置文件失败，异常信息：%v", err)
	}
	if err := vibe.Unmarshal(W); nil != err {
		log.Fatalf("赋值配置对象失败，异常信息：%v", err)
	}

	log.Println("读取配置文件结束！")
}
