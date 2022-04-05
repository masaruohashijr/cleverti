package service

import (
	"bytes"
	"cleverti/pkg/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallPrinter(ds domain.Dataset) {
	jsonStr, err := json.Marshal(ds)
	url := "http://localhost:8888/printer"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "datasets")
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
	fmt.Println("response Body:", string(body))

}
