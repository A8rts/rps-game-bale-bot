package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Update struct {
	Update_id int
}

type GetUpdatesRequestBody struct {
	Offset int
}

type GetUpdatesResponse struct {
	Ok     bool
	Result []Update
}

var API_URL = "https://tapi.bale.ai/bot1656532239:LtLZkHFiul7s7HmWoxABejYZQvIONHkcHfBlYHR4"
var offset = 0

func main() {
	for {
		checkForUpdates()
	}
}

func checkForUpdates() {
	body, err := json.Marshal(GetUpdatesRequestBody{0})

	if err != nil {
		fmt.Println(err)
	}

	requestBody := bytes.NewBuffer(body)

	res, err := http.Post(API_URL+"/getUpdates", "application/json", requestBody)

	if err != nil {
		fmt.Println(err)
	}

	var responseData GetUpdatesResponse
	err = json.NewDecoder(res.Body).Decode(&responseData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(responseData)
}
