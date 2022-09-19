package roundrobin_test

import (
	"fmt"
	"testing"
	"time"

	"go-staffs/concurrency/roundrobin"
)

func TestRoundRobin_Execute(t *testing.T) {
	cmd1 := func() error {
		fmt.Println("cmd1")
		time.Sleep(time.Second * 2)
		return nil
	}
	cmd2 := func() error {
		fmt.Println("cmd2")
		time.Sleep(time.Second * 3)
		return nil
	}
	cmd3 := func() error {
		fmt.Println("cmd3")
		time.Sleep(time.Second * 2)
		return nil
	}
	cmds := []roundrobin.Command{cmd1, cmd2, cmd3}

	rr := roundrobin.NewRoundRobin(cmds)
	go rr.Execute()

	time.Sleep(time.Second * 10)
	rr.Quit()
}

func TestSarasa(t *testing.T) {
	per := Path{
		1: 1,
		2: 1,
		3: 4,
	}
	fmt.Println(per.Distance())
}

type Path map[int]int

func (p Path) Distance() int {
	res := 0
	for k, _ := range p {
		res += k
	}
	return res
}
