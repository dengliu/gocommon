package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type ip struct {
	Origin string `json:"origin"`
}

type user struct {
	Name name `json:"name"`
	IPs  []ip `json:"ips"`
}

func main() {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	fmt.Println("Starting the application...")
	request, _ := http.NewRequest(http.MethodGet, "https://httpbin.org/ip", nil)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)

		return
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	d := ip{}
	if err := json.Unmarshal(data, &d); err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Println(d)
	fmt.Println(string(data))

	n := name{
		FirstName: "Nic",
		LastName:  "Raboy",
	}
	jsonValue, _ := json.Marshal(n)

	request, _ = http.NewRequest(http.MethodPost, "https://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	response, err = client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	myJson := bytes.NewBuffer([]byte(`{"name":"Maximilien"}`))
	resp, err := client.Post("https://www.google.com", "application/json", myJson)
	if err != nil {
		fmt.Errorf("Error %s", err)
		return
	}

	defer resp.Body.Close()
	data, _ = ioutil.ReadAll(response.Body)
	fmt.Println(data)
}
