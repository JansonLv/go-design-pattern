package main

import "fmt"

type call interface {
	onclick()
}

type buttion struct {
	name string
}

func (b *buttion) onclick() {
	fmt.Println("按钮点击了")
}

type report struct {
	call call
}

func newReportWarp(caller call) call {
	return &report{call: caller}
}

func (r *report) onclick() {
	r.call.onclick()
	fmt.Println("上报")
}

type monitor struct {
	call call
}

func newMonitorWarp(caller call) call {
	return &monitor{call: caller}
}

func (m *monitor) onclick() {
	m.call.onclick()
	fmt.Println("监控")
}

func main() {
	b := &buttion{}
	b1 := newMonitorWarp(b)
	b2 := newReportWarp(b1)
	b2.onclick()
}
