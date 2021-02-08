package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"./readconfig"
	"path/filepath"
)

type User struct {
	Login    string `json:"username"`
	Password string `json:"password"`
}

type Cmd struct {
	Login string `json:"username"`
}

/*
func login() {
	user := &User{Login: "****", Password: "****"}
	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonStr))

	url := "http://host:8080/signin"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	fmt.Println("req:>", req)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
*/
//
func WeatherGET(serv_url string) {
	//user := &User{Login: dev_name}
	//jsonStr, err := json.Marshal(user)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(jsonStr))

	//url := serv_url + "/v1/commands/?device=" + dev_name
	url := serv_url
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	cmds := string(body)
	fmt.Println("response Body cmds:", string(cmds))
}



func main() {
	//login()

	//http://api.openweathermap.org/data/2.5/weather?q=Moscow,ru&appid=48dc6dc3b49c57b1eb4be6a6ba245009
	readcfg := readconfig.Config_reader("./readconfig/dss.conf")
	//server_url :="http://"+readcfg.Host+":"+strconv.Itoa(readcfg.Port)
	server_url :="http://api.openweathermap.org/data/2.5/weather?q="+readconfig.City+","+readconfig.CCode+"&appid="+readconfig.token
	WeatherGET(server_url)


}