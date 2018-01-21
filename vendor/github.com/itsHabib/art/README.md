# art - API Resource Tool

## Overview
art is a command line tool that can be used to interact with a REST API and to ensure the proper
responses are sent back. Given a base URL the API struct maps resources to specific
RestResources. Each RestResource also contains a Router which maps status codes to
callback response functions. 

## Usage
The example shows how to interact with  an API with a base URL of https://httpbin.org and a resource enpoint of /get.
The callback response function in this example just returns a nil error when a 200 status code
is given and also prints the contents from the request. 

```golang
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/itsHabib/art"
)

var api = art.NewAPI("https://httpbin.org")

func main() {
	list := flag.Bool("list", false, "Get list of all API resources")
	flag.Parse()
	if *list {
		fmt.Println("Available Resources")
		for _, name := range api.ResourceNames() {
			fmt.Println(name)
		}
		return
	}
	resource := os.Args[1]
	if err := api.Call(resource, nil); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	router := art.NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(content))
		return nil
	})
	api.AddResource("get", art.NewResource("/get", "GET", router))
}
```
