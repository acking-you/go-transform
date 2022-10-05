package send

import (
	"bytes"
	"errors"
	"fmt"
	"logger/lg"
	"net"
	"time"
)

const kPingPeriod = 3
const kPort = 5666

var localAddrs *[]*net.IP
var controller chan bool
var pauseFlag bool
var ticker *time.Timer
var count int

func SendUDPByTicker() {
	//资源初始化
	pauseFlag = false
	controller = make(chan bool)
	count = 0
	ticker = time.NewTimer(kPingPeriod * time.Second)
	if localAddrs == nil {
		addrs, err := getLocalAddr()
		localAddrs = &addrs
		if err != nil {
			lg.L.Errorln("获取本地地址出错")
		}
	}

	//每隔kPingPeriod发送UDP消息
	defer func() {
		ticker.Stop()
		close(controller)
	}()

	for {
		select {
		case flag := <-controller:
			pauseFlag = flag
			//每次controller都重新唤醒Timer
			ticker.Reset(kPingPeriod * time.Second)
		case start := <-ticker.C:
			doRoundUDPInfo(start)
		}
	}
}

func Pause() {
	controller <- true
}

func ReStart() {
	controller <- false
}

func doRoundUDPInfo(start time.Time) {
	err := sendUDPInfo(*localAddrs)
	if err != nil {
		lg.L.Errorf("发送消息失败 %s \n", err.Error())
		lg.L.Error("尝试重新发送第二轮")
		_ = sendUDPInfo(*localAddrs)
		return
	}

	lg.L.Infof("第%d次发送扫描消息成功，耗时 %v", count, time.Since(start))
	count++
	//根据pauseFlag判断是否需要重置ticker
	if !pauseFlag {
		ticker.Reset(kPingPeriod * time.Second)
	}
}

func GetValidLocalAddr() (string, error) {
	addrs, err := getLocalAddr()
	if err != nil {
		return "", err
	}
	if len(addrs) == 0 {
		return "", errors.New("有效地址为空")
	}
	return addrs[0].To4().String(), nil
}

func isCMask(mask []byte) bool {
	return bytes.Compare(mask, []byte{255, 255, 255, 0}) == 0
}

func getLocalAddr() ([]*net.IP, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return nil, err
	}
	var ret []*net.IP
	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && isCMask(ipnet.Mask) && ipnet.IP.To4()[0] != 10 {
				ip := ipnet.IP.To4()
				lg.L.Infof("本机有效IP：%v", ip)
				ret = append(ret, &ip)
			}
		}
	}
	return ret, nil
}

func sendUDPInfo(ips []*net.IP) error {
	for _, ip := range ips {
		bts := []byte(*ip)
		if len(bts) != 4 {
			continue
		}
		for i := 0; i <= 255; i++ {
			ipaddr := fmt.Sprintf("%d.%d.%d.%d:%d", bts[0], bts[1], bts[2], i, kPort)
			if err := udpSend(ipaddr, "ping"); err != nil {
				return err
			}
		}
	}
	return nil
}

func udpSend(addr, text string) error {
	conn, err := net.Dial("udp", addr)
	defer func() {
		_ = conn.Close()
	}()
	if err != nil {
		return err
	}
	conn.Write([]byte(text))
	return nil
}
