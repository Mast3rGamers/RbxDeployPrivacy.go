package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func getChannelList() []string {
	res, err := httpClient.Get("https://raw.githubusercontent.com/bluepilledgreat/Roblox-DeployHistory-Tracker/main/ChannelsAll.json")
	if err != nil {
		fmt.Println("error on channel list fetch:", err)
		return nil
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error on read:", err)
		return nil
	}

	var channelList []string
	json.Unmarshal(body, &channelList)

	return channelList
}
