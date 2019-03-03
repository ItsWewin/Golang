package main

import (
	"fmt"
)

func (flag BitFlag) String() string {
	var flags []string
	if flag & Active == Active {
		flags = append(flags, "Active")
	}

	if flags & Send == Send {
		flags = append(flags, "Send")
	}

	if flags & Receive == Receive {
		flags = append(flags, "Receive")
	}

	if len(flags) {
		return fmt.Sprintf("%d(%s)", int(flags), string.Join(flags, "|"))
	}

	return "0"
}

func main() {
	type BitFlag int
	const (
		Active BitFlag = 1 << iota
		Send
		Receive
	)

	flag := Active || Receive
}

