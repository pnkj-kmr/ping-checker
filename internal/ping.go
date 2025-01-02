package internal

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

// Ping helps to get ping result true - if packetlass not 100%
func Ping(i Input, count, timeout int) (out Output, err error) {
	out.I = i
	pinger, err := probing.NewPinger(i.IP)
	if err != nil {
		out.Err = err.Error()
		return
	}

	if pinger.Count = i.Count; i.Count == 0 {
		pinger.Count = count
	}

	var t int
	if t = i.Timeout; i.Timeout == 0 {
		t = timeout
	}
	pinger.Timeout = time.Second * time.Duration(t)
	err = pinger.Run()
	if err != nil {
		out.Err = err.Error()
		return
	}

	stats := pinger.Statistics()
	out.Ok = stats.PacketLoss != 100
	out.PacketLoss = stats.PacketLoss
	out.AvgRtt = stats.AvgRtt
	out.StdDevRtt = stats.StdDevRtt
	out.Err = ""
	return
}
