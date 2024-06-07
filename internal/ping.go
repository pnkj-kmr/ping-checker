package internal

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

// Ping helps to get ping result true - if packetlass not 100%
func Ping(i Input, count, timeout int) (out Output, err error) {
	pinger, err := probing.NewPinger(i.IP)
	if err != nil {
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
		return
	}

	stats := pinger.Statistics()
	return Output{
		I: i, Ok: stats.PacketLoss != 100,
		Err: err, PacketLoss: stats.PacketLoss,
		AvgRtt: stats.AvgRtt, StdDevRtt: stats.StdDevRtt,
	}, err
}
