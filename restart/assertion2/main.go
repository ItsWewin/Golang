package main

import "fmt"

type Phone struct{}
type Mp3 struct{}
type Mp4 struct{}

type Usb interface {
	storage()
}

type Computer struct{}

func (p Phone) storage() {
	fmt.Printf("phone's storage is 32 GB\n")
}

func (p Mp3) storage() {
	fmt.Printf("mp3's storage is 8 GB\n")
}

func (p Mp4) storage() {
	fmt.Printf("mp4's storage is 16 GB\n")
}

func (p Mp4) video() {
	fmt.Println("mp4 can play video!")
}

// func (c Computer) readDeviceType(device interface{}) {
func (c Computer) readDeviceType(device Usb) {
	switch device.(type) {
	case Phone:
		fmt.Println("this device is phone")
	case Mp4:
		// device.video() // 错误： device.video undefined (type Usb has no field or method video)go
		device.(Mp4).video() // ok
		fmt.Println("this device is Mp4")
	case Mp3:
		fmt.Println("this device is Mp3")
	default:
		fmt.Println("don't know!")
	}
}

func main() {
	c := Computer{}
	p := Phone{}
	mp3 := Mp3{}
	mp4 := Mp4{}

	c.readDeviceType(p)
	c.readDeviceType(mp3)
	c.readDeviceType(mp4)
}
