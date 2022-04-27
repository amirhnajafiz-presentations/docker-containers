package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Test() {
	url := "https://api-nba-v1.p.rapidapi.com/seasons"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "d9ebf01ef3mshb3ed3ad38d37daep179ba1jsn50dda0c0097d")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
