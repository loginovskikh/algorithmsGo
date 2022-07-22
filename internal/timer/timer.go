package timer

import (
	"errors"
	"time"
)

type Timer struct {
	startTime uint
	endTime   uint
}

func (t *Timer) Start() {
	t.startTime = uint(time.Now().UnixNano())
}

func (t *Timer) Stop() {
	t.endTime = uint(time.Now().UnixNano())
}

func (t Timer) Diff() (uint, error) {
	if t.startTime == 0 || t.endTime == 0 {
		return 0, errors.New("timer not started or not stopped")
	}

	return t.endTime - t.startTime, nil
}
