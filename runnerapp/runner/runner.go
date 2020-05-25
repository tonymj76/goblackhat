package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner _
type Runner interface {
	Start() error
	Add(...func(int))
}

//ErrTimeOut _
var ErrTimeOut = errors.New("time out")

//ErrInterrupted _
var ErrInterrupted = errors.New("interrupted error")

type monitor struct {
	complete  chan error
	interrupt chan os.Signal
	time      <-chan time.Time
	tasks     []func(int)
}

//Start the app
func (m *monitor) Start() error {
	signal.Notify(m.interrupt, os.Interrupt)

	go func() {
		m.complete <- m.run()
	}()

	select {
	case err := <-m.complete:
		return err
	case <-m.time:
		return ErrTimeOut
	}
}

func (m *monitor) run() error {
	for id, task := range m.tasks {
		if m.gotInterruped() {

			return ErrInterrupted
		}
		task(id)
	}
	return nil
}

func (m *monitor) gotInterruped() bool {
	select {
	case <-m.interrupt:
		signal.Stop(m.interrupt)
		return true
	default:
		return false
	}
}

//Add more worker functions to the stack
func (m *monitor) Add(fn ...func(int)) {
	m.tasks = append(m.tasks, fn...)
}

//NewMonitor return new runner app
func NewMonitor(d time.Duration) Runner {
	return &monitor{
		time:      time.After(d),
		complete:  make(chan error),
		interrupt: make(chan os.Signal),
	}
}
