package main

import (
	"fmt"
)

func checkChannel(channel string) {
	if channel == "" {
		return
	}
	res, err := httpClient.Get("https://clientsettings.roblox.com/v2/client-version/WindowsPlayer/channel/" + channel)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	defer res.Body.Close()

	valueOfPublic := publicChannels[channel]
	valueOfPrivate := privateChannels[channel]

	status := res.StatusCode

	switch {
	case status == 401 && valueOfPublic != "":
		delete(publicChannels, channel)
		privateChannels[channel] = "private"
		fmt.Printf("[%s] CHANNEL IS NOW PRIVATE!\n", channel)
		sendMessage("||@everyone|| [" + channel + "] CHANNEL IS NOW PRIVATE!")
	case status == 200 && valueOfPrivate != "":
		delete(privateChannels, channel)
		publicChannels[channel] = "public"
		fmt.Printf("[%s] CHANNEL IS NOW PUBLIC\n", channel)
		sendMessage("||@everyone|| [" + channel + "] CHANNEL IS NOW PUBLIC!")
	case status == 200 && valueOfPublic == "":
		publicChannels[channel] = "public"
		fmt.Printf("[%s] Public channel\n", channel)
		sendMessage("||@everyone|| [" + channel + "] Public channel")
	case status == 401 && valueOfPrivate == "":
		privateChannels[channel] = "private"
	}

	return
}
