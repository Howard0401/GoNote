package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var created struct {
	API string `json:"api"`
	ID  string `json:"id"`
}

func main() {
	address := flag.String("server", "http://localhost:8080", "Http gateway Url, eg. http://localhost:8080")
	flag.Parse()

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)

	//Create
	Create(address, pfx)

	Read(address, pfx)

	Update(address, pfx)

	ReadAll(address, pfx)

	Delete(address, pfx)

}

func Create(address *string, pfx string) {
	res, err := http.Post(*address+"/v1/todo", "application/json", strings.NewReader(fmt.Sprintf(
		`{
			"api":"v1",
			"toDo": {
				"title":"title (%s)",
				"description":"description (%s)",
				"reminder":"%s"
			}
		}`, pfx, pfx, pfx,
	)))

	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var body string
	if err != nil {
		body = fmt.Sprintf("failed read Create response body:%v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n\n", res.StatusCode, body)

	err = json.Unmarshal(bodyBytes, &created)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
	}
}

func Read(address *string, pfx string) {
	res, err := http.Get(fmt.Sprintf("%s%s%s", *address, "/v1/todo", created.ID))
	if err != nil {
		log.Fatalf("failed to call Read method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	var body string
	if err != nil {
		body = fmt.Sprintf("failed to read response body:%v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Read response: Code=%d, Body=%s\n\n", res.StatusCode, body)
}

func Update(address *string, pfx string) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID),
		strings.NewReader(fmt.Sprintf(`
	{
		"api":"v1",
		"toDo":{
			"title":"title (%s) +updated",
			"description":"description (%s) + updated",
			"reminder": "%s"
		}
	}
`, pfx, pfx, pfx)))
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to call Update method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var body string
	if err != nil {
		body = fmt.Sprintf("failed read Update response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Update response: Code=%d. Body=%s\n\n", res.StatusCode, body)
}

func ReadAll(address *string, pfx string) {
	res, err := http.Get(*address + "/v1/todo/all")
	if err != nil {
		log.Fatalf("failed to call ReadAll method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var body string
	if err != nil {
		body = fmt.Sprintf("failed to read all response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("ReadAll response: Code=%d, Body=%s\n\n", res.StatusCode, body)
}

func Delete(address *string, pfx string) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.ID), nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to call Delete method: %v", err)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var body string
	if err != nil {
		body = fmt.Sprintf("failed read Delete response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Delete reponse: Code=%d, Body=%s\n\n", res.StatusCode, body)
}
