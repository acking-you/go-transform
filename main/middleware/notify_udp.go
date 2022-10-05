package middleware

import (
	"github.com/gin-gonic/gin"
	"udp_info/send"
)

func NotifyUDP(c *gin.Context) {
	send.Pause() //后台udp消息停止
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
	send.ReStart() //后台udp服务重启
}
