package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("RbxDeployPrivacy.go\n.:Developer: @Mast3rGamers on GitHub:.\n")

	for {
		list := getChannelList()

		for _, channel := range list {
			checkChannel(strings.ToLower(channel))
		}

		time.Sleep(12 * time.Second)
	}
}
