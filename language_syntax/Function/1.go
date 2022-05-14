package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	ID   int
	name string
}

type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {
	u := &user{
		1234,
		"huang",
	}

	if _, err := updateUser(u); err != nil {
		fmt.Println(err)
		return
	}

}

func updateUser(u *user) (*updateStats, error) {
	response := `{"Modified": 12, "Duration": 123.123, "Success": true, "Message": "Mission Success!"}`

	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	fmt.Println(&us)

	return &us, nil
}

// go install -v github.com/ramya-rao-a/go-outline@latest
// go install -v github.com/stamblerre/gocode@latest

// go env -w GO111MODULE=on
// go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
