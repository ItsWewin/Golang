package main

import "fmt"

type usb interface {
	start()
	stop()
}

type Phone struct{}
type Camera struct{}
type Computer struct{}

func (p Phone) start() {
	fmt.Println("phone start")
}

func (p Phone) stop() {
	fmt.Println("phone start")
}

func (p Camera) start() {
	fmt.Println("Camera start")
}

func (p Camera) stop() {
	fmt.Println("Camera start")
}

func (c Computer) useUsb(device usb) {
	device.start()
	device.stop()
}

func main() {

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	camera2 := Camera{}
	devices := [3]usb{camera, phone, camera2}

	for _, device := range devices {
		computer.useUsb(device)
	}
}
