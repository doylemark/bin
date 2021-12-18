package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type AuthResponse struct {
	Token string `json:"token"`
}

func main() {
	client := &http.Client{Timeout: time.Second * 10}

	creds := map[string]string{
		"username":      "admin",
		"password":      "password",
		"twoFactorCode": "string",
	}

	data, err := json.Marshal(creds)

	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(data)

	req, err := http.NewRequest("POST", "http://localhost:11300/api/v1/Session", r)
	req.Header.Add("Content-Type", "application/json-patch+json")

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	var auth AuthResponse

	err = json.Unmarshal(bytes, &auth)

	if err != nil {
		panic(err)
	}

	fmt.Println(auth.Token)
}
