package internal

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

// Ping helps to get ping result true - if packetlass not 100%
func Ping(addr string, timeout int) (ok bool, err error) {
	pinger, err := probing.NewPinger(addr)
	if err != nil {
		return
	}
	pinger.Count = 4
	pinger.Timeout = time.Second * time.Duration(timeout)
	err = pinger.Run()
	if err != nil {
		return
	}
	stats := pinger.Statistics()
	return stats.PacketLoss != 100, err
}
