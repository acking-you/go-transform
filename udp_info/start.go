package udp_info

import "udp_info/send"

func Start() {
	send.SendUDPByTicker()
}
