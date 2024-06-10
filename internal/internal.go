package internal

import "time"

// PingChecker interface help to freeze
// the input and outout function definations
type PingChecker interface {
	GetInput() []Input
	ProduceOutput(<-chan Output, chan<- struct{})
}

// Input represent the input
type Input struct {
	IP      string `json:"ip"`
	Tag     string `json:"tag,omitempty"`
	Count   int    `json:"count,omitempty"`
	Timeout int    `json:"timeout,omitempty"`
}

// Output represent the output
type Output struct {
	I          Input         `json:"input"`
	Ok         bool          `json:"ok"`
	Err        string        `json:"error,omitempty"`
	PacketLoss float64       `json:"packet_loss"`
	AvgRtt     time.Duration `json:"avg_rtt"`
	StdDevRtt  time.Duration `json:"std_dev_rtt"`
}
