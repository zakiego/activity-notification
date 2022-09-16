package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func SendNotif(message string) {
	const BOT_TOKEN = "################################" // ex: 1234567890:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
	const GROUP_ID = "#########"                         // ex: -123456789

	url := "https://api.telegram.org/bot" + BOT_TOKEN + "/sendMessage?chat_id=" + GROUP_ID + "&text=" + message

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func IsOnline() bool {
	timeout := time.Duration(5000 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}

	_, err := client.Get("https://google.com")

	if err != nil {
		return false
	}

	return true
}
