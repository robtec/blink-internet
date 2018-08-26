package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hink/go-blink1"
)

const (
	pingTimeout = time.Duration(10 * time.Second)
	pingEvery   = time.Duration(20 * time.Second)
)

var (
	blue  = blink1.State{Blue: 20}
	red   = blink1.State{Red: 20}
	green = blink1.State{Green: 20}
	amber = blink1.State{Red: 20, Green: 15}
	off   = blink1.State{}
)

func main() {

	device, err := blink1.OpenNextDevice()
	defer device.Close()

	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		stopBlink(device)
		os.Exit(1)
	}()

	fmt.Println("# monitoring the internet connection #")
	device.SetState(blue)

	for {

		err := connectedToInternet()

		switch err := err.(type) {

		case net.Error:
			if err.Timeout() {
				device.SetState(amber)
			} else {
				device.SetState(red)
			}

		case nil:
			device.SetState(green)
		}

		time.Sleep(pingEvery)
	}
}

func connectedToInternet() error {

	conn, err := net.DialTimeout("tcp", "www.google.com:80", pingTimeout)

	if err != nil {
		return err
	}

	conn.Close()

	return err
}

func stopBlink(d *blink1.Device) {

	fmt.Println("stopping blink1")
	d.SetState(off)
	d.Close()
}
