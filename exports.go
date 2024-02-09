package main

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 3 * time.Second}

var publicChannels = make(map[string]string)
var privateChannels = make(map[string]string)

var webhookId string = ""
var webhookToken string = ""
