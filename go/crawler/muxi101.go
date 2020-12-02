package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	requestURL := "http://pass.muxi-tech.xyz/auth/api/signin"

	//data := url.Values{}
	//data.Set("password", "略")
	//data.Set("username", "略")

	playload := strings.NewReader(`{"password":"略","username":"略"}`)

	request, err := http.NewRequest("POST", requestURL, playload)
	//didi, _ := ioutil.ReadAll(request.Body)
	//fmt.Println(string(didi))
	if err != nil {
		panic(err)
	}

	//request.Header.Add("Content-Type", "application/json")
	//request.Header.Add("User-Agent", "User-Agent")
	//request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	//request.Header.Add("Connection", "keep-alive")
	//request.Header.Add("")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)

	fmt.Println(string(body))

}
