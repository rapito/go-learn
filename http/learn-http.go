package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"github.com/parnurzeal/gorequest"
)

var _ = json.NewEncoder(os.Stdout)

func httpPlain() {
	resp, err := http.Get("http://ip.jsontest.com/")

	if err != nil {
		fmt.Println("Error on GET ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body ", err)
	}


	result := string(body[:])

	fmt.Println("Body result: ", result)

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error reading json from body ", err)
	}
	fmt.Println(data)

	ip := data["ip"].(string)
	fmt.Println("Json: ", string(ip))
}

func httpGoRequest() {
	request := gorequest.New()

	resp, body, errs := request.Get("https://ip.jsontest.com").End()
	fmt.Println(resp)
	fmt.Println(body)
	fmt.Println(errs)
}

func main() {
	//	httpPlain()
	httpGoRequest()

}
