package roundrobin

import (
	"fmt"
	"time"
)

type Command func() error

type RoundRobin struct {
	commands []Command
	tasks    int

	chanTask chan error
	quit     chan interface{}
}

func NewRoundRobin(cmds []Command) *RoundRobin {
	return &RoundRobin{
		commands: cmds,
		tasks:    len(cmds),
		chanTask: make(chan error),
		quit:     make(chan interface{}),
	}
}

func (rr RoundRobin) Execute() error {
	defer func() {
		close(rr.chanTask)
		close(rr.quit)
	}()

	task := 0
	go rr.balancer(rr.commands[task])
	for {
		select {
		case <-rr.chanTask:
			fmt.Println("enter after execute task")
			task++
		case <-time.After(time.Duration(3) * time.Second):
			fmt.Println("enter by time.after")
			<-rr.chanTask
			task++
		case <-rr.quit:
			fmt.Println("finish program")
			break
		}

		if task >= rr.tasks {
			task = 0
		}

		go rr.balancer(rr.commands[task])
	}
}

func (rr RoundRobin) balancer(cmd Command) {
	rr.chanTask <- cmd()
}

func (rr RoundRobin) Quit() {
	fmt.Println("finishing...")
	rr.quit <- 0
}
