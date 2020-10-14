package channel

import (
	"testing"
	"time"
)

func Test_spinner(t *testing.T) {
	var timeNum time.Duration
	timeNum = 1 * time.Second
	spinner(timeNum)
}

func Test_chan2(t *testing.T) {

	chan2()
}
