package main

import (
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"logger/lg"
	"net/http"
	"udp_info/send"
)

func main() {
	r := gin.Default()
	r.StaticFS("/", http.Dir("../web/"))

	addr, err := send.GetValidLocalAddr()
	if err != nil {
		lg.L.Errorf("获取本地地址出错 %v", err)
	} else {
		//打印二维码
		content := fmt.Sprintf("%s:%d", addr, 1314)
		obj := qrcodeTerminal.New()
		obj.Get(content).Print()
		color.Red("\nhttp://%s\n\n", content)
	}
	if err != nil {
		panic(err)
	}
	err = r.Run(":1314")
	if err != nil {
		panic(err)
	}
}
