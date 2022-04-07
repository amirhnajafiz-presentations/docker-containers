package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api-nba-v1.p.rapidapi.com/seasons"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "d9ebf01ef3mshb3ed3ad38d37daep179ba1jsn50dda0c0097d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
