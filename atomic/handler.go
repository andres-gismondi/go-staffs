package atomic

import (
	"fmt"
	"sync/atomic"
	"time"
)

type TimeSelect struct {
	flag int32
}

func NewTimeSelect() *TimeSelect {
	return &TimeSelect{flag: 0}
}

func (ts *TimeSelect) Execute() error {
	if atomic.LoadInt32(&ts.flag) > 0 {
		fmt.Println("Already executed")
		return nil
	}
	atomic.AddInt32(&ts.flag, 1)
	time.Sleep(time.Second * 4)
	fmt.Println("Execute!")
	atomic.AddInt32(&ts.flag, -1)

	return nil
}
